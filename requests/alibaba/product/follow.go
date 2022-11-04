package product

import (
	"strconv"

	"github.com/bububa/go1688"
)

// FollowRequest 关注商品 APIRequest
type FollowRequest struct {
	// ProductID 商品id
	ProductID uint64 `json:"productId,omitempty"`
}

// Name implement RequestData interface
func (r FollowRequest) Name() string {
	return "alibaba.product.follow"
}

// Map implement RequestData interface
func (r FollowRequest) Map() map[string]string {
	return map[string]string{
		"productId": strconv.FormatUint(r.ProductID, 10),
	}
}

// FollowResponse 关注商品 API Response
type FollowResponse struct {
	go1688.BaseResponse
	// Code 0表示成功
	Code int `json:"code,omitempty"`
	// Message 结果的描述
	Message string `json:"message,omitempty"`
}

// IsError check response success
func (f FollowResponse) IsError() bool {
	return f.Code != 0
}

// Error implement error interface
func (f FollowResponse) Error() string {
	if !f.IsError() {
		return f.BaseResponse.Error()
	}
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(strconv.Itoa(f.Code))
	builder.WriteString(", MESSAGE: ")
	builder.WriteString(f.Message)
	return builder.String()
}

// Follow 关注商品
func Follow(client *go1688.Client, req *FollowRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp FollowResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return err
	}
	if resp.BaseResponse.IsError() {
		return resp
	}
	return nil
}
