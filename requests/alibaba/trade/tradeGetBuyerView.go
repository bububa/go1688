package trade

import (
	"encoding/json"

	"github.com/bububa/go1688"
	"github.com/shopspring/decimal"
)

type TradeGetBuyerViewRequest struct {
	Website       APIWebsite `json:"webSite"`                 // 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
	OrderId       uint64     `json:"orderId"`                 // 交易id，订单号
	IncludeFields string     `json:"includeFields,omitempty"` // 查询结果中包含的域，GuaranteesTerms：保障条款，NativeLogistics：物流信息，RateDetail：评价详情，OrderInvoice：发票信息。默认返回GuaranteesTerms、NativeLogistics、OrderInvoice。
	AttributeKeys []string   `json:"attributeKeys,omitempty"` // 垂直表中的attributeKeys
}

type TradeGetBuyerViewRefinedRequest struct {
	Website       APIWebsite `json:"webSite"`                 // 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
	OrderId       uint64     `json:"orderId"`                 // 交易id，订单号
	IncludeFields string     `json:"includeFields,omitempty"` // 查询结果中包含的域，GuaranteesTerms：保障条款，NativeLogistics：物流信息，RateDetail：评价详情，OrderInvoice：发票信息。默认返回GuaranteesTerms、NativeLogistics、OrderInvoice。
	AttributeKeys string     `json:"attributeKeys,omitempty"` // 垂直表中的attributeKeys
}

func (this *TradeGetBuyerViewRequest) Refine() *TradeGetBuyerViewRefinedRequest {
	var attributeKeys []byte
	if len(this.AttributeKeys) > 0 {
		attributeKeys, _ = json.Marshal(this.AttributeKeys)
	}
	return &TradeGetBuyerViewRefinedRequest{
		Website:       this.Website,
		OrderId:       this.OrderId,
		IncludeFields: this.IncludeFields,
		AttributeKeys: string(attributeKeys),
	}
}

func (this *TradeGetBuyerViewRefinedRequest) Name() string {
	return "alibaba.trade.cancel"
}

type TradeGetBuyerViewResponse struct {
	go1688.BaseResponse
	Result *TradeInfo `json:"result,omitempty"`
}

type TradeInfo struct {
	BaseInfo             *OrderBaseInfo        `json:"baseInfo,omitempty"`             // 订单基础信息
	BizInfo              *OrderBizInfo         `json:"orderBizInfo,omitempty"`         // 订单业务信息
	TradeTerms           []TradeTermInfo       `json:"tradeTerms,omitempty"`           // 交易条款
	ProductItems         []ProductItemInfo     `json:"productItems,omitempty"`         // 商品条目信息
	NativeLogistics      *NativeLogisticsInfo  `json:"nativeLogistics,omitempty"`      // 国内物流
	InvoiceInfo          *OrderInvoiceInfo     `json:"orderInvoiceInfo,omitempty"`     // 发票信息
	GuaranteesTerms      *GuaranteesTermInfo   `json:"guaranteesTerms,omitempty"`      // 保障条款
	OrderRateInfo        *OrderRateInfo        `json:"orderRateInfo,omitempty"`        // 订单评价信息
	OverseasExtraAddress *OverseasExtraAddress `json:"overseasExtraAddress,omitempty"` // 跨境地址扩展信息
	Customs              *OrderCustoms         `json:"customs,omitempty"`              // 跨境报关信息
	QuoteList            []CaigouQuoteInfo     `json:"quoteList,omitempty"`            //采购单详情列表，为大企业采购订单独有域。
	ExtAttributes        []KeyValuePair        `json:"extAttributes,omitempty"`        // 订单扩展属性
}

