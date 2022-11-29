package trade

import (
	"strconv"

	"github.com/bububa/go1688"
	"github.com/bububa/go1688/requests/alibaba/logistics"
)

// RefundBuyerQueryOrderRefundListRequest 订单列表查看(买家视角) API Request
type RefundBuyerQueryOrderRefundListRequest struct {
	// OrderID 订单Id
	OrderID uint64 `json:"orderId,omitempty"`
	// ApplyStartTime 退款申请时间（起始）
	ApplyStartTime go1688.JsonTime `json:"applyStartTime,omitempty"`
	// ApplyEndTime 退款申请时间（截止）
	ApplyEndTime go1688.JsonTime `json:"applyEndTime,omitempty"`
	// RefundStatusSet 退款状态列表
	RefundStatusSet []RefundStatus `json:"refundStatusSet,omitempty"`
	// SellerMemberID 卖家memberId
	SellerMemberID string `json:"sellerMemberId,omitempty"`
	// Page 当前页码
	Page int `json:"currentPageNum,omitempty"`
	// PageSize 每页条数
	PageSize int `json:"pageSize,omitempty"`
	// LogisticsNo 退货物流单号（传此字段查询时，需同时传入sellerMemberId）
	LogisticsNo string `json:"logisticsNo,omitempty"`
	// ModifyStartTime 退款修改时间(起始)
	ModifyStartTime go1688.JsonTime `json:"modifyStartTime,omitempty"`
	// ModifyEndTime 退款修改时间(截止)
	ModifyEndTime go1688.JsonTime `json:"modifyEndTime,omitempty"`
	// DisputeType 1:售中退款，2:售后退款；0:所有退款单
	DisputeType int `json:"disputeType,omitempty"`
}

// Name implement RequestData interface
func (r RefundBuyerQueryOrderRefundListRequest) Name() string {
	return "alibaba.trade.refund.buyer.queryOrderRefundList"
}

// Map implement RequestData interface
func (r RefundBuyerQueryOrderRefundListRequest) Map() map[string]string {
	ret := make(map[string]string, 16)
	if r.OrderID > 0 {
		ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	}
	if !r.ApplyStartTime.IsZero() {
		ret["applyStartTime"] = r.ApplyStartTime.Format()
	}
	if !r.ApplyEndTime.IsZero() {
		ret["applyEndTime"] = r.ApplyEndTime.Format()
	}
	if !r.ModifyStartTime.IsZero() {
		ret["modifyStartTime"] = r.ModifyStartTime.Format()
	}
	if !r.ModifyEndTime.IsZero() {
		ret["modifyEndTime"] = r.ModifyEndTime.Format()
	}
	if r.SellerMemberID != "" {
		ret["sellerMemberId"] = r.SellerMemberID
	}
	if r.LogisticsNo != "" {
		ret["logisticsNo"] = r.LogisticsNo
	}
	if len(r.RefundStatusSet) > 0 {
		ret["refundStatusSet"] = go1688.JSONMarshal(r.RefundStatusSet)
	}
	if r.DisputeType > 0 {
		ret["disputeType"] = strconv.Itoa(r.DisputeType)
	}
	if r.Page > 1 {
		ret["currentPageNum"] = strconv.Itoa(r.Page)
	}
	if r.PageSize > 1 {
		ret["pageSize"] = strconv.Itoa(r.PageSize)
	}
	return ret
}

// RefundBuyerQueryOrderRefundListResponse 订单列表查看(买家视角) API Response
type RefundBuyerQueryOrderRefundListResponse struct {
	go1688.BaseResponse
	// Result 查询返回列表
	Result *OpQueryOrderRefundListResult `json:"result,omitempty"`
}

// OpQueryOrderRefundListResult 查询结果
type OpQueryOrderRefundListResult struct {
	// List 退款单列表
	List []OpOrderRefundModel `json:"opOrderRefundModels,omitempty"`
	// TotalCount 符合条件总的记录条数
	TotalCount int `json:"totalCount,omitempty"`
	// CurrentPageNum 查询的当前页码
	CurrentPageNum int `json:"currentPageNum,omitempty"`
}

