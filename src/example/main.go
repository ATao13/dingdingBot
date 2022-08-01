package main

import "dingdingBot/src/dingdingBot"

func main() {

	bot := dingdingBot.InitBot("https://oapi.dingtalk.com/robot/send?access_token=xxxxxx", "xxxxxxxxx")
	bot.SendTextMessage("hello", false, nil, nil)
}
