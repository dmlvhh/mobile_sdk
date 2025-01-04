package mobile_sdk

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

type SimDataUsageMonthlyBatchRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		DeadLine       string `json:"deadLine"`
		DataAmountList []struct {
			DataAmount        string `json:"dataAmount"`
			ApnDataAmountList []struct {
				ApnName       string `json:"apnName"`
				ApnDataAmount string `json:"apnDataAmount"`
			} `json:"apnDataAmountList"`
			Msisdn string `json:"msisdn"`
		} `json:"dataAmountList"`
	} `json:"result"`
}
type CardBindStatusRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Result    string `json:"result"`
		ErrorCode string `json:"errorCode"`
		ErrorDes  string `json:"errorDes"`
	} `json:"result"`
}

type SimRealNameStatusRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Reason         string `json:"reason"`
		RealNameStatus string `json:"realNameStatus"`
		SuccessTime    string `json:"successTime"`
	} `json:"result"`
}

type TPTokenRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}
