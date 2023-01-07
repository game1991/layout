package jwt

import "testing"

func TestJWT(t *testing.T) {
	jwtSecrect := "122344"
	info := make(map[string]interface{})
	info["username"] = "ut12344456"
	info["age"] = 123445
	// 加密
	code, err := JWTEncode(jwtSecrect, info)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(code)
	// 解密失败
	errJwtSecret := "1223456"
	res, err := JWTDecode(errJwtSecret, code)
	if err != nil {
		t.Log(err)
	}
	// 正确解密
	res, err = JWTDecode(jwtSecrect, code)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(res)

}
