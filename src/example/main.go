package main

import "dingdingBot/src/dingdingBot"

func main() {

	bot := dingdingBot.InitBot("https://oapi.dingtalk.com/robot/send?access_token=480b3337b3a161d54be35f63a6347f94ad4afcf011add595ecee71658aa41cc4", "SECf30050f98dea590e4b10bba8dc75dd0785eca91cc775290803fce46e2760e95a")
	bot.SendTextMessage("hello", false, nil, nil)
}
