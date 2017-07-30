package middleware

import (
	"github.com/Deseao/anon/internal/group"
	"github.com/gin-gonic/gin"
)

func GroupHandler(handler *group.GroupHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("groupHandler", handler)
	}
}
