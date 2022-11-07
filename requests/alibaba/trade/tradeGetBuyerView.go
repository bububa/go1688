package trade

import (
	"strconv"

	"github.com/shopspring/decimal"

	"github.com/bububa/go1688"
)

// TradeGetBuyerViewRequest 订单详情查看(买家视角) API Request
type TradeGetBuyerViewRequest struct {
	// Website 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
	Website APIWebsite `json:"webSite,omitempty"`
	// OrderID 交易id，订单号
	OrderID uint64 `json:"orderId"`
	// IncludeFields 查询结果中包含的域，GuaranteesTerms：保障条款，NativeLogistics：物流信息，RateDetail：评价详情，OrderInvoice：发票信息。默认返回GuaranteesTerms、NativeLogistics、OrderInvoice。
	IncludeFields string `json:"includeFields,omitempty"`
	// AttributeKeys 垂直表中的attributeKeys
	AttributeKeys []string `json:"attributeKeys,omitempty"`
}

// Name implement ReqeustData interface
func (r TradeGetBuyerViewRequest) Name() string {
	return "alibaba.trade.cancel"
}

// Map implement ReqeustData interface
func (r TradeGetBuyerViewRequest) Map() map[string]string {
	ret := make(map[string]string, 4)
	ret["webSite"] = r.Website
	ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	if r.IncludeFields != "" {
		ret["includeFields"] = r.IncludeFields
	}
	if len(r.AttributeKeys) > 0 {
		ret["attributeKeys"] = go1688.JSONMarshal(r.AttributeKeys)
	}
	return ret
}

// TradeGetBuyerViewResponse 订单详情查看(买家视角) API Response
type TradeGetBuyerViewResponse struct {
	go1688.BaseResponse
	Result *TradeInfo `json:"result,omitempty"`
}

// TradeInfo 交易信息
type TradeInfo struct {
	// BaseInfo 订单基础信息
	BaseInfo *OrderBaseInfo `json:"baseInfo,omitempty"`
	// BizInfo 订单业务信息
	BizInfo *OrderBizInfo `json:"orderBizInfo,omitempty"`
	// TradeTerms 交易条款
	TradeTerms []TradeTermInfo `json:"tradeTerms,omitempty"`
	// ProductItems 商品条目信息
	ProductItems []ProductItemInfo `json:"productItems,omitempty"`
	// NativeLogistics 国内物流
	NativeLogistics *NativeLogisticsInfo `json:"nativeLogistics,omitempty"`
	// InvoiceInfo 发票信息
	InvoiceInfo *OrderInvoiceInfo `json:"orderInvoiceInfo,omitempty"`
	// GuaranteesTerms 保障条款
	GuaranteesTerms *GuaranteesTermInfo `json:"guaranteesTerms,omitempty"`
	// OrderRateInfo 订单评价信息
	OrderRateInfo *OrderRateInfo `json:"orderRateInfo,omitempty"`
	// OverseasExtraAddress 跨境地址扩展信息
	OverseasExtraAddress *OverseasExtraAddress `json:"overseasExtraAddress,omitempty"`
	// Customs 跨境报关信息
	Customs *OrderCustoms `json:"customs,omitempty"`
	// QuoteList 采购单详情列表，为大企业采购订单独有域。
	QuoteList []CaigouQuoteInfo `json:"quoteList,omitempty"`
	// ExtAttributes 订单扩展属性
	ExtAttributes []KeyValuePair `json:"extAttributes,omitempty"`
}

