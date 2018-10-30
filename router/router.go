package router

import (
	"captchademo/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRouter(router *gin.Engine) {
	router.GET("/getCaptcha", controller.GetCaptcha)
	router.GET("/verifyCaptcha", controller.VerifyCaptcha)
	router.GET("/show/:source", controller.GetCaptchaPng)
}
