package p4p

import (
	"strconv"

	"github.com/bububa/go1688"
)

// QueryOfferDetailActivityRequest 获取营销活动价格等活动信息 API Request
type QueryOfferDetailActivityRequest struct {
	// OfferID 商品id
	OfferID uint64 `json:"offerId"`
}

// Name implement RequestData interface
func (r QueryOfferDetailActivityRequest) Name() string {
	return "alibaba.cps.queryOfferDetailActivity"
}

// Map implement RequestData interface
func (r QueryOfferDetailActivityRequest) Map() map[string]string {
	return map[string]string{
		"offerId": strconv.FormatUint(r.OfferID, 10),
	}
}

// QueryOfferDetailActivityResponse 获取营销活动价格等活动信息 API Response
type QueryOfferDetailActivityResponse struct {
	go1688.BaseResponse
	// Result 营销活动结果
	Result *QueryOfferDetailActivityResult `json:"result"`
}

// QueryOfferDetailActivityResult 营销活动结果
type QueryOfferDetailActivityResult struct {
	ErrorCode string                    `json:"errorCode,omitempty"`
	ErrorMsg  string                    `json:"errorMsg,omitempty"`
	Success   bool                      `json:"success,omitempty"`
	Result    *UnionActivityOfferDetail `json:"result,omitempty"`
}

// IsError check success
func (r QueryOfferDetailActivityResult) IsError() bool {
	return !r.Success
}

// Error implement error interface
func (r QueryOfferDetailActivityResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE:")
	builder.WriteString(r.ErrorCode)
	builder.WriteString(", MSG:")
	builder.WriteString(r.ErrorMsg)
	return builder.String()
}

// UnionActivityOfferDetail 营销活动
type UnionActivityOfferDetail struct {
	// OfferID 商品id
	OfferID uint64 `json:"offerId,omitempty"`
	// ActivityID 营销活动Id
	ActivityID uint64 `json:"activityId,omitempty"`
	// ActivityName 活动名称
	ActivityName string `json:"activityName,omitempty"`
	// HotTime 预热时间,活动未开始,不可用活动价下单; 为null表示无预热时间
	HotTime go1688.JsonTime `json:"hotTime,omitempty"`
	// StartTime 活动开始时间；大于now时，活动有效
	StartTime go1688.JsonTime `json:"startTime,omitempty"`
	// EndTime 活动结束时间；小于now时，活动有效
	EndTime go1688.JsonTime `json:"endTime,omitempty"`
	// BeginQuantity 活动起批量
	BeginQuantity int64 `json:"beginQuantity,omitempty"`
	// Stock 活动总库存，为null时使用offer原库存
	Stock int64 `json:"stock,omitempty"`
	// PersionLimitCount 商品本身限购数，非活动价可购买数；-1表示不限，0表示可购买数为0；3个*LimitCount字段都等于-1时，表示没有任何限购
	PersonLimitCount int64 `json:"personLimitCount,omitempty"`
	// PromotionLimitCount 限购数，等于0且personLimitCount>0时，可以以原价下单，但不能以活动价下单；-1表示不限数量；3个*LimitCount字段都等于-1时，表示没有任何限购
	PromotionLimitCount int64 `json:"promotionLimitCount,omitempty"`
	// ActivityLimitCount 活动限购数；该场内活动商品限购数，-1表示不限购；0表示不可购买该场活动所有商品；3个*LimitCount字段都等于-1时，表示没有任何限购
	ActivityLimitCount int64 `json:"activityLimitCount,omitempty"`
	// FreepostageStartTime 活动限时包邮开始时间；null 表示不限时
	FreepostageStartTime go1688.JsonTime `json:"freepostageStartTime,omitempty"`
	// FreepostageEndTime 活动限时包邮结束时间；null 表示不限时
	FreepostageEndTime go1688.JsonTime `json:"freepostageEndTime,omitempty"`
	// ExcludeAreaList 免包邮地区，与活动包邮配合使用
	ExcludeAreaList []UnionActivityArea `json:"excludeAreaList,omitemtpy"`
	// RangePrice 如果offer是范围报价，且价格优惠是折扣的情况，返回折扣计算后的价格范围;优先取该字段，该字段为空时，表示分sku报价，取promotionItemList
	RangePrice *UnionActivityRangePrice `json:"rangePrice,omitempty"`
	// PromotionItemList 优惠结果，根据优惠方式（PromotionInfo），结合offer的原价信息，计算出优惠结果：每个sku或者每个区间价的促销价，折扣率
	PromotionItemList []UnionActivityPromotionItem `json:"promotionItemList,omitempty"`
	// SkuStockList  sku维度的库存结果
	SkuStockList []UnionActivitySkuStock `json:"skuStockList,omitempty"`
	// IntroOrderFlow 这里平台会计算一个推荐使用的下单flow，可以用这个flow值调用下单接口
	IntroOrderFlow string `json:"introOrderFlow,omitempty"`
}

// UnionActivityArea 免包邮地区，与活动包邮配合使用
type UnionActivityArea struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// UnionActivityRangePrice 活动价格区间
type UnionActivityRangePrice struct {
	// Price 区间价数组，每个item为一个价格,以分为单位
	Price []int64 `json:"price,omitempty"`
	// BeginQuantity 区间对应的起批量
	BeginQuantity int64 `json:"beginQuantity,omitempty"`
}

// UnionActivityPromotionItem 优惠结果
type UnionActivityPromotionItem struct {
	// SkuID sku优惠结果时有意义；对于区间价的优惠结果，此字段无意义，可能为null
	SkuID uint64 `json:"skuId,omitempty"`
	// OriginalPrice 原价，单位分
	OriginalPrice int64 `json:"originalPrice,omitempty"`
	// PromotionPrice 优惠价，单位分
	PromotionPrice int64 `json:"promotionPrice,omitempty"`
}

// UnionActivitySkuStock  库存结果
type UnionActivitySkuStock struct {
	// Stock 库存
	Stock int64 `json:"stock,omitempty"`
	// SkuID skuId
	SkuID uint64 `json:"skuId,omitempty"`
}

// QueryOfferDetailActivity 获取营销活动价格等活动信息
func QueryOfferDetailActivity(client *go1688.Client, req *QueryOfferDetailActivityRequest, accessToken string) (*UnionActivityOfferDetail, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp QueryOfferDetailActivityResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	if resp.Result.IsError() {
		return nil, resp.Result
	}
	return resp.Result.Result, nil
}