// OrderBaseInfo 订单基础信息
type OrderBaseInfo struct {
	// AllDeliveredTime 完全发货时间
	AllDeliveredTime go1688.JsonTime `json:"allDeliveredTime,omitempty"`
	// SellerCreditLevel 卖家诚信等级
	SellerCreditLevel string `json:"sellerCreditLevel,omitempty"`
	// PayTime 付款时间，如果有多次付款，这里返回的是首次付款时间
	PayTime go1688.JsonTime `json:"payTime,omitempty"`
	// Discount 折扣信息，单位分
	Discount int64 `json:"discount,omitempty"`
	// AlipayTradeID 外部支付交易Id
	AlipayTradeID string `json:"alipayTradeId,omitempty"`
	// SumProductPayment 产品总金额(该订单产品明细表中的产品金额的和)，单位元
	SumProductPayment decimal.Decimal `json:"sumProductPayment,omitempty"`
	// BuyerFeedback 买家留言，不超过500字
	BuyerFeedback string `json:"buyerFeedback,omitempty"`
	// FlowTemplateCode 4.0交易流程模板code
	FlowTemplateCode string `json:"flowTemplateCode,omitempty"`
	// SellerOrder 是否自主订单（邀约订单）
	SellerOrder bool `json:"sellerOrder,omitempty"`
	// BuyerLoginID 买家loginId，旺旺Id
	BuyerLoginID string `json:"buyerLoginId,omitempty"`
	// ModifyTime 修改时间
	ModifyTime go1688.JsonTime `json:"modifyTime,omitempty"`
	// SubBuyerLoginID 买家子账号
	SubBuyerLoginID string `json:"subBuyerLoginId,omitempty"`
	// ID 交易id
	ID uint64 `json:"id,omitempty"`
	// CloseReason 关闭原因。buyerCancel:买家取消订单，sellerGoodsLack:卖家库存不足，other:其它
	CloseReason CancelReason `json:"closeReason,omitempty"`
	// BuyerContact 买家联系人
	BuyerContact *TradeContact `json:"buyerContact,omitempty"`
	// SellerAlipayID 卖家支付宝id
	SellerAlipayID string `json:"sellerAlipayId,omitempty"`
	// CompleteTime 完成时间
	CompleteTime go1688.JsonTime `json:"completeTime,omitempty"`
	// SellerLoginID 卖家oginId，旺旺Id
	SellerLoginID string `json:"sellerLoginId,omitempty"`
	// BuyerID 买家主账号id
	BuyerID string `json:"buyerID,omitempty"`
	// CloseOperateType 关闭订单操作类型。CLOSE_TRADE_BY_SELLER:卖家关闭交易,CLOSE_TRADE_BY_BOPS:BOPS后台关闭交易,CLOSE_TRADE_BY_SYSTEM:系统（超时）关闭交易,CLOSE_TRADE_BY_BUYER:买家关闭交易,CLOSE_TRADE_BY_CREADIT:诚信保障投诉关闭
	CloseOperateType TradeCloseType `json:"closeOperateType,omitempty"`
	// TotalAmount 应付款总金额，totalAmount = ∑itemAmount + shippingFee，单位为元
	TotalAmount decimal.Decimal `json:"totalAmount,omitempty"`
	// SellerID 卖家主账号id
	SellerID string `json:"sellerID,omitempty"`
	// ShippingFee 运费，单位为元
	ShippingFee decimal.Decimal `json:"shippingFee,omitempty"`
	// BuyerUserID 买家数字id
	BuyerUserID uint64 `json:"buyerUserId,omitempty"`
	// BuyerMemo 买家备忘信息
	BuyerMemo string `json:"buyerMemo,omitempty"`
	// Refund 退款金额，单位为元
	Refund decimal.Decimal `json:"refund,omitempty"`
	// Status 交易状态，waitbuyerpay:等待买家付款;waitsellersend:等待卖家发货;waitbuyerreceive:等待买家收货;confirm_goods:已收货;success:交易成功;cancel:交易取消;terminated:交易终止;未枚举:其他状态
	Status TradeStatus `json:"status,omitempty"`
	// RefundPayment 退款金额
	RefundPayment int64 `json:"refundPayment,omitempty"`
	// SellerContact 卖家联系人信息
	SellerContact *TradeContact `json:"sellerContact,omitempty"`
	// RefundStatus 订单的售中退款状态，等待卖家同意：waitselleragree ，待买家修改：waitbuyermodify，等待买家退货：waitbuyersend，等待卖家确认收货：waitsellerreceive，退款成功：refundsuccess，退款失败：refundclose
	RefundStatus RefundStatus `json:"refundStatus,omitempty"`
	// Remark 备注，1688指下单时的备注
	Remark string `json:"remark,omitempty"`
	// PreOrderID 预订单ID
	PreOrderID uint64 `json:"preOrderId,omitempty"`
	// ConfirmedTime 确认时间
	ConfirmedTime go1688.JsonTime `json:"confirmedTime,omitempty"`
	// CloseRemark 关闭订单备注
	CloseRemark string `json:"closeRemark,omitempty"`
	// TradeType 1:担保交易 2:预存款交易 3:ETC境外收单交易 4:即时到帐交易 5:保障金安全交易 6:统一交易流程 7:分阶段付款 8.货到付款交易 9.信用凭证支付交易 10.账期支付交易，50060 交易4.0
	TradeType string `json:"tradeType,omitempty"`
	// ReceivingTime 收货时间，这里返回的是完全收货时间
	ReceivingTime go1688.JsonTime `json:"receivingTime,omitempty"`
	// StepAgreementPath 分阶段法务协议地址
	StepAgreementPath string `json:"stepAgreementPath,omitempty"`
	// IDOfStr 交易id(字符串格式)
	IDOfStr string `json:"idOfStr,omitempty"`
	// RefundStatusForAs 订单的售后退款状态
	RefundStatusForAs string `json:"refundStatusForAs,omitempty"`
	// StepPayAll 是否一次性付款
	StepPayAll bool `json:"stepPayAll,omitempty"`
	// SellerUserID 卖家数字id
	SellerUserID uint64 `json:"sellerUserId,omitempty"`
	// StepOrderList [交易3.0]分阶段交易，分阶段订单list
	StepOrderList []StepOrder `json:"stepOrderList,omitempty"`
	// BuyerAlipayID 买家支付宝id
	BuyerAlipayID string `json:"buyerAlipayId,omitempty"`
	// CreateiIme 创建时间
	CreateTime go1688.JsonTime `json:"createTime,omitempty"`
	// BusinessType 业务类型。国际站：ta(信保),wholesale(在线批发)。 中文站：普通订单类型 = "cn"; 大额批发订单类型 = "ws"; 普通拿样订单类型 = "yp"; 一分钱拿样订单类型 = "yf"; 倒批(限时折扣)订单类型 = "fs"; 加工定制订单类型 = "cz"; 协议采购订单类型 = "ag"; 伙拼订单类型 = "hp"; 供销订单类型 = "supply"; 淘工厂订单 = "factory"; 快订下单 = "quick"; 享拼订单 = "xiangpin"; 当面付 = "f2f"; 存样服务 = "cyfw"; 代销订单 = "sp"; 微供订单 = "wg";零售通 = "lst";跨境='cb';分销='distribution';采源宝='cab';加工定制="manufact"
	BusinessType BusinessType `json:"businessType,omitempty"`
	// OverSeaOrder 是否海外代发订单，是：true
	OverSeaOrder bool `json:"overSeaOrder,omitempty"`
	// RefundID 退款单ID
	RefundID string `json:"refundId,omitempty"`
	// TradeTypeDesc 下单时指定的交易方式
	TradeTypeDesc string `json:"tradeTypeDesc,omitempty"`
	// PayChannelList 支付渠道名称列表。一笔订单可能存在多种支付渠道。枚举值：支付宝,网商银行信任付,诚e赊,对公转账,赊销宝,账期支付,合并支付渠道,支付平台,声明付款,网商电子银行承兑汇票,银行转账,跨境宝,红包,其它
	PayChannelList []string `json:"payChannelList,omitempty"`
	// TradeTypeCode 下单时指定的交易方式tradeType
	TradeTypeCode string `json:"tradeTypeCode,omitempty"`
	// PayTimeout 支付超时时间，定长情况时单位：秒，目前都是定长
	PayTimeout int64 `json:"payTimeout,omitempty"`
	// PayTimeoutType 支付超时TYPE，0：定长，1：固定时间
	PayTimeoutType uint `json:"payTimeoutType,omitempty"`
}

