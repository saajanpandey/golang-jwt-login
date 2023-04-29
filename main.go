package main

import (
	"example/go-jwt/initializers"
	"example/go-jwt/middleware"

	"github.com/gin-gonic/gin"

	"example/go-jwt/controllers"
)

func init(){
	 initializers.LoadEnvVariables()
	 initializers.DatabaseConnect()
	 initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	
	r.POST("/signup",controllers.SignUp)
	r.POST("/login",controllers.Login)
	r.GET("/validate",middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
