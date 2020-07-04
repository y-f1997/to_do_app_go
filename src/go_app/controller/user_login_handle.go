package controller

import (
	"encoding/json"
	"go_app/model"
	"io/ioutil"
	"net/http"
)

func UserLoginHandleService(r *http.Request) (*model.User, *JSONError) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, &JSONError{err, "リクエストの読み込みに失敗しました", http.StatusInternalServerError}
	}
	requestJsonBytes := []byte(body)

	var requestUser model.User
	if err := json.Unmarshal(requestJsonBytes, &requestUser); err != nil {
		return nil, &JSONError{err, "JSONリクエストのパースに失敗しました", http.StatusInternalServerError}
	}
	loginUser, err := requestUser.UserLogin()
	if err != nil {
		return nil, &JSONError{err, "ユーザーIDまたはパスワードが誤っています", http.StatusBadRequest}
	}
	return loginUser, nil

}
