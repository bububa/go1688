package trade

import (
	"encoding/json"

	"github.com/bububa/go1688"
)

// AlipayUrlGetRequest 批量获取订单的支付链接 API Request
type AlipayUrlGetRequest struct {
	// OrderIDs 订单Id列表,最多批量30个订单，订单过多会导致超时，建议一次10个订单
	OrderIDs []uint64 `json:"orderIdList"`
}

// Name implement RequestData interface
func (r AlipayUrlGetRequest) Name() string {
	return "alibaba.alipay.url.get"
}

// Map implement RequestData interface
func (r AlipayUrlGetRequest) Map() map[string]string {
	ids, _ := json.Marshal(r.OrderIDs)
	return map[string]string{
		"orderIdList": go1688.JSONMarshal(ids),
	}
}

// AlipayUrlGetResponse 批量获取订单的支付链接 API Response
type AlipayUrlGetResponse struct {
	go1688.BaseResponse
	// PayURL 支付链接
	PayURL string `json:"payUrl,omitempty"`
}

// AlipayUrlGet 批量获取订单的支付链接
func AlipayUrlGet(client *go1688.Client, req *AlipayUrlGetRequest, accessToken string) (string, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp AlipayUrlGetResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return "", err
	}
	return resp.PayURL, nil
}
