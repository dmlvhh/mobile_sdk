package mobile_sdk

import (
	"encoding/json"
	"fmt"
	"log"
)

// SimBasicInfo 单卡基本信息查询
func (c *Config) SimBasicInfo(req *SimBasicInfoReq) (res *SimBasicInfoRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-basic-info", req)
	if err != nil {
		log.Printf("SimBasicInfo: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimChangeHistory 单卡状态变更历史查询
func (c *Config) SimChangeHistory(req *SimBasicInfoReq) (res *SimChangeHistoryRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-change-history", req)
	if err != nil {
		log.Printf("SimChangeHistory: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimImei 单卡绑定IMEI实时查询
func (c *Config) SimImei(req *SimBasicInfoReq) (res *SimImeiRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-imei", req)
	if err != nil {
		log.Printf("SimImei: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimStatus 单卡状态查询
func (c *Config) SimStatus(req *SimBasicInfoReq) (res *SimStatusRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-status", req)
	if err != nil {
		log.Printf("SimStatus: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimStopReason 单卡停机原因查询
func (c *Config) SimStopReason(req *SimBasicInfoReq) (res *SimStopReasonRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-stop-reason", req)
	if err != nil {
		log.Printf("SimStopReason: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimDataMargin 单卡本月套餐内流量使用量实时查询 todo 异常
func (c *Config) SimDataMargin(req *SimBasicInfoReq) (res *SimDataMarginRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-data-margin", req)
	if err != nil {
		log.Printf("SimDataMargin: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimDataUsage 单卡本月流量累计使用量查询
func (c *Config) SimDataUsage(req *SimBasicInfoReq) (res *SimDataUsageRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-data-usage", req)
	if err != nil {
		log.Printf("SimDataUsage: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimDataUsageMonthlyBatch 物联卡单月GPRS流量使用量批量查询
func (c *Config) SimDataUsageMonthlyBatch(req *SimBasicInfoBatchReq) (res *SimDataUsageMonthlyBatchRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-data-usage-monthly/batch", req)
	if err != nil {
		log.Printf("SimDataUsageMonthlyBatch: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// ChangeSimStatus 单卡状态变更
func (c *Config) ChangeSimStatus(req *SimBasicInfoReq) (res *ChangeSimStatusRes, err error) {
	request, err := c.ApiRequest("/v5/ec/change/sim-status", req)
	if err != nil {
		log.Printf("ChangeSimStatus: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// ChangeSimStatusBatch 物联卡状态变更批量办理
func (c *Config) ChangeSimStatusBatch(req *SimBasicInfoBatchReq) (res *ChangeSimStatusBatchRes, err error) {
	request, err := c.ApiRequest("/v5/ec/change/sim-status/batch", req)
	if err != nil {
		log.Printf("ChangeSimStatusBatch: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimCardInfoBatch 码号信息批量查询
func (c *Config) SimCardInfoBatch(req *SimBasicInfoBatchReq) (res *SimBasicInfoBatchRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-card-info/batch", req)
	if err != nil {
		log.Printf("SimCardInfoBatch: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimPlatformBatch 物联卡归属平台批量查询
func (c *Config) SimPlatformBatch(req *SimBasicInfoBatchReq) (res *SimPlatformBatchRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-platform/batch", req)
	if err != nil {
		log.Printf("SimPlatformBatch: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// CardBindStatus 物联卡机卡分离状态查询
func (c *Config) CardBindStatus(req *SimBasicInfoReq) (res *CardBindStatusRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/card-bind-status", req)
	if err != nil {
		log.Printf("CardBindStatus: %s", err)
		return nil, err
	}
	//fmt.Println(string(request))
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimRealNameStatus 物联卡实名登记状态查询
func (c *Config) SimRealNameStatus(req *SimBasicInfoReq) (res *SimRealNameStatusRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-real-name-status", req)
	if err != nil {
		log.Printf("SimRealNameStatus: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimDataDiagnosis 物联卡达量限速档位查询
func (c *Config) SimDataDiagnosis(req *SimBasicInfoReq) (res *SimDataDiagnosisRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-data-diagnosis", req)
	if err != nil {
		log.Printf("SimDataDiagnosis: %s", err)
		return nil, err
	}
	//fmt.Println("SimDataDiagnosis", request)
	err = json.Unmarshal([]byte(request), &res)
	return
}

// NetworkSpeed 网卡限速
func (c *Config) NetworkSpeed(req *SimBasicInfoReq) (res *NetworkSpeedRes, err error) {
	request, err := c.ApiRequest("/v5/ec/operate/network-speed", req)
	if err != nil {
		log.Printf("NetworkSpeed: %s", err)
		return nil, err
	}
	fmt.Println("NetworkSpeed", request)
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimManageStopRestartStatus 物联卡管理停复机冻结状态查询
func (c *Config) SimManageStopRestartStatus(req *SimBasicInfoReq) (res *SimManageStopRestartStatusRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-manage-stop-restart-status", req)
	if err != nil {
		log.Printf("SimManageStopRestartStatus: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QuerySimImei 查询设别最新的IMEI
func (c *Config) QuerySimImei(req *SimBasicInfoReq) (res *NetworkSpeedRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-imei", req)
	if err != nil {
		log.Printf("QuerySimImei: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QueryOnOffStatus 单卡开关机状态实时查询 0:关机 1:开机
func (c *Config) QueryOnOffStatus(req *SimBasicInfoReq) (res *QueryOnOffStatusRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/on-off-status", req)
	if err != nil {
		log.Printf("QueryOnOffStatus: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// QuerySimSession 查询物联卡的在线信息，区分 APN，返回 APN 信息、IP 地址、会话创建时间
func (c *Config) QuerySimSession(req *SimBasicInfoReq) (res *SimSessionRes, err error) {
	request, err := c.ApiRequest("/v5/ec/query/sim-session", req)
	if err != nil {
		log.Printf("QuerySimSession: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}

// SimGprsStatusReset 物联卡 GPRS 上网功能重置
func (c *Config) SimGprsStatusReset(req *SimBasicInfoReq) (res *SimGprsStatusResetRes, err error) {
	request, err := c.ApiRequest("/v5/ec/operate/sim-gprs-status-reset", req)
	if err != nil {
		log.Printf("SimGprsStatusReset: %s", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(request), &res)
	return
}
