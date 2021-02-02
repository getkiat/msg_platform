package api

import (
	"msg_platform/server/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router   *gin.Engine
	HTTPPort int
}

func NewAPI(config config.Config) *API {
	return &API{
		Router:                   gin.New(),
		HTTPPort:                 config.HTTPPort,
	}
}


type HelloResponse struct {
	ContactNo string `json:"contactNo" binding:"required"`
	Message string `json:"message" binding:"required"`
	Source string `json:"source" binding:"required"`
}

func yourHelloWorld(c *gin.Context) {
	var json HelloResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": json.Message,
	})
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}