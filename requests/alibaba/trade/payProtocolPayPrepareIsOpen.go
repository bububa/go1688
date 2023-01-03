package trade

import (
	"github.com/bububa/go1688"
)

// PayProtocolPayIsOpenRequest 查询是否开通免密支付 API Request
type PayProtocolPayIsOpenRequest struct {
}

// Name implement RequestData interface
func (r PayProtocolPayIsOpenRequest) Name() string {
	return "alibaba.trade.pay.protocolPay.isopen"
}

// Map implement RequestData interface
func (r PayProtocolPayIsOpenRequest) Map() map[string]string {
	return make(map[string]string)
}

// PayProtocolPayIsOpenResponse 查询是否开通免密支付 API Response
type PayProtocolPayIsOpenResponse struct {
	go1688.BaseResponse
	// Result 免密支付结果
	Result *PayProtocolPayIsOpenResult `json:"result,omitempty"`
}

// IsError check success
func (r PayProtocolPayIsOpenResponse) IsError() bool {
	return r.Result != nil && r.Result.IsError()
}

// Error implement error interface
func (r PayProtocolPayIsOpenResponse) Error() string {
	if r.Result == nil {
		return ""
	}
	return r.Result.Error()
}

// PayProtocolPayIsOpenResult 免密支付结果
type PayProtocolPayIsOpenResult struct {
	// Success 是否成功
	Success go1688.Bool `json:"success,omitempty"`
	// Code 错误码
	Code string `json:"code,omitempty"`
	// Message 结果的描述
	Message string `json:"message,omitempty"`
	// Result 扣款返回值
	Result *TradeWithholdStatusResult `json:"result,omitempty"`
}

// IsError check success
func (r PayProtocolPayIsOpenResult) IsError() bool {
	return !r.Success.Bool()
}

// Error implement error interface
func (r PayProtocolPayIsOpenResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.Code)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.Message)
	return builder.String()
}

// TradeWithholdStatusResult 签约状态
type TradeWithholdStatusResult struct {
	// PaymentAgreements .
	PaymentAgreements []TradePaymentAgreement `json:"paymentAgreements,omitempty"`
}

// TradePaymentAgreement
type TradePaymentAgreement struct {
	// PayChannel 支付成功渠道
	PayChannel string `json:"payChannel,omitempty"`
	// BindingStatus 支付宝或者诚E赊是否已设置绑定，signedStatus和bindingStatus均为true才能发起代扣
	BindingStatus go1688.Bool `json:"bindingStatus,omitempty"`
	// SignedStatus 支付宝或者诚E赊是否已签约代扣，signedStatus和bindingStatus均为true才能发起代扣
	SignedStatus go1688.Bool `json:"signedStatus,omitempty"`
	// SignURL 签约URl
	SignURL string `json:"signUrl,omitempty"`
	// AgreementNo 签约单号
	AgreementNo string `json:"agreementNo,omitempty"`
}

// PayProtocolPayIsOpen 查询是否开通免密支付
func PayProtocolPayIsOpen(client *go1688.Client, req *PayProtocolPayIsOpenRequest, accessToken string) ([]TradePaymentAgreement, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp PayProtocolPayIsOpenResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	if resp.BaseResponse.IsError() {
		return nil, resp.BaseResponse
	}
	if resp.Result != nil && resp.Result.Result != nil {
		return resp.Result.Result.PaymentAgreements, nil
	}
	return nil, nil
}
