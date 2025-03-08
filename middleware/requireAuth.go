package middleware

import "github.com/gin-gonic/gin"

func RequireAuth(c *gin.Context) {
	// get cookie from request

	// decode/validate

	// check exporation

	// find user with token

	// attach req

	// continue
	c.Next()
}
