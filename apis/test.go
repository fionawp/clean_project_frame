package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/clean_project_frame/common"
	"github.com/clean_project_frame/context"
)

type item struct {
	Name string
}

func Test(router *gin.RouterGroup, conf *context.Config) {
	router.GET("/test", func(c *gin.Context) {
		var myLogger = conf.GetLog()
		name := c.Query("name")
		if name == "" {
		    common.FormatResponseWithoutData(c, 200001, "param name is must")
		    return
		}

		myLogger.Info("param " + name)
		common.FormatResponse(c, 10000, "", &item{
		    Name : name,
		})
	})
}
