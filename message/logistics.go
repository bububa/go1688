package message

// LogisticsMessage 物流消息
type LogisticsMessage struct {
	// OrderLogisticsTracing 物流单状态变更（买家视角）
	OrderLogisticsTracing *OrderLogisticsTracing `json:"OrderLogisticsTracingModel,omitempty"`
	// MailNoChangeModel 物流单号修改消息
	MailNoChangeModel *MailNoChangeModel `json:"MailNoChangeModel,omitempty"`
}

// OrderLogisticsTracing 物流单状态变更（买家视角）
type OrderLogisticsTracing struct {
	// LogisticsID 物流编号
	LogisticsID string `json:"logisticsId,omitempty"`
	// CpCode cp code
	CpCode string `json:"cpCode,omitempty"`
	// MailNo 运单号
	MailNo string `json:"mailNo,omitempty"`
	// StatusChanged 物流单发生变化的状态，包括发货（CONSIGN）、揽收（ACCEPT）、运输（TRANSPORT）、派送（DELIVERING）、签收（SIGN）
	StatusChanged LogisticsStatus `json:"statusChanged,omitempty"`
	// ChangeTime 发生变化的时间
	ChagneTime string `json:"changeTime,omitempty"`
	// OrderLogsItems 该物流单关联的订单信息
	OrderLogsItems []OrderLogsItem `json:"orderLogsItems,omitempty"`
}

// MailNoChangeModel 物流单号修改消息
type MailNoChangeModel struct {
	// LogisticsID 物流编号
	LogisticsID string `json:"logisticsId,omitempty"`
	// OldCpCode 更改前的cp code
	OldCpCode string `json:"oldCpCode,omitempty"`
	// NewCpCode 更改后的cp code
	NewCpCode string `json:"newCpCode,omitempty"`
	// OldMailNo 更改前的运单号
	OldMailNo string `json:"oldMailNo,omitempty"`
	// NewMailNo 更改后的运单号
	NewMailNo string `json:"newMailNo,omitempty"`
	// EventType 发生时间
	EventType string `json:"eventType,omitempty"`
	// OrderLogsItems 该物流单关联的订单信息
	OrderLogsItems []OrderLogsItem `json:"orderLogsItems,omitempty"`
}

// OrderLogsItem 该物流单关联的订单信息
type OrderLogsItem struct {
	// OrderID 交易主单id
	OrderID uint64 `json:"orderId,omitempty"`
	// OrderEntryID 交易子单id
	OrderEntryID uint64 `json:"orderEntryId,omitempty"`
}

// Types implement Message interface
func (m LogisticsMessage) Types() []MessageType {
	return []MessageType{
		LOGISTICS_BUYER_VIEW_TRACE,
		LOGISTICS_MAIL_NO_CHANGE,
	}
}
