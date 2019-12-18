package trade

import (
	"encoding/json"

	"github.com/bububa/go1688"
)

type AlipayUrlGetRequest struct {
	OrderIds []uint64 `json:"orderIdList"` // 订单Id列表,最多批量30个订单，订单过多会导致超时，建议一次10个订单
}

func (this *AlipayUrlGetRequest) Refine() *AlipayUrlGetRefinedRequest {
	ids, _ := json.Marshal(this.OrderIds)
	return &AlipayUrlGetRefinedRequest{
		OrderIds: string(ids),
	}
}

type AlipayUrlGetRefinedRequest struct {
	OrderIds string `json:"orderIdList"` // 订单Id列表,最多批量30个订单，订单过多会导致超时，建议一次10个订单
}

func (this *AlipayUrlGetRefinedRequest) Name() string {
	return "alibaba.alipay.url.get"
}

type AlipayUrlGetResponse struct {
	go1688.BaseResponse
	PayUrl string `json:"payUrl,omitempty"` // 支付链接
}

func AlipayUrlGet(client *go1688.Client, req *AlipayUrlGetRequest, accessToken string) (string, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &AlipayUrlGetResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", resp
	}
	return resp.PayUrl, nil
}
