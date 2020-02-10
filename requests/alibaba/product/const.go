package product

const (
	NAMESPACE = "com.alibaba.product"
)

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
