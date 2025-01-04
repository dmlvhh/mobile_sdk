package mobile_sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func (c *Config) ApiRequest(url string, data any) (res string, err error) {
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
	apiUrl := c.ApiUrl + url
	fmt.Println(apiUrl)
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
	fmt.Println(string(body))
	fmt.Println("111111111111111111111111111")
	return string(body), nil
}
