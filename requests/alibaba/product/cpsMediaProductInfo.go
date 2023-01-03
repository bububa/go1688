package product

import (
	"strconv"

	"github.com/shopspring/decimal"

	"github.com/bububa/go1688"
)

// CpsMediaProductInfoRequest 获取商品详情接口 API Request
type CpsMediaProductInfoRequest struct {
	// OfferID 1688商品ID，等同于productId
	OfferID uint64 `json:"offerId,omitempty"`
	// NeedCpsSuggestPrice 是否需要CPS建议价
	NeedCpsSuggestPrice bool `json:"needCpsSuggestPrice,omitempty"`
	//  NeedIntelligentInfo 是否返回算法改写的信息，包括标题、图片和详情图片
	NeedIntelligentInfo bool `json:"needIntelligentInfo,omitempty"`
}

// Name implement RequestData Interface
func (r CpsMediaProductInfoRequest) Name() string {
	return "alibaba.cpsMedia.productInfo"
}

// Map implement RequestData Interface
func (r CpsMediaProductInfoRequest) Map() map[string]string {
	ret := make(map[string]string, 2)
	ret["offerId"] = strconv.FormatUint(r.OfferID, 10)
	if r.NeedCpsSuggestPrice {
		ret["needCpsSuggestPrice"] = "true"
	}
	if r.NeedIntelligentInfo {
		ret["needIntelligentInfo"] = "true"
	}
	return ret
}

// CpsMediaProductInfoResponse 获取商品详情接口 API Response
type CpsMediaProductInfoResponse struct {
	go1688.BaseResponse
	// ProductInfo 商品详情
	ProductInfo *ProductInfo `json:"productInfo,omitempty"`
	// BizGroupInfos 业务信息
	BizGroupInfos []*BizGroupInfo `json:"bizGroupInfos,omitempty"`
}

// BizGroupInfo 业务信息
type BizGroupInfo struct {
	// Support 是否支持,isConsignMarketOffer=ture表示代销offer
	Support bool `json:"support,omitempty"`
	// Description 垂直市场名字
	Description string `json:"description,omitempty"`
	// Code 垂直市场标记, isConsignMarketOffer, 伙拼isTuanPiOffer
	Code string `json:"code,omitempty"`
}

// ProductInfo 商品详情
type ProductInfo struct {
	// ID 商品ID
	ID uint64 `json:"productID,omitempty"`
	// CategoryID 类目ID，标识商品所属类目
	CategoryID uint64 `json:"categoryID,omitempty"`
	// CategoryName 类目名
	CategoryName string `json:"categoryName,omitempty"`
	// GroupID 分组ID，确定商品所属分组。1688可传入多个分组ID，国际站同一个商品只能属于一个分组，因此默认只取第一个
	GroupID []uint64 `json:"groupID,omitempty"`
	// Subject 商品标题，最多128个字符
	Subject string `json:"subject,omitempty"`
	// Description 商品详情描述，可包含图片中心的图片URL
	Description string `json:"description,omitempty"`
	// PictureAuth 是否图片私密信息，国际站此字段无效
	PictureAuth bool `json:"pictureAuth,omitempty"`
	// Image 商品主图
	Image *struct {
		// Images 主图列表，使用相对路径，需要增加域名：https://cbu01.alicdn.com/
		Images []string `json:"images,omitempty"`
	} `json:"image,omitempty"`
	// SkuInfos sku信息
	SkuInfos []*SkuInfo `json:"skuInfos,omitempty"`
	// SaleInfo 商品销售信息
	SaleInfo *SaleInfo `json:"saleInfo,omitempty"`
	// ShippingInfo 商品物流信息
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// QualityLevel 质量星级(0-5)
	QualityLevel int `json:"qualityLevel,omitempty"`
	// SupplierLoginID 供应商loginId
	SupplierLoginID string `json:"supplierLoginId,omitempty"`
	// MainVideo 主图视频播放地址
	MainVedio string `json:"mainVedio,omitempty"`
	// ProductCargoNumber 商品货号，产品属性中的货号
	ProductCargoNumber string `json:"productCargoNumber,omitempty"`
	// ReferencePrice 参考价格，返回价格区间，可能为空
	ReferencePrice string `json:"referencePrice,omitempty"`
	// Attributes 产品属性列表
	Attributes []ProductAttribute `json:"attributes,omitempty"`
	// Status 商品状态。
	// - published: 上网状态;
	// - member expired: 会员撤销;
	// - auto expired: 自然过期;
	// - expired: 过期(包含手动过期与自动过期);
	// - member deleted: 会员删除;
	// - modified: 修改;
	// - new: 新发;
	// - deleted: 删除;
	// - TBD: to be delete;
	// - approved: 审批通过;
	// - auditing: 审核中;
	// - untread: 审核不通过;
	Status ProductStatus `json:"status,omitempty"`
}

// ProductAttribute 产品属性
type ProductAttribute struct {
	// ID 属性ID
	ID uint64 `json:"attributeID,omitempty"`
	// Value 值内容
	Value string `json:"value,omitempty"`
	// ValueID 属性值ID
	ValueID uint64 `json:"valueID,omitempty"`
	// Name 属性ID所对应的显示名，比如颜色，尺码
	Name string `json:"attributeName,omitempty"`
	// IsCustom 是否为自定义属性，国际站无需关注
	IsCustom bool `json:"isCustom,omitempty"`
}

