package middleware

import (
	"github.com/Deseao/anon/api/internal/group"
	"github.com/gin-gonic/gin"
)

func GroupHandler(c *gin.Context) {
	handler := group.GroupHandler{}
	c.Set("groupHandler", handler)
}
