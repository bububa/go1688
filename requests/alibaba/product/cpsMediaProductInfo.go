package product

import (
	"github.com/bububa/go1688"
	"github.com/shopspring/decimal"
)

type CpsMediaProductInfoRequest struct {
	OfferId             uint64 `json:"offerId"`
	NeedCpsSuggestPrice bool   `json:"needCpsSuggestPrice"`
}

func (this *CpsMediaProductInfoRequest) Name() string {
	return "alibaba.cpsMedia.productInfo"
}

type CpsMediaProductInfoResponse struct {
	go1688.BaseResponse
	ProductInfo   *ProductInfo    `json:"productInfo,omitempty"`   // 商品详情
	BizGroupInfos []*BizGroupInfo `json:"bizGroupInfos,omitempty"` // 业务信息
}

type BizGroupInfo struct {
	Support     bool   `json:"support"`     // 是否支持,isConsignMarketOffer=ture表示代销offer
	Description string `json:"description"` // 垂直市场名字
	Code        string `json:"code"`        // 垂直市场标记, isConsignMarketOffer, 伙拼isTuanPiOffer
}

type ProductInfo struct {
	Id           uint64   `json:"productID,omitempty"`    // 商品ID
	CategoryId   uint64   `json:"categoryID,omitempty"`   // 类目ID，标识商品所属类目
	CategoryName string   `json:"categoryName,omitempty"` // 类目名
	GroupId      []uint64 `json:"groupID,omitempty"`      // 分组ID，确定商品所属分组。1688可传入多个分组ID，国际站同一个商品只能属于一个分组，因此默认只取第一个
	Subject      string   `json:"subject,omitempty"`      // 商品标题，最多128个字符
	Description  string   `json:"description,omitempty"`  // 商品详情描述，可包含图片中心的图片URL
	PictureAuth  bool     `json:"pictureAuth,omitempty"`  // 是否图片私密信息，国际站此字段无效
	Image        *struct {
		Images []string `json:"images,omitempty"` // 主图列表，使用相对路径，需要增加域名：https://cbu01.alicdn.com/
	} `json:"image,omitempty"` // 商品主图
	SkuInfos           []*SkuInfo         `json:"skuInfos,omitempty"`           // sku信息
	SaleInfo           *SaleInfo          `json:"saleInfo,omitempty"`           // 商品销售信息
	ShippingInfo       *ShippingInfo      `json:"shippingInfo,omitempty"`       // 商品物流信息
	QualityLevel       uint               `json:"qualityLevel,omitempty"`       // 质量星级(0-5)
	SupplierLoginId    string             `json:"supplierLoginId,omitempty"`    // 供应商loginId
	MainVedio          string             `json:"mainVedio,omitempty"`          // 主图视频播放地址
	ProductCargoNumber string             `json:"productCargoNumber,omitempty"` // 商品货号，产品属性中的货号
	ReferencePrice     string             `json:"referencePrice,omitempty"`     // 参考价格，返回价格区间，可能为空
	Attributes         []ProductAttribute `json:"attributes,omitempty"`         // 产品属性列表
	Status             ProductStatus      `json:"status,omitempty"`             // 商品状态。published:上网状态;member expired:会员撤销;auto expired:自然过期;expired:过期(包含手动过期与自动过期);member deleted:会员删除;modified:修改;new:新发;deleted:删除;TBD:to be delete;approved:审批通过;auditing:审核中;untread:审核不通过;
}

type ProductStatus = string

const (
	PRODUCT_PUBLISHED      ProductStatus = "published"      // 上网状态
	PRODUCT_MEMBER_EXPIRED ProductStatus = "member expired" // 会员撤销
	PRODUCT_AUTO_EXPIRED   ProductStatus = "auto expired"   // 自然过期
	PRODUCT_EXPIRED        ProductStatus = "expired"        // 过期(包含手动过期与自动过期)
	PRODUCT_MEMBER_DELETED ProductStatus = "member deleted" // 会员删除
	PRODUCT_MODIFIED       ProductStatus = "modified"       // 修改
	PRODUCT_NEW            ProductStatus = "new"            // 新发
	PRODUCT_DELETED        ProductStatus = "deleted"        // 删除
	PRODUCT_TBD            ProductStatus = "TBD"            // to be delete
	PRODUCT_APPROVED       ProductStatus = "approved"       // 审批通过
	PRODUCT_AUDITING       ProductStatus = "auditing"       // 审核中
	PRODUCT_UNTREAD        ProductStatus = "untread"        // 审核不通过
)

type ProductAttribute struct {
	Id       uint64 `json:"attributeID,omitempty"`   // 属性ID
	Value    string `json:"value,omitempty"`         // 值内容
	ValueId  uint64 `json:"valueID,omitempty"`       // 属性值ID
	Name     string `json:"attributeName,omitempty"` // 属性ID所对应的显示名，比如颜色，尺码
	IsCustom bool   `json:"isCustom,omitempty"`      // 是否为自定义属性，国际站无需关注
}

