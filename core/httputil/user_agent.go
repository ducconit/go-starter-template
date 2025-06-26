package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
)

func ParsingUserAgent(c *gin.Context) useragent.UserAgent {
	userAgentString := c.GetHeader("User-Agent")
	return useragent.Parse(userAgentString)
}
