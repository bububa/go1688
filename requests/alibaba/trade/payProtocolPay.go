package trade

import (
	"github.com/bububa/go1688"
)

type PayProtocolPayRequest struct {
	OrderId uint64 `json:"orderId"` // 订单ID
}

func (this *PayProtocolPayRequest) Name() string {
	return "alibaba.trade.pay.protocolPay"
}

type PayProtocolPayResponse struct {
	go1688.BaseResponse
	Code    string `json:"code,omitempty"`    // 错误码
	Message string `json:"message,omitempty"` // 结果的描述
}

func PayProtocolPay(client *go1688.Client, req *PayProtocolPayRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &PayProtocolPayResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	if !resp.Success {
		return resp
	}
	return nil
}
