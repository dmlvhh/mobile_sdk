package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		20000,
	)
	//创建客户端
	cli := NewClient(cfg)

	//fmt.Println(sign, cli)
	res, err := GetToken(cli, &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		ChannelId: "3",
		Times:     timestamp,
		Sign:      sign,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Data)
}