// OpOrderRefundModel 退款单
type OpOrderRefundModel struct {
	// AftersaleAgreeTimeout 售后超时标记
	AftersaleAgreeTimeout bool `json:"aftersaleAgreeTimeout,omitempty"`
	// AftersaleAutoDisburse 售后自动打款
	AftersaleAutoDisburse bool `json:"aftersaleAutoDisburse,omitempty"`
	// AlipayPaymentID 支付宝交易号
	AlipayPaymentID string `json:"alipayPaymentId,omitempty"`
	// ApplyCarriage 运费的申请退款金额，单位：分
	ApplyCarriage int64 `json:"applyCarriage,omitempty"`
	// ApplyExpect 买家原始输入的退款金额(可以为空)
	ApplyExpect int64 `json:"applyExpect,omitempty"`
	// ApplyPayment 买家申请退款金额，单位：分
	ApplyPayment int64 `json:"applyPayment,omitempty"`
	// ApplyReason 申请原因
	ApplyReason string `json:"applyReason,omitempty"`
	// ApplyReasonID 申请原因ID
	ApplyReasonID uint64 `json:"applyReasonId,omitempty"`
	// ApplySubReason 二级退款原因
	ApplySubReason string `json:"applySubReason,omitempty"`
	// ApplySubReasonID 二级退款原因Id
	ApplySubReasonID uint64 `json:"applySubReasonId,omitempty"`
	// AsynErrCode
	AsynErrCode string `json:"asynErrCode,omitempty"`
	// AsynSubErrCode
	AsynSubErrCode string `json:"asynSubErrCode,omitempty"`
	// BuyerAlipayID 买家支付宝ID
	BuyerAlipayID string `json:"buyerAlipayId,omitempty"`
	// BuyerLogisticsName 买家退货物流公司名
	BuyerLogisticsName string `json:"buyerLogisticsName,omitempty"`
	// BuyerMemberID 买家会员ID
	BuyerMemberID string `json:"buyerMemberId,omitempty"`
	// BuyerSendGoods 买家是否已经发货（如果有退货的流程）
	BuyerSendGoods bool `json:"buyerSendGoods,omitempty"`
	// BuyerUserID 买家阿里帐号ID(包括淘宝帐号Id)
	BuyerUserID uint64 `json:"buyerUserId,omitempty"`
	// CanRefundPayment 最大能够退款金额，单位：分
	CanRefundPayment int64 `json:"canRefundPayment,omitempty"`
	// CrmModifyRefund 是否小二修改过退款单
	CrmModifyRefund bool `json:"crmModifyRefund,omitempty"`
	// DisburseChannel 极速到账打款渠道
	DisburseChannel string `json:"disburseChannel,omitempty"`
	// DisputeRequest 售后退款要求
	DisputeRequest int `json:"disputeRequest,omitempty"`
	// DisputeType 纠纷类型：售中退款 售后退款，默认为售中退款
	DisputeType int `json:"disputeType,omitempty"`
	// ExtInfo 扩展信息
	ExtInfo map[string]string `json:"extInfo,omitempty"`
	// FreightBill 运单号
	FreightBill string `json:"freightBill,omitempty"`
	// FrozenFund 实际冻结账户金额,单位：分
	FrozenFund int64 `json:"frozenFund,omitempty"`
	// GmtApply 申请退款时间
	GmtApply go1688.JsonTime `json:"gmtApply,omitempty"`
	// GmtCompleted 完成时间
	GmtCompleted go1688.JsonTime `json:"gmtCompleted,omitempty"`
	// GmtCreate 创建时间
	GmtCreate go1688.JsonTime `json:"gmtCreate,omitempty"`
	// GmtFreezed 该退款单超时冻结开始时间
	GmtFreezed go1688.JsonTime `json:"gmtFreezed,omitempty"`
	// GmtModified 修改时间
	GmtModified go1688.JsonTime `json:"gmtModified,omitempty"`
	// GmtTimeOut 该退款单超时完成的时间期限
	GmtTimeOut go1688.JsonTime `json:"gmtTimeOut,omitempty"`
	// GoodsReceived 买家是否已收到货
	GoodsReceived bool `json:"goodsReceived,omitempty"`
	// GoodsStatus 1：买家未收到货 2：买家已收到货 3：买家已退货
	GoodsStatus int `json:"goodsStatus,omitempty"`
	// ID 退款单编号
	ID uint64 `json:"id,omitempty"`
	// InstantRefundType 极速到账退款类型
	InstantRefundType int `json:"instantRefundType,omitempty"`
	// InsufficientAccount 交易4.0退款余额不足
	InsufficientAccount bool `json:"insufficientAccount,omitempty"`
	// InsufficientBail 极速到账退款保证金不足
	InsufficientBail bool `json:"insufficientBail,omitempty"`
	// NewRefundReturn 是否新流程创建的退款退货
	NewRefundReturn bool `json:"newRefundReturn,omitempty"`
	// OnlyRefund 是否仅退款
	OnlyRefund bool `json:"onlyRefund,omitempty"`
	// OrderEntryCountMap 子订单退货数量
	OrderEntryCountMap map[string]int64 `json:"orderEntryCountMap,omitempty"`
	// OrderEntryIDList 退款单包含的订单明细，时间逆序排列
	OrderEntryIDList []uint64 `json:"orderEntryIdList,omitempty"`
	// OrderID 退款单对应的订单编号
	OrderID uint64 `json:"orderId,omitempty"`
	// PrepaidBalance 极速退款垫资金额,该值不为空时,只代表该退款单可以走垫资流程,但不代表一定垫资成功
	PrepaidBalance int64 `json:"prepaidBalance,omitempty"`
	// ProductName 产品名称(退款单关联订单明细的货品名称)
	ProductName string `json:"productName,omitempty"`
	// RefundCarriage 运费的实际退款金额，单位：分
	RefundCarriage int64 `json:"refundCarriage,omitempty"`
	// RefundGoods 是否要求退货
	RefundGoods bool `json:"refundGoods,omitempty"`
	// RefundID 退款单逻辑主键
	RefundID string `json:"refundId,omitempty"`
	// RefundPayment 实际退款金额，单位：分
	RefundPayment int64 `json:"refundPayment,omitempty"`
	// RejectReason 卖家拒绝原因
	RejectReason string `json:"rejectReason,omitempty"`
	// RejectTimes 退款单被拒绝的次数
	RejectTimes int `json:"rejectTimes,omitempty"`
	// SellerAlipayID 卖家支付宝ID
	SellerAlipayID string `json:"sellerAlipayId,omitempty"`
	// SellerDelayDisburse 是否卖家延迟打款，即安全退款
	SellerDelayDisburse bool `json:"sellerDelayDisburse,omitempty"`
	// SellerMemberID 卖家会员ID
	SellerMemberID string `json:"sellerMemberId,omitempty"`
	// SellerMobile 收货人手机
	SellerMobile string `json:"sellerMobile,omitempty"`
	// SellerRealName 收货人姓名
	SellerRealName string `json:"sellerRealName,omitempty"`
	// SellerReceiveAddress 买家退货时卖家收货地址
	SellerReceiveAddress string `json:"sellerReceiveAddress,omitempty"`
	// SellerTel 收货人电话
	SellerTel string `json:"sellerTel,omitempty"`
	// SellerUserID 卖家阿里帐号ID(包括淘宝帐号Id)
	SellerUserID uint64 `json:"sellerUserId,omitempty"`
	// Status 退款状态
	Status string `json:"status,omitempty"`
	// SupportNewSteppay 是否支持交易4.0
	SupportNewSteppay bool `json:"supportNewSteppay,omitempty"`
	// TaskStatus 工单子状态，没有流到CRM创建工单时为空
	TaskStatus bool `json:"taskStatus,omitempty"`
	// TimeOutFreeze 是否超时系统冻结，true代表冻结，false代表不冻结
	TimeOutFreeze bool `json:"timeOutFreeze,omitempty"`
	// TimeOutOperateType 超时后执行的动作
	TimeOutOperateType string `json:"timeOutOperateType,omitempty"`
	// TradeTypeStr 交易类型，用来替换枚举类型的tradeType
	TradeTypeStr string `json:"tradeTypeStr,omitempty"`
	// Success 是否成功
	Success bool `json:"success,omitempty"`
	// RefundOperationList  操作记录列表
	RefundOperationList []OpOrderRefundOperationModal `json:"refundOperationList,omitempty"`
	// IsCrmModifyRefund 是否小二修改过退款单
	IsCrmModifyRefund bool `json:"isCrmModifyRefund,omitempty"`
	// IsTimeOutFreeze 是否超时系统冻结，true代表冻结，false代表不冻结
	IsTimeOutFreeze bool `json:"isTimeOutFreeze,omitempty"`
	// IsInsufficientAccount 交易4.0退款余额不足
	IsInsufficientAccount bool `json:"isInsufficientAccount,omitempty"`
	// IsGoodsReceived 买家是否已收到货
	IsGoodsReceived bool `json:"isGoodsReceived,omitempty"`
	// IsOnlyRefund 是否仅退款
	IsOnlyRefund bool `json:"isOnlyRefund,omitempty"`
	// IsRefundGoods 是否要求退货
	IsRefundGoods bool `json:"isRefundGoods,omitempty"`
	// IsSellerDelayDisburse 是否卖家延迟打款，即安全退款
	IsSellerDelayDisburse bool `json:"isSellerDelayDisburse,omitempty"`
	// IsAftersaleAutoDisburse 售后自动打款
	IsAftersaleAutoDisburse bool `json:"isAftersaleAutoDisburse,omitempty"`
	// IsSupportNewSteppay 是否支持交易4.0
	IsSupportNewSteppay bool `json:"isSupportNewSteppay,omitempty"`
	// IsNewRefundReturn 是否新流程创建的退款退货
	IsNewRefundReturn bool `json:"isNewRefundReturn,omitempty"`
	// IsBuyerSendGoods 买家是否已经发货（如果有退货的流程）
	IsBuyerSendGoods bool `json:"isBuyerSendGoods,omitempty"`
	// IsAftersaleAgreeTimeout 售后超时标记
	IsAftersaleAgreeTimeout bool `json:"isAftersaleAgreeTimeout,omitempty"`
	// IsInsufficientBail 极速到账退款保证金不足
	IsInsufficientBail bool `json:"isInsufficientBail,omitempty"`
}

