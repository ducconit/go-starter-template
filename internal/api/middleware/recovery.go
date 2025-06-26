package middleware

import (
	"core/httputil"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery middleware recovers from any panics and writes a 500 if there was one.
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		httputil.InternalServerError(c, "Internal server error")
		c.Abort()
	})
}