// TradeContact 联系人
type TradeContact struct {
	// Phone 联系电话
	Phone string `json:"phone,omitempty"`
	// Fax 传真
	Fax string `json:"fax,omitempty"`
	// Email 邮箱
	Email string `json:"email,omitempty"`
	// ImInPlatform 联系人在平台的IM账号
	ImInPlatform string `json:"imInPlatform,omitempty"`
	// Name 联系人名称
	Name string `json:"name,omitempty"`
	// Mobile 联系人手机号
	Mobile string `json:"mobile,omitempty"`
	// CompanyName 公司名称
	CompanyName string `json:"companyName,omitempty"`
	// WgSenderName 发件人名称，在微供等分销场景下由分销商设置
	WgSenderName string `json:"wgSenderName,omitempty"`
	// WgSenderPhone 发件人电话，在微供等分销场景下由分销商设置
	WgSenderPhone string `json:"wgSenderPhone,omitempty"`
}

// StepOrder [交易3.0]分阶段交易，分阶段订单
type StepOrder struct {
	// ID 阶段id
	ID uint64 `json:"stepOrderId,omitempty"`
	// Status waitactivate 未开始（待激活） waitsellerpush 等待卖家推进 success 本阶段完成 settlebill 分账 cancel 本阶段终止 inactiveandcancel 本阶段未开始便终止 waitbuyerpay 等待买家付款 waitsellersend 等待卖家发货 waitbuyerreceive 等待买家确认收货 waitselleract 等待卖家XX操作 waitbuyerconfirmaction 等待买家确认XX操作
	Status StepOrderStatus `json:"stepOrderStatus,omitempty"`
	// PayStatus 1 未冻结/未付款 2 已冻结/已付款 4 已退款 6 已转交易 8 交易未付款被关闭
	PayStatus int `json:"stepPayStatus,omitempty"`
	// No 阶段序列：1、2、3...
	No int `json:"stepNo,omitepty"`
	// LastStep 是否最后一个阶段
	LastStep bool `json:"lastStep,omitempty"`
	// HasDisbursed 是否已打款给卖家
	HasDisbursed bool `json:"hasDisbursed,omitempty"`
	// PayFee 创建时需要付款的金额，不含运费
	PayFee decimal.Decimal `json:"payFee,omitempty"`
	// ActualPayFee 应付款（含运费）= 单价×数量-单品优惠-店铺优惠+运费+修改的金额（除运费外，均指分摊后的金额）
	ActualPayFee decimal.Decimal `json:"actualPayFee,omitempty"`
	// DiscountFee 本阶段分摊的店铺优惠
	DiscountFee decimal.Decimal `json:"discountFee,omitempty"`
	// ItemDiscountFee 本阶段分摊的单品优惠
	ItemDiscountFee decimal.Decimal `json:"itemDiscountFee,omitempty"`
	// Price 本阶段分摊的单价
	Price decimal.Decimal `json:"price,omitempty"`
	// Amount 购买数量
	Amount int `json:"amount,omitempty"`
	// PostFee 运费
	PostFee decimal.Decimal `json:"postFee,omitempty"`
	// AdjustPostFee 修改价格修改的金额
	AdjustPostFee decimal.Decimal `json:"adjustPostFee,omitempty"`
	// GmtCreate 创建时间
	GmtCreate go1688.JsonTime `json:"gmtCreate,omitempty"`
	// GmtModified 修改时间
	GmtModified go1688.JsonTime `json:"gmtModified,omitempty"`
	// EnterTime 开始时间
	EnterTime go1688.JsonTime `json:"enterTime,omitempty"`
	// PayTime 付款时间
	PayTime go1688.JsonTime `json:"payTime,omitempty"`
	// SellerActionTime 卖家操作时间
	SellerActionTime go1688.JsonTime `json:"sellerActionTime,omitempty"`
	// EndTime 本阶段结束时间
	EndTime go1688.JsonTime `json:"endTime,omitempty"`
	// MessagePath 卖家操作留言路径
	MessagePath string `json:"messagePath,omitempty"`
	// PicturePath 卖家上传图片凭据路径
	PicturePath string `json:"picturePath,omitempty"`
	// Message 卖家操作留言
	Message string `json:"message,omitempty"`
	// TemplateID 使用的模板id
	TemplateID uint64 `json:"templateId,omitempty"`
	// Name 当前步骤的名称
	Name string `json:"stepName,omitempty"`
	// SellerActionName 卖家操作名称
	SellerActionName string `json:"sellerActionName,omitempty"`
	// BuyerPayTimeout 买家不付款的超时时间(秒)
	BuyerPayTimeout int64 `json:"buyerPayTimeout,omitempty"`
	// BuyerConfirmTimeout 买家不确认的超时时间
	BuyerConfirmTimeout int64 `json:"buyerConfirmTimeout,omitempty"`
	// NeedLogistics 是否需要物流
	NeedLogistics bool `json:"needLogistics,omitempty"`
	// NeedSellerAction 是否需要卖家操作和买家确认
	NeedSellerAction bool `json:"needSellerAction,omitempty"`
	// TransferAfterConfirm 阶段结束是否打款
	TransferAfterConfirm bool `json:"transferAfterConfirm,omitempty"`
	// NeedSellerCallNext 是否需要卖家推进
	NeedSellerCallNext bool `json:"needSellerCallNext,omitempty"`
	// InstantPay 是否允许即时到帐
	InstantPay bool `json:"instantPay,omitempty"`
}

