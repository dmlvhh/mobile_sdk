package main

type TokenRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Token           string `json:"token"`
		TokenExpiretime string `json:"token_expiretime"`
	} `json:"data"`
}
type SimBasicInfoRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Msisdn     string `json:"msisdn"`
		Imsi       string `json:"imsi"`
		Iccid      string `json:"iccid"`
		ActiveDate string `json:"activeDate"`
		OpenDate   string `json:"openDate"`
		Remark     string `json:"remark"`
	} `json:"result"`
}

type SimChangeHistoryRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		ChangeHistoryList []struct {
			DescStatus   string `json:"descStatus"`
			TargetStatus string `json:"targetStatus"`
			ChangeDate   string `json:"changeDate"`
		} `json:"changeHistoryList"`
	} `json:"result"`
}

type SimImeiRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Imei string `json:"imei"`
	} `json:"result"`
}

type SimStatusRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		CardStatus     string `json:"cardStatus"`
		LastChangeDate string `json:"lastChangeDate"`
	} `json:"result"`
}

type SimBasicInfoBatchRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Imsi    string `json:"imsi"`
		Msisdn  string `json:"msisdn"`
		Iccid   string `json:"iccid"`
	} `json:"result"`
}

type ChangeSimStatusBatchRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		JobId string `json:"jobId"`
	} `json:"result"`
}

type SimPlatformBatchRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		PlatformType string `json:"platformType"`
		Message      string `json:"message"`
		Msisdn       string `json:"msisdn"`
		Status       string `json:"status"`
	} `json:"result"`
}

type SimStopReasonRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		PlatformType string `json:"platformType"`
		StopReason   string `json:"stopReason"`
	} `json:"result"`
}

type SimDataMarginRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SimDataUsageRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		DataAmount       string `json:"dataAmount"`
		ApnUseAmountList []struct {
			ApnName              string        `json:"apnName"`
			ApnUseAmount         string        `json:"apnUseAmount"`
			PccCodeUseAmountList []interface{} `json:"pccCodeUseAmountList"`
		} `json:"apnUseAmountList"`
	} `json:"result"`
}
type ChangeSimStatusRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		OrderNum string `json:"orderNum"`
		Msisdn   string `json:"msisdn"`
	} `json:"result"`
}

//type Response struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg"`
//	Time string      `json:"time"`
//	Data interface{} `json:"data"`
//}

//const (
//	ERROR   = 7
//	SUCCESS = 0
//)
//

//func Result(code int, msg string, data interface{}, w http.ResponseWriter) {
//	// 设置响应头为 JSON
//	w.Header().Set("Content-Type", "application/json")
//
//	// 获取当前时间
//	currentTime := time.Now().Format("2006-01-02 15:04:05")
//
//	// 构建响应结构体
//	response := Response{
//		Code: code,
//		Msg:  msg,
//		Time: currentTime,
//		Data: data,
//	}
//	// 将响应编码为 JSON 并写入响应体
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, "服务错误", 7)
//		return
//	}
//}

//type GetTokenRes struct {
//	Token           string `json:"token"`
//	TokenExpiretime string `json:"token_expiretime"`
//}
