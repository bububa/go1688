package trade

import (
	"github.com/bububa/go1688"
)

// CreateOrder4CybMediaRequest 溢价模式订单创建接口 API Request
type CreateOrder4CybMediaRequest struct {
	// Address 收货地址
	Address *Address `json:"addressParam,omitempty"`
	// CargoList 购买的商品信息
	CargoList []Cargo `json:"cargoParamList,omitempty"`
	// Message 买家留言
	Message string `json:"message,omitempty"`
	// OuterOrderInfo 机构的订单信息，json格式，如果校验不通过，不能创建订单。业务线原因机构订单信息不能放在上面的下单参数结构体中，只能新增该字段用于机构订单信息回写，没有幂等校验，纯记录统计分析。mediaOrderId:机构订单号;phone:电话;offers.id:Long,1688商品id;offers.specId:String,1688商品specId(可能无);offers.price:Long,媒体溢价单价(单位分);offers.num:Long,售卖数量
	OuterOrderInfo *OuterOrderInfo `json:"outerOrderInfo,omitempty"`
	// TradeType 由于不同的商品支持的交易方式不同，没有一种交易方式是全局通用的，所以当前下单可使用的交易方式必须通过下单预览接口的tradeModeNameList获取。交易方式类型说明：fxassure（交易4.0通用担保交易），alipay（大市场通用的支付宝担保交易（目前在做切流，后续会下掉）），period（普通账期交易）, assure（大买家企业采购询报价下单时需要使用的担保交易流程）, creditBuy（诚E赊），bank（银行转账），631staged（631分阶段付款），37staged（37分阶段）；此字段不传则系统默认会选取一个可用的交易方式下单，如果开通了诚E赊默认是creditBuy（诚E赊），未开通诚E赊默认使用的方式是支付宝担宝交易。
	TradeType TradeType `json:"tradeType,omitempty"`
	// UseChannelPrice 是否走渠道专属价，传true优先使用渠道传享价，不传或者传false则不走渠道专享价，走普通的分销价
	UseChannelPrice bool `json:"useChannelPrice,omitempty"`
	// Flow 不传，价格逻辑由useChannelPrice控制；传入general使用批发价，即useChannelPrice的价格逻辑失效；传入paired使用火拼价，若该商品未参与伙拼，则下单失败。
	Flow TradeFlow `json:"flow,omitempty"`
	// ShopPromotionID .
	ShopPromotionID string `json:"shopPromotionId,omitempty"`
}

// Name implement RequestData interface
func (r CreateOrder4CybMediaRequest) Name() string {
	return "alibaba.trade.createOrder4CybMedia"
}

// Map implement RequestData interface
func (r CreateOrder4CybMediaRequest) Map() map[string]string {
	ret := make(map[string]string, 5)
	if r.Address != nil {
		ret["addressParam"] = go1688.JSONMarshal(r.Address)
	}
	if len(r.CargoList) > 0 {
		ret["cargoParamList"] = go1688.JSONMarshal(r.CargoList)
	}
	if r.Message != "" {
		ret["message"] = r.Message
	}
	if r.OuterOrderInfo != nil {
		ret["outerOrderInfo"] = go1688.JSONMarshal(r.OuterOrderInfo)
	}
	if r.TradeType != "" {
		ret["tradeType"] = r.TradeType
	}
	if r.UseChannelPrice {
		ret["useChannelPrice"] = "true"
	}
	if r.Flow != "" {
		ret["flow"] = r.Flow
	}
	if r.ShopPromotionID != "" {
		ret["shopPromotionId"] = r.ShopPromotionID
	}
	return ret
}

// OuterOrderInfo 机构的订单信息
type OuterOrderInfo struct {
	// MediaOrderID 机构订单号
	MediaOrderID uint64 `json:"mediaOrderId,omitempty"`
	// Phone 电话
	Phone string `json:"phone,omitempty"`
	// Offers
	Offers []OuterOrderOffer `json:"offers,omitempty"`
}

// OuterOrderOffer
type OuterOrderOffer struct {
	// ID 1688商品id
	ID uint64 `json:"id,omitempty"`
	// SpecID 1688商品specId(可能无)
	SpecID string `json:"specId,omitempty"`
	// Price 媒体溢价单价(单位分)
	Price int64 `json:"price,omitempty"`
	// Num 售卖数量
	Num int64 `json:"num,omitempty"`
}

// CreateOrder4CybMediaResponse 溢价模式订单创建接口 API Response
type CreateOrder4CybMediaResponse struct {
	go1688.BaseResponse
	// Result 返回结果
	Result *CreateOrder4CybMediaResult `json:"result,omitempty"`
}

// CreateOrder4CybMediaResult 溢价模式订单创建接口 API Result
type CreateOrder4CybMediaResult struct {
	// TotalSuccessAmount 下单成功的订单总金额，单位：分
	TotalSuccessAmount int64 `json:"totalSuccessAmount,omitempty"`
	// OrderID 下单成功后的订单id
	OrderID string `jsong:"orderId,omitempty"`
	// PostFee 原始运费，单位：分。注意：下单后卖家可能调整，因此该值可能不等于最终支付运费
	PostFee int64 `json:"postFee,omitempty"`
}

// CreateOrder4CybMedia 溢价模式订单创建接口
func CreateOrder4CybMedia(client *go1688.Client, req *CreateOrder4CybMediaRequest, accessToken string) (*CreateOrder4CybMediaResult, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp CreateOrder4CybMediaResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
