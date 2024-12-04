package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateSignature(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	fmt.Println("sign", sign)
}
