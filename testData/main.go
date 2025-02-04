package main

import (
	"fmt"
	"github.com/dmlvhh/mobile_sdk"
	"net/url"
)

func main() {
	//Tp := mobile_sdk.NewTP1Config("http://token.dctxiot.com/api/token/getToken", "29030", "a0674936bd251655ff8e14e18c74b879", "3")
	Tp := mobile_sdk.NewTP1Config("http://token.dctxiot.com/api/token/getToken", "29030", "a0674936bd251655ff8e14e18c74b879", "8")
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
	fmt.Println(request.Data.Token)
	//token := request.Data.Token
	//cf := mobile_sdk.NewConfig("https://api.iot.10086.cn", "29030", "a0674936bd251655ff8e14e18c74b879")
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
	//res, err := cf.CardBindStatus(&mobile_sdk.SimBasicInfoReq{
	//	Transid:  cf.TransID,
	//	Token:    token,
	//	TestType: "0",
	//	//Iccid:    "89860846162470274994",
	//	Msisdn: "1442161864994", // 可选字段
	//})
	//fmt.Println(res.Result)
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

//func main() {
//	Tp := mobile_sdk.NewTP3Config("http://www.llk.10086link.cn/web/api/iot/yijia/get/cmcc/token", "CMCC888888888888103", "1868309673338204160")
//
//	request, err := Tp.ApiPostRequest(map[string]string{
//		"appId":     Tp.AppID,
//		"accountId": Tp.AccountID,
//	})
//	if err != nil {
//		return
//	}
//	var res mobile_sdk.TP3TokenRes
//	json.Unmarshal([]byte(request), &res)
//	fmt.Println(res.Data)
//	cf := mobile_sdk.NewConfig("https://api.iot.10086.cn", "", "")
//	//res1, err2 := cf.ApiRequest("/v5/ec/query/sim-basic-info", &mobile_sdk.SimBasicInfoReq{
//	//	Transid: cf.TransID,
//	//	Token:   res.Data,
//	//	//Msisdn:  "1442078069363", // 可选字段
//	//	//Iccid: "89860837132490489025", // 可选字段
//	//	Iccid: "89860837132490570000", // 可选字段
//	//	//Imsi:    "460240261864998",      // 可选字段
//	//})
//	res1, err2 := cf.ApiRequest("/v5/ec/query/sim-real-name-status", &mobile_sdk.SimBasicInfoReq{
//		Transid: cf.TransID,
//		Token:   res.Data,
//		//Msisdn:  "1442078069363", // 可选字段
//		Iccid: "89860837132490563909", // 可选字段
//		//Iccid: "89860837132490570000", // 可选字段
//		//Imsi: "1442078067002", // 可选字段
//	})
//	if err2 != nil {
//		fmt.Printf(err2.Error())
//		return
//	}
//	fmt.Println(res1)
//	//res1, err2 := cf.ApiRequest("/v5/ec/query/card-bind-status", &mobile_sdk.SimBasicInfoReq{
//	//	Transid:  cf.TransID,
//	//	Token:    res.Data,
//	//	TestType: "0",
//	//	Msisdn:   "1442078069363", // 可选字段
//	//	//Iccid: "89860837132490563908", // 可选字段
//	//	//Iccid: "89860837132490570000", // 可选字段
//	//	//Imsi: "1442078067002", // 可选字段
//	//})
//	//if err2 != nil {
//	//	fmt.Printf(err2.Error())
//	//	return
//	//}
//	//fmt.Println(res1)
//
//	//res1, err2 := cf.ApiRequest("/v5/ec/query/sim-status", &mobile_sdk.SimBasicInfoReq{
//	//	Transid:  cf.TransID,
//	//	Token:    res.Data,
//	//	OperType: "0",
//	//	//Msisdn:   "1442078069363", // 可选字段
//	//	Iccid: "89860837132490567358", // 可选字段
//	//	//Iccid: "89860837132490570000", // 可选字段
//	//	//Imsi: "1442078067002", // 可选字段
//	//})
//	//if err2 != nil {
//	//	fmt.Printf(err2.Error())
//	//	return
//	//}
//	//res1, err2 := cf.ApiRequest("/v5/ec/query/sim-data-usage", &mobile_sdk.SimBasicInfoReq{
//	//	Transid:  cf.TransID,
//	//	Token:    res.Data,
//	//	OperType: "0",
//	//	//Msisdn:   "1442078069363", // 可选字段
//	//	Iccid: "89860837132490495120", // 可选字段
//	//	//Iccid: "89860837132490570000", // 可选字段
//	//	//Imsi: "1442078067002", // 可选字段
//	//})
//	//if err2 != nil {
//	//	fmt.Printf(err2.Error())
//	//	return
//	//}
//	//fmt.Println(res1)
//
//}
