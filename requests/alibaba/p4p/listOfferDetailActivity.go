package p4p

import (
	"strconv"

	"github.com/bububa/go1688"
)

// ListOfferDetailActivityRequest 获取所有可用营销活动列表(媒体选择要使用的最优活动) API Request
type ListOfferDetailActivityRequest struct {
	// OfferID 商品id
	OfferID uint64 `json:"offerId,omitempty"`
}

// Name implement RequestData interface
func (r ListOfferDetailActivityRequest) Name() string {
	return "alibaba.cps.listOfferDetailActivity"
}

// Map implement RequestData interface
func (r ListOfferDetailActivityRequest) Map() map[string]string {
	return map[string]string{
		"offerId": strconv.FormatUint(r.OfferID, 10),
	}
}

// ListOfferDetailActivity 获取所有可用营销活动列表(媒体选择要使用的最优活动)
func ListOfferDetailActivity(client *go1688.Client, req *ListOfferDetailActivityRequest, accessToken string) (*UnionActivityOfferDetail, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp QueryOfferDetailActivityResponse
	if err := client.Do(finalRequest, accessToken, resp); err != nil {
		return nil, err
	}
	if resp.Result.IsError() {
		return nil, resp.Result
	}
	return resp.Result.Result, nil
}
