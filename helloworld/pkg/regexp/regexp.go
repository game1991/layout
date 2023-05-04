package regexp

import (
	"regexp"
	"strconv"
	"time"
)

var (
	phoneNumberRegexpStrict = regexp.MustCompile(`^(13[0-9]|14[5-9]|15[0-3,5-9]|16[2,5,6,7]|17[0-8]|18[0-9]|19[1,3,5,8,9])\d{8}$`)
	phoneNumberRegexp       = regexp.MustCompile(`^1[3-9]\d{9}$`)
	emailRegexpStrict       = regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	emailRegexp             = regexp.MustCompile(`^[^@\s]+@[^@\s]+$`)
	userNameRegexp          = regexp.MustCompile(`^[a-z][a-zA-Z0-9-_]{2,31}$`)                           //小写字母开头,字母数字横杠下划线3到32位
	passwordRegexp          = regexp.MustCompile(`[A-Za-z].*[0-9]|[0-9].*[A-Za-z]`)                      //必须有字母数字,支持特殊字符汉字(密码长度未做限制,若有需要可自行len()代码控制)
	passwordSpecialRegexp   = regexp.MustCompile("^(\\w|\\+|\\[|\\]|\\-|[~`!@#$%^&()*={}:\"<>,.?|/])+$") //支持特殊所有特殊符号
	passwordCommonRegexp    = regexp.MustCompile(`.*[0-9]+.*[a-zA-Z]+.*|.*[a-zA-Z]+.*[0-9]+.*`)          //必须有字母数字

	nickNameRegexp = regexp.MustCompile(`^[^<>&'"\s]{1,32}$`) //不包含<>&'""空格

	wechatRegexp                    = regexp.MustCompile(`^[a-zA-Z]([-_a-zA-Z0-9]{5,19})+$`)
	qqRegexp                        = regexp.MustCompile(`^[1-9][0-9]{4,19}$`)
	realNameRegexp                  = regexp.MustCompile("^[\u4e00-\u9fa5●•·]*$|^[A-Za-z·\\s]*$")
	formatRegexp                    = regexp.MustCompile(`^(\\d{6})(18|19|20)?(\\d{2})(0\\d|10|11|12)([012]\\d|3[01])(\\d{3})(\\d|X)?$`)
	area                            = map[string]string{"11": "北京", "12": "天津", "13": "河北", "14": "山西", "15": "内蒙", "21": "辽宁", "22": "吉林", "23": "黑龙", "31": "上海", "32": "江苏", "33": "浙江", "34": "安徽", "35": "福建", "36": "江西", "37": "山东", "41": "河南", "42": "湖北", "43": "湖南", "44": "广东", "45": "广西", "46": "海南", "50": "重庆", "51": "四川", "52": "贵州", "53": "云南", "54": "西藏", "61": "陕西", "62": "甘肃", "63": "青海", "64": "宁夏", "65": "新疆", "71": "台湾", "81": "香港", "82": "澳门", "91": "国外"}
	minDate                         = time.Date(1890, 0, 0, 0, 0, 0, 0, time.Local)
	maxDate                         = time.Now()
	weight                          = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	code                            = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	businessLicenseRegexp           = regexp.MustCompile(`^[^_IOZSVa-z\W]{2}\d{6}[^_IOZSVa-z\W]{10}$`)
	websiteRegexp                   = regexp.MustCompile(`^\w+[^\s]+(\.[^\s]+)$`)
	websiteRegexpStrict             = regexp.MustCompile(`^(((ht|f)tps?):\/\/)?[\w-]+(\.[\w-]+)+([\w.,@?^=%&:/~+#-\(\)\-\.]*[\w@?^=%&/~+#-\(\)\-\.])?$`)
	websiteRegexpStrictWithProtocol = regexp.MustCompile(`^(((ht|f)tps?):\/\/)[\w-]+(\.[\w-]+)+([\w.,@?^=%&:/~+#-\(\)\-\.]*[\w@?^=%&/~+#-\(\)\-\.])?$`)
	landlineRegexp                  = regexp.MustCompile(`^\d{1,20}$`)
	// otherRegexp              = regexp.MustCompile("[`~!@$^&*=':;',.?~！@￥……&*‘；：”“'。、？ ]")
	specialCharRegexp = regexp.MustCompile(`^[^<>&'" ]+$`) //不包含这些特殊字符集
)

/*
以下的Valid校验函数,大部分是以入参为非空值前提作为条件判断
部分函数,例如昵称\微信\QQ\网站是允许入参为空值
使用时,请按自己的实际业务场景使用
*/

// ValidPhone 校验手机号格式
func ValidPhone(phone string) bool {
	if phone == "" || !phoneNumberRegexp.MatchString(phone) {
		return false
	}
	return true
}

