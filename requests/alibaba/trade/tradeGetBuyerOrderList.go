package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// TradeGetBuyerOrderListRequest 订单列表查看(买家视角) API Request
type TradeGetBuyerOrderListRequest struct {
	// BizTypes 业务类型，支持： "cn"(普通订单类型), "ws"(大额批发订单类型), "yp"(普通拿样订单类型), "yf"(一分钱拿样订单类型), "fs"(倒批(限时折扣)订单类型), "cz"(加工定制订单类型), "ag"(协议采购订单类型), "hp"(伙拼订单类型), "gc"(国采订单类型), "supply"(供销订单类型), "nyg"(nyg订单类型), "factory"(淘工厂订单类型), "quick"(快订下单), "xiangpin"(享拼订单), "nest"(采购商城-鸟巢), "f2f"(当面付), "cyfw"(存样服务), "sp"(代销订单标记), "wg"(微供订单), "factorysamp"(淘工厂打样订单), "factorybig"(淘工厂大货订单)
	BizTypes []string `json:"bizTypes,omitempty"`
	// CreateEndTime 下单结束时间
	CreateEndTime go1688.JsonTime `json:"createEndTime,omitempty"`
	// CreateStartTime 下单开始时间
	CreateStartTime go1688.JsonTime `json:"createStartTime,omitempty"`
	// IsHis 是否查询历史订单表,默认查询当前表，即默认值为false
	IsHis bool `json:"isHis,omitempty"`
	// ModifyEndTime 查询修改时间结束
	ModifyEndTime go1688.JsonTime `json:"modifyEndTime,omitempty"`
	// ModifyStartTime 查询修改时间开始
	ModifyStartTime go1688.JsonTime `json:"modifyStartTime,omitempty"`
	// OrderStatus 订单状态，值有 success, cancel(交易取消，违约金等交割完毕), waitbuyerpay(等待卖家付款)， waitsellersend(等待卖家发货), waitbuyerreceive(等待买家收货 )
	OrderStatus TradeStatus `json:"orderStatus,omitempty"`
	// Page 查询分页页码，从1开始
	Page int `json:"page,omitempty"`
	// PageSize 查询的每页的数量
	PageSize int `json:"pageSize,omitempty"`
	// RefundStatus 退款状态，支持： "waitselleragree"(等待卖家同意), "refundsuccess"(退款成功), "refundclose"(退款关闭), "waitbuyermodify"(待买家修改), "waitbuyersend"(等待买家退货), "waitsellerreceive"(等待卖家确认收货)
	RefundStatus RefundStatus `json:"refundStatus,omitempty"`
	// SellerMemberID 卖家memberId
	SellerMemberID string `json:"sellerMemberId,omitempty"`
	// SellerRateStatus 卖家评价状态 (4:已评价,5:未评价,6;不需要评价)
	SellerRateStatus int `json:"sellerRateStatus,omitempty"`
	// TradeType 交易类型: 担保交易(1), 预存款交易(2), ETC境外收单交易(3), 即时到帐交易(4), 保障金安全交易(5), 统一交易流程(6), 分阶段交易(7), 货到付款交易(8), 信用凭证支付交易(9), 账期支付交易(10), 1688交易4.0，新分阶段交易(50060), 当面付的交易流程(50070), 服务类的交易流程(50080)
	TradeType string `json:"tradeType,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"productName,omitempty"`
	// NeedBuyerAddressAndPhone 是否需要查询买家的详细地址信息和电话
	NeedBuyerAddressAndPhone bool `json:"needBuyerAddressAndPhone,omitempty"`
	// NeedMemoInfo 是否需要查询备注信息
	NeedMemoInfo bool `json:"needMemoInfo,omitempty"`
}

// Name implement RequestData interface
func (r TradeGetBuyerOrderListRequest) Name() string {
	return "alibaba.trade.getBuyerOrderList"
}

// Map implement RequestData interface
func (r TradeGetBuyerOrderListRequest) Map() map[string]string {
	ret := make(map[string]string, 16)
	if len(r.BizTypes) > 0 {
		ret["bizTypes"] = go1688.JSONMarshal(r.BizTypes)
	}
	if !r.CreateEndTime.IsZero() {
		ret["createEndTime"] = r.CreateEndTime.Format()
	}
	if !r.CreateStartTime.IsZero() {
		ret["createStartTime"] = r.CreateStartTime.Format()
	}
	if r.IsHis {
		ret["isHis"] = "true"
	}
	if !r.ModifyEndTime.IsZero() {
		ret["modifyEndTime"] = r.ModifyEndTime.Format()
	}
	if !r.ModifyStartTime.IsZero() {
		ret["modifyStartTime"] = r.ModifyStartTime.Format()
	}
	if r.OrderStatus != "" {
		ret["orderStatus"] = r.OrderStatus
	}
	if r.Page > 1 {
		ret["page"] = strconv.Itoa(r.Page)
	}
	if r.PageSize > 1 {
		ret["pageSize"] = strconv.Itoa(r.PageSize)
	}
	if r.RefundStatus != "" {
		ret["refundStatus"] = r.RefundStatus
	}
	if r.SellerMemberID != "" {
		ret["sellerMemberId"] = r.SellerMemberID
	}
	if r.SellerRateStatus != 0 {
		ret["sellerRateStatus"] = strconv.Itoa(r.SellerRateStatus)
	}
	if r.TradeType != "" {
		ret["tradeType"] = r.TradeType
	}
	if r.ProductName != "" {
		ret["productName"] = r.ProductName
	}
	if r.NeedBuyerAddressAndPhone {
		ret["needBuyerAddressAndPhone"] = "true"
	}
	if r.NeedMemoInfo {
		ret["needMemoInfo"] = "true"
	}
	return ret
}

// TradeGetBuyerOrderListResponse 订单列表查看(买家视角) API Response
type TradeGetBuyerOrderListResponse struct {
	go1688.BaseResponse
	// Result 查询返回列表
	Result []TradeInfo `json:"result,omitempty"`
	// TotalRecord 总记录数
	TotalRecord int `json:"totalRecord,omitempty"`
}

// TradeGetBuyerOrderList 订单列表查看(买家视角)
func TradeGetBuyerOrderList(client *go1688.Client, req *TradeGetBuyerOrderListRequest, accessToken string) (*TradeGetBuyerOrderListResponse, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp TradeGetBuyerOrderListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
