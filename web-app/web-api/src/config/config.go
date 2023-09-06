package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	Keycloak KeycloakStruct
}

type ConfigYaml struct {
	Keycloak KeycloakStruct `yaml:"keycloak"`
}

type KeycloakStruct struct {
	Url    string `yaml:"url"`
	Client string `yaml:"client"`
	Token  string `yaml:"token"`
}

// load configuration from env and config.yaml
func Load() (*ConfigStruct, error) {
	// read env variables
	keycloakClient := os.Getenv("KEYCLOAK_CLIENT")
	keycloakToken := os.Getenv("KEYCLOAK_TOKEN")

	// read config.json
	configYaml := &ConfigYaml{}
	err := defaults.Set(configYaml)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &configYaml)
	if err != nil {
		return nil, err
	}

	//fill config
	var config = new(ConfigStruct)
	config.Keycloak.Client = keycloakClient
	config.Keycloak.Token = keycloakToken

	return config, nil
}