// SkuInfo sku信息
type SkuInfo struct {
	// ID 该规格在所有商品中的唯一标记
	ID uint64 `json:"skuId,omitempty"`
	// Attributes SKU属性值，可填多组信息
	Attributes []*SkuAttribute `json:"attributes,omitempty"`
	// CargoNumber 指定规格的货号
	CargoNumber string `json:"cargoNumber,omitempty"`
	// AmountOnSale 可销售数量
	AmountOnSale int64 `json:"amountOnSale,omitempty"`
	// RetailPrice 建议零售价
	RetailPrice decimal.Decimal `json:"retailPrice,omitempty"`
	// Price 报价时该规格的单价
	Price decimal.Decimal `json:"price,omitempty"`
	// SpecID,该规格在本商品内的唯一标记
	SpecID string `json:"specId,omitempty"`
	// ConsignPrice 分销基准价。代销场景均使用该价格。无SKU商品查看saleInfo中的consignPrice
	ConsignPrice decimal.Decimal `json:"consignPrice,omitempty"`
	// CpsSuggestPrice CPS建议价（单位：元）
	CpsSuggestPrice decimal.Decimal `json:"cpsSuggestPrice,omitempty"`
}

// SaleInfo 商品销售信息
type SaleInfo struct {
	// SupportOnlineTrade 是否支持网上交易。true：支持 false：不支持
	SupportOnlineTrade bool `json:"supportOnlineTrade,omitempty"`
	// MixWholeSale 是否支持混批
	MixWholeSale bool `json:"mixWholeSale,omitempty"`
	// PriceAuth 是否价格私密信息
	PriceAuth bool `json:"priceAuth,omitempty"`
	// PriceRanges 区间价格。按数量范围设定的区间价格
	PriceRanges []ProductPriceRange `json:"priceRanges,omitempty"`
	// AmountOnSale 可售数量
	AmountOnSale float64 `json:"amountOnSale,omitempty"`
	// Unit 计量单位
	Unit string `json:"unit,omitempty"`
	// MinOrderQuantity 最小起订量，范围是1-99999。
	MinOrderQuantity int64 `json:"minOrderQuantity,omitempty"`
	// BatchNumber 每批数量，默认为空或者非零值，该属性不为空时sellunit为必填
	BatchNumber int64 `json:"batchNumber,omitempty"`
	// RetailPrice 建议零售价
	RetailPrice decimal.Decimal `json:"retailprice,omitempty"`
	// SellUnit 售卖单位，如果为批量售卖，代表售卖的单位，该属性不为空时batchNumber为必填，例如1"手"=12“件"的"手"
	SellUnit string `json:"sellunit,omitempty"`
	// QuoteType 0-无SKU按数量报价,1-有SKU按规格报价,2-有SKU按数量报价
	QuoteType QuoteType `json:"quoteType,omitempty"`
	// ConsignPrice 分销基准价。代销场景均使用该价格。有SKU商品查看skuInfo中的consignPrice
	ConsignPrice decimal.Decimal `json:"consignPrice,omitempty"`
	// CpsSuggestPrice CPS建议价（单位：元）
	CpsSuggestPrice decimal.Decimal `json:"cpsSuggestPrice,omitempty"`
}

// ProductPriceRange 区间价格。按数量范围设定的区间价格
type ProductPriceRange struct {
	// StartQuantity 起批量
	StartQuantity int64 `json:"startQuantity,omitempty"`
	// Price 价格
	Price decimal.Decimal `json:"price,omitempty"`
}

// ShippingInfo 商品物流信息
type ShippingInfo struct {
	// FreightTemplateID 运费模板ID，0表示运费说明，1表示卖家承担运费，其他值表示使用运费模版。此参数可调用运费模板相关API获取
	FreightTemplateID uint64 `json:"freightTemplateID,omitempty"`
	// UnitWeight 重量/毛重
	UnitWeight float64 `json:"unitWeight,omitempty"`
	// PackageSize 尺寸，单位是厘米，长宽高范围是1-9999999。1688无需关注此字段
	PackageSize string `json:"packageSize,omitempty"`
	// Volumn 体积，单位是立方厘米，范围是1-9999999，1688无需关注此字段
	Volumn int64 `json:"volumn,omitempty"`
	// HandlingTime 备货期，单位是天，范围是1-60。1688无需处理此字段
	HandlingTime int `json:"handlingTime,omitempty"`
	// SendGoodsAddressID 发货地址ID
	SendGoodsAddressID uint64 `json:"sendGoodsAddressId,omitempty"`
	// SendGoodsAddressText 发货地描述
	SendGoodsAddressText string `json:"sendGoodsAddressText,omitempty"`
	// ShuttleWeight 净重
	SuttleWeight float64 `json:"suttleWeight,omitempty"`
	// Height 高度
	Height float64 `json:"height,omitempty"`
	// Width 宽度
	Width float64 `json:"width,omitempty"`
	// Length 长度
	Length float64 `json:"length,omitempty"`
}

// SkuAttribute 商品属性和属性值
type SkuAttribute struct {
	// ID 属性ID
	ID uint64 `json:"attributeID,omitempty"`
	// Value 属性值内容
	Value string `json:"attributeValue,omitempty"`
	// ImageURL 图片名
	ImageURL string `json:"skuImageUrl,omitempty"`
	// DisplayName sku属性ID所对应的显示名，比如颜色，尺码
	DisplayName string `json:"attributeDisplayName,omitempty"`
	// Name sku属性ID所对应的显示名，比如颜色，尺码
	Name string `json:"attributeName,omitempty"`
}

// CpsMediaProductInfo 获取商品详情接口
func CpsMediaProductInfo(client *go1688.Client, req *CpsMediaProductInfoRequest, accessToken string) (*ProductInfo, []*BizGroupInfo, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp CpsMediaProductInfoResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, nil, err
	}
	return resp.ProductInfo, resp.BizGroupInfos, nil
}