type OrderBaseInfo struct {
	AllDeliveredTime  go1688.JsonTime `json:"allDeliveredTime,omitempty"`  // 完全发货时间
	sellerCreditLevel string          `json:"sellerCreditLevel,omitempty"` // 卖家诚信等级
	PayTime           go1688.JsonTime `json:"payTime,omitempty"`           // 付款时间，如果有多次付款，这里返回的是首次付款时间
	Discount          uint            `json:"discount,omitempty"`          // 折扣信息，单位分
	AlipayTradeId     string          `json:"alipayTradeId,omitempty"`     // 外部支付交易Id
	SumProductPayment decimal.Decimal `json:"sumProductPayment,omitempty"` // 产品总金额(该订单产品明细表中的产品金额的和)，单位元
	BuyerFeedback     string          `json:"buyerFeedback,omitempty"`     // 买家留言，不超过500字
	FlowTemplateCode  string          `json:"flowTemplateCode,omitempty"`  // 4.0交易流程模板code
	SellerOrder       bool            `json:"sellerOrder,omitempty"`       // 是否自主订单（邀约订单）
	BuyerLoginId      string          `json:"buyerLoginId,omitempty"`      // 买家loginId，旺旺Id
	ModifyTime        go1688.JsonTime `json:"modifyTime,omitempty"`        // 修改时间
	SubBuyerLoginId   string          `json:"subBuyerLoginId,omitempty"`   // 买家子账号
	Id                uint64          `json:"id,omitempty"`                // 交易id
	CloseReason       string          `json:"closeReason,omitempty"`       //关闭原因。buyerCancel:买家取消订单，sellerGoodsLack:卖家库存不足，other:其它
	BuyerContact      *TradeContact   `json:"buyerContact,omitempty"`      // 买家联系人
	SellerAlipayId    string          `json:"sellerAlipayId,omitempty"`    // 卖家支付宝id
	CompleteTime      go1688.JsonTime `json:"completeTime,omitempty"`      // 完成时间
	SellerLoginId     string          `json:"sellerLoginId,omitempty"`     // 卖家oginId，旺旺Id
	BuyerId           string          `json:"buyerID,omitempty"`           // 买家主账号id
	CloseOperateType  TradeCloseType  `json:"closeOperateType,omitempty"`  // 关闭订单操作类型。CLOSE_TRADE_BY_SELLER:卖家关闭交易,CLOSE_TRADE_BY_BOPS:BOPS后台关闭交易,CLOSE_TRADE_BY_SYSTEM:系统（超时）关闭交易,CLOSE_TRADE_BY_BUYER:买家关闭交易,CLOSE_TRADE_BY_CREADIT:诚信保障投诉关闭
	TotalAmount       decimal.Decimal `json:"totalAmount,omitempty"`       // 应付款总金额，totalAmount = ∑itemAmount + shippingFee，单位为元
	SellerId          string          `json:"sellerID,omitempty"`          // 卖家主账号id
	ShippingFee       decimal.Decimal `json:"shippingFee,omitempty"`       // 运费，单位为元
	BuyerUserId       uint64          `json:"buyerUserId,omitempty"`       // 买家数字id
	BuyerMemo         string          `json:"buyerMemo,omitempty"`         // 买家备忘信息
	Refund            decimal.Decimal `json:"refund,omitempty"`            // 退款金额，单位为元
	Status            TradeStatus     `json:"status,omitempty"`            // 交易状态，waitbuyerpay:等待买家付款;waitsellersend:等待卖家发货;waitbuyerreceive:等待买家收货;confirm_goods:已收货;success:交易成功;cancel:交易取消;terminated:交易终止;未枚举:其他状态
	RefundPayment     uint            `json:"refundPayment,omitempty"`     // 退款金额
	SellerContact     *TradeContact   `json:"sellerContact,omitempty"`     // 卖家联系人信息
	RefundStatus      RefundStatus    `json:"refundStatus,omitempty"`      // 订单的售中退款状态，等待卖家同意：waitselleragree ，待买家修改：waitbuyermodify，等待买家退货：waitbuyersend，等待卖家确认收货：waitsellerreceive，退款成功：refundsuccess，退款失败：refundclose
	Remark            string          `json:"remark,omitempty"`            // 备注，1688指下单时的备注
	PreOrderId        uint64          `json:"preOrderId,omitempty"`        // 预订单ID
	ConfirmedTime     go1688.JsonTime `json:"confirmedTime,omitempty"`     // 确认时间
	CloseRemark       string          `json:"closeRemark,omitempty"`       // 关闭订单备注
	TradeType         string          `json:"tradeType,omitempty"`         // 1:担保交易 2:预存款交易 3:ETC境外收单交易 4:即时到帐交易 5:保障金安全交易 6:统一交易流程 7:分阶段付款 8.货到付款交易 9.信用凭证支付交易 10.账期支付交易，50060 交易4.0
	ReceivingTime     go1688.JsonTime `json:"receivingTime,omitempty"`     // 收货时间，这里返回的是完全收货时间
	StepAgreementPath string          `json:"stepAgreementPath,omitempty"` // 分阶段法务协议地址
	IdOfStr           string          `json:"idOfStr,omitempty"`           // 交易id(字符串格式)
	RefundStatusForAs string          `json:"refundStatusForAs,omitempty"` // 订单的售后退款状态
	StepPayAll        bool            `json:"stepPayAll,omitempty"`        // 是否一次性付款
	SellerUserId      uint64          `json:"sellerUserId,omitempty"`      // 卖家数字id
	StepOrderList     []StepOrder     `json:"stepOrderList,omitempty"`     // [交易3.0]分阶段交易，分阶段订单list
	BuyerAlipayId     string          `json:"buyerAlipayId,omitempty"`     // 买家支付宝id
	CreateTime        go1688.JsonTime `json:"createTime,omitempty"`        // 创建时间
	BusinessType      BusinessType    `json:"businessType,omitempty"`      // 业务类型。国际站：ta(信保),wholesale(在线批发)。 中文站：普通订单类型 = "cn"; 大额批发订单类型 = "ws"; 普通拿样订单类型 = "yp"; 一分钱拿样订单类型 = "yf"; 倒批(限时折扣)订单类型 = "fs"; 加工定制订单类型 = "cz"; 协议采购订单类型 = "ag"; 伙拼订单类型 = "hp"; 供销订单类型 = "supply"; 淘工厂订单 = "factory"; 快订下单 = "quick"; 享拼订单 = "xiangpin"; 当面付 = "f2f"; 存样服务 = "cyfw"; 代销订单 = "sp"; 微供订单 = "wg";零售通 = "lst";跨境='cb';分销='distribution';采源宝='cab';加工定制="manufact"
	OverSeaOrder      bool            `json:"overSeaOrder,omitempty"`      // 是否海外代发订单，是：true
	RefundId          string          `json:"refundId,omitempty"`          // 退款单ID
	TradeTypeDesc     string          `json:"tradeTypeDesc,omitempty"`     // 下单时指定的交易方式
	PayChannelList    []string        `json:"payChannelList,omitempty"`    // 支付渠道名称列表。一笔订单可能存在多种支付渠道。枚举值：支付宝,网商银行信任付,诚e赊,对公转账,赊销宝,账期支付,合并支付渠道,支付平台,声明付款,网商电子银行承兑汇票,银行转账,跨境宝,红包,其它
	TradeTypeCode     string          `json:"tradeTypeCode,omitempty"`     // 下单时指定的交易方式tradeType
	PayTimeout        int64           `json:"payTimeout,omitempty"`        // 支付超时时间，定长情况时单位：秒，目前都是定长
	PayTimeoutType    uint            `json:"payTimeoutType,omitempty"`    // 支付超时TYPE，0：定长，1：固定时间
}

