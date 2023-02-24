package main

import (
	"github.com/ATao13/dingdingBot/src"
)

func main() {

	bot := src.InitBot("https://oapi.dingtalk.com/robot/send?access_token=xxxxxx", "xxxxxxxxx")
	bot.SendTextMessage("hello", false, nil, nil)
}