// OrderBizInfo 订单业务信息
type OrderBizInfo struct {
	// OdsCyd 是否采源宝订
	OdsCyd bool `json:"odsCyd,omitempty"`
	// AccountPeriodTime 账期交易订单的到账时间
	AccountPeriodTime string `json:"accountPeriodTime,omitempty"`
	// CreditOrder 为true，表示下单时选择了诚e赊交易方式。注意不等同于“诚e赊支付”，支付时有可能是支付宝付款，具体支付方式查询tradeTerms.payWay
	CreditOrder bool `json:"creditOrder,omitempty"`
	// CreditOrderDetail 诚e赊支付详情，只有使用诚e赊付款时返回
	CreditOrderDetail *CreditOrderDetail `json:"creditOrderDetail,omitempty"`
	// PreOrderInfo 预订单信息
	PreOrderInfo *PreOrderInfo `json:"preOrderInfo,omitempty"`
	// ListOrderInfo .
	LstOrderInfo *LstOrderInfo `json:"lstOrderInfo,omitempty"`
}

// CreditOrderDetail 诚e赊支付详情，只有使用诚e赊付款时返回
type CreditOrderDetail struct {
	// PayAmount 订单金额
	PayAmount int64 `json:"payAmount,omitempty"`
	// CreateTime 支付时间
	CreateTime string `json:"createTime,omitempty"`
	// Status 状态
	Status string `json:"status,omitempty"`
	// GracePeriodEndTime 最晚还款时间
	GracePeriodEndTime string `json:"gracePeriodEndTime,omitempty"`
	// StatusStr 状态描述
	StatusStr string `json:"statusStr,omitempty"`
	// RestRepayAmount 应还金额
	RestRepayAmount int64 `json:"restRepayAmount,omitempty"`
}

