package dingdingBot

import (
	"reflect"
	"testing"
)

func TestDingTalkBot_SendTextMessage(t *testing.T) {
	bot := InitBot("https://oapi.dingtalk.com/robot/send?access_token=xxxxxx", "xxxxxxxxx")
	type args struct {
		msg             string
		IsAtAll         bool
		at_dingtalk_ids []string
		atmobiles       []string
	}
	tests := []struct {
		name     string
		args     args
		wantResp interface{}
		wantErr  bool
	}{
		{
			name: "test1",
			args: args{
				"hello", false, nil, nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := bot.SendTextMessage(tt.args.msg, tt.args.IsAtAll, tt.args.at_dingtalk_ids, tt.args.atmobiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendTextMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendTextMessage() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
