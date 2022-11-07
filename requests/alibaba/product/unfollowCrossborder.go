package product

import (
	"strconv"

	"github.com/bububa/go1688"
)

// UnfollowCrossborderRequest 解除关注商品API Request
type UnfollowCrossborderRequest struct {
	ProductID uint64 `json:"productId,omitempty"`
}

// Name implement RequestData interface
func (r UnfollowCrossborderRequest) Name() string {
	return "alibaba.product.unfollow.crossborder"
}

// Map implement RequestData interface
func (r UnfollowCrossborderRequest) Map() map[string]string {
	return map[string]string{
		"productId": strconv.FormatUint(r.ProductID, 10),
	}
}

// UnfollowCrossborderResponse 解除关注商品 API Response
type UnfollowCrossborderResponse struct {
	go1688.BaseResponse
	// Code 0表示成功
	Code int `json:"code,omitempty"`
	// Message 结果的描述
	Message string `json:"message,omitempty"`
}

// IsError check response success
func (f UnfollowCrossborderResponse) IsError() bool {
	return f.Code != 0
}

// Error implement error interface
func (f UnfollowCrossborderResponse) Error() string {
	if !f.IsError() {
		return f.BaseResponse.Error()
	}
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE:")
	builder.WriteString(strconv.Itoa(f.Code))
	builder.WriteString(", MSG:")
	builder.WriteString(f.Message)
	return builder.String()
}

// UnfollowCrossborder 解除关注商品
func UnfollowCrossborder(client *go1688.Client, req *UnfollowCrossborderRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp UnfollowCrossborderResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return err
	}
	if resp.BaseResponse.IsError() {
		return resp.BaseResponse
	}
	return nil
}
