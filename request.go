package main

type GetTokenReq struct {
	Appid     string `json:"appid"`
	AppSecret string `json:"app_secret"`
	ChannelId string `json:"channel_id"`
	Times     string `json:"times"`
	Sign      string `json:"sign"`
}
