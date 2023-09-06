package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	Data any `json:"data"`
}

type LogList struct {
	UID   string `json:"uid"`
	Stamp string `json:"stamp"`
}

type LogContent struct {
	UID     string   `json:"uid"`
	Stamp   string   `json:"stamp"`
	Content []string `json:"content"`
}

type HeaderStruct struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

type PreviewStruct struct {
	Headers []HeaderStruct           `json:"headers"`
	Items   []map[string]interface{} `json:"items"`
	Key     string                   `json:"key"`
}

// list of all file processes on server
func (env *Env) keycloakLogList(c *gin.Context) {
	// need to get a list of process logs
	var res = []LogList{
		{UID: "test1", Stamp: "stamp"},
		{UID: "test2", Stamp: "stamp"},
		{UID: "test3", Stamp: "stamp"},
		{UID: "test4", Stamp: "stamp"},
		{UID: "test5", Stamp: "stamp"},
		{UID: "test6", Stamp: "stamp"},
		{UID: "test7", Stamp: "stamp"},
		{UID: "test8", Stamp: "stamp"},
	}

	c.JSON(http.StatusOK, res)
}

// details of file process
func (env *Env) keycloakLogContent(c *gin.Context) {
	// return log content
	uid := c.Params.ByName("uid")

	var res = LogContent{
		UID:   uid,
		Stamp: "test",
		Content: []string{
			"line1",
			"line2",
		},
	}

	c.JSON(http.StatusOK, res)
}

// upload new file to server
func (env *Env) keycloakUploadFile(c *gin.Context) {
	// upload file to server
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Upload the file to specific dst.
	uid := uuid.New().String()
	err = c.SaveUploadedFile(file, "files/"+uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	//return uid of uploaded file
	c.String(http.StatusOK, uid)
}

// preview file process
func (env *Env) keycloakPreviewProcess(c *gin.Context) {
	// preview file processing
	var json struct {
		UID       string `json:"uid" binding:"required"`
		Operation string `json:"operation" binding:"required"`
		Format    string `json:"format" binding:"required"`
	}

	err := c.Bind(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// parse file with format option\
	var headers = []HeaderStruct{
		{Key: "username", Title: "keycloak.username"},
	}
	var res = PreviewStruct{
		Headers: headers,
		Items: []map[string]interface{}{
			{"username": "test1", "number": 1, "password": "test"},
		},
		Key: "username",
	}
	// and return its content
	c.JSON(http.StatusOK, res)
}

// process file
func (env *Env) keycloakProcessFile(c *gin.Context) {
	// process file in keycloak
	var json struct {
		Operation string `json:"operation" binding:"required"`
		Format    string `json:"format" binding:"required"`
	}

	err := c.Bind(&json)
	if err == nil {
		//db[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
