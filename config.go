package main

type Config struct {
	APIBaseURL string `json:"api_base_url"`
	Appid      string `json:"appid"`
	AppSecret  string `json:"app_secret"`
	ChannelId  string `json:"channel_id"`
	Timeout    int    `json:"timeout"`
}

// NewConfig 创建一个新的配置实例
func NewConfig(apiBaseURL, appid, appSecret string, channelId string, timeout int) *Config {
	return &Config{
		APIBaseURL: apiBaseURL,
		Appid:      appid,
		AppSecret:  appSecret,
		ChannelId:  channelId,
		Timeout:    timeout,
	}
}