// OpOrderRefundOperationModal 操作记录
type OpOrderRefundOperationModal struct {
	// AfterOperateStatus 操作后的退款状态
	AfterOperateStatus string `json:"afterOperateStatus,omitempty"`
	// BeforeOperateStatus 操作前的退款状态
	BeforeOperateStatus string `json:"beforeOperateStatus,omitempty"`
	// CloseRefundStepId 分阶段订单正向操作关闭退款时的阶段ID
	CloseRefundStepId uint64 `json:"closeRefundStepId,omitempty"`
	// CrmModifyRefund 是否小二修改过退款单
	CrmModifyRefund bool `json:"crmModifyRefund,omitempty"`
	// Description 描述、说明
	Description string `json:"description,omitempty"`
	// Email 联系人EMAIL
	Email string `json:"email,omitempty"`
	// FreightBill 运单号
	FreightBill string `json:"freightBill,omitempty"`
	// GmtCreate 创建时间
	GmtCreate go1688.JsonTime `json:"gmtCreate,omitempty"`
	// GmtModified 修改时间
	GmtModified go1688.JsonTime `json:"gmtModified,omitempty"`
	// ID 主键，退款操作记录流水号
	ID uint64 `json:"id,omitempty"`
	// MessageStatus 凭证状态，1:正常 2:后台小二屏蔽
	MessageStatus int `json:"messageStatus,omitempty"`
	// Mobile 联系人手机
	Mobile string `json:"mobile,omitempty"`
	// MsgType 留言类型 3:小二留言给买家和卖家 4:给买家的留言 5:给卖家的留言 7:cbu的普通留言等同于淘宝的1
	MsgType int `json:"msgType,omitempty"`
	// OperateRemark 操作备注
	OperateRemark string `json:"operateRemark,omitempty"`
	// OperateTypeInt 操作类型 取代operateType
	OperateTypeInt int `json:"operateTypeInt,omitempty"`
	// OperatorID 操作者-memberID
	OperatorID string `json:"operatorId,omitempty"`
	// OperatorLoginID 操作者-loginID
	OperatorLoginID string `json:"operatorLoginId,omitempty"`
	// OperatorRoleID 操作者角色名称 买家 卖家 系统
	OperatorRoleID int `json:"operatorRoleID,omitempty"`
	// OperatorUserID 操作者-userID
	OperatorUserID uint64 `json:"operatorUserID,omitempty"`
	// Phone 联系人电话
	Phone string `json:"phone,omitempty"`
	// RefundAddress 退货地址
	RefundAddress string `json:"refundAddress,omitempty"`
	// RefundID 退款记录ID
	RefundID string `json:"refundId,omitempty"`
	// RejectReason 卖家拒绝退款原因
	RejectReason string `json:"rejectReason,omitempty"`
	// Vouchers 凭证图片地址
	Vouchers []string `json:"vouchers,omitempty"`
	// LogisticsCompany 物流公司详情
	LogisticsCompany []logistics.OpLogisticsCompany `json:"logisticsCompany,omitempty"`
	// BuyerLoginID 买家LoginId
	BuyerLoginID string `json:"buyerLoginId,omitempty"`
	// SellerLoginID 卖家LoginId
	SellerLoginID string `json:"sellerLoginId,omitempty"`
}

// RefundBuyerQueryOrderRefundList 订单列表查看(买家视角)
func RefundBuyerQueryOrderRefundList(client *go1688.Client, req *RefundBuyerQueryOrderRefundListRequest, accessToken string) (*OpQueryOrderRefundListResult, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp RefundBuyerQueryOrderRefundListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
