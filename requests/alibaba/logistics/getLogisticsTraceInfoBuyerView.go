package logistics

import (
	"github.com/bububa/go1688"
)

type GetLogisticsTraceInfosBuyerViewRequest struct {
	LogisticsId string     `json:"logisticsId,omitempty"` // 该订单下的物流编号
	OrderId     uint64     `json:"orderId"`               // 订单号
	Website     APIWebsite `json:"webSite"`               // 是1688业务还是icbu业务
}

func (this *GetLogisticsTraceInfosBuyerViewRequest) Name() string {
	return "alibaba.trade.getLogisticsTraceInfo.buyerView"
}

type GetLogisticsTraceInfosBuyerViewResponse struct {
	go1688.BaseResponse
	Traces []*LogisticsTrace `json:"logisticsTrace,omitempty"`
}

type LogisticsTrace struct {
	LogisticsId     string          `json:"logisticsId,omitempty"`     // 物流信息ID
	OrderId         uint64          `json:"orderId,omitempty"`         // 订单编号
	LogisticsBillNo string          `json:"logisticsBillNo,omitempty"` // 物流单号，运单号
	Steps           []LogisticsStep `json:"logisticsSteps,omitempty"`  // 物流跟踪步骤
}

type LogisticsStep struct {
	AcceptTime string `json:"acceptTime,omitempty"` // 物流跟踪单该步骤的时间
	Remark     string `json:"remark,omitempty"`     // 备注，如：“在浙江浦江县公司进行下级地点扫描，即将发往：广东深圳公司”
}

func GetLogisticsTraceInfosBuyerView(client *go1688.Client, req *GetLogisticsTraceInfosBuyerViewRequest, accessToken string) ([]*LogisticsTrace, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &GetLogisticsTraceInfosBuyerViewResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Traces, nil
}
