package message

// 1688订单部分发货（买家视角）
type OrderMessage struct {
	OrderId        uint64 `json:"orderId"`                // 订单ID
	Status         string `json:"currentStatus"`          // 当前订单状态，状态值为waitsellersend
	MsgSendTime    string `json:"msgSendTime"`            // 消息发送时间
	BuyerMemberId  string `json:"buyerMemberId"`          // 买家中文站会员ID
	SellerMemberId string `json:"sellerMemberId"`         // 卖家中文站会员ID
	refundAction   string `json:"refundAction,omitempty"` // 退款操作，具体状态为：BUYER_APPLY_REFUND(买家申请退款)、BUYER_RECEIVE_CLOSE(买家确认收货关闭)、SELLER_SEND_GOODS_CLOSE(卖家发货关闭)、BUYER_CANCEL_REFUND_CLOSE(买家撤销退款申请关闭)、BUYER_UPLOAD_BILL(买家上传凭证)、SELLER_UPLOAD_BILL(卖家上传凭证)、SELLER_REJECT_REFUND(卖家拒绝退款)、SELLER_AGREE_REFUND(卖家同意退款)、SELLER_RECEIVE_GOODS(卖家确认收货)、BUYER_SEND_GOODS(买家声明发货)、BUYER_MODIFY_REFUND_PROTOCOL(买家修改退款协议)、BUYER_APPLY_SUPPORT(买家申请客服介入)、SELLER_APPLY_SUPPORT(卖家申请客服介入)、SYSTEM_AGREE_REFUND_PROTOCOL(系统超时同意退款协议)、SYSTEM_AGREE_REFUND(系统超时同意退款)、SYSTEM_SEND_GOODS(系统超时退货，主交易流程的退货)、SYSTEM_MODIFY_REFUND_PROTOCOL(系统超时修改协议)、SYSTEM_NOTIFY_APPLY_SUPPORT(系统通知申请客服介入)、SELLER_AGREE_REFUND_PROCOTOL(卖家同意退款协议)、 SELLER_REJECT_REFUND_PROCOTOL(卖家拒绝退款协议)、 CRM_APPLY_TIMEOUT_CLOSE(申请客服介入、超时关闭、目前仅售后业务在用)、CRM_APPLY_SUPPORT(CRM申请介入)、 CRM_INTERVENE_TASK(CRM介入处理)、 CRM_DISMISS_TASK(CRM撤销工单)、 CRM_FINISH_TASK(CRM完结工单)、BUYER_STEP_PAY_ORDER_CLOSE(买家支付，退款关闭，分阶段订单情况)、BUYER_STEP_CONFIRM_CLOSE(买家确认，退款关闭，分阶段订单情况)、BUYER_CLOSE_TRADE_CLOSE(买家终止交易，退款关闭，分阶段订单情况)、SELLER_CONFIRM_ORDER_CLOSE(卖家确认，退款关闭，分阶段订单情况)、SELLER_STEP_PUSH_CLOSE(卖家推进，退款关闭，分阶段订单情况)
	Operator       string `json:"operator,omitempty"`     // 操作的发起人，buyer(买家)，seller(卖家)，system(系统)
}

func (this *OrderMessage) Types() []MessageType {
	return []MessageType{
		ORDER_BUYER_VIEW_PART_PART_SENDGOODS,
		ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS,
		ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS,
		ORDER_BUYER_VIEW_ORDER_SUCCESS,
		ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE,
		ORDER_BUYER_VIEW_ORDER_BUYER_REFUND_IN_SALES,
		ORDER_BUYER_VIEW_ORDER_REFUND_AFTER_SALES,
	}
}
