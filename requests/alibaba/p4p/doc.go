// Package p4p 营销活动相关API
package p4p

// BizType 经营模式
type BizType = string

const (
	// BizType_FACTORY 生产加工
	BizType_FACTORY BizType = "1"
	// BizType_WHOLESALE 经销批发
	BizType_WHOLESALE BizType = "2"
	// BizType_AGENT 招商代理
	BizType_AGENT BizType = "3"
	// BizType_SERVICE 商业服务
	BizType_SERVICE BizType = "4"
)

// BuyerProtection 买家保障
type BuyerProtection = string

const (
	// BuyerProtection7d 7天包换
	BuyerProtection7d BuyerProtection = "qtbh"
	// BuyerProtection15d 15天包换
	BuyerProtection15d BuyerProtection = "swtbh"
)

// DeliveryTimeType 发货时间
type DeliveryTimeType = string

const (
	// DeliveryTimeType24h 24小时发货
	DeliveryTimeType24h DeliveryTimeType = "1"
	// DeliveryTimeType48h 48小时发货
	DeliveryTimeType48h DeliveryTimeType = "2"
	// DeliveryTimeType72h 72小时发货
	DeliveryTimeType72h DeliveryTimeType = "3"
)
