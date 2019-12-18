package message

type ProductMessage struct {
	ProductIds          string                     `json:"productIds,omitempty"`               // 商品ID集合，至少有一个，用逗号分割
	MemberId            string                     `json:"memberId,omitempty"`                 // 1688会员ID
	Status              string                     `json:"status,omitempty"`                   // 消息类型，具体可为RELATION_VIEW_PRODUCT_EXPIRE、RELATION_VIEW_PRODUCT_NEW_OR_MODIFY、RELATION_VIEW_PRODUCT_DELETE、RELATION_VIEW_PRODUCT_REPOST
	MsgSendTime         string                     `json:"msgSendTime,omitempty"`              // 消息发送时间
	InventoryChangeList []OfferInventoryChangeList `json:"OfferInventoryChangeList,omitempty"` // 库存变更列表
	Products            []Product                  `json:"products,omitempty"`
}

type OfferInventoryChangeList struct {
	OfferId     uint64 `json:"offerId"`             // 1688商品id
	OfferOnSale uint   `json:"offerOnSale"`         // 在线可售offer数量
	SkuId       uint64 `json:"skuId,omitempty"`     // 商品skuId
	SkuOnSale   uint   `json:"skuOnSale,omitempty"` // 在线可售sku数量
	Quantity    int    `json:"quantity,omitempty"`  // 该offer整体库存变化数
	bizTime     string `json:"bizTime,omitempty"`   // 库存变更时间
}

type Product struct {
	ProductId         string `json:"productId"`         // 商品id
	SupportSuperPrice bool   `json:"supportSuperPrice"` // 是否支持超买价;true支持超买价
	Status            string `json:"status"`            // 商品状态;枚举,DELETE:删除商品,商品不在商品池;UPDATE:商品超买价状态可能有更新;机构可以只关心DELETE状态判断是否在商品池
}

func (this *ProductMessage) Types() []MessageType {
	return []MessageType{
		PRODUCT_RELATION_VIEW_PRODUCT_DELETE,
		PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE,
		PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY,
		PRODUCT_RELATION_VIEW_PRODUCT_REPOST,
		PRODUCT_PRODUCT_INVENTORY_CHANGE,
		PRODUCT_RELATION_VIEW_EXIT_SUPERBUYER,
		PRODUCT_RELATION_VIEW_PRODUCT_AUDIT,
	}
}
