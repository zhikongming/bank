package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhikongming/bank/config"
	"github.com/zhikongming/bank/dal"
)

func main() {
	r := gin.Default()

	config.InitConfig()
	dal.InitDB(config.GetConfig().DB.WriteDSN, config.GetConfig().DB.ReadDSN)

	RegisterRouter(r)
	r.Run(":8080")
}