// PreOrderInfo 预订单信息
type PreOrderInfo struct {
	// MarketName 创建预订单时传入的市场名
	MarketName string `json:"marketName,omitempty"`
	// CreatePreOrderApp 预订单是否为当前查询的通过当前查询的ERP创建
	CreatePreOrderApp bool `json:"createPreOrderApp,omitempty"`
}

type LstOrderInfo struct {
	// LstWarehouseType 零售通仓库类型。customer：虚仓；cainiao：实仓
	LstWarehouseType LstWarehouseType `json:"lstWarehouseType,omitempty"`
}

// TradeTermInfo 交易条款
type TradeTermInfo struct {
	// PayStatus 支付状态。国际站：WAIT_PAY(未支付),PAYER_PAID(已完成支付),PART_SUCCESS(部分支付成功),PAY_SUCCESS(支付成功),CLOSED(风控关闭),CANCELLED(支付撤销),SUCCESS(成功),FAIL(失败)。 1688:1(未付款);2(已付款);4(全额退款);6(卖家有收到钱，回款完成) ;7(未创建外部支付单);8 (付款前取消) ; 9(正在支付中);12(账期支付,待到账)
	PayStatus string `json:"payStatus,omitempty"`
	// PayTime 完成阶段支付时间
	PayTime go1688.JsonTime `json:"payTime,omitempty"`
	// PayWay 支付方式。 国际站：ECL(融资支付),CC(信用卡),TT(线下TT),ACH(echecking支付)。 1688:1-支付宝,2-网商银行信任付,3-诚e赊,4-银行转账,5-赊销宝,6-电子承兑票据,7-账期支付,8-合并支付渠道,9-无打款,10-零售通赊购,13-支付平台,12-声明付款
	PayWay string `json:"payWay,omitempty"`
	// PhasAmount 付款额
	PhasAmount decimal.Decimal `json:"phasAmount,omitempty"`
	// Phase 阶段单id
	Phase uint64 `json:"phase,omitempty"`
	// PhaseCondition 阶段条件，1688无此内容
	PhaseCondition string `json:"phaseCondition,omitempty"`
	// PhaseDate 阶段时间，1688无此内容
	PhaseDate string `json:"phaseDate,omitempty"`
	// CardPay 是否银行卡支付
	CardPay bool `json:"cardPay,omitempty"`
	// ExpressPay 是否快捷支付
	ExpressPay bool `json:"expressPay,omitempty"`
	// PayWayDesc 支付方式
	PayWayDesc string `json:"payWayDesc,omitempty"`
}

