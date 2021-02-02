package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) InitRoutes() {
	// router := gin.New();

	//Global middleware
	a.Router.Use(gin.Logger()) //Logger
	a.Router.Use(gin.Recovery()) //Recover frm panic

	// Service static files
	a.Router.Static("/static", "./static")
	a.Router.LoadHTMLFiles("index.html")


	// // Catch all unmatched routes here
	a.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	

	//Route here
	apiNameSpace := a.Router.Group("api") //version control
	apiNameSpace.GET("/healthcheck",healthcheck)
	apiNameSpace.POST("/helloworld",yourHelloWorld)
	apiNameSpace.GET("/helloworld",helloWorld)

	// return router
}

func (a *API) Run() {
	a.InitRoutes()

	// Begin blocking to listen for incoming requests
	a.Router.Run(fmt.Sprintf(":%v", a.HTTPPort))
}


