package cmd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// bearer token structure
type Token struct {
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// response structure
type Response struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// user response structure
type ResponseUser struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Enabled  bool   `json:"enabled"`
}

// realm payload structure
type PayloadRealm struct {
	ID                  string `json:"id"`
	Realm               string `json:"realm"`
	AccessTokenLifespan int    `json:"accessTokenLifespan"`
	Enabled             bool   `json:"enabled"`
	SslRequired         string `json:"sslRequired"`
	BruteForceProtected bool   `json:"bruteForceProtected"`
	LoginTheme          string `json:"loginTheme"`
	EventsEnabled       bool   `json:"eventsEnabled"`
	AdminEventsEnabled  bool   `json:"adminEventsEnabled"`
}

// user payload structure
type PayloadUser struct {
	FirstName   string            `json:"firstName"`
	LastName    string            `json:"lastName"`
	UserName    string            `json:"username"`
	Email       string            `json:"email"`
	Enabled     bool              `json:"enabled"`
	Credentials []PayloadPassword `json:"credentials"`
}

// password payload structure
type PayloadPassword struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary string `json:"temporary"`
}

// secret returns a 32 bytes AES key.
var secret = []byte{47, 46, 57, 24, 85, 35, 24, 47, 87, 35, 88, 98, 66, 32, 14, 47}

// client id
var clientId string = "admin-cli"

// client token
var clientToken string = ""

// keycloak url
var keycloakURL string = ""

// file name to process
var textFile string = ""

// format of text file
var textFormat string = ""

// delimiter of text file
var textDelim string = ""

// encrypt encrypts plain string with a secret key and returns encrypt string.
func encrypt(plainData string, secret []byte) (string, error) {
	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(plainData), nil)), nil
}

// decrypt decrypts encrypt string with a secret key and returns plain string.
func decrypt(encodedData string, secret []byte) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}

// get bearer token
func getToken() (string, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	params.Add("client_id", clientId)
	params.Add("client_secret", clientToken)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", keycloakURL+"/realms/master/protocol/openid-connect/token", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	token := new(Token)
	err = json.NewDecoder(resp.Body).Decode(token)
	if err != nil {
		return "", err
	}
	if len(token.Error) > 0 {
		return "", errors.New(token.ErrorDescription)
	}

	return token.AccessToken, nil
}

// save configuration to ini file
func saveIni() error {
	inidata := ini.Empty()
	sec, err := inidata.NewSection("keycloak")
	if err != nil {
		return err
	}
	_, err = sec.NewKey("url", keycloakURL)
	if err != nil {
		return err
	}
	sec, err = inidata.NewSection("client")
	if err != nil {
		return err
	}
	_, err = sec.NewKey("id", clientId)
	if err != nil {
		return err
	}
	encrypted, err := encrypt(clientToken, secret)
	if err != nil {
		return err
	}
	_, err = sec.NewKey("token", encrypted)
	if err != nil {
		return err
	}
	err = inidata.SaveTo(".keycloak")
	if err != nil {
		return err
	}
	return nil
}

// read configuration from ini file
func loadIni() error {
	inidata, err := ini.Load(".keycloak")
	if err != nil {
		fmt.Println("Failed to read authorization config, run keycloak auth before running other commands")
		os.Exit(1)
	}

	sec := inidata.Section("keycloak")
	keycloakURL = sec.Key("url").String()
	sec = inidata.Section("client")
	clientId = sec.Key("id").String()

	encrypted := sec.Key("token").String()
	decrypted, err := decrypt(encrypted, secret)
	if err != nil {
		return err
	}
	clientToken = decrypted
	return nil
}

// post data to keycloak
func postData(token string, path string, data interface{}, target interface{}) error {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", keycloakURL+"/admin/realms"+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// get data from keycloak
func getData(token string, path string, target interface{}) error {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", keycloakURL+"/admin/realms"+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// check error
func check(err error) {
	if err != nil {
		panic(err)
	}
}