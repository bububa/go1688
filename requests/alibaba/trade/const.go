package trade

const (
	NAMESPACE = "com.alibaba.trade"
)

// APIWebsite 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
type APIWebsite = string

const (
	// API_ALIBABA alibaba
	API_ALIBABA APIWebsite = "alibaba"
	// API_1688 1688
	API_1688 APIWebsite = "1688"
)

// TradeType 交易方式类型
type TradeType = string

const (
	// FXASSURE_TRADE 交易4.0通用担保交易
	FXASSURE_TRADE TradeType = "fxassure"
	// ALIPAY_TRADE 大市场通用的支付宝担保交易（目前在做切流，后续会下掉）
	ALIPAY_TRADE TradeType = "alipay"
	// PERIOD_TRADE 普通账期交易
	PERIOD_TRADE TradeType = "period"
	// ASSURE_TRADE 大买家企业采购询报价下单时需要使用的担保交易流程
	ASSURE_TRADE TradeType = "assure"
	// CREDIT_BUY_TRADE 诚E赊
	CREDIT_BUY_TRADE TradeType = "creditBuy"
	// BANK_TRADE 银行转账
	BANK_TRADE TradeType = "bank"
	// STAGED631_TRADE 631分阶段付款
	STAGED631_TRADE TradeType = "631staged"
	// STAGED37_TRADE 37分阶段
	STAGED37_TRADE TradeType = "37staged"
)

// TradeFlow 价格逻辑由useChannelPrice控制；传入general使用批发价，即useChannelPrice的价格逻辑失效；传入paired使用火拼价，若该商品未参与伙拼，则下单失败。
type TradeFlow = string

const (
	// TRADE_FLOW_GENERAL 使用批发价
	TRADE_FLOW_GENERAL TradeFlow = "general"
	//使用火拼价，若该商品未参与伙拼，则下单失败。
	TRADE_FLOW_PAIR TradeFlow = "pair"
)

// CancelReason 原因描述
type CancelReason = string

const (
	// BUYER_CANCEL 买家取消订单
	BUYER_CANCEL CancelReason = "buyerCancel"
	// SELLER_GOODS_LACK 卖家库存不足
	SELLER_GOODS_LACK CancelReason = "sellerGoodsLack"
	// OTHER_REASON 其它
	OTHER_REASON CancelReason = "other"
)

// TradeCloseType 关闭订单操作类型
type TradeCloseType = string

const (
	// CLOSE_TRADE_BY_SELLER 卖家关闭交易
	CLOSE_TRADE_BY_SELLER TradeCloseType = "CLOSE_TRADE_BY_SELLER"
	// CLOSE_TRADE_BY_BOPS 后台关闭交
	CLOSE_TRADE_BY_BOPS TradeCloseType = "CLOSE_TRADE_BY_BOPS"
	// CLOSE_TRADE_BY_SYSTEM 系统（超时）关闭交易
	CLOSE_TRADE_BY_SYSTEM TradeCloseType = "CLOSE_TRADE_BY_SYSTEM"
	// CLOSE_TRADE_BY_BUYER 买家关闭交易
	CLOSE_TRADE_BY_BUYER TradeCloseType = "CLOSE_TRADE_BY_BUYER"
	// CLOSE_TRADE_BY_CREADIT 诚信保障投诉关闭
	CLOSE_TRADE_BY_CREADIT TradeCloseType = "CLOSE_TRADE_BY_CREADT"
)

type TradeStatus = string

const (
	// TRADE_WAIT_BUYER_PAY 等待买家付款
	TRADE_WAIT_BUYER_PAY TradeStatus = "waitbuyerpay"
	// TRADE_WAIT_SELLER_SEND 等待卖家发货
	TRADE_WAIT_SELLER_SEND TradeStatus = "waitsellersend"
	// TRADE_WAIT_BUYER_RECEIVE 等待买家收货
	TRADE_WAIT_BUYER_RECEIVE TradeStatus = "waitbuyerreceive"
	// TRADE_CONFIRM_GOODS 已收货
	TRADE_CONFIRM_GOODS TradeStatus = "confirm_goods"
	// TRADE_SUCCESS 交易成功
	TRADE_SUCCESS TradeStatus = "success"
	// TRADE_CANCEL 交易取消
	TRADE_CANCEL TradeStatus = "cancel"
	// TRADE_TERMINATED 交易终止
	TRADE_TERMINATED TradeStatus = "terminated"
	// TRADE_UNKNOWN 其它
	TRADE_UNKNOWN TradeStatus = ""
)

