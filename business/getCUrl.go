package business

import (
	"io/ioutil"
	"net/http"
)

//  模拟http请求
func GetCurl(hostUrl string,request *http.Request) (resp *http.Response,body []byte,err error) {
	// 获取uid
	uidPre,err := request.Cookie("uid")
	if err != nil {
		return
	}
	// 获取sign
	uidSign, err := request.Cookie("sign")
	if err != nil {
		return
	}
	// 模拟接口访问
	client := &http.Client{}
	req, err := http.NewRequest("GET", hostUrl, nil)
	if err != nil {
		return
	}
	// 手动指定， 排查多余cookies
	cookieUid := &http.Cookie{Name: "uid", Value: uidPre.Value, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: uidSign.Value, Path: "/"}
	// 添加cookie到模拟的请求中
	req.AddCookie(cookieUid)
	req.AddCookie(cookieSign)
	// 获取返回结果
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
