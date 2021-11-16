package common

import (
	"2022/ginseckill/config"
	"2022/ginseckill/tools"
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type ISmsServer interface {
	ImitateSms(phone string,tplId string,key string) (bizId,codeStr string,err error)
	AliSendSms(phone string, tplId string, key string) (bizId ,code string,err error)
}

type SmsServer struct {}

func NewSmsServer() ISmsServer {
	return &SmsServer{}
}


// 模拟发送验证码
func (ss *SmsServer) ImitateSms(phone string,tplId string,key string) (bizId,codeStr string,err error) {
	var smsConf = config.Conf.AliSms
	tpl := "yunsong-sms-" + tplId
	k := "yunsong-sms-" + key

	if tpl != smsConf.TplId || k != smsConf.Key {
		err = errors.New("发送验证码特征错误")
		return "","",err
	}
	// 产生验证码
	codeStr = tools.GetRandomCode()
	bizId = "bizId"
	return
}
// 阿里云发送验证码
func (ss *SmsServer) AliSendSms(phone string, tplId string, key string) (bizId ,code string,err error) {
	var smsConf = config.Conf.AliSms
	tpl := "yunsong-sms-" + tplId
	k := "yunsong-sms-" + key
	if tpl != smsConf.TplId || k != smsConf.Key {
		err = errors.New("发送验证码特征错误")
		return "","",err
	}
	// 产生验证码
	client,err := dysmsapi.NewClientWithAccessKey(smsConf.RegionId,smsConf.AppKey,smsConf.AppSecret)
	if err != nil {
		return "","",err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsConf.SmsName
	request.TemplateCode = smsConf.TemplateCode
	request.PhoneNumbers = phone
	//固定格式
	par, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)
	// 获取返回数据
	response,err := client.SendSms(request)
	if err != nil {
		return "","",err
	}
	if response.Code == "OK" {
		// 发送成功
		bizId = response.BizId
		code = response.Code
		return
	}
	//发送失败
	err = errors.New("发送短信失败")
	return "","",err
}