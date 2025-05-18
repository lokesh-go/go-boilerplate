package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lokesh-go/go-api-service/api/rest/middleware"
	"github.com/lokesh-go/go-api-service/internal/dependencies"
)

// Handler
type Handler struct {
	deps *dependencies.Deps
}

// New handler
func New(d *dependencies.Deps) *Handler {
	return &Handler{
		deps: d,
	}
}

// Ping ...
func (h *Handler) Ping(c *gin.Context) {
	// Get meta data
	var metadata middleware.RequestMetaData
	if ctxMetadata, exists := c.Get(middleware.CtxKeyReqMeta); exists {
		metadata = ctxMetadata.(middleware.RequestMetaData)
	}

	// Prepare response
	appConfig := h.deps.Config.Get().App
	res := map[string]interface{}{
		"appName":     appConfig.Name,
		"appVersion":  appConfig.Version,
		"description": appConfig.Description,
		"auther":      appConfig.Author,
		"clientIP":    metadata.ClientIP,
	}

	// Returns
	c.JSON(http.StatusOK, res)
}