// ProductItemInfo 商品条目信息
type ProductItemInfo struct {
	// CargoNumber 指定单品货号，国际站无需关注。该字段不一定有值，仅仅在下单时才会把货号记录(如果卖家设置了单品货号的话)。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。
	CargoNumber string `json:"cargoNumber,omitempty"`
	// Description 描述,1688无此信息
	Description string `json:"description,omitempty"`
	// ItemAmount 实付金额，单位为元
	ItemAmount decimal.Decimal `json:"itemAmount,omitempty"`
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// Price 原始单价，以元为单位
	Price decimal.Decimal `json:"price,omitempty"`
	// ID 产品ID（非在线产品为空）
	ID uint64 `json:"productID,omitempty"`
	// ImgURL 商品图片url
	ImgURL []string `json:"productImgUrl,omitempty"`
	// SnapshotURL 产品快照url，交易订单产生时会自动记录下当时的商品快照，供后续纠纷时参考
	SnapshotURL string `json:"productSnapshotUrl,omitempty"`
	// Quantity 以unit为单位的数量，例如多少个、多少件、多少箱、多少吨
	Quantity decimal.Decimal `json:"quantity,omitempty"`
	// Refund 退款金额，单位为元
	Refund decimal.Decimal `json:"refund,omitempty"`
	// SkuID
	SkuID uint64 `json:"skuID,omitempty"`
	// Sort 排序字段，商品列表按此字段进行排序，从0开始，1688不提供
	Sort int `json:"sort,omitempty"`
	// Status 子订单状态
	Status string `json:"status,omitempty"`
	// SubItemID 子订单号，或商品明细条目ID
	SubItemID uint64 `json:"subItemId,omitempty"`
	// Type 类型，国际站使用，供卖家标注商品所属类型
	Type string `json:"type,omitempty"`
	// Unit 售卖单位 例如：个、件、箱、吨
	Unit string `json:"unit,omitempty"`
	// Weight 重量 按重量单位计算的重量，例如：100
	Weight string `json:"weight,omitempty"`
	// WeightUnit 重量单位 例如：g，kg，t
	WeightUnit string `json:"weightUnit,omitempty"`
	// GuaranteesTerms 保障条款，此字段仅针对1688
	GuaranteesTerms []GuaranteesTermInfo `json:"guaranteesTerms,omitempty"`
	// ProductCargoNumber 指定商品货号，该字段不一定有值，在下单时才会把货号记录。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。该字段和cargoNUmber的区别是：该字段是定义在商品级别上的货号，cargoNUmber是定义在单品级别的货号
	ProductCargoNumber string `json:"productCargoNumber,omitempty"`
	// SkuInfos SKU详情
	SkuInfos []SkuItemDesc `json:"skuInfos,omitempty"`
	// EntryDiscount 订单明细涨价或降价的金额
	EntryDiscount int64 `json:"entryDiscount,omitempty"`
	// SpecID 订单销售属性ID
	SpecID string `json:"specId,omitempty"`
	// QuantityFactor 以unit为单位的quantity精度系数，值为10的幂次，例如:quantityFactor=1000,unit=吨，那么quantity的最小精度为0.001吨
	QuantityFactor decimal.Decimal `json:"quantityFactor,omitempty"`
	// StatusStr 子订单状态描述
	StatusStr string `json:"statusStr,omitempty"`
	// RefundStatus WAIT_SELLER_AGREE 等待卖家同意 REFUND_SUCCESS 退款成功 REFUND_CLOSED 退款关闭 WAIT_BUYER_MODIFY 待买家修改 WAIT_BUYER_SEND 等待买家退货 WAIT_SELLER_RECEIVE 等待卖家确认收货
	RefundStatus string `json:"refundStatus,omitempty"`
	// CloseReason 关闭原因
	CloseReason string `json:"closeReason,omitempty"`
	// LogisticsStatus 1 未发货 2 已发货 3 已收货 4 已经退货 5 部分发货 8 还未创建物流订单
	LogisticsStatus int `json:"logisticsStatus,omitempty"`
	// GmtCreate 创建时间
	GmtCreate go1688.JsonTime `json:"gmtCreate,omitempty"`
	// GmtModified 修改时间
	GmtModified go1688.JsonTime `json:"gmtModified,omitempty"`
	// GmtCompleted 明细完成时间
	GmtCompleted go1688.JsonTime `json:"gmtCompleted,omitempty"`
	// GmtPayExpireTime 库存超时时间，格式为“yyyy-MM-dd HH:mm:ss”
	GmtPayExpireTime string `json:"gmtPayExpireTime,omitempty"`
	// RefundID 售中退款单号
	RefundID string `json:"refundId,omitempty"`
	// SubItemIDString 子订单号，或商品明细条目ID(字符串类型，由于Long类型的ID可能在JS和PHP中处理有问题，所以可以用字符串类型来处理)
	SubItemIDString string `json:"subItemIDString,omitempty"`
	// RefundIDForAs 售后退款单号
	RefundIDForAs string `json:"refundIdForAs,omitempty"`
}

// SkuItemDesc SKU属性
type SkuItemDesc struct {
	// Name 属性名
	Name string `json:"name,omitempty"`
	// Value 属性值
	Value string `json:"value,omitempty"`
}

// GuaranteesTermInfo 保障条款
type GuaranteesTermInfo struct {
	// AssuranceInfo 保障条款
	AssuranceInfo string `json:"assuranceInfo,omitempty"`
	// AssuranceType 保障方式。国际站：TA(信保)
	AssuranceType string `json:"assuranceType,omitempty"`
	// QualityAssuranceType 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后)
	QualityAssuranceType QualityAssuranceType `json:"qualityAssuranceType,omitempty"`
}

