package trade

import (
	"encoding/json"

	"github.com/bububa/go1688"
)

type CreateOrder4CybMediaRequest struct {
	Address        *Address        `json:"address"`
	CargoList      []*Cargo        `json:"cargo"`
	Message        string          `json:"message,omitempty"`
	OuterOrderInfo *OuterOrderInfo `json:"outerOrderInfo,omitempty"`
	TradeType      TradeType       `json:"tradeType"`
}

type OuterOrderInfo struct {
	MediaOrderId uint64            `json:"mediaOrderId,omitempty"` // 机构订单号
	Phone        string            `json:"phone,omitempty"`        // 电话
	Offers       []OuterOrderOffer `json:"offers,omitempty"`
}

type OuterOrderOffer struct {
	Id     uint64 `json:"id,omitempty"`     // 1688商品id
	SpecId string `json:"specId,omitempty"` // 1688商品specId(可能无)
	Price  uint   `json:"price,omitempty"`  // 媒体溢价单价(单位分)
	Num    uint   `json:"num,omitempty"`    // 售卖数量
}

type CreateOrder4CybMediaRefinedRequest struct {
	Address        string    `json:"addressParam"`   // 收货地址信息
	CargoList      string    `json:"cargoParamList"` // 商品信息
	Message        string    `json:"message,omitempty"`
	OuterOrderInfo string    `json:"outerOrderInfo"`
	TradeType      TradeType `json:"tradeType"`
}

func (this *CreateOrder4CybMediaRefinedRequest) Name() string {
	return "alibaba.trade.createOrder4CybMedia"
}

func (this *CreateOrder4CybMediaRequest) Refine() *CreateOrder4CybMediaRefinedRequest {
	address, _ := json.Marshal(this.Address)
	cargoList, _ := json.Marshal(this.CargoList)
	outerOrderInfo, _ := json.Marshal(this.OuterOrderInfo)
	return &CreateOrder4CybMediaRefinedRequest{
		Address:        string(address),
		CargoList:      string(cargoList),
		OuterOrderInfo: string(outerOrderInfo),
		Message:        this.Message,
		TradeType:      this.TradeType,
	}
}

type CreateOrder4CybMediaResponse struct {
	go1688.BaseResponse
	Result *CreateOrder4CybMediaResult `json:"result,omitempty"` // 返回结果
}

type CreateOrder4CybMediaResult struct {
	TotalSuccessAmount uint   `json:"totalSuccessAmount,omitempty"` // 下单成功的订单总金额，单位：分
	OrderId            uint64 `jsong:"orderId,omitempty"`           // 下单成功后的订单id
	PostFee            uint   `json:"postFee,omitempty"`            // 原始运费，单位：分。注意：下单后卖家可能调整，因此该值可能不等于最终支付运费
}

func CreateOrder4CybMedia(client *go1688.Client, req *CreateOrder4CybMediaRequest, accessToken string) (*CreateOrder4CybMediaResult, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &CreateOrder4CybMediaResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Result, nil
}
