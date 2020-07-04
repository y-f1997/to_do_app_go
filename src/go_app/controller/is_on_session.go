package controller

import (
	"encoding/json"
	"go_app/model"
	"log"
	"net/http"
)

// map, err
func IsOnSessionService(r *http.Request) ([]byte, *JSONError) {
	cookie, err := r.Cookie("isOnSession")
	isOnSessionMessage := make(map[string]bool)

	isOnSessionMessage["onSession"] = false

	if err != nil {
		resJson, err := json.Marshal(isOnSessionMessage)
		if err != nil {
			return nil, &JSONError{err, "JSONのパースに失敗しました", 500}
		}
		return resJson, nil
	}
	log.Println(model.UserCookieCache)
	for _, v := range model.UserCookieCache {
		if v == cookie.Value {
			isOnSessionMessage["onSession"] = true
			resJson, err := json.Marshal(isOnSessionMessage)
			if err != nil {
				return nil, &JSONError{err, "JSONのパースに失敗しました", 500}
			}
			return resJson, nil
		}
	}

	resJson, err := json.Marshal(isOnSessionMessage)
	if err != nil {
		return nil, &JSONError{err, "JSONのパースに失敗しました", 500}
	}
	return resJson, nil

}
