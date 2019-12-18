package trade

import (
	"encoding/json"
	"errors"

	"github.com/bububa/go1688"
)

type GetRefundReasonListRequest struct {
	OrderId       uint64            `json:"orderId,omitempty"`       // 主订单
	OrderEntryIds []uint64          `json:"orderEntryIds,omitempty"` // 子订单
	GoodsStatus   RefundGoodsStatus `json:"goodsStatus,omitempty"`   // 货物状态, 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived"
}

type GetRefundReasonListRefinedRequest struct {
	OrderId       uint64            `json:"orderId,omitempty"`       // 主订单
	OrderEntryIds string            `json:"orderEntryIds,omitempty"` // 子订单
	GoodsStatus   RefundGoodsStatus `json:"goodsStatus,omitempty"`   // 货物状态, 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived"
}

func (this *GetRefundReasonListRequest) Refine() *GetRefundReasonListRefinedRequest {
	var entryIds []byte
	if len(this.OrderEntryIds) > 0 {
		entryIds, _ = json.Marshal(this.OrderEntryIds)
	}

	return &GetRefundReasonListRefinedRequest{
		OrderId:       this.OrderId,
		OrderEntryIds: string(entryIds),
		GoodsStatus:   this.GoodsStatus,
	}
}

func (this *GetRefundReasonListRefinedRequest) Name() string {
	return "alibaba.trade.getRefundReasonList"
}

type GetRefundReasonListResponse struct {
	go1688.BaseResponse
	Result *GetRefundReasonListResult `json:"result,omitempty"`
}

type GetRefundReasonListResult struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
	Result  struct {
		Reasons []OrderRefundReason `json:"reasons,omitempty"` // 原因列表
	} `json:"result,omitempty"` // 结果
}

type OrderRefundReason struct {
	Id               uint64 `json:"id,omitempty"`               // 原因id
	Name             string `json:"name,omitempty"`             // 原因
	NeedVoucher      bool   `json:"needVoucher,omitempty"`      // 凭证是否必须上传
	NoRefundCarriage bool   `json:"noRefundCarriage,omitempty"` // 是否支持退运费
	Tip              string `json:"tip,omitempty"`              // 提示
}

func GetRefundReasonList(client *go1688.Client, req *GetRefundReasonListRequest, accessToken string) ([]OrderRefundReason, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &GetRefundReasonListResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	if resp.Result.Code != "" {
		return nil, errors.New(resp.Result.Message)
	}
	return resp.Result.Result.Reasons, nil
}