// ValidPhoneStrict 校验手机号格式严格模式
func ValidPhoneStrict(phone string) bool {
	if phone == "" || !phoneNumberRegexpStrict.MatchString(phone) {
		return false
	}
	return true
}

// ValidEmailStrict 校验手机号格式
func ValidEmailStrict(email string) bool {
	if email == "" || !emailRegexpStrict.MatchString(email) {
		return false
	}
	return true
}

// ValidEmail 校验手机号格式
func ValidEmail(email string) bool {
	if email == "" || !emailRegexp.MatchString(email) {
		return false
	}
	return true
}

// ValidUserName 校验用户名格式
func ValidUserName(username string) bool {
	if username == "" || !userNameRegexp.MatchString(username) {
		return false
	}
	return true
}

// ValidPassword 校验密码格式
func ValidPassword(password string) bool {
	// if password != "" && numberRegexp.MatchString(password) && charRegexp.MatchString(password) && vaildRegexp.MatchString(password) {
	// 	return false
	// }
	if password == "" || !passwordRegexp.MatchString(password) {
		return false
	}
	return true
}

// ValidPasswordNew 校验密码格式新版，允许符号，不再允许汉字
func ValidPasswordNew(password string) bool {
	passwordRaw := []rune(password)
	if len(passwordRaw) < 8 || len(passwordRaw) > 64 {
		return false
	}
	if password == "" || !passwordSpecialRegexp.MatchString(password) || !passwordCommonRegexp.MatchString(password) {
		return false
	}
	return true
}

// ValidRealName 校验姓名
func ValidRealName(realname string) bool {
	if realname == "" || !realNameRegexp.MatchString(realname) {
		return false
	}
	return true
}

// ValidIdentificationID 校验身份证
func ValidIdentificationID(identificationID string) bool {
	if identificationID == "" || !formatRegexp.MatchString(identificationID) {
		return false
	}
	if !validateArea(identificationID) {
		return false
	}
	if !validateBirth(identificationID) {
		return false
	}
	if !validateSum(identificationID) {
		return false
	}

	return true
}

func validateArea(id string) bool {
	if _, ok := area[id[0:2]]; ok {
		return true
	}
	return false
}

func validateBirth(id string) bool {
	birth := id[6:14]
	if date, err := time.Parse("20060102", birth); err != nil {
		return false
	} else if date.After(maxDate) && date.Before(minDate) {
		return false
	}
	return true
}

func validateSum(id string) bool {
	sum := 0
	for i, char := range id[:len(id)-1] {
		charF, _ := strconv.ParseFloat(string(char), 64)
		sum += int(charF) * weight[i]
	}
	return code[sum%11] == id[len(id)-1]
}

// ValidBusinessLicense 校验 统一社会信用代码
func ValidBusinessLicense(businessLicense string) bool {
	if businessLicense == "" || !businessLicenseRegexp.MatchString(businessLicense) {
		return false
	}
	return true
}

// ValidLandline 校验 公司座机
func ValidLandline(landline string) bool {
	if landline == "" || !landlineRegexp.MatchString(landline) {
		return false
	}
	return true
}

// ValidNickname 校验昵称格式
func ValidNickname(nickaname string) bool {
	//此处判断昵称允许为空
	if nickaname == "" || nickNameRegexp.MatchString(nickaname) {
		return true
	}
	return false
}

// ValidWechat 微信校验
func ValidWechat(wechat string) bool {
	//此处判断微信允许为空
	if wechat == "" || wechatRegexp.MatchString(wechat) {
		return true
	}
	return false
}

// ValidQQ QQ校验
func ValidQQ(qq string) bool {
	//此处判断QQ允许为空
	if qq == "" || qqRegexp.MatchString(qq) {
		return true
	}
	return false
}

// ValidWebsite 校验网站
func ValidWebsite(website string) bool {
	if website == "" || websiteRegexp.MatchString(website) {
		return true
	}
	return false
}

// ValidWebsiteStrict 校验网站严格版
func ValidWebsiteStrict(website string) bool {
	if website == "" || websiteRegexpStrict.MatchString(website) {
		return true
	}
	return false
}

// ValidWebsiteStrictWithProtocol 校验网站必须带协议
func ValidWebsiteStrictWithProtocol(website string) bool {
	if website == "" || websiteRegexpStrictWithProtocol.MatchString(website) {
		return true
	}
	return false
}

// ValidChar 特殊字符校验
func ValidChar(char string) bool {
	//如果入参是空对象，则视为通过，当对象不为空时进行正则匹配
	if char == "" || specialCharRegexp.MatchString(char) {
		return true
	}
	return false
}