type SkuInfo struct {
	Id              uint64          `json:"skuId,omitempty"`           // skuId,该规格在所有商品中的唯一标记
	Attributes      []*SkuAttribute `json:"attributes,omitempty"`      // SKU属性值，可填多组信息
	CargoNumber     string          `json:"cargoNumber,omitempty"`     // 指定规格的货号
	AmountOnSale    uint            `json:"amountOnSale,omitempty"`    // 可销售数量
	RetailPrice     decimal.Decimal `json:"retailPrice,omitempty"`     // 建议零售价
	Price           decimal.Decimal `json:"price,omitempty"`           // 报价时该规格的单价
	SpecId          string          `json:"specId,omitempty"`          // specId,该规格在本商品内的唯一标记
	ConsignPrice    decimal.Decimal `json:"consignPrice,omitempty"`    // 分销基准价。代销场景均使用该价格。无SKU商品查看saleInfo中的consignPrice
	CpsSuggestPrice decimal.Decimal `json:"cpsSuggestPrice,omitempty"` // CPS建议价（单位：元）
}

type SaleInfo struct {
	SupportOnlineTrade bool `json:"supportOnlineTrade,omitempty"` // 是否支持网上交易。true：支持 false：不支持
	MixWholeSale       bool `json:"mixWholeSale,omitempty"`       // 是否支持混批
	PriceAuth          bool `json:"priceAuth,omitempty"`          // 是否价格私密信息
	PriceRanges        []struct {
		StartQuantity uint            `json:"startQuantity,omitempty"` // 起批量
		Price         decimal.Decimal `json:"price,omitempty"`         // 价格
	} `json:"priceRanges,omitempty"` // 区间价格。按数量范围设定的区间价格
	AmountOnSale     float64         `json:"amountOnSale,omitempty"`     // 可售数量
	Unit             string          `json:"unit,omitempty"`             // 计量单位
	MinOrderQuantity uint            `json:"minOrderQuantity,omitempty"` // 最小起订量，范围是1-99999。
	BatchNumber      uint            `json:"batchNumber,omitempty"`      // 每批数量，默认为空或者非零值，该属性不为空时sellunit为必填
	RetailPrice      decimal.Decimal `json:"retailPrice,omitempty"`      // 建议零售价
	SellUnit         string          `json:"sellunit,omitempty"`         // 售卖单位，如果为批量售卖，代表售卖的单位，该属性不为空时batchNumber为必填，例如1"手"=12“件"的"手"
	QuoteType        uint            `json:"quoteType,omitempty"`        // 0-无SKU按数量报价,1-有SKU按规格报价,2-有SKU按数量报价
	ConsignPrice     decimal.Decimal `json:"consignPrice,omitempty"`     // 分销基准价。代销场景均使用该价格。有SKU商品查看skuInfo中的consignPrice
	CpsSuggestPrice  decimal.Decimal `json:"cpsSuggestPrice,omitempty"`  // CPS建议价（单位：元）
}

type ShippingInfo struct {
	FreightTemplateId    uint64  `json:"freightTemplateID,omitempty"`    // 运费模板ID，0表示运费说明，1表示卖家承担运费，其他值表示使用运费模版。此参数可调用运费模板相关API获取
	UnitWeight           float64 `json:"unitWeight,omitempty"`           // 重量/毛重
	PackageSize          string  `json:"packageSize,omitempty"`          // 尺寸，单位是厘米，长宽高范围是1-9999999。1688无需关注此字段
	Volumn               uint    `json:"volumn,omitempty"`               // 体积，单位是立方厘米，范围是1-9999999，1688无需关注此字段
	HandlingTime         uint    `json:"handlingTime,omitempty"`         // 备货期，单位是天，范围是1-60。1688无需处理此字段
	SendGoodsAddressId   uint64  `json:"sendGoodsAddressId,omitempty"`   // 发货地址ID
	SendGoodsAddressText string  `json:"sendGoodsAddressText,omitempty"` // 发货地描述
	SuttleWeight         float64 `json:"suttleWeight,omitempty"`         // 净重
	Height               float64 `json:"height,omitempty"`               // 高度
	Width                float64 `json:"width,omitempty"`                // 宽度
	Length               float64 `json:"length,omitempty"`               // 长度
}

type SkuAttribute struct {
	Id          uint64 `json:"attrubuteID,omitempty"`          // sku属性ID
	Value       string `json:"attributeValue,omitempty"`       // sku值内容
	ImageUrl    string `json:"skuImageUrl,omitempty"`          // sku图片名
	DisplayName string `json:"attributeDisplayName,omitempty"` // sku属性ID所对应的显示名，比如颜色，尺码
	Name        string `json:"attributeName,omitempty"`        // sku属性ID所对应的显示名，比如颜色，尺码
}

func CpsMediaProductInfo(client *go1688.Client, req *CpsMediaProductInfoRequest, accessToken string) (*ProductInfo, []*BizGroupInfo, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &CpsMediaProductInfoResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, nil, err
	}
	if resp.IsError() {
		return nil, nil, resp
	}
	return resp.ProductInfo, resp.BizGroupInfos, nil
}
