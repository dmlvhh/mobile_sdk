package main

import (
	"fmt"
	"github.com/dmlvhh/mobile_sdk"
	"net/url"
)

func main() {
	Tp := mobile_sdk.NewTP1Config("http://token.dctxiot.com/api/token/getToken", "29030", "a0674936bd251655ff8e14e18c74b879", "3")
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
	cf := mobile_sdk.NewConfig("https://api.iot.10086.cn", "29030", "a0674936bd251655ff8e14e18c74b879")
	//res, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &mobile_sdk.SimBasicInfoReq{
	//	Transid: cf.TransID,
	//	Token:   token,
	//	//Msisdn:  "1442161864994", // 可选字段
	//	Iccid: "89860846162470274994", // 可选字段
	//	//Imsi:    "460240261864998",      // 可选字段
	//})
	//if err2 != nil {
	//	fmt.Printf(err2.Error())
	//	return
	//}
	//res, err := cf.SimRealNameStatus(&mobile_sdk.SimBasicInfoReq{
	//	Transid: cf.TransID,
	//	Token:   token,
	//	Iccid:   "89860846162470274994",
	//})
	res, err := cf.CardBindStatus(&mobile_sdk.SimBasicInfoReq{
		Transid:  cf.TransID,
		Token:    token,
		TestType: "0",
		//Iccid:    "89860846162470274994",
		Msisdn: "1442161864994", // 可选字段
	})
	fmt.Println(res.Result)
}

//func main() {
//	Tp := mobile_sdk.NewTP2Config("https://ac.buleideiot.com/api/get-chinamobile-onelink-token", "sj4xACNePCXgZARj", "22")
//	params := url.Values{}
//	params.Add("channel_id", Tp.ChannelID)
//	params.Add("sign", Tp.Sign)
//	request, err := Tp.ApiGetRequest(params)
//	if err != nil {
//		return
//	}
//	fmt.Println(request.Data.Token)
//	token := request.Data.Token
//	cf := mobile_sdk.NewConfig("https://api.iot.10086.cn", "PMMxeJGqamJw7Jrh13361", "mUzJa4znlE7f9ZtdIReVFMFubUzx2vs6")
//	res, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &mobile_sdk.SimBasicInfoReq{
//		Transid: cf.TransID,
//		Token:   token,
//		//Msisdn:  "1442077770000", // 可选字段
//		Iccid: "89860837132490270001", // 可选字段
//		//Imsi:    "460240261864998",      // 可选字段
//	})
//	if err2 != nil {
//		fmt.Printf(err2.Error())
//		return
//	}
//	fmt.Println(res)
//}

//func main() {
//	Tp := mobile_sdk.NewTP2Config("https://ac.buleideiot.com/api/get-chinamobile-onelink-token", "CVDBTTnfZLrMgeyn", "13")
//	params := url.Values{}
//	params.Add("channel_id", Tp.ChannelID)
//	params.Add("sign", Tp.Sign)
//	request, err := Tp.ApiGetRequest(params)
//	if err != nil {
//		return
//	}
//	fmt.Println(request.Data.Token)
//	token := request.Data.Token
//	cf := mobile_sdk.NewConfig("https://api.iot.10086.cn", "PMMxeJGqamJw7Jrh13361", "mUzJa4znlE7f9ZtdIReVFMFubUzx2vs6")
//	res, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &mobile_sdk.SimBasicInfoReq{
//		Transid: cf.TransID,
//		Token:   token,
//		//Msisdn:  "898604451424D0102274", // 可选字段
//		Iccid: "898604451424D0102274", // 可选字段
//		//Imsi: "1440452640897", // 可选字段
//	})
//	if err2 != nil {
//		fmt.Printf(err2.Error())
//		return
//	}
//	fmt.Println(res)
//}
