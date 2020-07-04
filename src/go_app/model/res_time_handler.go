package model

import (
	"go_app/consts"
	"time"
)

//timestampがjs側でNullで登録される場合の値をレスポンスで返すときに中身をnullに変える
func ResTimeHandler(t *time.Time) *time.Time {
	time, _ := time.Parse(time.RFC3339, consts.Consts.ClientNullTimeStr)
	if t == nil || t.Equal(time) {
		return nil
	}
	return t
}
