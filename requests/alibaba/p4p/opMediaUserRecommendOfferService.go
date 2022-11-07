package p4p

import (
	"strconv"

	"github.com/shopspring/decimal"

	"github.com/bububa/go1688"
)

// OpMediaUserRecommendOfferServiceRequest 机构用户个性化商品推荐接口 API Request
type OpMediaUserRecommendOfferServiceRequest struct {
	// DeviceIDMd5 设备id求md5(32位小写)(手机号与设备号至少一个)
	DeviceIDMd5 string `json:"deviceIdMd5,omitempty"`
	// PhoneMd5 手机号求md5(32位小写)(手机号与设备号至少一个)
	PhoneMd5 string `json:"phoneMd5,omitempty"`
	// PageNo 页码
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 每页数量
	PageSize int `json:"pageSize,omitempty"`
}

// Name implement RequestData interface
func (r OpMediaUserRecommendOfferServiceRequest) Name() string {
	return "alibaba.cps.op.mediaUserRecommendOfferService"
}

// Map implement RequestData interface
func (r OpMediaUserRecommendOfferServiceRequest) Map() map[string]string {
	ret := make(map[string]string, 4)
	if r.DeviceIDMd5 != "" {
		ret["deviceIdMd5"] = r.DeviceIDMd5
	}
	if r.PhoneMd5 != "" {
		ret["phoneMd5"] = r.PhoneMd5
	}
	if r.PageNo > 0 {
		ret["pageNo"] = strconv.Itoa(r.PageNo)
	}
	if r.PageSize > 0 {
		ret["pageSize"] = strconv.Itoa(r.PageSize)
	}
	return ret
}

// OpMediaUserRecommendOfferServiceResponse 机构用户个性化商品推荐接口 API Response
type OpMediaUserRecommendOfferServiceResponse struct {
	go1688.BaseResponse
	// Result 营销活动结果
	Result *OpMediaUserRecommendOfferServiceResult `json:"result"`
}

// OpMediaUserRecommendOfferServiceResult 营销活动结果
type OpMediaUserRecommendOfferServiceResult struct {
	ErrorCode string           `json:"errorCode,omitempty"`
	ErrorMsg  string           `json:"errorMsg,omitempty"`
	Success   bool             `json:"success,omitempty"`
	Result    []RecommendOffer `json:"result,omitempty"`
}

// IsError check success
func (r OpMediaUserRecommendOfferServiceResult) IsError() bool {
	return !r.Success
}

// Error implement error interface
func (r OpMediaUserRecommendOfferServiceResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE:")
	builder.WriteString(r.ErrorCode)
	builder.WriteString(", MSG:")
	builder.WriteString(r.ErrorMsg)
	return builder.String()
}

// RecommendOffer 推荐商品
type RecommendOffer struct {
	// OfferID 商品id
	OfferID uint64 `json:"offerId,omitempty"`
	// ImgURL 商品首图
	ImgURL string `json:"imgUrl,omitempty"`
	// RecommendPrice 建议零售价
	RecommendPrice decimal.Decimal `json:"recommendPrice,omitempty"`
	// RecommendTitle 智能标题
	RecommendTitle string `json:"recommendTitle,omitempty"`
}

// OpMediaUserRecommendOfferService 机构用户个性化商品推荐接口
func OpMediaUserRecommendOfferService(client *go1688.Client, req *OpMediaUserRecommendOfferServiceRequest, accessToken string) ([]RecommendOffer, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp OpMediaUserRecommendOfferServiceResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	if resp.Result.IsError() {
		return nil, resp.Result
	}
	return resp.Result.Result, nil
}
