package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateSignature(appid, appSecret, timestamp string) string {
	signatureString := fmt.Sprintf("appid=%s&app_secret=%s&times=%s", appid, appSecret, timestamp)
	hash := md5.New()
	hash.Write([]byte(signatureString))
	hashString := hex.EncodeToString(hash.Sum(nil))
	signature := strings.ToUpper(hashString)
	return signature
}

//func GenerateSignature(appid, appSecret, channelId string, timestamp string) string {
//	signatureString := fmt.Sprintf("appid=%s&app_secret=%s&channelId=%s&times=%s", appid, appSecret, channelId, timestamp)
//	hash := md5.New()
//	hash.Write([]byte(signatureString))
//	hashString := hex.EncodeToString(hash.Sum(nil))
//	signature := strings.ToUpper(hashString)
//	return signature
//}

//// GetAuthHeader 获取认证头
//func GetAuthHeader(cfg *Config, timestamp string) string {
//	message := fmt.Sprintf("app_key=%s&timestamp=%s", cfg.AppKey, timestamp)
//	signature := GenerateSignature(cfg.AppSecret, message)
//	return fmt.Sprintf("AppKey %s, Signature %s, Timestamp %s", cfg.AppKey, signature, timestamp)
//}
//
//func GetToken(params map[string]string) (res Response) {
//	baseURL, err := url.Parse(global.GVA_CONFIG.Mobile.ApiUrl)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Add query parameters
//	queryParams := baseURL.Query()
//	for key, value := range params {
//		queryParams.Add(key, value)
//	}
//
//	// Attach the query parameters to the base URL
//	baseURL.RawQuery = queryParams.Encode()
//
//	// Send a GET request
//	resp, err := http.Get(baseURL.String())
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	// Read the response body
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	_ = json.Unmarshal(body, &res)
//	return
//}
