package controller

import (
	"encoding/json"
	"fmt"
	"go_app/model"
	"io/ioutil"
	"net/http"
)

func UserSignUpHandleService(r *http.Request) *JSONError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &JSONError{err, "リクエストの読み込みに失敗しました", http.StatusInternalServerError}
	}
	reqByte := []byte(body)
	var reqUser model.User
	if err := json.Unmarshal(reqByte, &reqUser); err != nil {
		return &JSONError{err, "リクエストのパースに失敗しました", http.StatusInternalServerError}
	}
	if err := reqUser.UserSignUp(); err != nil {
		return &JSONError{err, fmt.Sprintf("ユーザー登録に失敗しました。エラーメッセージ: %v", err), http.StatusBadRequest}
	}
	return nil
}
