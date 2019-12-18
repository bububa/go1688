package trade

const (
	NAMESPACE = "com.alibaba.trade"
)

type APIWebsite = string

const (
	API_ALIBABA APIWebsite = "alibaba"
	API_1688    APIWebsite = "1688"
)

type TradeType = string

const (
	FXASSURE_TRADE   TradeType = "fxassure"  // 交易4.0通用担保交易
	ALIPAY_TRADE     TradeType = "alipay"    // 大市场通用的支付宝担保交易（目前在做切流，后续会下掉）
	PERIOD_TRADE     TradeType = "period"    // 普通账期交易
	ASSURE_TRADE     TradeType = "assure"    // 大买家企业采购询报价下单时需要使用的担保交易流程
	CREDIT_BUY_TRADE TradeType = "creditBuy" // 诚E赊
	BANK_TRADE       TradeType = "bank"      // 银行转账
	STAGED631_TRADE  TradeType = "631staged" // 631分阶段付款
	STAGED37_TRADE   TradeType = "37staged"  // 37分阶段
)

type CancelReason = string

const (
	BUYER_CANCEL      CancelReason = "buyerCancel"     // 买家取消订单
	SELLER_GOODS_LACK CancelReason = "sellerGoodsLack" // 卖家库存不足
	OTHER_REASON      CancelReason = "other"           // 其它
)

type TradeCloseType = string

const (
	CLOSE_TRADE_BY_SELLER  TradeCloseType = "CLOSE_TRADE_BY_SELLER" // 卖家关闭交易
	CLOSE_TRADE_BY_BOPS    TradeCloseType = "CLOSE_TRADE_BY_BOPS"   // 后台关闭交
	CLOSE_TRADE_BY_SYSTEM  TradeCloseType = "CLOSE_TRADE_BY_SYSTEM" // 系统（超时）关闭交易
	CLOSE_TRADE_BY_BUYER   TradeCloseType = "CLOSE_TRADE_BY_BUYER"  // 买家关闭交易
	CLOSE_TRADE_BY_CREADIT TradeCloseType = "CLOSE_TRADE_BY_CREADT" // 诚信保障投诉关闭
)

type TradeStatus = string

const (
	TRADE_WAIT_BUYER_PAY     TradeStatus = "waitbuyerpay"     // 等待买家付款
	TRADE_WAIT_SELLER_SEND   TradeStatus = "waitsellersend"   // 等待卖家发货
	TRADE_WAIT_BUYER_RECEIVE TradeStatus = "waitbuyerreceive" // 等待买家收货
	TRADE_CONFIRM_GOODS      TradeStatus = "confirm_goods"    // 已收货
	TRADE_SUCCESS            TradeStatus = "success"          // 交易成功
	TRADE_CANCEL             TradeStatus = "cancel"           // 交易取消
	TRADE_TERMINATED         TradeStatus = "terminated"       // 交易终止
	TRADE_UNKNOWN            TradeStatus = ""                 //其它
)

type RefundStatus = string

const (
	REFUND_WAIT_SELLER_AGREE   RefundStatus = "waitselleragree"   // 等待卖家同意
	REFUND_WAIT_BUYER_MODIFY   RefundStatus = "waitbuyermodify"   // 待买家修改
	REFUND_WAIT_BUYER_SEND     RefundStatus = "waitbuyersend"     // 等待买家退货
	REFUND_WAIT_SELLER_RECEIVE RefundStatus = "waitsellerreceive" // 等待卖家确认收货
	REFUND_SUCCESS             RefundStatus = "refundsuccess"     // 退款成功
	REFUND_CLOSE               RefundStatus = "refundclose"       // 退款成功
)

type BusinessType = string

