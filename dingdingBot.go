package dingdingBot

//钉钉Webhook机器人使用，支持 text，link、actionCard、markdown 格式消息发送
// 兼容sercet模式

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type textT struct {
	Content string `json:"content"`
}
type atInfo struct {
	AtMobiles []string `json:"atmobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isatall"`
}
type textMessage struct {
	At      atInfo `json:"at"`
	Text    textT  `json:"text"`
	Msgtype string `json:"msgtype"`
}

type DingTalkBot struct {
	client *http.Client
	url    string
}

func getSign(secret string) (sign string, timeStamp string) {
	//获取签名和时间戳
	timeStamp = strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	sign_param := timeStamp + "\n" + secret
	hmacCode := hmac.New(sha256.New, []byte(secret))
	hmacCode.Write([]byte(sign_param))
	sign = base64.StdEncoding.EncodeToString(hmacCode.Sum(nil))
	return
}

func InitBot(url string, secret string) *DingTalkBot {
	//初始化 client
	// ：param url 钉钉机器人调用链接
	// ：param secret 签名，如果不使用加签验证可穿空值
	if len(secret) != 0 {
		sign, timeStamp := getSign(secret)
		url = fmt.Sprintf("%s&timestamp=%s&sign=%s", url, timeStamp, sign)
	}
	return &DingTalkBot{
		&http.Client{},
		url,
	}
}
func (bot *DingTalkBot) SendTextMessage(msg string, IsAtAll bool, at_dingtalk_ids []string, atmobiles []string) (resp interface{}, err error) {
	//  text类型
	// :param msg: 消息内容
	// :param IsAtAll: @所有人时：true，否则为false（可选）
	// :param atmobiles: 被@人的手机号（注意：可以在msg内容里自定义@手机号的位置，也支持同时@多个手机号，可选）
	// :param at_dingtalk_ids: 被@人的dingtalkId（可选）
	// :return: 返回消息发送结果
	text_info := new(textMessage)
	text_info.Text.Content = msg
	text_info.At.IsAtAll = IsAtAll
	text_info.At.AtUserIds = at_dingtalk_ids
	text_info.At.AtMobiles = atmobiles
	text_info.Msgtype = "text"
	text, err := json.Marshal(text_info)
	req, _ := http.NewRequest("POST", bot.url, strings.NewReader(string(text)))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	re, err := bot.client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer re.Body.Close()
	body, err := ioutil.ReadAll(re.Body)
	resp = string(body)
	return resp, err
}
