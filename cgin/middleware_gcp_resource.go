package cgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgcp"
)

func MiddlewareGCPResource(generator *cgcp.GCPContextResourceGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := generator.Gen(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, struct {
				Message string `json:"message"`
			}{
				Message: "Internal server error",
			})
			return
		}
		SetGCPResource(c, r)
		c.Next()
	}
}
