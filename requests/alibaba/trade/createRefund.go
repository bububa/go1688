package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// CreateRefundRequest 创建退款退货申请 API Request
type CreateRefundRequest struct {
	// OrderID 主订单
	OrderID uint64 `json:"orderId,omitempty"`
	// OrderEntryIDs 子订单
	OrderEntryIDs []uint64 `json:"orderEntryIds,omitempty"`
	// DisputeRequest 退款/退款退货。只有已收到货，才可以选择退款退货。退款:"refund"; 退款退货:"returnRefund"
	DisputeRequest DisputeRequest `json:"disputeRequest,omitempty"`
	// ApplyPayment 退款金额（单位：分）。不大于实际付款金额；等待卖家发货时，必须为商品的实际付款金额。
	ApplyPayment int64 `json:"applyPayment,omitempty"`
	// ApplyCarriage 退运费金额（单位：分）。
	ApplyCarriage int64 `json:"applyCarriage,omitempty"`
	// ApplyReasonID 退款原因id（从API getRefundReasonList获取）
	ApplyReasonID uint64 `json:"applyReasonId,omitempty"`
	// Description 退款申请理由，2-150字
	Description string `json:"description,omitempty"`
	// GoodsStatus 货物状态, 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived"
	GoodsStatus RefundGoodsStatus `json:"goodsStatus,omitempty"`
	// Vouchers 凭证图片URLs。1-5张，必须使用API uploadRefundVoucher返回的“图片域名/相对路径”
	Vouchers []string `json:"vouchers,omitempty"`
	// OrderEntryCountList 子订单退款数量。仅在售中买家已收货（退款退货）时，可指定退货数量；默认，全部退货。
	OrderEntryCountList []OrderEntryCount `json:"orderEntryCountList,omitempty"`
}

// Name implement RequestData interface
func (r CreateRefundRequest) Name() string {
	return "alibaba.trade.createRefund"
}

// Map implement RequestData interface
func (r CreateRefundRequest) Map() map[string]string {
	ret := make(map[string]string, 10)
	ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	if len(r.OrderEntryIDs) > 0 {
		ret["orderEntryIds"] = go1688.JSONMarshal(r.OrderEntryIDs)
	}
	ret["disputeRequest"] = r.DisputeRequest
	ret["applyPayment"] = strconv.FormatInt(r.ApplyPayment, 10)
	ret["applyCarriage"] = strconv.FormatInt(r.ApplyCarriage, 10)
	ret["applyReasonId"] = strconv.FormatUint(r.ApplyReasonID, 10)
	ret["description"] = r.Description
	ret["goodsStatus"] = r.GoodsStatus
	if len(r.Vouchers) > 0 {
		ret["vouchers"] = go1688.JSONMarshal(r.Vouchers)
	}
	if len(r.OrderEntryCountList) > 0 {
		ret["orderEntryCountList"] = go1688.JSONMarshal(r.OrderEntryCountList)
	}
	return ret
}

// OrderEntryCount 子订单退款数量。仅在售中买家已收货（退款退货）时，可指定退货数量；默认，全部退货。
type OrderEntryCount struct {
	// ID 子订单id
	ID uint64 `json:"id,omitempty"`
	// Count 子订单购买商品数量
	Count int `json:"count,omitempty"`
}

// CreateRefundResponse 创建退款退货申请 API Response
type CreateRefundResponse struct {
	go1688.BaseResponse
	Result *CreateRefundResult `json:"result,omitempty"`
}

// CreateRefundResult 创建退款退货申请 API Result
type CreateRefundResult struct {
	// Code 错误码
	Code string `json:"code"`
	// Message 错误信息
	Message string `json:"message"`
	// Result 结果
	Result struct {
		// RefundID 创建成功，退款id
		RefundID string `json:"refundId,omitempty"`
	} `json:"result,omitempty"`
	// Success 是否成功
	Success bool `json:"success,omitempty"`
}

// IsError check create refund success
func (r CreateRefundResult) IsError() bool {
	return !r.Success
}

// Error implement error interface
func (r CreateRefundResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.Code)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.Message)
	return builder.String()
}

// CreateRefund 创建退款退货申请
func CreateRefund(client *go1688.Client, req *CreateRefundRequest, accessToken string) (string, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp CreateRefundResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return "", err
	}
	if resp.Result.IsError() {
		return "", resp.Result
	}
	return resp.Result.Result.RefundID, nil
}
