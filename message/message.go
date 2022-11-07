package message

import (
	"encoding/json"
	"errors"

	"github.com/bububa/go1688"
)

type PushMessage struct {
	// Message 对接方获取这个参数的值，然后通过json的反序列化，就得到了消息模型
	Message string `form:"message" json:"message" binding:"required"`
	// Sign 针对消息的一个签名，可防篡改
	Sign string `form:"_aop_signature" json:"_aop_signature" binding:"required"`
}

func (p PushMessage) Valid(clt *go1688.Client) bool {
	return p.Sign == clt.Sign("", map[string]string{"message": p.Message})
}

// Message 消息接口
type Message interface {
	Types() []MessageType
}

// MessageTemplate 消息模板
type MessageTemplate struct {
	// ID 消息ID，消息唯一性标识
	ID uint64 `json:"msgId"`
	// GmtBorn 消息推送时间
	GmtBorn int64 `json:"gmtBorn"`
	// Data 具体推送的业务消息数据，json格式，字段说明，参考各个业务消息说明
	Data json.RawMessage `json:"data"`
	// UserInfo memberId
	UserInfo string `json:"userInfo"`
	// Type 消息类型，每个业务消息都唯一对应一个类型，参考业务消息的类型定义
	Type MessageType `json:"type"`
	// ExtraInfo 扩展字段，暂未启用
	ExtraInfo string `json:"extraInfo"`
}

// Message decode message from MessageTemplate
func (m MessageTemplate) Message() (Message, error) {
	var orderMsg OrderMessage
	for _, msgType := range orderMsg.Types() {
		if msgType == m.Type {
			err := json.Unmarshal(m.Data, &orderMsg)
			return &orderMsg, err
		}
	}
	var productMsg ProductMessage
	for _, msgType := range productMsg.Types() {
		if msgType == m.Type {
			err := json.Unmarshal(m.Data, &productMsg)
			return &productMsg, err
		}
	}
	return nil, errors.New("not found")
}
