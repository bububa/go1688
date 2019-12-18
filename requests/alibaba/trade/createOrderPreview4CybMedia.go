package trade

import (
	"encoding/json"

	"github.com/bububa/go1688"
	"github.com/shopspring/decimal"
)

type CreateOrderPreview4CybMediaRequest struct {
	Address   *Address `json:"address"`
	CargoList []*Cargo `json:"cargo"`
}

type CreateOrderPreview4CybMediaRefinedRequest struct {
	Address   string `json:"addressParam"`   // 收货地址信息
	CargoList string `json:"cargoParamList"` // 商品信息
}

func (this *CreateOrderPreview4CybMediaRefinedRequest) Name() string {
	return "alibaba.createOrder.preview4CybMedia"
}

type Address struct {
	Id           uint64 `json:"addressId,omitempty"`    // 收货地址id
	FullName     string `json:"fullName,omitempty"`     // 收货人姓名
	Mobile       string `json:"mobile,omitempty"`       // 手机
	Phone        string `json:"phone,omitempty"`        // 电话
	PostCode     string `json:"postCode,omitempty"`     // 邮编
	CityText     string `json:"cityText,omitempty"`     // 市文本
	ProvinceText string `json:"provinceText,omitempty"` // 省份文本
	AreaText     string `json:"areaText,omitempty"`     // 区文本
	TownText     string `json:"townText,omitempty"`     // 镇文本
	Address      string `json:"address,omitempty"`      // 街道地址
	DistrictCode string `json:"districtCode,omitempty"` // 地址编码
}

type Cargo struct {
	OfferId  uint64  `json:"offerId,omitempty"`  // 商品对应的offer id
	SpecId   string  `json:"specId,omitempty"`   // 商品规格id, 无sku商品不传
	Quantify float64 `json:"quantity,omitempty"` // 商品数量(计算金额用)
}

func (this *CreateOrderPreview4CybMediaRequest) Refine() *CreateOrderPreview4CybMediaRefinedRequest {
	address, _ := json.Marshal(this.Address)
	cargoList, _ := json.Marshal(this.CargoList)
	return &CreateOrderPreview4CybMediaRefinedRequest{
		Address:   string(address),
		CargoList: string(cargoList),
	}
}

type CreateOrderPreview4CybMediaResponse struct {
	go1688.BaseResponse
	PostFeeByDescOfferList []uint64        `json:"postFeeByDescOfferList,omitempty"` // 运费说明的商品列表
	ConsignOfferList       []uint64        `json:"consignOfferList,omitempty"`       // 代销商品列表
	OrderPreview           []*OrderPreview `json:"orderPreviewResuslt,omitempty"`    // 订单预览结果，过自动拆单会返回多个记录
}

type OrderPreview struct {
	DiscountFee              uint        `json:"discountFee,omitempty"`              // 计算完货品金额后再次进行的减免金额. 单位: 分
	TradeModeNameList        []string    `jsong:"tradeModeNameList,omitempty"`       // 当前交易可以支持的交易方式列表。某些场景的创建订单接口需要使用。
	Status                   bool        `json:"status,omitempty"`                   // 状态
	TaoSampleSinglePromotion bool        `json:"taoSampleSinglePromotion,omitempty"` // 是否有淘货源单品优惠 false:有单品优惠 true：没有单品优惠
	SumPayment               uint        `json:"sumPayment,omitempty"`               // 订单总费用, 单位为分.
	Message                  string      `json:"message,omitempty"`                  // 返回信息
	SumCarriage              uint        `json:"sumCarriage,omitempty"`              // 总运费信息, 单位为分.
	ResultCode               string      `json:"resultCode,omitempty"`               // 返回码
	SumPaymentNoCarriage     uint        `json:"sumPaymentNoCarriage,omitempty"`     // 不包含运费的货品总费用, 单位为分.
	AdditionalFee            uint        `json:"additionalFee,omitempty"`            // 附加费,单位，分
	FlowFlag                 string      `json:"flowFlag,omitempty"`                 // 订单下单流程
	CargoList                []CargoInfo `json:"cargoList,omitempty"`                // 规格信息
	ShopPromotionList        []Promotion `json:"shopPromotionList,omitempty"`        // 可用店铺级别优惠列表
	TradeModeList            []TradeMode `json:"tradeModeList,omitempty"`            // 当前交易可以支持的交易方式列表。结果可以参照1688下单预览页面的交易方式。
}

type CargoInfo struct {
	OfferId        uint64          `json:"offerId,omitempty"`            // 商品ID
	Amount         decimal.Decimal `json:"amount,omitempty"`             // 产品总金额
	Message        string          `json:"message,omitempty"`            // 返回信息
	FinalUnitPrice decimal.Decimal `json:"finalUnitPrice,omitempty"`     // 最终单价
	SpecId         string          `json:"specId,omitempty"`             // 规格ID，offer内唯一
	SkuId          uint64          `json:"skuId,omitempty"`              // 规格ID，全局唯一
	ResultCode     string          `json:"resultCode,omitempty"`         // 返回码
	PromotionList  []Promotion     `json:"cargoPromotionList,omitempty"` // 商品优惠列表
}

type Promotion struct {
	Id          string `json:"promotionId,omitempty"` // 优惠券ID
	Selected    bool   `json:"selected,omitempty"`    // 是否默认选中
	Text        string `json:"text,omitempty"`        // 优惠券名称
	Desc        string `json:"desc,omitempty"`        // 优惠券描述
	FreePostage bool   `json:"freePostage,omitempty"` // 是否免邮
	DiscountFee uint   `json:"discountFee,omitempty"` // 减去金额，单位为分
}

type TradeMode struct {
	TradeWay    string    `json:"tradeWay,omitempty"`    // 交易方式
	Name        string    `json:"name,omitempty"`        // 交易方式名称
	TradeType   TradeType `json:"tradeType,omitempty"`   // 开放平台下单时候传入的tradeType
	Description string    `json:"description,omitempty"` // 交易描述
	OpSupport   bool      `json:"opSupport,omitempty"`   // 是否支持
}

func CreateOrderPreview4CybMedia(client *go1688.Client, req *CreateOrderPreview4CybMediaRequest, accessToken string) ([]*OrderPreview, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &CreateOrderPreview4CybMediaResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.OrderPreview, nil
}