type TradeContact struct {
	Phone         string `json:"phone,omitempty"`         // 联系电话
	Fax           string `json:"fax,omitempty"`           // 传真
	Email         string `json:"email,omitempty"`         // 邮箱
	ImInPlatform  string `json:"imInPlatform,omitempty"`  //联系人在平台的IM账号
	Name          string `json:"name,omitempty"`          // 联系人名称
	Mobile        string `json:"mobile,omitempty"`        // 联系人手机号
	CompanyName   string `json:"companyName,omitempty"`   // 公司名称
	WgSenderName  string `json:"wgSenderName,omitempty"`  // 发件人名称，在微供等分销场景下由分销商设置
	WgSenderPhone string `json:"wgSenderPhone,omitempty"` // 发件人电话，在微供等分销场景下由分销商设置
}

type StepOrder struct {
	Id                   uint64          `json:"stepOrderId,omitempty"`          // 阶段id
	Status               StepOrderStatus `json:"stepOrderStatus,omitempty"`      // waitactivate 未开始（待激活） waitsellerpush 等待卖家推进 success 本阶段完成 settlebill 分账 cancel 本阶段终止 inactiveandcancel 本阶段未开始便终止 waitbuyerpay 等待买家付款 waitsellersend 等待卖家发货 waitbuyerreceive 等待买家确认收货 waitselleract 等待卖家XX操作 waitbuyerconfirmaction 等待买家确认XX操作
	PayStatus            uint            `json:"stepPayStatus,omitempty"`        // 1 未冻结/未付款 2 已冻结/已付款 4 已退款 6 已转交易 8 交易未付款被关闭
	No                   uint            `json:"stepNo,omitepty"`                // 阶段序列：1、2、3...
	LastStep             bool            `json:"lastStep,omitempty"`             // 是否最后一个阶段
	HasDisbursed         bool            `json:"hasDisbursed,omitempty"`         // 是否已打款给卖家
	PayFee               decimal.Decimal `json:"payFee,omitempty"`               // 创建时需要付款的金额，不含运费
	ActualPayFee         decimal.Decimal `json:"actualPayFee,omitempty"`         // 应付款（含运费）= 单价×数量-单品优惠-店铺优惠+运费+修改的金额（除运费外，均指分摊后的金额）
	DiscountFee          decimal.Decimal `json:"discountFee,omitempty"`          // 本阶段分摊的店铺优惠
	ItemDiscountFee      decimal.Decimal `json:"itemDiscountFee,omitempty"`      // 本阶段分摊的单品优惠
	Price                decimal.Decimal `json:"price,omitempty"`                // 本阶段分摊的单价
	Amount               uint            `json:"amount,omitempty"`               // 购买数量
	PostFee              decimal.Decimal `json:"postFee,omitempty"`              // 运费
	AdjustPostFee        decimal.Decimal `json:"adjustPostFee,omitempty"`        // 修改价格修改的金额
	GmtCreate            go1688.JsonTime `json:"gmtCreate,omitempty"`            // 创建时间
	GmtModified          go1688.JsonTime `json:"gmtModified,omitempty"`          // 修改时间
	EnterTime            go1688.JsonTime `json:"enterTime,omitempty"`            // 开始时间
	PayTime              go1688.JsonTime `json:"payTime,omitempty"`              // 付款时间
	SellerActionTime     go1688.JsonTime `json:"sellerActionTime,omitempty"`     // 卖家操作时间
	EndTime              go1688.JsonTime `json:"endTime,omitempty"`              // 本阶段结束时间
	MessagePath          string          `json:"messagePath,omitempty"`          // 卖家操作留言路径
	PicturePath          string          `json:"picturePath,omitempty"`          // 卖家上传图片凭据路径
	Message              string          `json:"message,omitempty"`              // 卖家操作留言
	TemplateId           uint64          `json:"templateId,omitempty"`           // 使用的模板id
	Name                 string          `json:"stepName,omitempty"`             // 当前步骤的名称
	SellerActionName     string          `json:"sellerActionName,omitempty"`     // 卖家操作名称
	BuyerPayTimeout      int64           `json:"buyerPayTimeout,omitempty"`      // 买家不付款的超时时间(秒)
	BuyerConfirmTimeout  int64           `json:"buyerConfirmTimeout,omitempty"`  // 买家不确认的超时时间
	NeedLogistics        bool            `json:"needLogistics,omitempty"`        // 是否需要物流
	NeedSellerAction     bool            `json:"needSellerAction,omitempty"`     // 是否需要卖家操作和买家确认
	TransferAfterConfirm bool            `json:"transferAfterConfirm,omitempty"` // 阶段结束是否打款
	NeedSellerCallNext   bool            `json:"needSellerCallNext,omitempty"`   // 是否需要卖家推进
	InstantPay           bool            `json:"instantPay,omitempty"`           // 是否允许即时到帐
}

