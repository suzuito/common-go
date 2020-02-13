package cgin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Usecors ...
func Usecors(app ApplicationGin, root *gin.Engine) {
	root.Use(cors.New(cors.Config{
		AllowOrigins: app.AllowedOrigins(),
		AllowMethods: app.AllowedMethods(),
		AllowHeaders: app.AllowedHeaders(),
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: app.AllowedCredential(),
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))
}