// NativeLogisticsInfo 国内物流
type NativeLogisticsInfo struct {
	Address       string `json:"address,omitempty"`
	Area          string `json:"area,omitempty"`
	AreaCode      string `json:"areaCode,omitempty"`
	City          string `json:"city,omitempty"`
	ContactPerson string `json:"contactPerson,omitempty"`
	Fax           string `json:"fax,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	Province      string `json:"province,omitempty"`
	Telephone     string `json:"telephone,omitempty"`
	Zip           string `json:"zip,omitempty"`
	// LogisticsItems 运单明细
	LogisticsItems []NativeLogisticsItemInfo `json:"logisticsItems,omitempty"`
	// TownCode 镇，街道地址码
	TownCode string `json:"townCode,omitempty"`
	// Town 镇，街道
	Town string `json:"town,omitempty"`
}

// NativeLogisticsItemInfo 运单明细
type NativeLogisticsItemInfo struct {
	// DeliveredTime 发货时间
	DeliveredTime go1688.JsonTime `json:"deliveredTime,omitempty"`
	// LogisticsCode 物流编号
	LogisticsCode string `json:"logisticsCode,omitempty"`
	// Type SELF_SEND_GOODS("0")自行发货，在线发货ONLINE_SEND_GOODS("1"，不需要物流的发货 NO_LOGISTICS_SEND_GOODS("2")
	Type string `json:"type,omitempty"`
	// ID 主键id
	ID uint64 `json:"id,omitempty"`
	// Status 状态
	Status string `json:"status,omitempty"`
	// GmtModified 修改时间
	GmtModified go1688.JsonTime `json:"gmtModified,omitempty"`
	// GmtCreate 创建时间
	GmtCreate go1688.JsonTime `json:"gmtCreate,omitempty"`
	// Carriage 运费(单位为元)
	Carriage decimal.Decimal `json:"carriage,omitempty"`
	// FromProvince 发货省
	FromProvince string `json:"fromProvince,omitempty"`
	// FromCity 发货市
	FromCity string `json:"fromCity,omitempty"`
	// FromArea 发货区
	FromArea string `json:"fromArea,omitempty"`
	// FromAddress 发货街道地址
	FromAddress string `json:"fromAddress,omitempty"`
	// FromPhone 发货联系电话
	FromPhone string `json:"fromPhone,omitempty"`
	// FromMobile 发货联系手机
	FromMobile string `json:"fromMobile,omitempty"`
	// FromPost 发货地址邮编
	FromPost string `json:"fromPost,omitempty"`
	// LogisticsCompanyID 物流公司Id
	LogisticsCompanyID uint64 `json:"logisticsCompanyId,omitempty"`
	// LogisticsCompanyNo 物流公司编号
	LogisticesCompanyNo string `json:"logisticsCompanyNo,omitempty"`
	// LogisticsCompanyName 物流公司名称
	LogisticsCompanyName string `json:"logisticsCompanyName,omitempty"`
	// LogisticsBillNo 物流公司运单号
	LogisticsBillNo string `json:"logisticsBillNo,omitempty"`
	// SubItemIDs 商品明细条目id，如有多个以,分隔
	SubItemIDs string `json:"subItemIds,omitempty"`
	// ToProvince 收货省
	ToProvince string `json:"toProvince,omitempty"`
	// ToCity 收货市
	ToCity string `json:"toCity,omitempty"`
	// ToArea 收货区
	ToArea string `json:"toArea,omitempty"`
	// ToAddress 收货街道地址
	ToAddress string `json:"toAddress,omitempty"`
	// ToPhone 收货联系电话
	ToPhone string `json:"toPhone,omitempty"`
	// ToMobile 收货联系手机
	ToMobile string `json:"toMobile,omitempty"`
	// ToPost 收货地址邮编
	ToPost string `json:"toPost,omitempty"`
	// NoLogisticsName 物流姓名
	NoLogisticsName string `json:"noLogisticsName,omitempty"`
	// NoLogisticsTel 联系方式
	NoLogisticsTel string `json:"noLogisticsTel,omitempty"`
	// NoLogisticsBillNo 无需物流业务单号
	NoLogisticsBillNo string `json:"noLogisticsBillNo,omitempty"`
	// NoLogisticsCondition 无需物流类别,noLogisticsCondition=1， 表示其他第三方物流、小型物充商、车队等, noLogisticsCondition=2 表示补运费、差价, noLogisticsCondition=3 表示卖家配送, noLogisticsCondition=4 表示买家自提 noLogisticsCondition=5 表示其他原因
	NoLogisticsCondition string `json:"noLogisticsCondition,omitempty"`
	// IsTimePromise 是否使用限时达物流
	IsTimePromise bool `json:"isTimePromise,omitempty"`
	// ArriveTime 限时达物流，预计到达时间
	ArriveTime go1688.JsonTime `json:"arriveTime,omitempty"`
}

// OrderInvoiceInfo 发票信息
type OrderInvoiceInfo struct {
	// InvoiceCompanyName 发票公司名称(即发票抬头-title)
	InvoiceCompanyName string `json:"invoiceCompanyName,omitempty"`
	// InvoiceType 发票类型. 0：普通发票，1:增值税发票，9未知类型
	InvoiceType int `json:"invoiceType,omitempty"`
	// LocalInvoiceID 本地发票号
	LocalInvoiceID uint64 `json:"localInvoiceId,omitempty"`
	// OrderID 订单Id
	OrderID uint64 `json:"orderId,omitempty"`
	// ReceiveCode (收件人)址区域编码
	ReceiveCode string `json:"receiveCode,omitempty"`
	// ReceiveCodeText (收件人) 省市区编码对应的文案(增值税发票信息)
	ReceiveCodeText string `json:"receiveCodeText,omitempty"`
	// ReceiveMobile （收件者）发票收货人手机
	ReceiveMobile string `json:"receiveMobile,omitempty"`
	// ReceiveName （收件者）发票收货人
	ReceiveName string `json:"receiveName,omitempty"`
	// ReceivePhone （收件者）发票收货人电话
	ReceivePhone string `json:"receivePhone,omitempty"`
	// ReceivePost （收件者）发票收货地址邮编
	ReceivePost string `json:"receivePost,omitempty"`
	// ReceiveStreet (收件人) 街道地址(增值税发票信息)
	ReceiveStreet string `json:"receiveStreet,omitempty"`
	// RegisterAccountID (公司)银行账号
	RegisterAccountID string `json:"registerAccountId,omitempty"`
	// RegisterBank (公司)开户银行
	RegisterBank string `json:"registerBank,omitempty"`
	// RegisterCode (注册)省市区编码
	RegisterCode string `json:"registerCode,omitempty"`
	// RegisterCodeText (注册)省市区文本
	RegisterCodeText string `json:"registerCodeText,omitempty"`
	// RegisterPhone （公司）注册电话
	RegisterPhone string `json:"registerPhone,omitempty"`
	// RegisterStreet (注册)街道地址
	RegisterStreet string `json:"registerStreet,omitempty"`
	// TaxPayerIdentify 纳税人识别号
	TaxPayerIdentify string `json:"taxpayerIdentify,omitempty"`
}

// OrderRateInfo 订单评价信息
type OrderRateInfo struct {
	// BuyerRateStatus 买家评价状态(4:已评论,5:未评论,6;不需要评论)
	BuyerRateStatus int `json:"buyerRateStatus,omitempty"`
	// SellerRateStatus 卖家评价状态(4:已评论,5:未评论,6;不需要评论)
	SellerRateStatus int `json:"sellerRateStatus,omitempty"`
	// BuyerRateList 卖家給买家的评价
	BuyerRateList []OrderRateDetail `json:"buyerRateList,omitempty"`
	// SellerRateList 买家給卖家的评价
	SellerRateList []OrderRateDetail `json:"sellerRateList,omitempty"`
}

// OrderRateDetail 评价详情
type OrderRateDetail struct {
	// StarLevel 评价星级
	StarLevel int `json:"starLevel,omitempty"`
	// Content 评价详情
	Content string `json:"content,omitempty"`
	// ReceiverNick 收到评价的用户昵称
	ReceiverNick string `json:"receiverNick,omitempty"`
	// PosterNick 发送评价的用户昵称
	PosterNick string `json:"posterNick,omitempty"`
	// PublishTime 评价上线时间
	PublishTime go1688.JsonTime `json:"publishTime,omitempty"`
}

// OverseasExtraAddress 跨境地址扩展信息
type OverseasExtraAddress struct {
	// ChannelName 路线名称
	ChannelName string `json:"channelName,omitempty"`
	// ChannelID 路线id
	ChannelID string `json:"channelId,omitempty"`
	// ShippingCompanyID 货代公司id
	ShippingCompanyID string `json:"shippingCompanyId,omitempty"`
	// ShippingCompanyName 货代公司名称
	ShippingCompanyName string `json:"shoppingCompanyName,omitempty"`
	// CountryCode 国家code
	CountryCode string `json:"countryCode,omitempty"`
	// Country 国家
	Country string `json:"country,omitempty"`
	// Email 买家邮箱
	Email string `json:"email,omitempty"`
}

// OrderCustoms 跨境报关信息
type OrderCustoms struct {
	ID          uint64          `json:"id,omitempty"`
	GmtCreate   go1688.JsonTime `json:"gmtCreate,omitempty"`
	GmtModified go1688.JsonTime `json:"gmtModified,omitempty"`
	BuyerID     uint64          `json:"buyerId,omitempty"`
	OrderID     uint64          `json:"orderId,omitempty"`
	// Type 业务数据类型,默认1：报关单
	Type int `json:"type,omitempty"`
	// Attributes 报关信息列表
	Attributes []CustomerAttributeInfo `json:"attributes,omitempty"`
}

// CustomerAttributeInfo 报关信息
type CustomerAttributeInfo struct {
	// Sku sku标识
	Sku      string  `json:"sku,omitempty"`
	CnName   string  `json:"cnName,omitempty"`
	EnName   string  `json:"enName,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Quantity float64 `json:"quantify,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// CaigouQuoteInfo 采购单详情
type CaigouQuoteInfo struct {
	// ProductQuoteName 供应单项的名称
	ProductQuoteName string `json:"productQuoteName,omitempty"`
	// Price 价格，单位：元
	Price decimal.Decimal `json:"price,omitempty"`
	// Count 购买数量
	Count float64 `json:"count,omitempty"`
}

// KeyValuePair 扩展属性
type KeyValuePair struct {
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

// TradeGetBuyerView 订单详情查看(买家视角)
func TradeGetBuyerView(client *go1688.Client, req *TradeGetBuyerViewRequest, accessToken string) (*TradeInfo, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp TradeGetBuyerViewResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
