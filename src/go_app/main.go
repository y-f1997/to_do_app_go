package main

import (
	"crypto/rand"
	"encoding/binary"
	"go_app/consts"
	"go_app/controller"
	"strconv"
	"time"
)

func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func main() {
	consts.Consts.SecretKey = random()
	consts.Consts.CookieKey = "isOnSession"
	consts.Consts.Location = "Asia/Tokyo"
	consts.Consts.TimeFormat = map[string]string{"1": "2020-01-01", "2": "2020-02-01", "3": "2020-03-01", "4": "2020-04-01", "5": "2020-05-01", "6": "2020-06-01", "7": "2020-07-01", "8": "2020-08-01", "9": "2020-09-01", "10": "2020-10-01", "11": "2020-11-01", "12": "2020-12-01"}
	// js側でtimestampがNullで登録される場合の値(structにparseする段階でデフォルトの値が入ってしまう)
	consts.Consts.ClientNullTimeStr = "0001-01-01 09:18:59 +0918 LMT"
	loc, _ := time.LoadLocation(consts.Consts.Location)
	time.Local = loc
	controller.WebServer()
}

func CookieKey() string {
	return "isOnSession"
}
