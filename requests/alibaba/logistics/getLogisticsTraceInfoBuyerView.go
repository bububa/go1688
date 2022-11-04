package logistics

import (
	"strconv"

	"github.com/bububa/go1688"
)

// GetLogisticsTraceInfosBuyerViewRequest 获取交易订单的物流跟踪信息(买家视角) API Request
type GetLogisticsTraceInfosBuyerViewRequest struct {
	// LogisticsID 该订单下的物流编号
	LogisticsID string `json:"logisticsId,omitempty"`
	// OrderID 订单号
	OrderID uint64 `json:"orderId,omitempty"`
	// Website 是1688业务还是icbu业务
	Website APIWebsite `json:"webSite,omitempty"`
}

// Name implement RequestData interface
func (r GetLogisticsTraceInfosBuyerViewRequest) Name() string {
	return "alibaba.trade.getLogisticsTraceInfo.buyerView"
}

// Map implement RequestData interface
func (r GetLogisticsTraceInfosBuyerViewRequest) Map() map[string]string {
	ret := make(map[string]string, 3)
	if r.LogisticsID != "" {
		ret["logisticsId"] = r.LogisticsID
	}
	ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	ret["webSite"] = r.Website
	return ret
}

// GetLogisticsTraceInfosBuyerViewResponse 获取交易订单的物流跟踪信息(买家视角) API Response
type GetLogisticsTraceInfosBuyerViewResponse struct {
	go1688.BaseResponse
	// Traces 跟踪单详情
	Traces []LogisticsTrace `json:"logisticsTrace,omitempty"`
}

// LogisticsTrace 跟踪单详情
type LogisticsTrace struct {
	// LogisticsID 物流信息ID
	LogisticsID string `json:"logisticsId,omitempty"`
	// OrderID 订单编号
	OrderID uint64 `json:"orderId,omitempty"`
	// LogisticsBillNo 物流单号，运单号
	LogisticsBillNo string `json:"logisticsBillNo,omitempty"`
	// Steps 物流跟踪步骤
	Steps []LogisticsStep `json:"logisticsSteps,omitempty"`
}

// LogisticsStep 物流跟踪步骤
type LogisticsStep struct {
	// AcceptTime 物流跟踪单该步骤的时间
	AcceptTime string `json:"acceptTime,omitempty"`
	// Remark 备注，如：“在浙江浦江县公司进行下级地点扫描，即将发往：广东深圳公司”
	Remark string `json:"remark,omitempty"`
}

// GetLogisticsTraceInfosBuyerView 获取交易订单的物流跟踪信息(买家视角)
// 该接口需要获取订单买家的授权，获取买家的订单的物流跟踪信息，在采购或者分销场景中，作为买家也有获取物流详情的需求。该接口能查能根据物流单号查看物流单跟踪信息。由于物流单录入的原因，可能跟踪信息的API查询会有延迟。该API需要向开放平台申请权限才能访问。In the procurement or distribution scenario, buyers can obtain information on logistics tracking. The interface can view the logistics tracking information according to the logistics tacking number. Depending on the logistics information entry time, there may be a delay in API queries regarding the information tracking.
func GetLogisticsTraceInfosBuyerView(client *go1688.Client, req *GetLogisticsTraceInfosBuyerViewRequest, accessToken string) ([]LogisticsTrace, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp GetLogisticsTraceInfosBuyerViewResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Traces, nil
}
