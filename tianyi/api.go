package tianyi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type TianyiClient struct {
	BaseURL string
	UserID  string
	APIKey  string
}

func NewTianyiClient(baseURL string, userID string, apiKey string) *TianyiClient {
	return &TianyiClient{
		BaseURL: baseURL,
		UserID:  userID,
		APIKey:  apiKey,
	}
}

// 发起校验请求
func (c *TianyiClient) PostVerification() (string, error) {
	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("api_key", c.APIKey)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/verification", form)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

// 获取卡板信息
func (c *TianyiClient) GetCardInfo(iccid string) (res GetCardInfoRes, err error) {
	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("api_key", c.APIKey)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("iccid", iccid)

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/cardInfo", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

// 获取卡板信息
func (c *TianyiClient) GetFreezeInfo(iccid string) (res GetCardInfoRes, err error) {
	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("api_key", c.APIKey)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("iccid", iccid)

	resp, err := http.PostForm(c.BaseURL+"//index/ka_ban/getRestartStatus", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

// 修改SIM卡状态：1=停机，2=复机
func (c *TianyiClient) SetSimStatus(iccid string, status int) (res SetSimStatusRes, err error) {
	if iccid == "" || (status != 1 && status != 2) {
		return res, errors.New("参数错误：iccid不能为空，status只能是1或2")
	}

	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("iccid", iccid)
	form.Set("status", strconv.Itoa(status))

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/SimStatus", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

// 机卡分离
func (c *TianyiClient) UnbindSimCard(iccid string) (res UnbindSimCardRes, err error) {
	if iccid == "" {
		return res, errors.New("iccid 不能为空")
	}

	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("iccid", iccid)

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/ReRecord", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

func (c *TianyiClient) PlaceOrder(iccid string, packageID string, monthType int, orderNo string) (res PlaceOrderRes, err error) {
	if iccid == "" || packageID == "" {
		log.Println("iccid 和 package_id 是必须的")
		return res, errors.New("iccid 和 package_id 是必须的")
	}
	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("iccid", iccid)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("package_id", packageID)

	if monthType != 0 {
		form.Set("month_type", strconv.Itoa(monthType)) // 1本月 2次月
	}
	if orderNo != "" {
		form.Set("order_no", orderNo)
	}

	endpoint := c.BaseURL + "/apiv1/tianyi/orderCreate"
	resp, err := http.PostForm(endpoint, form)
	if err != nil {
		log.Println("下单请求失败:", err)
		return res, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败:", err)
		return res, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("解析响应失败:", err)
		return res, err
	}
	return res, nil
}

// 查询官方实名状态
func (c *TianyiClient) QueryOfficialRealStatus(iccid string) (res QueryOfficialRealStatusRes, err error) {
	if iccid == "" {
		return res, errors.New("iccid 不能为空")
	}

	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("iccid", iccid)

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/officialReal", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

// 刷新卡信息
func (c *TianyiClient) RefreshSim(iccid string) (res RefreshSimRes, err error) {
	if iccid == "" {
		return res, errors.New("iccid 不能为空")
	}

	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)
	form.Set("iccid", iccid)

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/refreshSim", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

// 获取移动通道的 token
func (c *TianyiClient) GetYDToken(appid string, chlID int) (string, error) {
	if appid == "" || chlID == 0 {
		return "", errors.New("appid 和 chl_id 是必须的")
	}

	// 使用 GET 请求
	urlStr := fmt.Sprintf("%s/index/index/ydToken?appid=%s&chl_id=%d", c.BaseURL, appid, chlID)
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

// 批量获取卡板信息（iccids 或 cards 传一个）
func (c *TianyiClient) GetBatchCardInfo(iccids []string, cards []string) (res GetBatchCardInfoRes, err error) {
	if len(iccids) == 0 && len(cards) == 0 {
		return res, errors.New("必须提供 iccids 或 cards 中的一个")
	}

	requestTime := time.Now().Unix()
	token := c.generateToken(requestTime)

	form := url.Values{}
	form.Set("user_id", c.UserID)
	form.Set("api_key", c.APIKey)
	form.Set("request_time", strconv.FormatInt(requestTime, 10))
	form.Set("request_token", token)

	if len(iccids) > 0 {
		form.Set("iccids", strings.Join(iccids, ","))
	}
	if len(cards) > 0 {
		form.Set("cards", strings.Join(cards, ","))
	}

	resp, err := http.PostForm(c.BaseURL+"/apiv1/tianyi/cardInfos", form)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res, nil
}

func (c *TianyiClient) generateToken(ts int64) string {
	base := strconv.FormatInt(ts, 10) + md5Sum(c.APIKey)
	return md5Sum(base)
}

func md5Sum(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