type OrderBizInfo struct {
	OdsCyd            bool               `json:"odsCyd,omitempty"`            // 是否采源宝订
	AccountPeriodTime string             `json:"accountPeriodTime,omitempty"` // 账期交易订单的到账时间
	CreditOrder       bool               `json:"creditOrder,omitempty"`       // 为true，表示下单时选择了诚e赊交易方式。注意不等同于“诚e赊支付”，支付时有可能是支付宝付款，具体支付方式查询tradeTerms.payWay
	CreditOrderDetail *CreditOrderDetail `json:"creditOrderDetail,omitempty"` //诚e赊支付详情，只有使用诚e赊付款时返回
	PreOrderInfo      *PreOrderInfo      `json:"preOrderInfo,omitempty"`      // 预订单信息
	LstOrderInfo      *LstOrderInfo      `json:"lstOrderInfo,omitempty"`
}

type CreditOrderDetail struct {
	PayAmount          uint   `json:"payAmount,omitempty"`          // 订单金额
	CreateTime         string `json:"createTime,omitempty"`         // 支付时间
	Status             string `json:"status,omitempty"`             // 状态
	GracePeriodEndTime string `json:"gracePeriodEndTime,omitempty"` // 最晚还款时间
	StatusStr          string `json:"statusStr,omitempty"`          // 状态描述
	RestRepayAmount    uint   `json:"restRepayAmount,omitempty"`    // 应还金额
}

