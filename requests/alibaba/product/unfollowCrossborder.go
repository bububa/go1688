package product

import (
	"errors"

	"github.com/bububa/go1688"
)

type UnfollowCrossborderRequest struct {
	ProductId uint64 `json:"productId"`
}

func (this *UnfollowCrossborderRequest) Name() string {
	return "alibaba.product.unfollow.crossborder"
}

type UnfollowCrossborderResponse struct {
	go1688.BaseResponse
	Code    int    `json:"code,omitempty"`    // 0表示成功
	Message string `json:"message,omitempty"` // 结果的描述
}

func UnfollowCrossborder(client *go1688.Client, req *UnfollowCrossborderRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &FollowResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	if resp.Code != 0 {
		return errors.New(resp.Message)
	}
	return nil
}
