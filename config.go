package mobile_sdk

import (
	"encoding/json"
)

//	type Config struct {
//		TokenAPIBaseURL string `json:"token_api_base_url"`
//		APIBaseURL      string `json:"api_base_url"`
//		Appid           string `json:"appid"`
//		AppSecret       string `json:"app_secret"`
//		ChannelId       string `json:"channel_id"`
//		Timeout         int    `json:"timeout"`
//	}
type Config struct {
	ApiUrl    string `json:"api_url"`
	Appid     string `json:"appid"`
	AppSecret string `json:"app_secret"`
	Timeout   int    `json:"timeout"`
	TransID   string `json:"trans_id"`
	Token     string `json:"token"`
}

// NewConfig 创建一个新的配置实例
func NewConfig(apiBaseURL, appid, appSecret string) *Config {
	transID := NewTransIDGenerator(appid).Generate()
	return &Config{
		ApiUrl:    apiBaseURL,
		Appid:     appid,
		AppSecret: appSecret,
		TransID:   transID,
	}
}

func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}

type TPConfig struct {
	ApiUrl    string `json:"api_url"`
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Apikey    string `json:"apikey"`
	ChannelID string `json:"channel_id"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Times     string `json:"times"`
}

// http://token.dctxiot.com/api/token/getToken
func NewTP1Config(apiUrl, appid, appSecret, ChannelID string) *TPConfig {
	times := GetTimestamp()
	sign := GenerateTP1Sign(appid, appSecret, times)
	return &TPConfig{
		ApiUrl:    apiUrl,
		AppID:     appid,
		AppSecret: appSecret,
		ChannelID: ChannelID,
		Times:     times,
		Sign:      sign,
	}
}

// https://ac.buleideiot.com/api/get-chinamobile-onelink-token
func NewTP2Config(apiUrl, apikey, ChannelID string) *TPConfig {
	sign := GenerateTP2Sign(ChannelID, apikey)
	return &TPConfig{
		ApiUrl:    apiUrl,
		ChannelID: ChannelID,
		Apikey:    apikey,
		Sign:      sign,
	}
}

func (c *TPConfig) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}