type PreOrderInfo struct {
	MarketName        string `json:"marketName,omitempty"`        // 创建预订单时传入的市场名
	CreatePreOrderApp bool   `json:"createPreOrderApp,omitempty"` // 预订单是否为当前查询的通过当前查询的ERP创建
}

type LstOrderInfo struct {
	LstWarehouseType LstWarehouseType `json:"lstWarehouseType,omitempty"` // 零售通仓库类型。customer：虚仓；cainiao：实仓
}

type TradeTermInfo struct {
	PayStatus      string          `json:"payStatus,omitempty"`      // 支付状态。国际站：WAIT_PAY(未支付),PAYER_PAID(已完成支付),PART_SUCCESS(部分支付成功),PAY_SUCCESS(支付成功),CLOSED(风控关闭),CANCELLED(支付撤销),SUCCESS(成功),FAIL(失败)。 1688:1(未付款);2(已付款);4(全额退款);6(卖家有收到钱，回款完成) ;7(未创建外部支付单);8 (付款前取消) ; 9(正在支付中);12(账期支付,待到账)
	PayTime        go1688.JsonTime `json:"payTime,omitempty"`        // 完成阶段支付时间
	PayWay         string          `json:"payWay,omitempty"`         // 支付方式。 国际站：ECL(融资支付),CC(信用卡),TT(线下TT),ACH(echecking支付)。 1688:1-支付宝,2-网商银行信任付,3-诚e赊,4-银行转账,5-赊销宝,6-电子承兑票据,7-账期支付,8-合并支付渠道,9-无打款,10-零售通赊购,13-支付平台,12-声明付款
	PhasAmount     decimal.Decimal `json:"phasAmount,omitempty"`     // 付款额
	Phase          uint64          `json:"phase,omitempty"`          // 阶段单id
	PhaseCondition string          `json:"phaseCondition,omitempty"` // 阶段条件，1688无此内容
	PhaseDate      string          `json:"phaseDate,omitempty"`      // 阶段时间，1688无此内容
	CardPay        bool            `json:"cardPay,omitempty"`        // 是否银行卡支付
	ExpressPay     bool            `json:"expressPay,omitempty"`     // 是否快捷支付
	PayWayDesc     string          `json:"payWayDesc,omitempty"`     // 支付方式
}

