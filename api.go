package mobile_sdk

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

// GetToken 获取token
func (c *Client) GetToken(r *GetTokenReq) (*TokenRes, error) {
	path := "/token/getToken"
	// 构建查询参数
	params := url.Values{}
	params.Add("appid", r.Appid)
	params.Add("app_secret", r.AppSecret)
	params.Add("channel_id", r.ChannelId)
	params.Add("times", r.Times)
	params.Add("sign", r.Sign)

	// 发送GET请求
	resp, err := c.Do("GET", path, params, nil, c.config.TokenAPIBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	// 解析响应
	var apiResp TokenRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Code != 1 {
		return nil, fmt.Errorf("API error: %s", apiResp.Msg)
	}
	// 更新客户端的 token 字段
	c.token = apiResp.Data.Token
	c.tokenExpireTime = apiResp.Data.TokenExpiretime
	return &apiResp, nil
}

// SimBasicInfo 单卡基本信息查询
func (c *Client) SimBasicInfo(r *SimBasicInfoReq) (*SimBasicInfoRes, error) {
	path := "/query/sim-basic-info"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimBasicInfoRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimChangeHistory 单卡状态变更历史查询
func (c *Client) SimChangeHistory(r *SimBasicInfoReq) (*SimChangeHistoryRes, error) {
	path := "/query/sim-change-history"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimChangeHistoryRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimImei 单卡绑定IMEI实时查询
func (c *Client) SimImei(r *SimBasicInfoReq) (*SimImeiRes, error) {
	path := "/query/sim-imei"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimImeiRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimStatus 单卡状态查询
func (c *Client) SimStatus(r *SimBasicInfoReq) (*SimStatusRes, error) {
	path := "/query/sim-status"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimStatusRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimStopReason 单卡停机原因查询
func (c *Client) SimStopReason(r *SimBasicInfoReq) (*SimStopReasonRes, error) {
	path := "/query/sim-stop-reason"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimStopReasonRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimDataMargin 单卡本月套餐内流量使用量实时查询 todo 异常
func (c *Client) SimDataMargin(r *SimBasicInfoReq) (*SimDataMarginRes, error) {
	path := "/query/sim-data-margin"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimDataMarginRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimDataUsage 单卡本月流量累计使用量查询
func (c *Client) SimDataUsage(r *SimBasicInfoReq) (*SimDataUsageRes, error) {
	path := "/query/sim-data-usage"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimDataUsageRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimDataUsageMonthlyBatch 物联卡单月GPRS流量使用量批量查询
func (c *Client) SimDataUsageMonthlyBatch(r *SimBasicInfoBatchReq) (*SimDataUsageMonthlyBatchRes, error) {
	path := "/query/sim-data-usage-monthly/batch"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimDataUsageMonthlyBatchRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// ChangeSimStatus 单卡状态变更
func (c *Client) ChangeSimStatus(r *SimBasicInfoReq) (*ChangeSimStatusRes, error) {
	path := "/change/sim-status"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp ChangeSimStatusRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}
	if apiResp.Status != "13012" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// ChangeSimStatusBatch 物联卡状态变更批量办理
func (c *Client) ChangeSimStatusBatch(r *SimBasicInfoBatchReq) (*ChangeSimStatusBatchRes, error) {
	path := "change/sim-status/batch"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp ChangeSimStatusBatchRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimCardInfoBatch 码号信息批量查询
func (c *Client) SimCardInfoBatch(r *SimBasicInfoBatchReq) (*SimBasicInfoBatchRes, error) {
	path := "/query/sim-card-info/batch"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimBasicInfoBatchRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}

// SimPlatformBatch 物联卡归属平台批量查询
func (c *Client) SimPlatformBatch(r *SimBasicInfoBatchReq) (*SimPlatformBatchRes, error) {
	path := "/query/sim-platform/batch"

	reqBody, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("Request Body: %s\n", reqBody)
	resp, err := c.DoWithoutToken("POST", path, r, "")
	if err != nil {
		return nil, fmt.Errorf("failed to call SimBasicInfo API: %w", err)
	}
	// 解析API响应
	var apiResp SimPlatformBatchRes
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	// 检查API返回的状态码
	if apiResp.Status != "0" {
		return nil, fmt.Errorf("API error: %s", apiResp.Message)
	}

	return &apiResp, nil

}
