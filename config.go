package mobile_sdk

import "encoding/json"

type Config struct {
	TokenAPIBaseURL string `json:"token_api_base_url"`
	APIBaseURL      string `json:"api_base_url"`
	Appid           string `json:"appid"`
	AppSecret       string `json:"app_secret"`
	ChannelId       string `json:"channel_id"`
	Timeout         int    `json:"timeout"`
}

// NewConfig 创建一个新的配置实例
func NewConfig(tokenAPIBaseURL, apiBaseURL, appid, appSecret string, channelId string, timeout int) *Config {
	return &Config{
		TokenAPIBaseURL: tokenAPIBaseURL,
		APIBaseURL:      apiBaseURL,
		Appid:           appid,
		AppSecret:       appSecret,
		ChannelId:       channelId,
		Timeout:         timeout,
	}
}
func (c *Config) String() string {
	data, _ := json.MarshalIndent(c, "", "  ")
	return string(data)
}
