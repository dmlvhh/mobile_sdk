package main

import (
	"encoding/json"
	"net/http"
	"time"
)

//type Response struct {
//	Code int    `json:"code"`
//	Msg  string `json:"msg"`
//	Time string `json:"time"`
//	Data struct {
//		Token           string `json:"token"`
//		TokenExpiretime string `json:"token_expiretime"`
//	} `json:"data"`
//}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Time string      `json:"time"`
	Data interface{} `json:"data"`
}

//const (
//	ERROR   = 7
//	SUCCESS = 0
//)
//

func Result(code int, msg string, data interface{}, w http.ResponseWriter) {
	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")

	// 获取当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 构建响应结构体
	response := Response{
		Code: code,
		Msg:  msg,
		Time: currentTime,
		Data: data,
	}
	// 将响应编码为 JSON 并写入响应体
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "服务错误", 7)
		return
	}
}

type GetTokenRes struct {
	Token           string `json:"token"`
	TokenExpiretime string `json:"token_expiretime"`
}
