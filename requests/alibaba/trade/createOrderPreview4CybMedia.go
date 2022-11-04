package trade

import (
	"github.com/shopspring/decimal"

	"github.com/bububa/go1688"
)

// CreateOrderPreview4CybMediaRequest 溢价模式创建订单前预览数据接口 API Request
type CreateOrderPreview4CybMediaRequest struct {
	// Address 收货地址信息
	Address *Address `json:"addressParam,omitempty"`
	// CargoList 商品信息
	CargoList []Cargo `json:"cargoParamList,omitempty"`
	// UseChannelPrice 是否走渠道专属价，传true优先使用渠道传享价，不传或者传false则不走渠道专享价，走普通的分销价
	UseChannelPrice bool `json:"useChannelPrice,omitempty"`
	// Flow 不传，价格逻辑由useChannelPrice控制；传入general使用批发价，即useChannelPrice的价格逻辑失效；传入paired使用火拼价，若该商品未参与伙拼，则下单失败。
	Flow TradeFlow `json:"flow,omitempty"`
}

// Name implement RequestData interface
func (r CreateOrderPreview4CybMediaRequest) Name() string {
	return "alibaba.createOrder.preview4CybMedia"
}

// Map implement RequestData interface
func (r CreateOrderPreview4CybMediaRequest) Map() map[string]string {
	ret := make(map[string]string, 2)
	if r.Address != nil {
		ret["addressParam"] = go1688.JSONMarshal(r.Address)
	}
	if len(r.CargoList) > 0 {
		ret["cargoParamList"] = go1688.JSONMarshal(r.CargoList)
	}
	if r.UseChannelPrice {
		ret["useChannelPrice"] = "true"
	}
	if r.Flow != "" {
		ret["flow"] = r.Flow
	}
	return ret
}

// Address 收货地址信息
type Address struct {
	// ID 收货地址id
	ID uint64 `json:"addressId,omitempty"`
	// FullName 收货人姓名
	FullName string `json:"fullName,omitempty"`
	// Mobile 手机
	Mobile string `json:"mobile,omitempty"`
	// Phone 电话
	Phone string `json:"phone,omitempty"`
	// PostCode 邮编
	PostCode string `json:"postCode,omitempty"`
	// CityText 市文本
	CityText string `json:"cityText,omitempty"`
	// ProvinceText 省份文本
	ProvinceText string `json:"provinceText,omitempty"`
	// AreaText 区文本
	AreaText string `json:"areaText,omitempty"`
	// TownText 镇文本
	TownText string `json:"townText,omitempty"`
	// Address 街道地址
	Address string `json:"address,omitempty"`
	// DistrictCode 地址编码
	DistrictCode string `json:"districtCode,omitempty"`
}

// Cargo 商品信息
type Cargo struct {
	// OfferID 商品对应的offer id
	OfferID uint64 `json:"offerId,omitempty"`
	// SpecID 商品规格id, 无sku商品不传
	SpecID string `json:"specId,omitempty"`
	// Quantity 商品数量(计算金额用)
	Quantify float64 `json:"quantity,omitempty"`
}

// CreateOrderPreview4CybMediaResponse 溢价模式创建订单前预览数据接口 API Response
type CreateOrderPreview4CybMediaResponse struct {
	go1688.BaseResponse
	// PostFeeByDescOfferList 运费说明的商品列表
	PostFeeByDescOfferList []uint64 `json:"postFeeByDescOfferList,omitempty"`
	// ConsignOfferList 代销商品列表
	ConsignOfferList []uint64 `json:"consignOfferList,omitempty"`
	// OrderPreview 订单预览结果，过自动拆单会返回多个记录
	OrderPreview []OrderPreview `json:"orderPreviewResuslt,omitempty"`
}

