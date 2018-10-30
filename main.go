package main

import (
	"captcha/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.New()
	ginRouter.GET("/getCaptcha", controller.GetCaptcha)
	ginRouter.GET("/verifyCaptcha", controller.VerifyCaptcha)
	ginRouter.GET("/show/:source", controller.GetCaptchaPng)
	ginRouter.Run(":8080")
}
