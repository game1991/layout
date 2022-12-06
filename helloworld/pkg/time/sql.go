package time

import (
	"database/sql"
	"time"
)

// ParseStringToNullTime 解析字符串到sql nulltime
func ParseStringToNullTime(in *string) (nt sql.NullTime, err error) {
	if in == nil {
		return
	}

	var t time.Time
	if t, err = time.ParseInLocation("2006-01-02 15:04:05", *in, time.Local); err != nil {
		return
	}
	err = nt.Scan(t)
	return
}
