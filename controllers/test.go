package controllers

import (
	"github.com/ANPez/gogsi/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	user := c.MustGet("user").(*types.User)
	c.JSON(http.StatusOK, gin.H{"ok": true, "hi": user.Email})
}
