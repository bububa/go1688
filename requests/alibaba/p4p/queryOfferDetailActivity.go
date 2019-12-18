package p4p

import (
	"errors"

	"github.com/bububa/go1688"
)

type QueryOfferDetailActivityRequest struct {
	OfferId uint64 `json:"offerId"` // 商品id
}

func (this *QueryOfferDetailActivityRequest) Name() string {
	return "alibaba.cps.queryOfferDetailActivity"
}

type QueryOfferDetailActivityResponse struct {
	go1688.BaseResponse
	Result *QueryOfferDetailActivityResult `json:"result"` // 营销活动结果
}

type QueryOfferDetailActivityResult struct {
	ErrorCode string                    `json:"errorCode,omitempty"`
	ErrorMsg  string                    `json:"errorMsg,omitempty"`
	Success   bool                      `json:"success,omitempty"`
	Result    *UnionActivityOfferDetail `json:"result,omitempty"`
}

type UnionActivityOfferDetail struct {
	OfferId              uint64                       `json:"offerId,omitempty"`              // 商品id
	ActivityId           uint64                       `json:"activityId,omitempty"`           // 营销活动Id
	ActivityName         string                       `json:"activityName,omitempty"`         // 活动名称
	HotTime              go1688.JsonTime              `json:"hotTime,omitempty"`              // 预热时间,活动未开始,不可用活动价下单; 为null表示无预热时间
	StartTime            go1688.JsonTime              `json:"startTime,omitempty"`            // 活动开始时间；大于now时，活动有效
	EndTime              go1688.JsonTime              `json:"endTime,omitempty"`              // 活动结束时间；小于now时，活动有效
	BeginQuantity        uint                         `json:"beginQuantity,omitempty"`        // 活动起批量
	Stock                uint                         `json:"stock,omitempty"`                // 活动总库存，为null时使用offer原库存
	PersonLimitCount     int                          `json:"personLimitCount,omitempty"`     // 商品本身限购数，非活动价可购买数；-1表示不限，0表示可购买数为0；3个*LimitCount字段都等于-1时，表示没有任何限购
	PromotionLimitCount  int                          `json:"promotionLimitCount,omitempty"`  // 限购数，等于0且personLimitCount>0时，可以以原价下单，但不能以活动价下单；-1表示不限数量；3个*LimitCount字段都等于-1时，表示没有任何限购
	ActivityLimitCount   int                          `json:"activityLimitCount,omitempty"`   // 活动限购数；该场内活动商品限购数，-1表示不限购；0表示不可购买该场活动所有商品；3个*LimitCount字段都等于-1时，表示没有任何限购
	FreepostageStartTime go1688.JsonTime              `json:"freepostageStartTime,omitempty"` // 活动限时包邮开始时间；null 表示不限时
	freepostageEndTime   go1688.JsonTime              `json:"freepostageEndTime,omitempty"`   // 活动限时包邮结束时间；null 表示不限时
	ExcludeAreaList      []UnionActivityArea          `json:"excludeAreaList,omitemtpy"`      // 免包邮地区，与活动包邮配合使用
	RangePrice           *UnionActivityRangePrice     `json:"rangePrice,omitempty"`           // 如果offer是范围报价，且价格优惠是折扣的情况，返回折扣计算后的价格范围;优先取该字段，该字段为空时，表示分sku报价，取promotionItemList
	PromotionItemList    []UnionActivityPromotionItem `json:"promotionItemList,omitempty"`    // 优惠结果，根据优惠方式（PromotionInfo），结合offer的原价信息，计算出优惠结果：每个sku或者每个区间价的促销价，折扣率
	SkuStockList         []UnionActivitySkuStock      `json:"skuStockList,omitempty"`         // sku维度的库存结果
	IntroOrderFlow       string                       `json:"introOrderFlow,omitempty"`       // 这里平台会计算一个推荐使用的下单flow，可以用这个flow值调用下单接口
}

type UnionActivityArea struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type UnionActivityRangePrice struct {
	Price         []uint `json:"price,omitempty"`         // 区间价数组，每个item为一个价格,以分为单位
	BeginQuantity uint   `json:"beginQuantity,omitempty"` // 区间对应的起批量
}

type UnionActivityPromotionItem struct {
	SkuId          uint64 `json:"skuId,omitempty"`          // sku优惠结果时有意义；对于区间价的优惠结果，此字段无意义，可能为null
	OriginalPrice  uint   `json:"originalPrice,omitempty"`  // 原价，单位分
	PromotionPrice uint   `json:"promotionPrice,omitempty"` // 优惠价，单位分
}

type UnionActivitySkuStock struct {
	Stock uint   `json:"stock,omitempty"` // 库存
	SkuId uint64 `json:"skuId,omitempty"` // skuId
}

func QueryOfferDetailActivity(client *go1688.Client, req *QueryOfferDetailActivityRequest, accessToken string) (*UnionActivityOfferDetail, error) {
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
