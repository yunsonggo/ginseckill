package business

import (
	"2022/ginseckill/tools"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func Auth(w http.ResponseWriter,r *http.Request) error {
	fmt.Println("执行权限验证")
	err := checkUserInfo(r)
	if err != nil {return err}

	return nil
}

// 执行验证cookie
func checkUserInfo(r *http.Request) error {
	// 获取cookie
	uidCookie,err := r.Cookie("uid")
	if err != nil {
		return errors.New("用户UID Cookie 获取失败！")
	}
	fmt.Printf("checkUserInfo cookie:%v,err:%v\n",uidCookie.Value,err)
	// 获取加密子串
	signCookie,err := r.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串 Cookie 获取失败！")
	}
	fmt.Printf("sign cookie:%v\n",signCookie.Value)
	// 解密
	signValue ,_ := url.PathUnescape(signCookie.Value)
	signByte,err := tools.UrlBase64Decode(signValue)
	if err != nil {
		fmt.Printf("base64 decode err:%s\n",err)
		return err
	}
	signStr := string(tools.AesCtrDecrypt(signByte))
	fmt.Println("用户ID：",uidCookie.Value)
	fmt.Println("解密后用户ID：",signStr)
	if checkInfo(uidCookie.Value,signStr) {
		return nil
	}
	return err
}

// 比对加密字串 可以实现其他的逻辑这里简单对比
func checkInfo(checkStr,signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}