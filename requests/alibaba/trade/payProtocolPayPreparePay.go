package trade

import (
	"github.com/bububa/go1688"
)

// PayProtocolPayPreparePayRequest 发起免密支付 API Request
type PayProtocolPayPreparePayRequest struct {
	// OrderID 订单ID
	OrderID uint64 `json:"orderId,omitempty"`
}

// Name implement RequestData interface
func (r PayProtocolPayPreparePayRequest) Name() string {
	return "alibaba.trade.pay.protocolPay.preparePay"
}

// Map implement RequestData interface
func (r PayProtocolPayPreparePayRequest) Map() map[string]string {
	return map[string]string{
		"tradeWithholdPreparePayParam": go1688.JSONMarshal(r),
	}
}

// PayProtocolPayPreparePayResponse 发起免密支付 API Response
type PayProtocolPayPreparePayResponse struct {
	go1688.BaseResponse
	// Result 免密支付结果
	Result *PayProtocolPayPreparePayResult `json:"result,omitempty"`
}

// IsError check success
func (r PayProtocolPayPreparePayResponse) IsError() bool {
	return r.Result != nil && r.Result.IsError()
}

// Error implement error interface
func (r PayProtocolPayPreparePayResponse) Error() string {
	if r.Result == nil {
		return ""
	}
	return r.Result.Error()
}

// PayProtocolPayPreparePayResult 免密支付结果
type PayProtocolPayPreparePayResult struct {
	// Success 是否成功
	Success go1688.Bool `json:"success,omitempty"`
	// Code 错误码
	Code string `json:"code,omitempty"`
	// Message 结果的描述
	Message string `json:"message,omitempty"`
	// Result 扣款返回值
	Result *TradeWithholdPreparePayResult `json:"result,omitempty"`
}

// TradeWithholdPreparePayResult
type TradeWithholdPreparePayResult struct {
	// PayChannel 支付成功渠道
	PayChannel string `json:"payChannel,omitempty"`
	// PaySuccess 支付是否成功
	PaySuccess bool `json:"paySuccess,omitempty"`
}

// IsError check success
func (r PayProtocolPayPreparePayResult) IsError() bool {
	return !r.Success.Bool()
}

// Error implement error interface
func (r PayProtocolPayPreparePayResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.Code)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.Message)
	return builder.String()
}

// PayProtocolPayPreparePay 发起免密支付
func PayProtocolPayPreparePay(client *go1688.Client, req *PayProtocolPayPreparePayRequest, accessToken string) (*TradeWithholdPreparePayResult, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp PayProtocolPayPreparePayResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	if resp.BaseResponse.IsError() {
		return nil, resp.BaseResponse
	}
	if resp.Result != nil {
		return resp.Result.Result, nil
	}
	return nil, nil
}
