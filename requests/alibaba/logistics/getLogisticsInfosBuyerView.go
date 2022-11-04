package logistics

import (
	"strconv"

	"github.com/bububa/go1688"
)

// GetLogisticsInfosBuyerViewRequest 获取交易订单的物流信息(买家视角) API Request
type GetLogisticsInfosBuyerViewRequest struct {
	// OrderID 订单ID
	OrderID uint64 `json:"orderId,omitempty"`
	// Fields 需要返回的字段，目前有:company.name,sender,receiver,sendgood。返回的字段要用英文逗号分隔开
	Fields string `json:"fields,omitempty"`
	// Website 是1688业务还是icbu业务
	Website APIWebsite `json:"webSite,omitempty"`
}

// Name implement RequestData interface
func (r GetLogisticsInfosBuyerViewRequest) Name() string {
	return "alibaba.trade.getLogisticsInfos.buyerView"
}

// Map implement RequestData interface
func (r GetLogisticsInfosBuyerViewRequest) Map() map[string]string {
	ret := make(map[string]string)
	if r.OrderID > 0 {
		ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	}
	ret["fields"] = r.Fields
	if r.Website != "" {
		ret["webSite"] = r.Website
	}
	return ret
}

// GetLogisticsInfosBuyerViewResponse 获取交易订单的物流信息(买家视角) API Response
type GetLogisticsInfosBuyerViewResponse struct {
	go1688.BaseResponse
	// Result 返回结果
	Result []*LogisticsOrder `json:"result,omitempty"`
}

// LogisticsOrder 物流信息
type LogisticsOrder struct {
	// Sender 发件人信息
	Sender *LogisticsSender `json:"sender,omitempty"`
	// ServiceFeature .
	ServiceFeature string `json:"serviceFeature,omitempty"`
	// OrderEntryIDs 订单号列表，无子订单的等于主订单编号，否则为对应子订单列表
	OrderEntryIDs string `json:"orderEntryIds,omitempty"`
	// LogisticsBillNo 物流单号，运单号
	LogisticsBillNo string `json:"logisticsBillNo,omitempty"`
	// LogisticsID 物流信息ID
	LogisticsID string `json:"logisticsId,omitempty"`
	// Receiver 收件人信息
	Receiver *LogisticsReceiver `json:"receiver,omitempty"`
	// LogisticsCompanyID 物流公司ID
	LogisticsCompanyID string `json:"logisticsCompanyId,omitempty"`
	// LogisticsCompanyName 物流公司编码
	LogisticsCompanyName string `json:"logisticsCompanyName,omitempty"`
	// Status 物流状态。WAITACCEPT:未受理;CANCEL:已撤销;ACCEPT:已受理;TRANSPORT:运输中;NOGET:揽件失败;SIGN:已签收;UNSIGN:签收异常
	Status LogisticsStatus `json:"status,omitempty"`
	// SendGoods 商品信息
	SendGoods []LogisticsSendGood `json:"sendGoods,omitempty"`
	// GmtSystemSend .
	GmtSystemSend string `json:"gmtSystemSend,omitempty"`
	// Remarks 备注
	Remarks string `json:"remarks,omitempty"`
}

// LogisticsSender 发件人信息
type LogisticsSender struct {
	// Name 发件人姓名
	Name string `json:"senderName,omitempty"`
	// Phone 发件人电话
	Phone string `json:"senderPhone,omitempty"`
	// Mobile 发件人电话
	Mobile string `json:"senderMobile,omitempty"`
	// Encrypt .
	Encrypt string `json:"encrypt,omitempty"`
	// ProvinceCode 省编码
	ProvinceCode string `json:"senderProvinceCode,omitempty"`
	// CityCode 城市编码
	CityCode string `json:"senderCityCode,omitempty"`
	// CountyCode 国家编码
	CountyCode string `json:"senderCountryCode,omitempty"`
	// Address 发货人地址
	Address string `json:"senderAddress,omitempty"`
	// Province 省份
	Province string `json:"senderProvince,omitempty"`
	// City 城市
	City string `json:"senderCity,omitempty"`
	// Country 国家
	County string `json:"senderCounty,omitempty"`
}

// LogisticsReceiver 收件人信息
type LogisticsReceiver struct {
	// Name 收件人名字
	Name string `json:"receiverName,omitempty"`
	// Phone 收件人电话
	Phone string `json:"receiverPhone,omitempty"`
	// Mobile 收件人电话
	Mobile string `json:"receiverMobile,omitempty"`
	// Encrypt .
	Encrypt string `json:"encrypt,omitempty"`
	// ProvinceCode 省编码
	ProvinceCode string `json:"receiverProvinceCode,omitempty"`
	// CityCode 城市编码
	CityCode string `json:"receiverCityCode,omitempty"`
	// CountyCode 国家编码
	CountyCode string `json:"receiverCountryCode,omitempty"`
	// Address 发货人地址
	Address string `json:"receiverAddress,omitempty"`
	// Province 省份
	Province string `json:"receiverProvince,omitempty"`
	// City 城市
	City string `json:"receiverCity,omitempty"`
	// Country 国家
	County string `json:"receiverCounty,omitempty"`
}

// LogisticsSendGood 商品信息
type LogisticsSendGood struct {
	// Name 商品名
	Name string `json:"goodName,omitempty"`
	// Quantity 商品数量
	Quantity string `json:"quantity,omitempty"`
	// Unit 商品单位
	Unit string `json:"unit,omitempty"`
}

// GetLogisticsInfosBuyerView 获取交易订单的物流信息(买家视角)
// 该接口需要获得订单买家的授权，获取买家的订单的物流详情，在采购或者分销场景中，作为买家也有获取物流详情的需求。该接口能查能根据
// 订单号查看物流详情，包括发件人，收件人，所发货物明细等。由于物流单录入的原因，可能跟踪信息的API查询会有延迟。该API需要向开放平台申请权限才能访问。In the procurement or distribution scenario, buyers can ask for obtaining the logistics details. The interface can check the logistics details according to the order ID, including the sender, the recipient, the details of the goods sent, and so on. Depending on the logistics information entry time, there may be a delay in API queries regarding the information tracking.
func GetLogisticsInfosBuyerView(client *go1688.Client, req *GetLogisticsInfosBuyerViewRequest, accessToken string) ([]*LogisticsOrder, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp GetLogisticsInfosBuyerViewResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