type ProductItemInfo struct {
	CargoNumber        string               `json:"cargoNumber,omitempty"`        // 指定单品货号，国际站无需关注。该字段不一定有值，仅仅在下单时才会把货号记录(如果卖家设置了单品货号的话)。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。
	Description        string               `json:"description,omitempty"`        // 描述,1688无此信息
	ItemAmount         decimal.Decimal      `json:"itemAmount,omitempty"`         // 实付金额，单位为元
	Name               string               `json:"name,omitempty"`               // 商品名称
	Price              decimal.Decimal      `json:"price,omitempty"`              // 原始单价，以元为单位
	Id                 uint64               `json:"productID,omitempty"`          // 产品ID（非在线产品为空）
	ImgUrl             []string             `json:"productImgUrl,omitempty"`      // 商品图片url
	SnapshotUrl        string               `json:"productSnapshotUrl,omitempty"` // 产品快照url，交易订单产生时会自动记录下当时的商品快照，供后续纠纷时参考
	Quantity           decimal.Decimal      `json:"quantity,omitempty"`           // 以unit为单位的数量，例如多少个、多少件、多少箱、多少吨
	Refund             decimal.Decimal      `json:"refund,omitempty"`             // 退款金额，单位为元
	SkuId              uint64               `json:"skuID,omitempty"`              // skuID
	Sort               uint                 `json:"sort,omitempty"`               // 排序字段，商品列表按此字段进行排序，从0开始，1688不提供
	Status             string               `json:"status,omitempty"`             // 子订单状态
	SubItemId          uint64               `json:"subItemId,omitempty"`          // 子订单号，或商品明细条目ID
	Type               string               `json:"type,omitempty"`               // 类型，国际站使用，供卖家标注商品所属类型
	Unit               string               `json:"unit,omitempty"`               // 售卖单位 例如：个、件、箱、吨
	Weight             string               `json:"weight,omitempty"`             // 重量 按重量单位计算的重量，例如：100
	WeightUnit         string               `json:"weightUnit,omitempty"`         // 重量单位 例如：g，kg，t
	GuaranteesTerms    []GuaranteesTermInfo `json:"guaranteesTerms,omitempty"`    // 保障条款，此字段仅针对1688
	ProductCargoNumber string               `json:"productCargoNumber,omitempty"` // 指定商品货号，该字段不一定有值，在下单时才会把货号记录。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。该字段和cargoNUmber的区别是：该字段是定义在商品级别上的货号，cargoNUmber是定义在单品级别的货号
	SkuInfos           []SkuItemDesc        `json:"skuInfos,omitempty"`
	EntryDiscount      int                  `json:"entryDiscount,omitempty"`    // 订单明细涨价或降价的金额
	SpecId             string               `json:"specId,omitempty"`           // 订单销售属性ID
	QuantityFactor     decimal.Decimal      `json:"quantityFactor,omitempty"`   // 以unit为单位的quantity精度系数，值为10的幂次，例如:quantityFactor=1000,unit=吨，那么quantity的最小精度为0.001吨
	StatusStr          string               `json:"statusStr,omitempty"`        // 子订单状态描述
	RefundStatus       string               `json:"refundStatus,omitempty"`     // WAIT_SELLER_AGREE 等待卖家同意 REFUND_SUCCESS 退款成功 REFUND_CLOSED 退款关闭 WAIT_BUYER_MODIFY 待买家修改 WAIT_BUYER_SEND 等待买家退货 WAIT_SELLER_RECEIVE 等待卖家确认收货
	CloseReason        string               `json:"closeReason,omitempty"`      // 关闭原因
	LogisticsStatus    uint                 `json:"logisticsStatus,omitempty"`  // 1 未发货 2 已发货 3 已收货 4 已经退货 5 部分发货 8 还未创建物流订单
	GmtCreate          go1688.JsonTime      `json:"gmtCreate,omitempty"`        // 创建时间
	GmtModified        go1688.JsonTime      `json:"gmtModified,omitempty"`      // 修改时间
	GmtCompleted       go1688.JsonTime      `json:"gmtCompleted,omitempty"`     // 明细完成时间
	GmtPayExpireTime   string               `json:"gmtPayExpireTime,omitempty"` // 库存超时时间，格式为“yyyy-MM-dd HH:mm:ss”
	RefundId           string               `json:"refundId,omitempty"`         // 售中退款单号
	SubItemIdString    string               `json:"subItemIDString,omitempty"`  // 子订单号，或商品明细条目ID(字符串类型，由于Long类型的ID可能在JS和PHP中处理有问题，所以可以用字符串类型来处理)
	RefundIdForAs      string               `json:"refundIdForAs,omitempty"`    // 售后退款单号
}

type SkuItemDesc struct {
	Name  string `json:"name,omitempty"`  // 属性名
	Value string `json:"value,omitempty"` // 属性值
}

type GuaranteesTermInfo struct {
	AssuranceInfo        string               `json:"assuranceInfo,omitempty"`        // 保障条款
	AssuranceType        string               `json:"assuranceType,omitempty"`        // 保障方式。国际站：TA(信保)
	QualityAssuranceType QualityAssuranceType `json:"qualityAssuranceType,omitempty"` // 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后)
}

type NativeLogisticsInfo struct {
	Address        string                    `json:"address,omitempty"`
	Area           string                    `json:"area,omitempty"`
	AreaCode       string                    `json:"areaCode,omitempty"`
	City           string                    `json:"city,omitempty"`
	ContactPerson  string                    `json:"contactPerson,omitempty"`
	Fax            string                    `json:"fax,omitempty"`
	Mobile         string                    `json:"mobile,omitempty"`
	Province       string                    `json:"province,omitempty"`
	Telephone      string                    `json:"telephone,omitempty"`
	Zip            string                    `json:"zip,omitempty"`
	LogisticsItems []NativeLogisticsItemInfo `json:"logisticsItems,omitempty"` // 运单明细
	TownCode       string                    `json:"townCode,omitempty"`       // 镇，街道地址码
	Town           string                    `json:"town,omitempty"`           // 镇，街道
}