// RefundStatus 退款状态
type RefundStatus = string

const (
	// REFUND_WAIT_SELLER_AGREE 等待卖家同意
	REFUND_WAIT_SELLER_AGREE RefundStatus = "waitselleragree"
	// REFUND_WAIT_BUYER_MODIFY 待买家修改
	REFUND_WAIT_BUYER_MODIFY RefundStatus = "waitbuyermodify"
	// REFUND_WAIT_BUYER_SEND 等待买家退货
	REFUND_WAIT_BUYER_SEND RefundStatus = "waitbuyersend"
	// REFUND_WAIT_SELLER_RECEIVE 等待卖家确认收货
	REFUND_WAIT_SELLER_RECEIVE RefundStatus = "waitsellerreceive"
	// REFUND_SUCCESS 退款成功
	REFUND_SUCCESS RefundStatus = "refundsuccess"
	// REFUND_CLOSE 退款成功
	REFUND_CLOSE RefundStatus = "refundclose"
)

// BusinessType 业务类型
type BusinessType = string

const (
	// TA_BT 国际站 - 信保
	TA_BT BusinessType = "ta"
	// WHOLESALE_BT 国际站 - 在线批发
	WHOLESALE_BT BusinessType = "wholesale"
	// CN_BT 普通订单类型
	CN_BT BusinessType = "cn"
	// WS_BT 大额批发订单类型
	WS_BT BusinessType = "ws"
	// YP_BT 普通拿样订单类型
	YP_BT BusinessType = "yp"
	// YF_BT 一分钱拿样订单类型
	YF_BT BusinessType = "yf"
	// FS_BT 倒批(限时折扣)订单类型
	FS_BT BusinessType = "fs"
	// CZ_BT 加工定制订单类型
	CZ_BT BusinessType = "cz"
	// AG_BT 协议采购订单类型
	AG_BT BusinessType = "ag"
	// HP_BT 伙拼订单类型
	HP_BT BusinessType = "hp"
	// SUPPLY_BT 供销订单类型
	SUPPLY_BT BusinessType = "supply"
	// FACTORY_BT 淘工厂订单
	FACTORY_BT BusinessType = "factory"
	// QUICK_BT 快订下单
	QUICK_BT BusinessType = "quick"
	// XIANGPIN_BT 享拼订单
	XIANGPIN_BT BusinessType = "xiangpin"
	// F2F_BT 当面付
	F2F_BT BusinessType = "f2f"
	// CYFW_BT 存样服务
	CYFW_BT BusinessType = "cyfw"
	// SP_BT 代销订单
	SP_BT BusinessType = "sp"
	// WG_BT 微供订单
	WG_BT BusinessType = "wg"
	// LST_BT 零售通
	LST_BT BusinessType = "lst"
	// CB_BT 跨境
	CB_BT BusinessType = "cb"
	// DISTRIBUTION_BT 分销
	DISTRIBUTION_BT BusinessType = "distribution"
	// CAB_BT 采源宝
	CAB_BT BusinessType = "cab"
	// MANUFACT_BT 加工定制
	MANUFACT_BT BusinessType = "manufact"
)

// LstWarehouseType 零售通仓库类型。customer：虚仓；cainiao：实仓
type LstWarehouseType = string

const (
	// CUSTOMER_LST_WAREHOUSE 虚仓
	CUSTOMER_LST_WAREHOUSE LstWarehouseType = "customer"
	// CAINIAO_LST_WAREHOUSE 实仓
	CAINIAO_LST_WAREHOUSE LstWarehouseType = "cainiao"
)

