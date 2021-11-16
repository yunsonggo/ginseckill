package tools

import (
	"regexp"
)

//email verify
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	//pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	//regular := `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$`
	regular := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// password verify
func VerifyPassword(password string) bool {
	//regular := `^[a-zA-Z]([-_a-zA-Z0-9]{7,19})+$`
	regular := `^[a-zA-Z]([-_a-zA-Z0-9]{7,19})+$`
	reg := regexp.MustCompile(regular)
	res := reg.MatchString(password)
	return res
}

// 中文
func VerifyChineseWord(wordStr string) bool {
	regular := `^[\u4E00-\u9FA5]`
	reg := regexp.MustCompile(regular)
	res := reg.MatchString(wordStr)
	return res
}
//用户名： /^[a-zA-Z0-9_]{4,16}$/  	/^[a-z0-9_-]{3,16}$/
//密码：	    /^[a-z0-9_]{6,18}$/

//十六进制值：	/^#?([a-f0-9]{6}|[a-f0-9]{3})$/
//电子邮箱	  /^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$/
//Email正则 /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
///^[a-z\d]+(\.[a-z\d]+)*@([\da-z](-[\da-z])?)+(\.{1,2}[a-z]+)+$/
//URL： 	    /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/
//IP 地址：	/((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)/
///^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
//HTML 标签：	/^<([a-z]+)([^<]+)*(?:>(.*)<\/\1>|\s+\/>)$/
//删除代码\\注释：      	(?<!http:|\S)//.*$
//Unicode编码中的汉字范围：	/^[\u2E80-\u9FFF]+$/
//正整数正则  /^\d+$/;   ^\d{4}$
//负整数正则 /^-\d+$/;
//整数正则 /^-?\d+$/;
//正数正则  /^\d*\.?\d+$/;
//负数正则 /^-\d*\.?\d+$/;
//数字正则 /^-?\d*\.?\d+$/;
//手机号正则 /^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$/; ^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$
//身份证号（18位）正则 /^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/;
//ipv4地址正则 /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/;
//RGB Hex颜色正则 /^#?([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$/;
//日期正则，简单判定,未做月份及日期的判定 /^\d{4}(\-)\d{1,2}\1\d{1,2}$/;
//日期正则，复杂判定 /^(?:(?!0000)[0-9]{4}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1[0-9]|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[0-9]{2}(?:0[48]|[2468][048]|[13579][26])|(?:0[48]|[2468][048]|[13579][26])00)-02-29)$/;
//QQ号正则，5至11位 /^[1-9][0-9]{4,10}$/;
//微信号正则，6至20位，以字母开头，字母，数字，减号，下划线 /^[a-zA-Z]([-_a-zA-Z0-9]{5,19})+$/;
//车牌号正则 /^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}[A-Z0-9]{4}[A-Z0-9挂学警港澳]{1}$/;
//包含中文正则 /[\u4E00-\u9FA5]/;

