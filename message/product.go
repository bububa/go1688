package message

// ProductMessage 商品消息
type ProductMessage struct {
	// ProductIDs 商品ID集合，至少有一个，用逗号分割
	ProductIDs string `json:"productIds,omitempty"`
	// MemberID 1688会员ID
	MemberID string `json:"memberId,omitempty"`
	// Status 消息类型，具体可为RELATION_VIEW_PRODUCT_EXPIRE、RELATION_VIEW_PRODUCT_NEW_OR_MODIFY、RELATION_VIEW_PRODUCT_DELETE、RELATION_VIEW_PRODUCT_REPOST
	Status ProductMessageStatus `json:"status,omitempty"`
	// MsgSendTime 消息发送时间
	MsgSendTime string `json:"msgSendTime,omitempty"`
	// InventoryChangeList 库存变更列表
	InventoryChangeList []OfferInventoryChangeList `json:"OfferInventoryChangeList,omitempty"`
	// Products 商品列表
	Products []Product `json:"products,omitempty"`
}

// OfferInventoryChangeList 库存变更列表
type OfferInventoryChangeList struct {
	// OfferID 1688商品id
	OfferID uint64 `json:"offerId"`
	// OfferOnSale 在线可售offer数量
	OfferOnSale int64 `json:"offerOnSale"`
	// SkuID 商品skuId
	SkuID uint64 `json:"skuId,omitempty"`
	// SkuOnSale 在线可售sku数量
	SkuOnSale int64 `json:"skuOnSale,omitempty"`
	// Quantity 该offer整体库存变化数
	Quantity int64 `json:"quantity,omitempty"`
	// BizTime 库存变更时间
	BizTime string `json:"bizTime,omitempty"`
}

// Product 商品
type Product struct {
	// ProductID 商品id
	ProductID string `json:"productId"`
	// SupportSuperPrice 是否支持超买价;true支持超买价
	SupportSuperPrice bool `json:"supportSuperPrice"`
	// Status 商品状态;枚举,DELETE:删除商品,商品不在商品池;UPDATE:商品超买价状态可能有更新;机构可以只关心DELETE状态判断是否在商品池
	Status ProductStatus `json:"status"`
}

// Types implement Message interface
func (m ProductMessage) Types() []MessageType {
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
