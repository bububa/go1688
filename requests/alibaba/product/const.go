package product

const (
	NAMESPACE = "com.alibaba.product"
)

// ProductStatus 商品状态。
type ProductStatus = string

const (
	// PRODUCT_PUBLISHED 上网状态;
	PRODUCT_PUBLISHED ProductStatus = "published"
	// PRODUCT_MEMBER_EXPIRED 会员撤销
	PRODUCT_MEMBER_EXPIRED ProductStatus = "member expired"
	// PRODUCT_AUTO_EXPIRED 自然过期
	PRODUCT_AUTO_EXPIRED ProductStatus = "auto expired"
	// PRODUCT_EXPIRED 过期(包含手动过期与自动过期)
	PRODUCT_EXPIRED ProductStatus = "expired"
	// PRODUCT_MEMBER_DELETED 会员删除
	PRODUCT_MEMBER_DELETED ProductStatus = "member deleted"
	// PRODUCT_MODIFIED 修改
	PRODUCT_MODIFIED ProductStatus = "modified"
	// PRODUCT_NEW 新发
	PRODUCT_NEW ProductStatus = "new"
	// PRODUCT_DELETED 删除
	PRODUCT_DELETED ProductStatus = "deleted"
	// PRODUCT_TBD to be delete
	PRODUCT_TBD ProductStatus = "TBD"
	// PRODUCT_APPROVED 审批通过
	PRODUCT_APPROVED ProductStatus = "approved"
	// PRODUCT_AUDITING 审核中
	PRODUCT_AUDITING ProductStatus = "auditing"
	// PRODUCT_UNTREAD 审核不通过
	PRODUCT_UNTREAD ProductStatus = "untread"
)

// QuoteType 0-无SKU按数量报价,1-有SKU按规格报价,2-有SKU按数量报价
type QuoteType = int

const (
	QUOTE_TYPE_NONE      QuoteType = 0
	QUOTE_TYPE_BY_SKU    QuoteType = 1
	QUOTE_TYPE_BY_AMOUNT QuoteType = 2
)
