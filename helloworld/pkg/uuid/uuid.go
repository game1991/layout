package uuid

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/google/uuid"
)

// UUID 生成唯一id，目前使用uuid
func UUID() string {
	return uuid.New().String()
}

// UUID32 生成没有 - 的uuid
func UUID32() string {
	u := uuid.New()
	return uuid32(&u)
}

func uuid32(u *uuid.UUID) string {
	var buf = make([]byte, 32)
	hex.Encode(buf, u[:4])
	hex.Encode(buf[8:12], u[4:6])
	hex.Encode(buf[12:16], u[6:8])
	hex.Encode(buf[16:20], u[8:10])
	hex.Encode(buf[20:], u[10:])
	return string(buf[:])
}

var uuid22 *base64.Encoding

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// UUID22 使用base64编码的uuid字符串(精简10Byte的存储空间)
func UUID22() string {
	u := uuid.New()
	return uuid22.EncodeToString(u[:])
}

func init() {
	uuid22 = base64.StdEncoding.WithPadding(base64.NoPadding)
}
