package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("id")
		if sessionID == nil {
			c.JSON(http.StatusNotFound,
				gin.H{"message": "unauthorized"})
			c.Abort()
		}
	}
}
