package regexp

import "testing"

func TestValidPhone(t *testing.T) {
	t.Log(ValidPhone("18812345678"))
}
func TestErrValidPhone(t *testing.T) {
	t.Log(ValidPhone(""))
	t.Log(ValidPhone("28812345678"))
	t.Log(ValidPhone("1812345678"))
	t.Log(ValidPhone("188112345678"))
	t.Log(ValidPhone("1212345678"))
}

func TestValidPhoneStrict(t *testing.T) {
	t.Log(ValidPhoneStrict("18812345678"))
	t.Log(ValidPhoneStrict("18612345678"))
}
func TestErrValidPhoneStrict(t *testing.T) {
	t.Log(ValidPhoneStrict(""))
	t.Log(ValidPhoneStrict("14123456789"))
	t.Log(ValidPhoneStrict("15423456789"))
	t.Log(ValidPhoneStrict("16323456789"))
	t.Log(ValidPhoneStrict("16423456789"))
	t.Log(ValidPhoneStrict("16823456789"))
	t.Log(ValidPhoneStrict("16923456789"))
	t.Log(ValidPhoneStrict("17923456789"))
	t.Log(ValidPhoneStrict("19023456789"))
	t.Log(ValidPhoneStrict("19223456789"))
	t.Log(ValidPhoneStrict("19423456789"))
	t.Log(ValidPhoneStrict("19623456789"))
	t.Log(ValidPhoneStrict("19723456789"))
	t.Log(ValidPhoneStrict("19223456789"))
	t.Log(ValidPhoneStrict("188112345678"))
	t.Log(ValidPhoneStrict("1212345678"))
}

func TestValidEmailStrict(t *testing.T) {
	t.Log(ValidEmailStrict("123@uniontech.com"))
	t.Log(ValidEmailStrict("w-_2-s.3@h-1.com"))
}
func TestErrValidEmailStrict(t *testing.T) {
	t.Log(ValidEmailStrict(""))
	t.Log(ValidEmailStrict("1asd758@as.com@2d.cn"))
}

func TestValidEmail(t *testing.T) {
	t.Log(ValidEmail("123@uniontech.com"))
	t.Log(ValidEmail("w-_2-s.3@h-1.com"))
	t.Log(ValidEmail("'s@'s'"))
	t.Log(ValidEmail("?@?"))
	t.Log(ValidEmail("'@'"))
	t.Log(ValidEmail("\"@\""))
}
func TestErrValidEmail(t *testing.T) {
	t.Log(ValidEmail(""))
	t.Log(ValidEmail("@as.com@2d.cn"))
	t.Log(ValidEmail("@2d.cn"))
	t.Log(ValidEmail(" @2d.cn"))
	t.Log(ValidEmail("a@ 2d.cn"))
	t.Log(ValidEmail("a@	2d.cn"))
	t.Log(ValidEmail("a @2d.cn"))

}

func TestValidUserName(t *testing.T) {
	t.Log(ValidUserName("a2221aas151asdsad1d51as5d_asdssa1"))
}

func TestErrValidUserName1(t *testing.T) {
	t.Log(ValidUserName("a13"))
}

func TestValidPassword(t *testing.T) {
	t.Log(ValidPassword("!$%$12a@#@"))
	t.Log(ValidPasswordNew("aaa!@#$%^&*()_+-{}[]|/?,.<>a"))
	t.Log(ValidPasswordNew("aaaaaa1111"))
	t.Log(ValidPasswordNew("~`!@#$a%^&*()_-+[]{}:\",./?|1"))
	t.Log(ValidPasswordNew("~`!@#$a%^&*()_-+[]{}:\",./?|a"))
}

func TestValidPassword1(t *testing.T) {
	t.Log(ValidPassword(" a 123"))
	t.Log(ValidPassword(""))
	t.Log(ValidPassword("111sssssaaa1223"))
	t.Log(ValidPassword("$%$ a!@#@"))
	t.Log(ValidPassword("1234567812345678123456781234567812345678123456781234567812345678a"))
	t.Log(ValidPassword("123a4不可能5678"))
	t.Log(ValidPassword("9t～！＠＃￥％＾＆＊"))
	t.Log(ValidPassword("9t～！＠＃￥％＾＆＊（）＿＋Ｐ｛｝：＂｜＜＞？"))
}

func TestValidNickname(t *testing.T) {
	t.Log(ValidNickname("ss2d卅回i的啊深!@$^%^12^%**.?阿asisdu后"))
}

func TestErrValidNickname(t *testing.T) {
	t.Log(ValidNickname("1112"))
	t.Log(ValidNickname(""))
	t.Log(ValidNickname("<s你d"))
	t.Log(ValidNickname("s<s哈"))
	t.Log(ValidNickname("s是>d"))
	t.Log(ValidNickname("\"ssd"))
	t.Log(ValidNickname("'s走d"))
	t.Log(ValidNickname("s&sd"))
	t.Log(ValidNickname("s天2 d"))
	t.Log(ValidNickname("ss2d卅回i的啊深!@$^%^12^%**.?阿asisdu后s"))
	t.Log(ValidNickname("怎么肥四"))
}

