package p4p

import (
	"errors"

	"github.com/bububa/go1688"
)

type ListOfferDetailActivityRequest struct {
	OfferId uint64 `json:"offerId"` // 商品id
}

func (this *ListOfferDetailActivityRequest) Name() string {
	return "alibaba.cps.listOfferDetailActivity"
}

func ListOfferDetailActivity(client *go1688.Client, req *ListOfferDetailActivityRequest, accessToken string) (*UnionActivityOfferDetail, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &QueryOfferDetailActivityResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	if resp.Result.Success {
		return nil, errors.New(resp.Result.ErrorMsg)
	}
	return resp.Result.Result, nil
}
