package tool

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/game1991/layout/helloworld/internal/conf"
	"github.com/gorilla/securecookie"
)

func TestDecodeCookie(t *testing.T) {
	type args struct {
		cookie      string
		sessionName string
		keyPairs    [][]byte
	}

	c := &conf.Session{Secret: []string{"secretsession"}}
	encode, err := securecookie.EncodeMulti("api", "sessionID", securecookie.CodecsFromPairs(c.SessionSecret()...)...)
	if err != nil {
		t.Error(err)
	}

	cookie := &http.Cookie{
		Name:       "api",
		Value:      encode,
		Path:       "/",
		Expires:    time.Now().Add(20 * time.Minute),
		RawExpires: time.Now().Add(20 * time.Minute).Format("Mon, 02-Jan-2006 15:04:05 MST"),
		MaxAge:     int(time.Now().Add(20 * time.Minute).Unix()),
		Secure:     false,
		HttpOnly:   true,
		SameSite:   http.SameSiteStrictMode,
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试拿到sessionID",
			args: args{
				cookie:      cookie.String(),
				sessionName: "api",
				keyPairs:    c.SessionSecret(),
			},
			want:    "sessionID",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeCookie(tt.args.cookie, tt.args.sessionName, tt.args.keyPairs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("DecodeCookie() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
