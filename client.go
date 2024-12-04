package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
	config     *Config
}

// NewClient 创建一个新的 HTTP 客户端
func NewClient(cfg *Config) *Client {
	if cfg == nil {
		cfg = &Config{Timeout: 30} // 默认超时时间为 30 秒
	}
	return &Client{
		httpClient: &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second},
		config:     cfg,
	}
}

// Do 发送 HTTP 请求并返回响应
func (c *Client) Do(req *http.Request, resp interface{}) error {
	// 添加日志记录
	log.Printf("Sending %s request to %s", req.Method, req.URL)
	response, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return err
	}
	log.Printf("Response status: %d, Body: %s", response.StatusCode, string(body))
	if response.StatusCode != http.StatusOK {
		return NewAPIError(response.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, resp); err != nil {
		log.Printf("Failed to unmarshal JSON response: %v", err)
		return err
	}

	return nil
}

// DoRequest 发送 HTTP 请求并返回通用的 API 响应
func (c *Client) DoRequest(method, path string, reqBody interface{}, resp interface{}) error {
	// 构建请求
	req, err := c.BuildRequest(method, path, reqBody, nil)
	if err != nil {
		log.Printf("Failed to build request: %v", err)
		return err
	}

	// 发送请求并获取响应
	c.Do(req, resp)
	return nil
}

// BuildRequest 构建 HTTP 请求
func (c *Client) BuildRequest(method, path string, reqBody interface{}, params url.Values) (*http.Request, error) {
	// 构建基础 URL
	baseURL := c.config.APIBaseURL + path

	// 如果有查询参数，将其附加到 URL 上
	if params != nil && len(params) > 0 {
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		u.RawQuery = params.Encode()
		baseURL = u.String()
	}

	// 创建请求
	var bodyReader io.Reader
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		// 对于 POST 请求，将请求体序列化为 JSON
		if reqBody != nil {

			body, err := json.Marshal(reqBody)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewReader(body)
		}
	}

	req, err := http.NewRequest(method, baseURL, bodyReader)
	if err != nil {
		return nil, err
	}

	// 请求时设置 Content-Type
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		if reqBody != nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}
	return req, nil
}
