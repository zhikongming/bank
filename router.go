package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhikongming/bank/handler"
)

func RegisterRouter(r *gin.Engine) {
	r.LoadHTMLGlob("template/*")
	r.GET("/ping", handler.Ping)

	p := r.Group("/platform")
	{
		p.GET("/ping", handler.Ping)
	}

	s := r.Group("/open/v1")
	{
		s.GET("/ping", handler.Ping)
	}
}
