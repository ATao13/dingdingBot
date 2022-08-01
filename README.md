#### 钉钉机器人项目


- ###   简介
    因工作需要使用钉钉机器人发送监控信息，编写此程序以方面调用。
    本程序支持关键词和加签来两种验证方式，使用者可自由选择
    目前仅支持text消息格式的发送，其他格式将后续补充

- ###   调用步骤

```go
//1创建客户端
bot := dingdingBot.InitBot("https://oapi.dingtalk.com/robot/send?access_token=xxxx", "SECxxxxxxxxx")
//发送消息
bot.SendTextMessage("hello", false, nil, nil)
```
- ###   代码示例
- ```go
package main

import "dingdingBot/src/dingdingBot"

func main() {

	bot := dingdingBot.InitBot("https://oapi.dingtalk.com/robot/send?access_token=xxxx", "xxxxx")
	bot.SendTextMessage("hello", false, nil, nil)
}

```
