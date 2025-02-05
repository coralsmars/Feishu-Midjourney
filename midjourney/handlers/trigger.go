package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestTrigger struct {
	Type         string `json:"type"`
	DiscordMsgId string `json:"discordMsgId,omitempty"`
	MsgHash      string `json:"msgHash,omitempty"`
	Prompt       string `json:"prompt,omitempty"`
	Index        int64  `json:"index,omitempty"`
}

func MidjourneyBot(c *gin.Context) {
	var body RequestTrigger
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var err error
	switch body.Type {
	case "generate":
		err = GenerateImage(body.Prompt)
	case "upscale":
		err = ImageUpscale(body.Index, body.DiscordMsgId, body.MsgHash)
	case "variation":
		err = ImageVariation(body.Index, body.DiscordMsgId, body.MsgHash)
	case "maxUpscale":
		err = ImageMaxUpscale(body.DiscordMsgId, body.MsgHash)
	case "reset":
		err = ImageReset(body.DiscordMsgId, body.MsgHash)
	case "describe":
		err = ImageDescribe(body.Prompt)

	case "blend":
		//fmt.Printf("body.Prompt %s", body.Prompt)
		strNameList := strings.Split(body.Prompt, ";")
		err = ImageBlend(strNameList)
	default:
		err = errors.New("invalid type")
	}

	if err != nil {
		fmt.Println("response err", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
