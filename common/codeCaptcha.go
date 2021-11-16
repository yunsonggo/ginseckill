package common

import "github.com/mojocn/base64Captcha"

type GenerateCodeCaptchaBody struct {
	Id string
	B64s string
}

type VerifyCaptchaBody struct {
	Id string
	VerifyValue string
}

var codeCaptchaStore = base64Captcha.DefaultMemStore

// 生成验证码
func GenerateCaptchaCode() (captchaCode *GenerateCodeCaptchaBody,err error) {
	dc := base64Captcha.DriverDigit{
		Height: 38,
		Width: 118,
		Length: 4,
		MaxSkew: 0.4,
		DotCount: 30,
	}
	driverDigit := base64Captcha.NewDriverDigit(dc.Height,dc.Width,dc.Length,dc.MaxSkew,dc.DotCount)
	var driver base64Captcha.Driver = driverDigit
	c := base64Captcha.NewCaptcha(driver,codeCaptchaStore)
	id,b64s,err := c.Generate()
	if err != nil {
		return nil,err
	}
	var Code GenerateCodeCaptchaBody
	Code.Id = id
	Code.B64s = b64s
	return &Code,nil
}


// 验证验证码
func VerifyCaptchaCode(code VerifyCaptchaBody) (res bool) {
	return codeCaptchaStore.Verify(code.Id, code.VerifyValue, true)
}