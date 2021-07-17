package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.HTML(http.StatusOK, "ping.html", gin.H{
		"msg": "ping",
	})
}
