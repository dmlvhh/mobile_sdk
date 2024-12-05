package mobile_sdk

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	// 打印Token和过期时间
	fmt.Printf("Token: %s\n", tokenResp.Data.Token)
	fmt.Printf("Token Expire Time: %s\n", tokenResp.Data.TokenExpiretime)
}

func TestSimBasicInfo(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdn:  "1442161864998", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimBasicInfo(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	for _, result := range simResp.Result {
		fmt.Printf("MSISDN: %s\n", result.Msisdn)
		fmt.Printf("IMSI: %s\n", result.Imsi)
		fmt.Printf("ICCID: %s\n", result.Iccid)
		fmt.Printf("Active Date: %s\n", result.ActiveDate)
		fmt.Printf("Open Date: %s\n", result.OpenDate)
		fmt.Printf("Remark: %s\n", result.Remark)
	}
}
func TestSimChangeHistory(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdn:  "1442161864998", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimChangeHistory(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}

func TestSimImei(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdn:  "1442161864998", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimImei(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}

func TestSimStatus(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdn:  "1442161864998", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimStatus(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}

func TestSimStopReason(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdn:  "1442161864998", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimStopReason(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}
func TestSimDataUsage(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdn:  "1442161864998", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimDataUsage(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}

func TestChangeSimStatus(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoReq{
		Transid:  generator.Generate(),
		OperType: "0",
		Token:    tokenResp.Data.Token,
		Msisdn:   "1442161864993", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.ChangeSimStatus(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}
func TestSimCardInfoBatch(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoBatchReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdns: "1442161864998_1442161864995_1442161864994", // 可选字段
		//Iccids:   "89860846162470274998", // 可选字段
		//Imsis:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimCardInfoBatch(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}

func TestSimPlatformBatch(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoBatchReq{
		Transid: generator.Generate(),
		Token:   tokenResp.Data.Token,
		Msisdns: "1442161864998_1442161864995_1442161864994", // 可选字段
		//Iccids:   "89860846162470274998", // 可选字段
		//Imsis:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.SimPlatformBatch(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}

func TestChangeSimStatusBatch(t *testing.T) {
	appid := "29030"
	appSecret := "a0674936bd251655ff8e14e18c74b879"
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sign := GenerateSignature(appid, appSecret, timestamp)
	//fmt.Println(sign)
	cfg := NewConfig(
		"http://token.dctxiot.com/api",
		"https://api.iot.10086.cn/v5/ec",
		appid,
		appSecret,
		"3",
		10,
	)
	//创建客户端
	client := NewClient(cfg)
	// 准备请求参数
	tokenReq := &GetTokenReq{
		Appid:     appid,
		AppSecret: appSecret,
		Times:     timestamp,
		ChannelId: "3",
		Sign:      sign,
	}
	// 获取Token
	tokenResp, err := client.GetToken(tokenReq)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	generator := NewTransIDGenerator(appid)
	simReq := &SimBasicInfoBatchReq{
		Transid:  generator.Generate(),
		Token:    tokenResp.Data.Token,
		OperType: "9",
		// "operType": "11",
		Reason:  "01",
		Msisdns: "1442161864998_1442161864995_1442161864994", // 可选字段
		//Iccids:   "89860846162470274998", // 可选字段
		//Imsis:    "460240261864998",      // 可选字段
	}

	// 调用 SimBasicInfo 方法查询SIM卡基本信息
	simResp, err := client.ChangeSimStatusBatch(simReq)
	if err != nil {
		log.Fatalf("Failed to query SIM card info: %v", err)
	}

	// 打印查询结果
	fmt.Println("simResp", simResp)
}
