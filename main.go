package main

import (
	"auction/utils"
	"time"

	"auction/config"
	database "auction/db"
	"auction/logs"
	common "auction/routers/comm"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {

	// init database
	db := database.Init()
	config.InitConfig("")

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	// start timer
	utils.Timer()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v := r.Group("/api/v1")
	common.HostManager(v)
	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}
