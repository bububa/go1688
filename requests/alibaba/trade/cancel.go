package trade

import (
	"github.com/bububa/go1688"
)

type CancelRequest struct {
	Website APIWebsite   `json:"webSite"`          // 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
	TradeId uint64       `json:"tradeID"`          // 交易id，订单号
	Reason  CancelReason `json:"cancelReason"`     // 原因描述；buyerCancel:买家取消订单;sellerGoodsLack:卖家库存不足;other:其它
	Remark  string       `json:"remark,omitempty"` // 备注
}

func (this *CancelRequest) Name() string {
	return "alibaba.trade.cancel"
}

type CancelResponse struct {
	go1688.BaseResponse
}

func Cancel(client *go1688.Client, req *CancelRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &CancelResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
