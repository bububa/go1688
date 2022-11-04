package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// OrderMemoAddRequest 修改订单备忘 API Request
type OrderMemoAddRequest struct {
	// OrderID 交易id，订单号
	OrderID uint64 `json:"orderId,omitempty"`
	// Memo 备忘信息
	Memo string `json:"memo,omitempty"`
	// RemarkIcon 备忘图标，目前仅支持数字。1位红色图标，2为蓝色图标，3为绿色图标，4为黄色图标
	RemarkIcon string `json:"remarkIcon,omitempty"`
}

// Name implement RequestData interface
func (r OrderMemoAddRequest) Name() string {
	return "alibaba.order.memoAdd"
}

// Map implement RequestData interface
func (r OrderMemoAddRequest) Map() map[string]string {
	ret := make(map[string]string, 3)
	ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	ret["memo"] = r.Memo
	if r.RemarkIcon != "" {
		ret["remarkIcon"] = r.RemarkIcon
	}
	return ret
}

// OrderMemoAddResponse 修改订单备忘
type OrderMemoAddResponse struct {
	go1688.BaseResponse
	// ErrorCode 错误编码
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorMsg 错误信息
	ErrorMsg string `json:"errorMsg,omitempty"`
	// Success 是否成功
	Success bool `json:"success,omitempty"`
}

// IsError check success
func (r OrderMemoAddResponse) IsError() bool {
	return !r.Success
}

func (r OrderMemoAddResponse) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.ErrorCode)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.ErrorMsg)
	return builder.String()
}

// OrderMemoAdd 修改订单备忘
func OrderMemoAdd(client *go1688.Client, req *OrderMemoAddRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp OrderMemoAddResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return err
	}
	if resp.BaseResponse.IsError() {
		return resp.BaseResponse
	}
	return nil
}
