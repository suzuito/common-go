package cgin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// UseCORS ...
func UseCORS(app ApplicationGin, root *gin.Engine) {
	root.Use(cors.New(cors.Config{
		AllowOrigins:     app.AllowedOrigins(),
		AllowMethods:     app.AllowedMethods(),
		AllowHeaders:     app.AllowedHeaders(),
		ExposeHeaders:    app.ExposeHeaders(),
		AllowCredentials: app.AllowedCredential(),
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))
}
