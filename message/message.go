package message

import (
	"encoding/json"
	"errors"

	"github.com/bububa/go1688"
)

type PushMessage struct {
	Message string `form:"message" json:"message" binding:"required"`               //对接方获取这个参数的值，然后通过json的反序列化，就得到了消息模型
	Sign    string `form:"_aop_signature" json:"_aop_signature" binding:"required"` // 针对消息的一个签名，可防篡改
}

func (this *PushMessage) Valid(clt *go1688.Client) bool {
	return this.Sign == clt.Sign("", map[string]string{"message": this.Message})
}

type Message interface {
	Types() []MessageType
}

type MessageTemplate struct {
	Id        uint64      `json:"msgId"`     // 消息ID，消息唯一性标识
	GmtBorn   int64       `json:"gmtBorn"`   // 消息推送时间
	Data      string      `json:"data"`      // 具体推送的业务消息数据，json格式，字段说明，参考各个业务消息说明
	UserInfo  string      `json:"userInfo"`  // memberId
	Type      MessageType `json:"type"`      // 消息类型，每个业务消息都唯一对应一个类型，参考业务消息的类型定义
	ExtraInfo string      `json:"extraInfo"` // 扩展字段，暂未启用
}

func (this *MessageTemplate) Message() (Message, error) {
	var orderMsg OrderMessage
	for _, msgType := range orderMsg.Types() {
		if msgType == this.Type {
			err := json.Unmarshal([]byte(this.Data), &orderMsg)
			return &orderMsg, err
		}
	}
	var productMsg ProductMessage
	for _, msgType := range productMsg.Types() {
		if msgType == this.Type {
			err := json.Unmarshal([]byte(this.Data), &productMsg)
			return &productMsg, err
		}
	}
	return nil, errors.New("not found")
}
