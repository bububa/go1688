package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// CancelRequest 取消交易 API Request
type CancelRequest struct {
	// Website 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
	Website APIWebsite `json:"webSite,omitempty"`
	// TradeID 交易id，订单号
	TradeID uint64 `json:"tradeID,omitempty"`
	// Reason 原因描述；buyerCancel:买家取消订单;sellerGoodsLack:卖家库存不足;other:其它
	Reason CancelReason `json:"cancelReason,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
}

// Name implement RequestData interface
func (r CancelRequest) Name() string {
	return "alibaba.trade.cancel"
}

// Map implement RequestData interface
func (r CancelRequest) Map() map[string]string {
	ret := make(map[string]string, 4)
	ret["webSite"] = r.Website
	ret["tradeID"] = strconv.FormatUint(r.TradeID, 10)
	ret["cancelReason"] = r.Reason
	if r.Remark != "" {
		ret["remark"] = r.Remark
	}
	return ret
}

// Cancel 取消交易
func Cancel(client *go1688.Client, req *CancelRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	return client.Do(finalRequest, accessToken, nil)
}
