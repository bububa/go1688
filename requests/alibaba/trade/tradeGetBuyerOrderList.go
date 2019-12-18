package trade

import (
	"encoding/json"

	"github.com/bububa/go1688"
)

type TradeGetBuyerOrderListRequest struct {
	BizTypes                 []string        `json:"bizTypes,omitempty"`                 // 业务类型，支持： "cn"(普通订单类型), "ws"(大额批发订单类型), "yp"(普通拿样订单类型), "yf"(一分钱拿样订单类型), "fs"(倒批(限时折扣)订单类型), "cz"(加工定制订单类型), "ag"(协议采购订单类型), "hp"(伙拼订单类型), "gc"(国采订单类型), "supply"(供销订单类型), "nyg"(nyg订单类型), "factory"(淘工厂订单类型), "quick"(快订下单), "xiangpin"(享拼订单), "nest"(采购商城-鸟巢), "f2f"(当面付), "cyfw"(存样服务), "sp"(代销订单标记), "wg"(微供订单), "factorysamp"(淘工厂打样订单), "factorybig"(淘工厂大货订单)
	CreateEndTime            go1688.JsonTime `json:"createEndTime,omitempty"`            // 下单结束时间
	CreateStartTime          go1688.JsonTime `json:"createStartTime,omitempty"`          // 下单开始时间
	IsHis                    bool            `json:"isHis,omitempty"`                    // 是否查询历史订单表,默认查询当前表，即默认值为false
	ModifyEndTime            go1688.JsonTime `json:"modifyEndTime,omitempty"`            // 查询修改时间结束
	ModifyStartTime          go1688.JsonTime `json:"modifyStartTime,omitempty"`          // 查询修改时间开始
	OrderStatus              TradeStatus     `json:"orderStatus,omitempty"`              // 订单状态，值有 success, cancel(交易取消，违约金等交割完毕), waitbuyerpay(等待卖家付款)， waitsellersend(等待卖家发货), waitbuyerreceive(等待买家收货 )
	Page                     uint            `json:"page,omitempty"`                     // 查询分页页码，从1开始
	PageSize                 uint            `json:"pageSize,omitempty"`                 // 查询的每页的数量
	RefundStatus             RefundStatus    `json:"refundStatus,omitempty"`             // 退款状态，支持： "waitselleragree"(等待卖家同意), "refundsuccess"(退款成功), "refundclose"(退款关闭), "waitbuyermodify"(待买家修改), "waitbuyersend"(等待买家退货), "waitsellerreceive"(等待卖家确认收货)
	SellerMemberId           string          `json:"sellerMemberId,omitempty"`           // 卖家memberId
	SellerRateStatus         uint            `json:"sellerRateStatus,omitempty"`         // 卖家评价状态 (4:已评价,5:未评价,6;不需要评价)
	TradeType                string          `json:"tradeType,omitempty"`                // 交易类型: 担保交易(1), 预存款交易(2), ETC境外收单交易(3), 即时到帐交易(4), 保障金安全交易(5), 统一交易流程(6), 分阶段交易(7), 货到付款交易(8), 信用凭证支付交易(9), 账期支付交易(10), 1688交易4.0，新分阶段交易(50060), 当面付的交易流程(50070), 服务类的交易流程(50080)
	ProductName              string          `json:"productName,omitempty"`              // 商品名称
	NeedBuyerAddressAndPhone bool            `json:"needBuyerAddressAndPhone,omitempty"` // 是否需要查询买家的详细地址信息和电话
	NeedMemoInfo             bool            `json:"needMemoInfo,omitempty"`             // 是否需要查询备注信息
}

