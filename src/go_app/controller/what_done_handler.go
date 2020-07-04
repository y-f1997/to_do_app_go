package controller

import (
	"encoding/json"
	"fmt"
	"go_app/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetListByMonthService(r *http.Request, cookieUserId string) ([]model.WhatDoneRes, *JSONError) {
	querySpanStart := r.URL.Query()["spanStart"]
	querySpanEnd := r.URL.Query()["spanEnd"]

	var month string
	var spanStart time.Time
	var spanEnd time.Time
	if len(querySpanStart) == 0 || len(querySpanStart) == 0 {
		spanStart = time.Now()
		spanEnd = time.Now()
	} else {
		spanStart, _ = time.Parse(time.RFC3339, querySpanStart[0])
		spanEnd, _ = time.Parse(time.RFC3339, querySpanEnd[0])
	}
	log.Println(month)
	whatDoneRes, err := model.FindListByMonth(spanStart, spanEnd, cookieUserId)
	if err != nil {
		return nil, &JSONError{err, "検索に失敗しました", http.StatusNotFound}
	}
	return whatDoneRes, nil
}

func DeleteWhatDoneService(r *http.Request, cookieUserId string) *JSONError {
	body, _ := ioutil.ReadAll(r.Body)
	bodyBytes := []byte(body)
	var deletingWhatDone *model.WhatDoneDB
	json.Unmarshal(bodyBytes, &deletingWhatDone)

	deletingWhatDone.UserId = cookieUserId
	if err := deletingWhatDone.Delete(); err != nil {
		return &JSONError{err, "削除に失敗しました", http.StatusBadRequest}
	}

	return nil
}

func GetWhatDoneDetailService(r *http.Request) (*model.WhatDoneDB, *JSONError) {
	query := r.URL.Query()
	m, err401 := AuthByJwt(r)
	if err401 != nil {
		return nil, err401
	}
	cookieUserId := m["user_id"].(string)

	timeKeyStr := query["whatDoneKey"][0]
	timeKeyTime, _ := time.Parse(time.RFC3339, timeKeyStr)
	var reqWhatDone model.WhatDoneDB
	reqWhatDone.CrtTimestamp = timeKeyTime
	res, err := reqWhatDone.FindById(timeKeyTime, cookieUserId)
	if err != nil {
		return nil, &JSONError{err, "検索に失敗しました", http.StatusNotFound}
	}
	return res, nil
}

func SaveWhatDoneService(r *http.Request, cookieUserId string) *JSONError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &JSONError{err, "リクエストの読み込みに失敗しました", http.StatusInternalServerError}
	}
	reqByte := []byte(body)
	var reqWhatDone *model.WhatDoneDB
	if err := json.Unmarshal(reqByte, &reqWhatDone); err != nil {
		return &JSONError{err, "リクエストのパースに失敗しました", http.StatusInternalServerError}
	}
	reqWhatDone.UserId = cookieUserId
	reqWhatDone.CrtTimestamp = time.Now()
	if err := reqWhatDone.Insert(); err != nil {
		return &JSONError{err, fmt.Sprintf("データのインサートに失敗しました。エラーメッセージ: %s", err), http.StatusInternalServerError}
	}
	return nil
}

func UpdateWhatDoneService(r *http.Request, cookieUserId string) *JSONError {
	body, _ := ioutil.ReadAll(r.Body)
	reqBodyBytes := []byte(body)
	var whatDone *model.WhatDoneDB
	if err := json.Unmarshal(reqBodyBytes, &whatDone); err != nil {
		return &JSONError{err, "リクエストデータのパースに失敗しました", http.StatusInternalServerError}
	}
	whatDone.UserId = cookieUserId
	if err := whatDone.Update(); err != nil {
		return &JSONError{err, "アップデートに失敗しました", http.StatusBadRequest}
	}
	return nil
}
