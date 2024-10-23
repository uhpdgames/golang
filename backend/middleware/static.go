package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeStaticOrAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api/") {
			c.Next()
			return
		}

		if strings.HasPrefix("static", "./static"+path) {
			c.File("./static/index.html")
			return
		}
	}
}
