package p4p

const (
	NAMESPACE = "com.alibaba.p4p"
)

// ActionType 行为类型
type ActionType = string

const (
	// ActionType_CLICK CLICK(点击商品详情)
	ActionType_CLICK ActionType = "CLICK"
	// ActionType_COMMENT COMMENT(评论)
	ActionType_COMMENT ActionType = "COMMENT"
	// ActionTYpe_STORE STORE(收藏)
	ActionType_STORE ActionType = "STORE"
	// ActionType_CART CART(加购物车)
	ActionType_CART ActionType = "CART"
	// ActionType_SEARCH SEARCH(关键词搜索，配合actionDetail)
	ActionType_SEARCH ActionType = "SEARCH"
	// ActionType_VIEW 曝光(VIEW)
	ActionType_VIEW ActionType = "VIEW"
)

// FeedType 商品类型
type FeedType = int

const (
	// FeedType_1688 1688商品
	FeedType_1688 FeedType = 1
	// FeedType_MEDIA 机构商品
	FeedType_MEDIA FeedType = 2
)
