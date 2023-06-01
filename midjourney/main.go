package main

import (
	"midjourney/handlers"
	"midjourney/initialization"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	cfg := pflag.StringP("config", "c", "./config.yaml", "api server config file path.")
	pflag.Parse()

	initialization.LoadConfig(*cfg)
	initialization.LoadDiscordClient(handlers.DiscordMsgCreate, handlers.DiscordMsgUpdate)

	r := gin.Default()

	r.POST("/v1/trigger/midjourney-bot", handlers.MidjourneyBot)
	r.POST("/v1/trigger/upload", handlers.UploadFile)

	r.Run(":16007")
}
