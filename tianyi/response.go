package tianyi

type GetCardInfoRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info struct {
		Iccid         string      `json:"iccid"`
		TotalFlow     string      `json:"total_flow"`
		UsedFlow      string      `json:"used_flow"`
		Status        int         `json:"status"`
		OfficialReal  int         `json:"official_real"`
		ExpiredAt     interface{} `json:"expired_at"`
		RealUsedFlow  string      `json:"real_used_flow"`
		IsPartyReal   interface{} `json:"is_party_real"`
		IsFreeze      int         `json:"is_freeze"`
		IsFreezeNotes interface{} `json:"is_freeze_notes"`
	} `json:"info"`
}
type SetSimStatusRes struct {
	Code      int           `json:"code"`
	Msg       string        `json:"msg"`
	Info      []interface{} `json:"info"`
	Attribute string        `json:"attribute"`
	Pname     string        `json:"pname"`
}
type UnbindSimCardRes struct {
	Code      int           `json:"code"`
	Msg       string        `json:"msg"`
	Info      []interface{} `json:"info"`
	Attribute string        `json:"attribute"`
	Pname     string        `json:"pname"`
}
type QueryOfficialRealStatusRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info struct {
		OfficialReal int `json:"official_real"`
		IsPartyReal  int `json:"is_party_real"`
	} `json:"info"`
	Attribute string `json:"attribute"`
	Pname     string `json:"pname"`
}
type RefreshSimRes struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	OfficialUrl string `json:"official_url"`
}
type PlaceOrderRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info struct {
		OrderNo string `json:"order_no"`
	} `json:"info"`
}

type GetBatchCardInfoRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info []struct {
		Iccid        string      `json:"iccid"`
		TotalFlow    string      `json:"total_flow"`
		UsedFlow     interface{} `json:"used_flow"`
		Status       int         `json:"status"`
		OfficialReal int         `json:"official_real"`
		ExpiredAt    string      `json:"expired_at"`
		RealUsedFlow string      `json:"real_used_flow"`
		IsPartyReal  int         `json:"is_party_real"`
	} `json:"info"`
}
