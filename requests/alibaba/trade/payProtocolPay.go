package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// PayProtocolPayRequest 支付宝协议代扣支付 API Request
type PayProtocolPayRequest struct {
	// OrderID 订单ID
	OrderID uint64 `json:"orderId,omitempty"`
}

// Name implement RequestData interface
func (r PayProtocolPayRequest) Name() string {
	return "alibaba.trade.pay.protocolPay"
}

// Map implement RequestData interface
func (r PayProtocolPayRequest) Map() map[string]string {
	return map[string]string{
		"orderId": strconv.FormatUint(r.OrderID, 10),
	}
}

// PayProtocolPayResponse 支付宝协议代扣支付 API Response
type PayProtocolPayResponse struct {
	go1688.BaseResponse
	// Code 错误码
	Code string `json:"code,omitempty"`
	// Message 结果的描述
	Message string `json:"message,omitempty"`
}

// IsError check success
func (r PayProtocolPayResponse) IsError() bool {
	return !r.Success.Bool()
}

// Error implement error interface
func (r PayProtocolPayResponse) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.Code)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.Message)
	return builder.String()
}

// PayProtocolPay 支付宝协议代扣支付
func PayProtocolPay(client *go1688.Client, req *PayProtocolPayRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp PayProtocolPayResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return err
	}
	if resp.BaseResponse.IsError() {
		return resp.BaseResponse
	}
	return nil
}
