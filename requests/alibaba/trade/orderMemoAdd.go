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

// OrderMemoAdd 修改订单备忘
func OrderMemoAdd(client *go1688.Client, req *OrderMemoAddRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	return client.Do(finalRequest, accessToken, nil)
}
