package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// list of all file processes on server
func (env *Env) keycloakLogList(c *gin.Context) {
	// need to get a list of process logs

	c.String(http.StatusOK, "success")
}

// details of file process
func (env *Env) keycloakLogContent(c *gin.Context) {
	// return log content

	c.String(http.StatusOK, "success")
}

// upload new file to server
func (env *Env) keycloakUploadFile(c *gin.Context) {
	// upload file to server
	// save it with datetime+uid name

	// and return datetime+uid

	c.String(http.StatusOK, "success")
}

// preview file process
func (env *Env) keycloakPreviewProcess(c *gin.Context) {
	// preview file processing
	// parse file with format option

	// and return its content

	c.String(http.StatusOK, "success")
}

// process file
func (env *Env) keycloakProcessFile(c *gin.Context) {
	// process file in keycloak

	// and save and return log

	c.String(http.StatusOK, "success")
}
