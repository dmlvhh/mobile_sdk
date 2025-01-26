package mobile_sdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func (c *TPConfig) ApiGetRequest(params url.Values) (res TPTokenRes, err error) {
	apiUrl := fmt.Sprintf("%s?%s", c.ApiUrl, params.Encode())
	//fmt.Println("Request URL:", apiUrl)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(apiUrl)
	if err != nil {
		log.Printf("Error making GET request: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	json.Unmarshal(body, &res)
	return
}

func (c *TPConfig) ApiPostRequest2(data any) (res string, err error) {
	reqData, err := json.Marshal(&data)
	fmt.Println(string(reqData))
	//fmt.Println("c", c)
	if err != nil {
		log.Printf("Error Marshal: %s", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	fmt.Println(c.ApiUrl)
	client := &http.Client{Transport: tr}
	resp, err := client.Post(c.ApiUrl, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		log.Printf("Error making POST request: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	//fmt.Println(string(body))
	return string(body), nil
}
func (c *TPConfig) ApiPostRequest(data any) (res string, err error) {
	// 将数据序列化为 JSON
	reqData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling data: %s", err)
		return "", err
	}
	fmt.Println(string(reqData))

	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 创建 HTTP 客户端
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	client := &http.Client{Transport: tr}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ApiUrl, bytes.NewBuffer(reqData))
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making POST request: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return "", err
	}

	// 返回响应
	return string(body), nil
}
func (c *Config) ApiRequest(url string, data any) (res string, err error) {
	reqData, err := json.Marshal(&data)
	//fmt.Println(string(reqData))
	//fmt.Println("c", c)
	if err != nil {
		log.Printf("Error Marshal: %s", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不安全，仅用于测试
	}
	apiUrl := c.ApiUrl + url
	//fmt.Println(apiUrl)
	client := &http.Client{Transport: tr}
	resp, err := client.Post(apiUrl, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		log.Printf("Error making POST request: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	//fmt.Println(string(body))
	return string(body), nil
}