const (
	TA_BT           BusinessType = "ta"           // 国际站 - 信保
	WHOLESALE_BT    BusinessType = "wholesale"    // 国际站 - 在线批发
	CN_BT           BusinessType = "cn"           // 普通订单类型
	WS_BT           BusinessType = "ws"           // 大额批发订单类型
	YP_BT           BusinessType = "yp"           // 普通拿样订单类型
	YF_BT           BusinessType = "yf"           // 一分钱拿样订单类型
	FS_BT           BusinessType = "fs"           // 倒批(限时折扣)订单类型
	CZ_BT           BusinessType = "cz"           // 加工定制订单类型
	AG_BT           BusinessType = "ag"           // 协议采购订单类型
	HP_BT           BusinessType = "hp"           // 伙拼订单类型
	SUPPLY_BT       BusinessType = "supply"       // 供销订单类型
	FACTORY_BT      BusinessType = "factory"      // 淘工厂订单
	QUICK_BT        BusinessType = "quick"        // 快订下单
	XIANGPIN_BT     BusinessType = "xiangpin"     // 享拼订单
	F2F_BT          BusinessType = "f2f"          // 当面付
	CYFW_BT         BusinessType = "cyfw"         // 存样服务
	SP_BT           BusinessType = "sp"           // 代销订单
	WG_BT           BusinessType = "wg"           // 微供订单
	LST_BT          BusinessType = "lst"          // 零售通
	CB_BT           BusinessType = "cb"           // 跨境
	DISTRIBUTION_BT BusinessType = "distribution" // 分销
	CAB_BT          BusinessType = "cab"          // 采源宝
	MANUFACT_BT     BusinessType = "manufact"     // 加工定制
)

type LstWarehouseType = string

const (
	CUSTOMER_LST_WAREHOUSE LstWarehouseType = "customer"
	CAINIAO_LST_WAREHOUSE  LstWarehouseType = "cainiao"
)

type StepOrderStatus = string

const (
	STEP_ORDER_WAIT_ACTIVATE           StepOrderStatus = "waitactive"             // 未开始（待激活）
	STEP_ORDER_WAIT_SELLER_PUSH        StepOrderStatus = "waitsellerpush"         // 等待卖家推进
	STEP_ORDER_SUCCESS                 StepOrderStatus = "success"                // 本阶段完成
	STEP_ORDER_SETTLE_BILL             StepOrderStatus = "settlebill"             // 分账
	STEP_ORDER_CANCEL                  StepOrderStatus = "cancel"                 // 本阶段终止
	STEP_ORDER_INACTIVE_AND_CANCEL     StepOrderStatus = "inactiveandcancel"      // 本阶段未开始便终止
	STEP_ORDER_WAIT_BUYER_PAY          StepOrderStatus = "waitbuyerpay"           // 等待买家付款
	STEP_ORDER_WAIT_SELLER_SEND        StepOrderStatus = "waitsellersend"         // 等待卖家发货
	STEP_ORDER_WAIT_BUYER_RECEIVE      StepOrderStatus = "waitbuyerreceive"       // 等待买家确认收货
	STEP_ORDER_WAIT_SELLER_ACT         StepOrderStatus = "waitselleract"          // 等待卖家XX操作
	STEP_ORDER_WAIT_BUYER_CONFIRMATION StepOrderStatus = "waitbuyerconfirmaction" // 等待买家确认XX操作
)

type QualityAssuranceType = string

const (
	PRE_SHIPMENT_QA  QualityAssuranceType = "pre_shipment"  // 发货前
	POST_DELIVERY_QA QualityAssuranceType = "post_delivery" // 发货后
)

type DisputeRequest = string

const (
	DISPUTE_REFUND        DisputeRequest = "refund"
	DISPUTE_RETURN_REFUND DisputeRequest = "returnRefund"
)

type RefundGoodsStatus = string

const (
	REFUND_WAIT_SELLER_SEND       RefundGoodsStatus = "refundWaitSellerSend"      // 售中等待卖家发货
	REFUND_WAIT_BUYER_RECEIVE     RefundGoodsStatus = "refundWaitBuyerReceive"    // 售中等待买家收货
	REFUND_BUYER_RECEIVED         RefundGoodsStatus = "refundBuyerReceived"       // 售中已收货（未确认完成交易）
	AFTER_SALE_BUYER_NOT_RECEIVED RefundGoodsStatus = "aftersaleBuyerNotReceived" // 售后未收货
	AFTER_SALE_BUYER_RECEIVED     RefundGoodsStatus = "aftersaleBuyerReceived"    // 售后已收到货
)
