package mobile_sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

var Tp *TPConfig
var cf *Config

func TestTPConfig_NewTP1Config(t *testing.T) {
	Tp = NewTP1Config("http://token.dctxiot.com/api/token/getToken", "29030", "a0674936bd251655ff8e14e18c74b879", "8")
	params := url.Values{}
	params.Add("appid", Tp.AppID)
	params.Add("app_secret", Tp.AppSecret)
	params.Add("channel_id", Tp.ChannelID)
	params.Add("times", Tp.Times)
	params.Add("sign", Tp.Sign)
	request, err := Tp.ApiGetRequest(params)
	if err != nil {
		return
	}
	//fmt.Println(request.Data.Token)
	token := request.Data.Token
	cf = NewConfig("https://api.iot.10086.cn", "29030", "a0674936bd251655ff8e14e18c74b879")
	res, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &SimBasicInfoReq{
		Transid: cf.TransID,
		Token:   token,
		//Msisdn:  "1442161864994", // 可选字段
		Iccid: "89860839092420107303", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	})
	if err2 != nil {
		fmt.Printf(err2.Error())
		return
	}
	fmt.Println(res)
}

func TestTPConfig_NewTP2Config(t *testing.T) {
	//Tp = NewTP2Config("https://ac.buleideiot.com/api/get-chinamobile-onelink-token", "sj4xACNePCXgZARj", "22")
	Tp = NewTP2Config("https://ac.buleideiot.com/api/get-chinamobile-onelink-token", "NNIYDQcitiDJngWA", "17")
	params := url.Values{}
	params.Add("channel_id", Tp.ChannelID)
	params.Add("sign", Tp.Sign)
	request, err := Tp.ApiGetRequest(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request.Data.Token)
	token := request.Data.Token
	cf = NewConfig("https://api.iot.10086.cn", "PMMxeJGqamJw7Jrh13361", "mUzJa4znlE7f9ZtdIReVFMFubUzx2vs6")
	//res, err2 := cf.ApiRequest("/v5/ec/query/sim-manage-stop-restart-status", &SimBasicInfoReq{
	//res, err2 := cf.ApiRequest("/v5/ec/operate/sim-gprs-status-reset", &SimBasicInfoReq{ //0:关机 1:开机
	//res, err2 := cf.ApiRequest("/v5/ec/query/on-off-status", &SimBasicInfoReq{ //0:关机 1:开机
	//res, err2 := cf.ApiRequest("/v5/ec/query/sim-session", &SimBasicInfoReq{ //0:关机 1:开机
	//res, err2 := cf.ApiRequest("/v5/ec/secure/sim-real-name-reg", &SimBasicInfoReq{
	res, err2 := cf.ApiRequest("/v5/ec/secure/sim-real-name-reg", &SimBasicInfoReq{
		//res, err2 := cf.ApiRequest("/v5/ec/query/card-bind-status", &SimBasicInfoReq{
		//res, err2 := cf.ApiRequest("/v5/ec/query/sim-imei", &SimBasicInfoReq{
		Transid: cf.TransID,
		Token:   token,
		//TestType: "0",
		//Msisdn:  "1442360450001", // 可选字段
		UrlType: "6",
		Iccid:   "89860866182590580849", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	})
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(res)
}

func TestNewTP3Config(t *testing.T) {
	Tp = NewTP3Config("http://www.llk.10086link.cn/web/api/iot/yijia/get/cmcc/token", "CMCC888888888888103", "1868309673338204160")
	request, err := Tp.ApiPostRequest(map[string]string{
		"appId":     Tp.AppID,
		"accountId": Tp.AccountID,
	})
	if err != nil {
		return
	}
	var res TP3TokenRes
	json.Unmarshal([]byte(request), &res)
	fmt.Println(res)
	cf = NewConfig("https://api.iot.10086.cn", "", "")
	res1, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &SimBasicInfoReq{
		Transid: cf.TransID,
		Token:   res.Data,
		//Msisdn:  "1442077770000", // 可选字段
		//Iccid: "89860837132490489025", // 可选字段
		Iccid: "89860837132490567358", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	})
	if err2 != nil {
		fmt.Printf(err2.Error())
		return
	}
	fmt.Println(res1)
}

func TestNewTP4Config(t *testing.T) {
	TP := NewTP4Config("http://acc-api-token.nnkj77.com/iot/xt1000/token.php")
	params := url.Values{}
	request, err := TP.ApiGetRequest(params)
	if err != nil {
		return
	}
	fmt.Println(request.Token)
	token := request.Token
	cf = NewConfig("https://api.iot.10086.cn", "", "")
	res2, err2 := cf.ApiRequest("/v5/ec/query/sim-status", &SimBasicInfoReq{
		Transid:  cf.TransID,
		Token:    token,
		OperType: "0",
		//Msisdn:   "1442078069363", // 可选字段
		Iccid: "89860814172480061846", // 可选字段
		//Iccid: "89860837132490570000", // 可选字段
		//Imsi: "1442078067002", // 可选字段
	})
	if err2 != nil {
		fmt.Printf(err2.Error())
		return
	}
	//res1, err2 := cf.ApiRequest("/v5/ec/query/sim-real-name-status", &SimBasicInfoReq{
	//	Transid: cf.TransID,
	//	Token:   token,
	//	//Msisdn:  "1442078069363", // 可选字段
	//	Iccid: "89860814172480062979", // 可选字段
	//	//Iccid: "89860837132490570000", // 可选字段
	//	//Imsi: "1442078067002", // 可选字段
	//})
	//res1, err2 := cf.ApiRequest("/v5/ec/query/card-bind-status", &SimBasicInfoReq{
	//	Transid:  cf.TransID,
	//	Token:    token,
	//	TestType: "0",
	//	Msisdn:   "1442000252275", // 可选字段
	//	//Iccid: "89860837132490563908", // 可选字段
	//	//Iccid: "89860837132490570000", // 可选字段
	//	//Imsi: "1442078067002", // 可选字段
	//})
	if err2 != nil {
		fmt.Printf(err2.Error())
		return
	}
	fmt.Println(res2)
}
