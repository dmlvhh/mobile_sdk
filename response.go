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

type SimDataDiagnosisRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		APNList struct {
			APNInfo []struct {
				APNName              string `json:"APNName"`
				ModifyTime           string `json:"modifyTime"`
				APNStatus            string `json:"APNStatus"`
				APNServiceUsageState string `json:"APNServiceUsageState"`
			} `json:"APNInfo"`
		} `json:"APNList"`
	} `json:"result"`
}
type NetworkSpeedRes struct {
}

type SimManageStopRestartStatusRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		ManageStopRestartStatus     string `json:"  manageStopRestartStatus"`   //0否1是
		Reason                      string `json:"reason"`                      //冻结原因
		ManageStOpRestartStatusTime string `json:"manageStopRestartStatusTime"` //  管理停机冻结时间或解冻操作时间
	} `json:"result"`
}
type QueryOnOffStatusRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Status string `json:"  status"`
	} `json:"result"`
}
type SimSessionRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		SImSessionList []struct {
			ApnId      string `json:"apnId"`
			Status     string `json:"status"`
			Ip         string `json:"ip"`
			CreateDate string `json:"createDate"`
			Rat        string `json:"rat"`
			Ipv6Prefix string `json:"ipv6Prefix"`
			Ipv6       string `json:"ipv6"`
		} `json:"s  imSessionList"`
	} `json:"result"`
}

type SimGprsStatusResetRes struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []interface{} `json:"result"`
}
type SimRealNameRegRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		MiniParam interface{} `json:"miniParam"`
		BusiSeq   string      `json:"busiSeq"`
		Url       string      `json:"url"`
	} `json:"result"`
}
type TPTokenRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
	Token string `json:"token"`
}
type TP3TokenRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
	Data    string `json:"data"`
	IntCode int    `json:"intCode"`
}

type TP4TokenRes struct {
	Time    int    `json:"time"`
	Code    int    `json:"code"`
	Token   string `json:"token"`
	Outtime int    `json:"outtime"`
	Out     int    `json:"out"`
}