type NativeLogisticsItemInfo struct {
	DeliveredTime        go1688.JsonTime `json:"deliveredTime,omitempty"`        // 发货时间
	LogisticsCode        string          `json:"logisticsCode,omitempty"`        // 物流编号
	Type                 string          `json:"type,omitempty"`                 // SELF_SEND_GOODS("0")自行发货，在线发货ONLINE_SEND_GOODS("1"，不需要物流的发货 NO_LOGISTICS_SEND_GOODS("2")
	Id                   uint64          `json:"id,omitempty"`                   // 主键id
	Status               string          `json:"status,omitempty"`               // 状态
	GmtModified          go1688.JsonTime `json:"gmtModified,omitempty"`          // 修改时间
	GmtCreate            go1688.JsonTime `json:"gmtCreate,omitempty"`            // 创建时间
	Carriage             decimal.Decimal `json:"carriage,omitempty"`             // 运费(单位为元)
	FromProvince         string          `json:"fromProvince,omitempty"`         // 发货省
	FromCity             string          `json:"fromCity,omitempty"`             // 发货市
	FromArea             string          `json:"fromArea,omitempty"`             // 发货区
	FromAddress          string          `json:"fromAddress,omitempty"`          // 发货街道地址
	FromPhone            string          `json:"fromPhone,omitempty"`            // 发货联系电话
	FromMobile           string          `json:"fromMobile,omitempty"`           // 发货联系手机
	FromPost             string          `json:"fromPost,omitempty"`             // 发货地址邮编
	LogisticsCompanyId   uint64          `json:"logisticsCompanyId,omitempty"`   // 物流公司Id
	LogisticesCompanyNo  string          `json:"logisticsCompanyNo,omitempty"`   // 物流公司编号
	LogisticsCompanyName string          `json:"logisticsCompanyName,omitempty"` // 物流公司名称
	LogisticsBillNo      string          `json:"logisticsBillNo,omitempty"`      // 物流公司运单号
	SubItemIds           string          `json:"subItemIds,omitempty"`           // 商品明细条目id，如有多个以,分隔
	ToProvince           string          `json:"toProvince,omitempty"`           // 收货省
	ToCity               string          `json:"toCity,omitempty"`               // 收货市
	ToArea               string          `json:"toArea,omitempty"`               // 收货区
	ToAddress            string          `json:"toAddress,omitempty"`            // 收货街道地址
	ToPhone              string          `json:"toPhone,omitempty"`              // 收货联系电话
	ToMobile             string          `json:"toMobile,omitempty"`             // 收货联系手机
	ToPost               string          `json:"toPost,omitempty"`               // 收货地址邮编
	NoLogisticsName      string          `json:"noLogisticsName,omitempty"`      // 物流姓名
	NoLogisticsTel       string          `json:"noLogisticsTel,omitempty"`       // 联系方式
	NoLogisticsBillNo    string          `json:"noLogisticsBillNo,omitempty"`    // 无需物流业务单号
	NoLogisticsCondition string          `json:"noLogisticsCondition,omitempty"` // 无需物流类别,noLogisticsCondition=1， 表示其他第三方物流、小型物充商、车队等, noLogisticsCondition=2 表示补运费、差价, noLogisticsCondition=3 表示卖家配送, noLogisticsCondition=4 表示买家自提 noLogisticsCondition=5 表示其他原因
	IsTimePromise        bool            `json:"isTimePromise,omitempty"`        // 是否使用限时达物流
	ArriveTime           go1688.JsonTime `json:"arriveTime,omitempty"`           // 限时达物流，预计到达时间
}

