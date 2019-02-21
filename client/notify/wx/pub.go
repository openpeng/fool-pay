package wxnotify

import (
	"github.com/openpeng/fool-pay/common"
)

type WechatNotify struct {
	*NotifyClient
}

func NewWechatNotify(config common.BaseConfig) *WechatNotify {
	temp := &WechatNotify{}
	temp.NotifyClient = NewNotifyClient(config, temp)
	return temp
}

func (wn *WechatNotify) BuildResData() string {
	return "<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>"
}

func (wpc *WechatNotify) BuildData() string {
	return "<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>"
}
