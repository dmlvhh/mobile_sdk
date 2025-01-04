package mobile_sdk

import (
	"fmt"
	"net/url"
	"testing"
)

var Tp *TPConfig
var cf *Config

func TestTPConfig_NewTP1Config(t *testing.T) {
	Tp = NewTP1Config("http://token.dctxiot.com/api/token/getToken", "29030", "a0674936bd251655ff8e14e18c74b879", "3")
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
		Msisdn:  "1442161864994", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	})
	if err2 != nil {
		fmt.Printf(err2.Error())
		return
	}
	fmt.Println(res)
}

func TestTPConfig_NewTP2Config(t *testing.T) {
	Tp = NewTP2Config("https://ac.buleideiot.com/api/get-chinamobile-onelink-token", "sj4xACNePCXgZARj", "22")
	params := url.Values{}
	params.Add("channel_id", Tp.ChannelID)
	params.Add("sign", Tp.Sign)
	request, err := Tp.ApiGetRequest(params)
	if err != nil {
		return
	}
	fmt.Println(request.Data.Token)
	token := request.Data.Token
	cf = NewConfig("https://api.iot.10086.cn", "PMMxeJGqamJw7Jrh13361", "mUzJa4znlE7f9ZtdIReVFMFubUzx2vs6")
	res, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &SimBasicInfoReq{
		Transid: cf.TransID,
		Token:   token,
		Msisdn:  "1442077770000", // 可选字段
		//Iccid:   "89860846162470274998", // 可选字段
		//Imsi:    "460240261864998",      // 可选字段
	})
	if err2 != nil {
		fmt.Printf(err2.Error())
		return
	}
	fmt.Println(res)
}
