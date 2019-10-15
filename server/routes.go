package server

import (
	"github.com/gin-gonic/gin"
	"github.com/clean_project_frame/apis"
	"github.com/clean_project_frame/context"
)

func registerRoutes(app *gin.Engine, conf *context.Config) {
	//routes
	searchPrefix := app.Group("/mytest")
	{
		apis.Test(searchPrefix, conf)
	}
}