// StepOrderStatus 分阶段交易状态
type StepOrderStatus = string

const (
	// STEP_ORDER_WAIT_ACTIVATE 未开始（待激活）
	STEP_ORDER_WAIT_ACTIVATE StepOrderStatus = "waitactive"
	// STEP_ORDER_WAIT_SELLER_PUSH 等待卖家推进
	STEP_ORDER_WAIT_SELLER_PUSH StepOrderStatus = "waitsellerpush"
	// STEP_ORDER_SUCCESS 本阶段完成
	STEP_ORDER_SUCCESS StepOrderStatus = "success"
	// STEP_ORDER_SETTLE_BILL 分账
	STEP_ORDER_SETTLE_BILL StepOrderStatus = "settlebill"
	// STEP_ORDER_CANCEL 本阶段终止
	STEP_ORDER_CANCEL StepOrderStatus = "cancel"
	// STEP_ORDER_INACTIVE_AND_CANCEL 本阶段未开始便终止
	STEP_ORDER_INACTIVE_AND_CANCEL StepOrderStatus = "inactiveandcancel"
	// STEP_ORDER_WAIT_BUYER_PAY 等待买家付款
	STEP_ORDER_WAIT_BUYER_PAY StepOrderStatus = "waitbuyerpay"
	// STE_ORDER_WAIT_SELLER_SEND 等待卖家发货
	STEP_ORDER_WAIT_SELLER_SEND StepOrderStatus = "waitsellersend"
	// STEP_ORDER_WAIT_BUYER_RECEIVE 等待买家确认收货
	STEP_ORDER_WAIT_BUYER_RECEIVE StepOrderStatus = "waitbuyerreceive"
	// STEP_ORDER_WAIT_SELLER_ACT 等待卖家XX操作
	STEP_ORDER_WAIT_SELLER_ACT StepOrderStatus = "waitselleract"
	// STEP_ORDER_WAIT_BUYER_CONFIRMATION 等待买家确认XX操作
	STEP_ORDER_WAIT_BUYER_CONFIRMATION StepOrderStatus = "waitbuyerconfirmaction"
)

// QualityAssuranceType 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后)
type QualityAssuranceType = string

const (
	// PRE_SHIPMENT_QA 发货前
	PRE_SHIPMENT_QA QualityAssuranceType = "pre_shipment"
	// POST_DELIVERY_QA 发货后
	POST_DELIVERY_QA QualityAssuranceType = "post_delivery"
)

// DisputeRequest 退款/退款退货。只有已收到货，才可以选择退款退货。退款:"refund"; 退款退货:"returnRefund"
type DisputeRequest = string

const (
	// DISPUTE_REFUND 退款
	DISPUTE_REFUND DisputeRequest = "refund"
	// DISPUTE_RETURN_REFUND 退款退货
	DISPUTE_RETURN_REFUND DisputeRequest = "returnRefund"
)

// RefundGoodsStatus 货物状态
type RefundGoodsStatus = string

const (
	// REFUND_WAIT_SELLER_SEND 售中等待卖家发货
	REFUND_WAIT_SELLER_SEND RefundGoodsStatus = "refundWaitSellerSend"
	// REFUND_WAIT_BUYER_RECEIVE 售中等待买家收货
	REFUND_WAIT_BUYER_RECEIVE RefundGoodsStatus = "refundWaitBuyerReceive"
	// REFUND_BUYER_RECEIVED 售中已收货（未确认完成交易）
	REFUND_BUYER_RECEIVED RefundGoodsStatus = "refundBuyerReceived"
	// AFTER_SALE_BUYER_NOT_RECEIVED 售后未收货
	AFTER_SALE_BUYER_NOT_RECEIVED RefundGoodsStatus = "aftersaleBuyerNotReceived"
	// AFTER_SALE_BUYER_RECEIVED 售后已收到货
	AFTER_SALE_BUYER_RECEIVED RefundGoodsStatus = "aftersaleBuyerReceived"
)
