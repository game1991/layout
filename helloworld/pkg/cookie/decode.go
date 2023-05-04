package tool

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

// DecodeCookie 解析http请求头信息cookie获取sessionID
func DecodeCookie(cookie, sessionName string, keyPairs ...[]byte) (string, error) {
	req := &http.Request{Header: make(http.Header)}
	req.Header.Add("Cookie", cookie)

	c, err := req.Cookie(sessionName)
	if err != nil {
		return "", err
	}
	var sessionID string
	if err := securecookie.DecodeMulti(sessionName, c.Value, &sessionID, securecookie.CodecsFromPairs(keyPairs...)...); err != nil {
		return "", err
	}
	return sessionID, nil
}