type OrderInvoiceInfo struct {
	InvoiceCompanyName string `json:"invoiceCompanyName,omitempty"` // 发票公司名称(即发票抬头-title)
	InvoiceType        uint   `json:"invoiceType,omitempty"`        // 发票类型. 0：普通发票，1:增值税发票，9未知类型
	LocalInvoiceId     uint64 `json:"localInvoiceId,omitempty"`     // 本地发票号
	OrderId            uint64 `json:"orderId,omitempty"`            // 订单Id
	ReceiveCode        string `json:"receiveCode,omitempty"`        // (收件人)址区域编码
	ReceiveCodeText    string `json:"receiveCodeText,omitempty"`    // (收件人) 省市区编码对应的文案(增值税发票信息)
	ReceiveMobile      string `json:"receiveMobile,omitempty"`      // （收件者）发票收货人手机
	ReceiveName        string `json:"receiveName,omitempty"`        // （收件者）发票收货人
	ReceivePhone       string `json:"receivePhone,omitempty"`       // （收件者）发票收货人电话
	ReceivePost        string `json:"receivePost,omitempty"`        // （收件者）发票收货地址邮编
	ReceiveStreet      string `json:"receiveStreet,omitempty"`      // (收件人) 街道地址(增值税发票信息)
	RegisterAccountId  string `json:"registerAccountId,omitempty"`  // (公司)银行账号
	RegisterBank       string `json:"registerBank,omitempty"`       // (公司)开户银行
	RegisterCode       string `json:"registerCode,omitempty"`       // (注册)省市区编码
	RegisterCodeText   string `json:"registerCodeText,omitempty"`   // (注册)省市区文本
	RegisterPhone      string `json:"registerPhone,omitempty"`      // （公司）注册电话
	RegisterStreet     string `json:"registerStreet,omitempty"`     // (注册)街道地址
	TaxPayerIdentify   string `json:"taxpayerIdentify,omitempty"`   // 纳税人识别号
}

type OrderRateInfo struct {
	BuyerRateStatus  uint              `json:"buyerRateStatus,omitempty"`  // 买家评价状态(4:已评论,5:未评论,6;不需要评论)
	SellerRateStatus uint              `json:"sellerRateStatus,omitempty"` // 卖家评价状态(4:已评论,5:未评论,6;不需要评论)
	BuyerRateList    []OrderRateDetail `json:"buyerRateList,omitempty"`    // 卖家給买家的评价
	SellerRateList   []OrderRateDetail `json:"sellerRateList,omitempty"`   // 买家給卖家的评价
}

type OrderRateDetail struct {
	StarLevel    int             `json:"starLevel,omitempty"`    // 评价星级
	Content      string          `json:"content,omitempty"`      // 评价详情
	ReceiverNick string          `json:"receiverNick,omitempty"` // 收到评价的用户昵称
	PosterNick   string          `json:"posterNick,omitempty"`   // 发送评价的用户昵称
	PublishTime  go1688.JsonTime `json:"publishTime,omitempty"`  // 评价上线时间
}

type OverseasExtraAddress struct {
	ChannelName         string `json:"channelName,omitempty"`         // 路线名称
	ChannelId           string `json:"channelId,omitempty"`           // 路线id
	ShippingCompanyId   string `json:"shippingCompanyId,omitempty"`   // 货代公司id
	ShippingCompanyName string `json:"shoppingCompanyName,omitempty"` // 货代公司名称
	CountryCode         string `json:"countryCode,omitempty"`         // 国家code
	Country             string `json:"country,omitempty"`             // 国家
	Email               string `json:"email,omitempty"`               // 买家邮箱
}

type OrderCustoms struct {
	Id          uint64                  `json:"id,omitempty"`
	GmtCreate   go1688.JsonTime         `json:"gmtCreate,omitempty"`
	GmtModified go1688.JsonTime         `json:"gmtModified,omitempty"`
	BuyerId     uint64                  `json:"buyerId,omitempty"`
	OrderId     uint64                  `json:"orderId,omitempty"`
	Type        int                     `json:"type,omitempty"`       // 业务数据类型,默认1：报关单
	Attributes  []CustomerAttributeInfo `json:"attributes,omitempty"` // 报关信息列表
}

type CustomerAttributeInfo struct {
	Sku      string  `json:"sku,omitempty"` // sku标识
	CnName   string  `json:"cnName,omitempty"`
	EnName   string  `json:"enName,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Quantity float64 `json:"quantify,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type CaigouQuoteInfo struct {
	ProductQuoteName string          `json:"productQuoteName,omitempty"` // 供应单项的名称
	Price            decimal.Decimal `json:"price,omitempty"`            // 价格，单位：元
	Count            float64         `json:"count,omitempty"`            // 购买数量
}

type KeyValuePair struct {
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

func TradeGetBuyerView(client *go1688.Client, req *TradeGetBuyerViewRequest, accessToken string) (*TradeInfo, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &TradeGetBuyerViewResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Result, nil
}
