package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	config          *Config
	client          *resty.Client
	token           string
	tokenExpireTime string
}

// NewClient 创建一个新的 HTTP 客户端
func NewClient(config *Config) *Client {
	client := resty.New()
	client.SetTimeout(time.Duration(config.Timeout) * time.Second)
	client.SetRetryCount(3) // 设置重试次数
	client.SetRetryWaitTime(1 * time.Second)
	client.SetRetryMaxWaitTime(3 * time.Second)

	return &Client{
		config: config,
		client: client,
	}
}

// Do 发送 HTTP 请求并返回响应
func (c *Client) Do(method, path string, params url.Values, body interface{}, baseURL string) (*resty.Response, error) {
	// 获取当前时间戳
	timestamp := GetTimestamp()

	// 生成签名
	signature := GenerateSignature(c.config.Appid, c.config.AppSecret, timestamp)

	// 构建请求URL
	var url string
	if baseURL != "" {
		url = fmt.Sprintf("%s%s", baseURL, path)
	} else {
		url = fmt.Sprintf("%s%s", c.config.APIBaseURL, path)
	}
	fmt.Println("url", url)
	// 设置请求头部
	req := c.client.R().
		SetQueryParam("appid", c.config.Appid).
		SetQueryParam("app_secret", c.config.AppSecret).
		SetQueryParam("channel_id", c.config.ChannelId).
		SetQueryParam("timestamp", timestamp).
		SetQueryParam("sign", signature)
	fmt.Println("req", req)
	// 如果有查询参数，添加到请求中
	for key, values := range params {
		for _, value := range values {
			req.SetQueryParam(key, value)
		}
	}

	// 如果有请求体，设置请求体
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		req.SetBody(jsonData)
	}
	// 如果有请求体，设置请求体，并自动添加 token
	//if body != nil {
	//	// 将请求体转换为 map[string]interface{} 类型
	//	bodyMap, ok := body.(map[string]interface{})
	//	if !ok {
	//		// 如果不是 map[string]interface{}，直接设置请求体
	//		jsonData, err := json.Marshal(body)
	//		if err != nil {
	//			return nil, fmt.Errorf("failed to marshal request body: %w", err)
	//		}
	//		req.SetBody(jsonData)
	//	} else {
	//		// 如果是 map[string]interface{}，添加 token 到请求体中
	//		if c.token != "" {
	//			bodyMap["token"] = c.token
	//		}
	//		jsonData, err := json.Marshal(bodyMap)
	//		if err != nil {
	//			return nil, fmt.Errorf("failed to marshal request body: %w", err)
	//		}
	//		req.SetBody(jsonData)
	//	}
	//	fmt.Println("bodyMap", bodyMap)
	//}

	// 发送请求
	resp, err := req.Execute(method, url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	// 检查API返回的状态码
	var apiResp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}
	fmt.Println(apiResp.Code)
	if apiResp.Code != 1 {
		// 如果API返回错误，可能是 token 无效或过期，尝试重新获取 token 并重试
		if apiResp.Code == 12021 { // 假设 401 表示 token 无效或过期
			_, err := c.GetToken(&GetTokenReq{
				Appid:     c.config.Appid,
				AppSecret: c.config.AppSecret,
				ChannelId: c.config.ChannelId,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to refresh token: %w", err)
			}

			// 重新发送请求
			return c.Do(method, path, params, body, baseURL)
		}
		return nil, fmt.Errorf("API error: %s", apiResp.Msg)
	}
	return resp, nil
}
func (c *Client) DoWithoutToken(method, path string, body interface{}, baseURL string) (*resty.Response, error) {

	// 构建请求URL
	var url string
	if baseURL != "" {
		url = fmt.Sprintf("%s%s", baseURL, path)
	} else {
		url = fmt.Sprintf("%s%s", c.config.APIBaseURL, path)
	}
	fmt.Println("Request URL:", url)
	req := c.client.R().SetHeader("Content-Type", "application/json")
	// 如果有请求体，设置请求体
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		req.SetBody(jsonData)

		// 打印请求体，便于调试
		fmt.Printf("Request Body: %s\n", jsonData)
	}

	// 发送请求
	resp, err := req.Execute(method, url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// 打印响应体，便于调试
	fmt.Printf("Response Body: %s\n", resp.Body())

	// 检查HTTP状态码
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return resp, nil
}
