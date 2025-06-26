package v1

import (
	"core/httputil"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping API
// @Description Ping API
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Ping successfully"
// @Router /ping [get]
func Ping(c *gin.Context) {
	httputil.Success(c, gin.H{
		"status": "ok",
	})
}

// GetConfig godoc
// @Summary Get config API
// @Description Get config
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Get config successfully"
// @Router /config [get]
func GetConfig(c *gin.Context) {
	httputil.Success(c, gin.H{})
}
