package main

import (
	"captcha/router"
	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.New()
	router.ConfigRouter(ginRouter)
	ginRouter.Run(":8080")
}
