package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notifier/pkg/alertmanager"
	"notifier/pkg/grafana"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func Handler() *gin.Engine {
	router := gin.Default()

	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(logger.SetLogger())
	router.Use(helmet.Default())
	router.SetTrustedProxies(nil)

	groupRouter := router.Group("api/v1")
	groupRouter.POST("/grafana", grafanaAlertHandler)
	groupRouter.POST("/alertmanager", alertManagerHandler)

	return router

}

func grafanaAlertHandler(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	var grb grafana.GrafanaResponseBody
	err := decoder.Decode(&grb)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
		})
		fmt.Println(err)
	}
	err = grafana.GrafanaAlert(grb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}

}

func alertManagerHandler(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	var amr alertmanager.AlertManagerRequestBody
	err := decoder.Decode(&amr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "BadRequest",
		})
		fmt.Println(err)
	} else {
		err = alertmanager.AlertManagerReq(amr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	}

}