func TestValidRealName(t *testing.T) {
	t.Log(ValidRealName("Machal Jordon"))
	t.Log(ValidRealName("Machal·Jackson"))
	t.Log(ValidRealName("迪丽热巴●古力娜扎"))
	t.Log(ValidRealName(".Machal"))
	t.Log(ValidRealName("迪丽热巴●"))
	t.Log(ValidRealName("迪丽热巴哈迪斯"))
	t.Log(ValidRealName("a哈哈"))
	t.Log(ValidRealName("所啊@13.、‘？"))
	t.Log(ValidRealName("e?!@#"))
	t.Log(ValidRealName(" "))
}

func TestValidWebsite(t *testing.T) {
	t.Log(ValidWebsite(".s1:2wa-s_2/"))
	t.Log(ValidWebsite("1s1:2wa-s_2./"))
	t.Log(ValidWebsite("1s1:2wa-s_2./ "))
	t.Log(ValidWebsite("WWW.娃哈哈。com"))
	t.Log(ValidWebsite(" "))
	t.Log(ValidWebsite("http://lizhouquan.abcdefgh123456789.com"))
	t.Log(ValidWebsite("www.lizhouquan.a@?.com"))
}

func TestValidWebsiteStrict(t *testing.T) {
	t.Log(ValidWebsiteStrict(".s1:2wa-s_2/"))
	t.Log(ValidWebsiteStrict("http://1s12wa-s_2./"))
	t.Log(ValidWebsiteStrict("http://1s1:2wa-s_2./ "))
	t.Log(ValidWebsiteStrict("WWW.娃哈哈。com"))
	t.Log(ValidWebsiteStrict(" "))
	t.Log(ValidWebsiteStrict("http://lizhouquan.abcdefgh123456789.com"))
	t.Log(ValidWebsiteStrict("https://中国.com"))
	t.Log(ValidWebsiteStrict("http://www.lizh.@&?#+~:^=ouquan.a@?.com"))
	t.Log(ValidWebsiteStrict("https://cooperation.uniontech.com:443/file/mdoc/doc/202109/odSKdbmAsbWMuPD_2556239816.deb?attname=signed_com.decard.dcrf32-1.0.0(1).deb%26e=1651259594%26token=storage:wyR-fE_rRWY67cH3OnwDThfj8n8="))
	t.Log(ValidWebsiteStrict("ftps://xxx.xxx.xxx:443/file/sad2/2sdx@3/+o@S$#FS^.de-b?attname=signed_com.decard.dcrf32-1.0.0(1).deb%26e=1651259594%26token=storage:wyR-fE_rRWY67cH3OnwDThfj8n8="))
}

func TestValidWebsiteStrictWithProtocol(t *testing.T) {
	t.Log(ValidWebsiteStrictWithProtocol("WWW.娃哈哈。com"))
	t.Log(ValidWebsiteStrictWithProtocol("hehe.hello.com"))
	t.Log(ValidWebsiteStrictWithProtocol(" "))
	t.Log(ValidWebsiteStrictWithProtocol("http://www.lizh.@&?#+~:^=ouquan.a@?.com"))
	t.Log(ValidWebsiteStrictWithProtocol("https://cooperation.uniontech.com:443/file/mdoc/doc/202109/odSKdbmAsbWMuPD_2556239816.deb?attname=signed_com.decard.dcrf32-1.0.0(1).deb%26e=1651259594%26token=storage:wyR-fE_rRWY67cH3OnwDThfj8n8="))
	t.Log(ValidWebsiteStrictWithProtocol("ftps://xxx.xxx.xxx:443/file/sad2/2sdx@3/+o@S$#FS^.de-b?attname=signed_com.decard.dcrf32-1.0.0(1).deb%26e=1651259594%26token=storage:wyR-fE_rRWY67cH3OnwDThfj8n8="))
}

func TestIdentificationID(t *testing.T) {
	t.Log(ValidIdentificationID("110101199609249876"))
}

func TestValidBussinessLicense(t *testing.T) {
	t.Log(ValidBusinessLicense("125478874558774681"))
}

func TestValidChar(t *testing.T) {
	t.Log(ValidChar("test测试"))
	t.Log(ValidChar("测试测试测试"))
	t.Log(ValidChar("测试测试测试 "))
	t.Log(ValidChar(" 测试测试测试"))
	t.Log(ValidChar("测试测 试测试"))
	t.Log(ValidChar("测<测试测试测试"))
	t.Log(ValidChar("测>测试测试测试"))
	t.Log(ValidChar("测&测试测试测试"))
	t.Log(ValidChar("测试'试测试测试"))
	t.Log(ValidChar("测试\"测试测试"))
	t.Log(ValidChar("测试“”测试测试"))
	t.Log(ValidChar("测试“”测试\n测试"))
	t.Log(ValidChar("测试“”测试\r测试"))
	t.Log(ValidChar("测试“”测试\t测试"))
}
