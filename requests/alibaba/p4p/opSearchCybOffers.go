package p4p

import (
	"strconv"

	"github.com/shopspring/decimal"

	"github.com/bububa/go1688"
)

// OpSearchCybOffersRequest 商品列表搜索接口(替代alibaba.cps.listOverPricedOffer) API Request
type OpSearchCybOffersRequest struct {
	// BizType 枚举;经营模式;1:生产加工,2:经销批发,3:招商代理,4:商业服务
	BizType BizType `json:"bizType,omitempty"`
	// BuyerProtection 枚举;买家保障,多个值用逗号分割;qtbh:7天包换;swtbh:15天包换
	BuyerProtection string `json:"buyerProtection,omitempty"`
	// City 所在地区- 市
	City string `json:"city,omitempty"`
	// DeliveryTimeType 枚举;发货时间;1:24小时发货;2:48小时发货;3:72小时发货
	DeliveryTimeType DeliveryTimeType `json:"deliveryTimeType,omitempty"`
	// DescendOrder 是否倒序;正序: false;倒序:true
	DescendOrder bool `json:"descendOrder,omitempty"`
	// HolidayTagID 商品售卖类型筛选;枚举,多个值用分号分割;免费赊账:50000114
	HolidayTagID string `json:"holidayTagId,omitempty"`
	// KeyWords 搜索关键词
	KeyWords string `json:"keyWords,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面数量;最大20
	PageSize int `json:"pageSize,omitempty"`
	// PostCategoryID 类目id;4 纺织、皮革 5 电工电气 10 能源 12 交通运输 16 医药、保养 17 工艺品、礼品 57 电子元器件 58 照明工业 64 环保 66 医药、保养 67 办公、文教 69 商务服务 96 家纺家饰 311 童装 312 内衣 1813 玩具 2805 加工 2829 二手设备转让 10165 男装 1038378 鞋 1042954 箱包皮具 127380009 运动服饰 130822002 餐饮生鲜 130823000 性保健品 200514001 床上用品 201128501 直播 1 农业 2 食品酒水 7 数码、电脑 9 冶金矿产 15 日用百货 18 运动装备 33 汽摩及配件 53 传媒、广电 54 服饰配件、饰品 59 五金、工具 68 包装 70 安全、防护 96 家居饰品 97 美妆日化 97 美容护肤/彩妆 1501 母婴用品 10166 女装 10208 仪器仪表 122916001 宠物及园艺 123614001 钢铁 130822220 个护/家清 6 家用电器 8 化工 13 家装、建材 21 办公、文教 55 橡塑 65 机械及行业设备 71 汽摩及配件 72 印刷 73 项目合作 509 通信产品 1426 机床 1043472 毛巾、巾类 122916002 汽车用品
	PostCategoryID uint64 `json:"postCategoryId,omitempty"`
	// PriceStart 最低价
	PriceStart float64 `json:"priceStart,omitempty"`
	// PriceEnd 最高价
	PriceEnd float64 `json:"priceEnd,omitempty"`
	// PriceFilterFields 价格类型;默认分销价;agent_price:分销价;
	PriceFilterFields string `json:"priceFilterFields,omitempty"`
	// Province 所在地区- 省
	Province string `json:"province,omitempty"`
	// SortType 枚举;排序字段;normal:综合;
	SortType string `json:"sortType,omitempty"`
	// Tags 历史遗留，可不用
	Tags string `json:"tags,omitempty"`
	// OfferTags 枚举;1387842:渠道专享价商品
	OfferTags string `json:"offerTags,omitempty"`
	// OfferIDs 商品id搜索，多个id用逗号分割
	OfferIDs string `json:"offerIds,omitempty"`
}

// Name implement RequestData interface
func (r OpSearchCybOffersRequest) Name() string {
	return "alibaba.cps.op.searchCybOffers"
}

// Map implement RequestData interface
func (r OpSearchCybOffersRequest) Map() map[string]string {
	ret := make(map[string]string, 18)
	if r.BizType != "" {
		ret["bizType"] = r.BizType
	}
	if r.BuyerProtection != "" {
		ret["buyerProtection"] = r.BuyerProtection
	}
	if r.City != "" {
		ret["city"] = r.City
	}
	if r.DeliveryTimeType != "" {
		ret["deliveryTimeType"] = r.DeliveryTimeType
	}
	if r.DescendOrder {
		ret["descendOrder"] = "true"
	} else {
		ret["descendOrder"] = "false"
	}
	if r.HolidayTagID != "" {
		ret["holidayTagId"] = r.HolidayTagID
	}
	if r.KeyWords != "" {
		ret["keyWords"] = r.KeyWords
	}
	ret["page"] = strconv.Itoa(r.Page)
	ret["pageSize"] = strconv.Itoa(r.PageSize)
	if r.PostCategoryID > 0 {
		ret["postCategoryId"] = strconv.FormatUint(r.PostCategoryID, 10)
	}
	if r.PriceStart > 1e-15 {
		ret["priceStart"] = strconv.FormatFloat(r.PriceStart, 'f', 2, 64)
	}
	if r.PriceEnd > 1e-15 {
		ret["priceEnd"] = strconv.FormatFloat(r.PriceEnd, 'f', 2, 64)
	}
	if r.Province != "" {
		ret["province"] = r.Province
	}
	if r.PriceFilterFields != "" {
		ret["priceFilterFields"] = r.PriceFilterFields
	}
	if r.SortType != "" {
		ret["sortType"] = r.SortType
	}
	if r.Tags != "" {
		ret["tags"] = r.Tags
	}
	if r.OfferTags != "" {
		ret["offerTags"] = r.OfferTags
	}
	if r.OfferIDs != "" {
		ret["offerIds"] = r.OfferIDs
	}
	return ret
}

// OpSearchCybOffersResponse 商品列表搜索接口 API Response
type OpSearchCybOffersResponse struct {
	go1688.BaseResponse
	TotalCount int64                      `json:"totalCount,omitempty"`
	Result     []OverPricedCybSearchOffer `json:"result,omitempty"`
}

// OverPricedCybSearchOffer
type OverPricedCybSearchOffer struct {
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// ImgURL 商品首图
	ImgURL string `json:"imgUrl,omitempty"`
	// OfferID 商品id
	OfferID uint64 `json:"offerId,omitempty"`
	// SoldOut 销量
	SoldOut int64 `json:"soldOut,omitempty"`
	// SuperBuyerPrice 超买价，单位元
	SuperBuyerPrice decimal.Decimal `json:"superBuyerPrice,omitempty"`
	// Enable 是否有效
	Enable bool `json:"enable,omitempty"`
	// Profit 利润空间; - :表示无
	Profit string `json:"profit,omitempty"`
	// CurrentPrice 分销价
	CurrentPrice decimal.Decimal `json:"currentPrice,omitempty"`
	// OfferTags 标签数组
	OfferTags []string `json:"offerTags,omitempty"`
	// ChannelPrice 渠道专属价,可能无价格，需要调用方兼容处理
	ChannelPrice decimal.Decimal `json:"channelPrice,omitempty"`
}

// OpSearchCybOffers 商品列表搜索接口(替代alibaba.cps.listOverPricedOffer)
func OpSearchCybOffers(client *go1688.Client, req *OpSearchCybOffersRequest, accessToken string) (int64, []OverPricedCybSearchOffer, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp OpSearchCybOffersResponse
	if err := client.Do(finalRequest, accessToken, resp); err != nil {
		return 0, nil, err
	}
	return resp.TotalCount, resp.Result, nil
}
