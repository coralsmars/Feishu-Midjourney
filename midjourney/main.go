package main

import (
	"fmt"
	"midjourney/handlers"
	"midjourney/initialization"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/pflag"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	cfg := pflag.StringP("config", "c", "./config.yaml", "api server config file path.")
	f, err := os.OpenFile("discord.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	//strP: = pp.Sprint(Person{Name: "Alice", Age: 20})
	fmt.Print(f, pp.Sprint(Person{Name: "Alice", Age: 20}))
	pflag.Parse()

	initialization.LoadConfig(*cfg)
	initialization.LoadDiscordClient(handlers.DiscordMsgCreate, handlers.DiscordMsgUpdate)

	r := gin.Default()

	r.POST("/v1/trigger/midjourney-bot", handlers.MidjourneyBot)
	r.POST("/v1/trigger/upload", handlers.UploadFile)

	r.Run(":16007")
}
