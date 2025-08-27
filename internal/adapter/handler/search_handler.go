package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchHandler handles incoming HTTP requests for the /search endpoint.
type SearchHandler struct {
	// This will hold a reference to the use case in a future task.
}

// NewSearchHandler creates a new SearchHandler.
func NewSearchHandler() *SearchHandler {
	return &SearchHandler{}
}

// Search is the placeholder method for the search endpoint.
func (h *SearchHandler) Search(c *gin.Context) {
	// Placeholder response as per the task requirements.
	c.JSON(http.StatusOK, gin.H{
		"status": "Search endpoint is up!",
	})
}

// RegisterRoutes sets up the routing for the search handler.
func (h *SearchHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/search", h.Search)
}