type TradeGetBuyerOrderListRefinedRequest struct {
	BizTypes                 string          `json:"bizTypes,omitempty"`                 // 业务类型，支持： "cn"(普通订单类型), "ws"(大额批发订单类型), "yp"(普通拿样订单类型), "yf"(一分钱拿样订单类型), "fs"(倒批(限时折扣)订单类型), "cz"(加工定制订单类型), "ag"(协议采购订单类型), "hp"(伙拼订单类型), "gc"(国采订单类型), "supply"(供销订单类型), "nyg"(nyg订单类型), "factory"(淘工厂订单类型), "quick"(快订下单), "xiangpin"(享拼订单), "nest"(采购商城-鸟巢), "f2f"(当面付), "cyfw"(存样服务), "sp"(代销订单标记), "wg"(微供订单), "factorysamp"(淘工厂打样订单), "factorybig"(淘工厂大货订单)
	CreateEndTime            go1688.JsonTime `json:"createEndTime,omitempty"`            // 下单结束时间
	CreateStartTime          go1688.JsonTime `json:"createStartTime,omitempty"`          // 下单开始时间
	IsHis                    bool            `json:"isHis,omitempty"`                    // 是否查询历史订单表,默认查询当前表，即默认值为false
	ModifyEndTime            go1688.JsonTime `json:"modifyEndTime,omitempty"`            // 查询修改时间结束
	ModifyStartTime          go1688.JsonTime `json:"modifyStartTime,omitempty"`          // 查询修改时间开始
	OrderStatus              TradeStatus     `json:"orderStatus,omitempty"`              // 订单状态，值有 success, cancel(交易取消，违约金等交割完毕), waitbuyerpay(等待卖家付款)， waitsellersend(等待卖家发货), waitbuyerreceive(等待买家收货 )
	Page                     uint            `json:"page,omitempty"`                     // 查询分页页码，从1开始
	PageSize                 uint            `json:"pageSize,omitempty"`                 // 查询的每页的数量
	RefundStatus             RefundStatus    `json:"refundStatus,omitempty"`             // 退款状态，支持： "waitselleragree"(等待卖家同意), "refundsuccess"(退款成功), "refundclose"(退款关闭), "waitbuyermodify"(待买家修改), "waitbuyersend"(等待买家退货), "waitsellerreceive"(等待卖家确认收货)
	SellerMemberId           string          `json:"sellerMemberId,omitempty"`           // 卖家memberId
	SellerRateStatus         uint            `json:"sellerRateStatus,omitempty"`         // 卖家评价状态 (4:已评价,5:未评价,6;不需要评价)
	TradeType                string          `json:"tradeType,omitempty"`                // 交易类型: 担保交易(1), 预存款交易(2), ETC境外收单交易(3), 即时到帐交易(4), 保障金安全交易(5), 统一交易流程(6), 分阶段交易(7), 货到付款交易(8), 信用凭证支付交易(9), 账期支付交易(10), 1688交易4.0，新分阶段交易(50060), 当面付的交易流程(50070), 服务类的交易流程(50080)
	ProductName              string          `json:"productName,omitempty"`              // 商品名称
	NeedBuyerAddressAndPhone bool            `json:"needBuyerAddressAndPhone,omitempty"` // 是否需要查询买家的详细地址信息和电话
	NeedMemoInfo             bool            `json:"needMemoInfo,omitempty"`             // 是否需要查询备注信息
}

func (this *TradeGetBuyerOrderListRequest) Refine() *TradeGetBuyerOrderListRefinedRequest {
	var bizTypes []byte
	if len(this.BizTypes) > 0 {
		bizTypes, _ = json.Marshal(this.BizTypes)
	}
	return &TradeGetBuyerOrderListRefinedRequest{
		BizTypes:                 string(bizTypes),
		CreateEndTime:            this.CreateEndTime,
		CreateStartTime:          this.CreateStartTime,
		IsHis:                    this.IsHis,
		ModifyEndTime:            this.ModifyEndTime,
		ModifyStartTime:          this.ModifyStartTime,
		OrderStatus:              this.OrderStatus,
		Page:                     this.Page,
		PageSize:                 this.PageSize,
		RefundStatus:             this.RefundStatus,
		SellerMemberId:           this.SellerMemberId,
		SellerRateStatus:         this.SellerRateStatus,
		TradeType:                this.TradeType,
		ProductName:              this.ProductName,
		NeedBuyerAddressAndPhone: this.NeedBuyerAddressAndPhone,
		NeedMemoInfo:             this.NeedMemoInfo,
	}
}

func (this *TradeGetBuyerOrderListRefinedRequest) Name() string {
	return "alibaba.trade.getBuyerOrderList"
}

type TradeGetBuyerOrderListResponse struct {
	go1688.BaseResponse
	Result      []*TradeInfo `json:"result,omitempty"`
	TotalRecord uint         `json:"totalRecord,omitempty"`
}

func TradeGetBuyerOrderList(client *go1688.Client, req *TradeGetBuyerOrderListRequest, accessToken string) (uint, []*TradeInfo, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &TradeGetBuyerOrderListResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return 0, nil, err
	}
	if resp.IsError() {
		return 0, nil, resp
	}
	return resp.TotalRecord, resp.Result, nil
}