// OrderPreview 订单预览结果
type OrderPreview struct {
	// DiscountFee 计算完货品金额后再次进行的减免金额. 单位: 分
	DiscountFee uint `json:"discountFee,omitempty"`
	// TradeModeNameList 当前交易可以支持的交易方式列表。某些场景的创建订单接口需要使用。
	TradeModeNameList []string `jsong:"tradeModeNameList,omitempty"`
	// Status 状态
	Status bool `json:"status,omitempty"`
	// TaoSampleSinglePromotion 是否有淘货源单品优惠 false:有单品优惠 true：没有单品优惠
	TaoSampleSinglePromotion bool `json:"taoSampleSinglePromotion,omitempty"`
	// SumPayment 订单总费用, 单位为分.
	SumPayment int64 `json:"sumPayment,omitempty"`
	// Message 返回信息
	Message string `json:"message,omitempty"`
	// SumCarriage 总运费信息, 单位为分.
	SumCarriage int64 `json:"sumCarriage,omitempty"`
	// ResultCode 返回码
	ResultCode string `json:"resultCode,omitempty"`
	// SumPaymentNoCarriage 不包含运费的货品总费用, 单位为分.
	SumPaymentNoCarriage int64 `json:"sumPaymentNoCarriage,omitempty"`
	// AdditionalFee 附加费,单位，分
	AdditionalFee int64 `json:"additionalFee,omitempty"`
	// FlowFlag 订单下单流程
	FlowFlag string `json:"flowFlag,omitempty"`
	// CargoList 规格信息
	CargoList []CargoInfo `json:"cargoList,omitempty"`
	// ShopPromotionList 可用店铺级别优惠列表
	ShopPromotionList []Promotion `json:"shopPromotionList,omitempty"`
	// TradeModeList 当前交易可以支持的交易方式列表。结果可以参照1688下单预览页面的交易方式。
	TradeModeList []TradeMode `json:"tradeModeList,omitempty"`
}

// CargoInfo 规格信息
type CargoInfo struct {
	// OfferID 商品ID
	OfferID uint64 `json:"offerId,omitempty"`
	// Ammount 产品总金额
	Amount decimal.Decimal `json:"amount,omitempty"`
	// Message 返回信息
	Message string `json:"message,omitempty"`
	// FinalUnitPrice 最终单价
	FinalUnitPrice decimal.Decimal `json:"finalUnitPrice,omitempty"`
	// SpecID 规格ID，offer内唯一
	SpecID string `json:"specId,omitempty"`
	// SkuID 规格ID，全局唯一
	SkuID uint64 `json:"skuId,omitempty"`
	// ResultCode 返回码
	ResultCode string `json:"resultCode,omitempty"`
	// PromotionList 商品优惠列表
	PromotionList []Promotion `json:"cargoPromotionList,omitempty"`
}

// Promotion 商品优惠
type Promotion struct {
	// ID 优惠券ID
	ID string `json:"promotionId,omitempty"`
	// Selected 是否默认选中
	Selected bool `json:"selected,omitempty"`
	// Text 优惠券名称
	Text string `json:"text,omitempty"`
	// Desc 优惠券描述
	Desc string `json:"desc,omitempty"`
	// FreePostage 是否免邮
	FreePostage bool `json:"freePostage,omitempty"`
	// DiscountFee 减去金额，单位为分
	DiscountFee int64 `json:"discountFee,omitempty"`
}

// TradeMode 当前交易可以支持的交易方式
type TradeMode struct {
	// TradeWay 交易方式
	TradeWay string `json:"tradeWay,omitempty"`
	// Name 交易方式名称
	Name string `json:"name,omitempty"`
	// TradeType 开放平台下单时候传入的tradeType
	TradeType TradeType `json:"tradeType,omitempty"`
	// Description 交易描述
	Description string `json:"description,omitempty"`
	// OpSupport 是否支持
	OpSupport bool `json:"opSupport,omitempty"`
}

// CreateOrderPreview4CybMedia 溢价模式创建订单前预览数据接口
// 溢价模式创建订单前预览数据接口; 订单创建只允许购买同一个供应商的商品。本接口返回创建订单相关的优惠等信息。 1、校验商品数据是否允许订购。 2、校验代销关系 3、校验库存、起批量、是否满足混批条件
func CreateOrderPreview4CybMedia(client *go1688.Client, req *CreateOrderPreview4CybMediaRequest, accessToken string) (*CreateOrderPreview4CybMediaResponse, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp CreateOrderPreview4CybMediaResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
