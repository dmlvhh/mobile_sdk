package mobile_sdk

type GetTokenReq struct {
	Appid     string `json:"appid"`
	AppSecret string `json:"app_secret"`
	ChannelId string `json:"channel_id"`
	Times     string `json:"times"`
	Sign      string `json:"sign"`
}
type Request struct {
	TransID string      `json:"trans_id"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}
type SimBasicInfoReq struct {
	Transid           string `json:"transid"`
	Token             string `json:"token"`
	Msisdn            string `json:"msisdn,omitempty"`
	Iccid             string `json:"iccid,omitempty"`
	Imsi              string `json:"imsi,omitempty"`
	OperType          string `json:"operType,omitempty"`
	TestType          string `json:"testType,omitempty"`
	ApnName           string `json:"apnName,omitempty"`
	ServiceUsageState string `json:"serviceUsageState,omitempty"`
}

type SimBasicInfoBatchReq struct {
	Transid   string `json:"transid"`
	Token     string `json:"token"`
	Msisdns   string `json:"msisdns,omitempty"`
	Iccids    string `json:"iccids,omitempty"`
	Imsis     string `json:"imsis,omitempty"`
	OperType  string `json:"operType,omitempty"`
	Reason    string `json:"reason,omitempty"`
	QueryDate string `json:"queryDate,omitempty"`
}

type SimDataDiagnosisReq struct {
}
