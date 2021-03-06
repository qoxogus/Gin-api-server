package main

import (
	"Gin-api-server/config"
	"Gin-api-server/controller"
	"Gin-api-server/database"
	"Gin-api-server/library/jwt"
	"Gin-api-server/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type serverMethod interface {
	main()
}

func main() {
	config.InitConfig()
	database.Connect()

	r := gin.Default()

	r.Use(middleware.CORSmiddleware())

	r.GET("/server-test", controller.ServerTest)
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.Signin)
	r.POST("logout", controller.Logout)
	r.POST("/token-test", jwt.VerifyAccessToken, controller.TokenTest)
	r.POST("/re-token", jwt.VerifyRefreshToken, jwt.CreateReissuanceToken, controller.TokenTest)

	r.Run(":3000")
}
