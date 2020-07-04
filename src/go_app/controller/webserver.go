package controller

import (
	"encoding/json"
	"go_app/model"
	"log"
	"net/http"
)

func initialHandle() http.Handler {

	return http.FileServer(http.Dir("./views/dist"))
}

func isOnSession(w http.ResponseWriter, r *http.Request) {
	// service層はjsonまたはエラーを返すところまで
	// service層はリクエストハンドラー
	// cookieフィールドがなかった場合の処理
	resJson, err := IsOnSessionService(r)
	if err != nil {
		APIError(w, err)
		return
	}

	w.Write(resJson)

}

func userLoginHandle(w http.ResponseWriter, r *http.Request) {

	loginUser, loginErr := UserLoginHandleService(r)

	if loginErr != nil {
		APIError(w, loginErr)
	} else {
		// token生成, tokenはどこかにキャッシュされる？
		if err := GenerateJWT(w, loginUser.UserId); err != nil {
			APIError(w, err)
		}
	}

}

func userSignUpHandle(w http.ResponseWriter, r *http.Request) {
	if err := UserSignUpHandleService(r); err != nil {
		APIError(w, err)
	}
}

func getList(w http.ResponseWriter, r *http.Request) {
	m, err := AuthByJwt(r)
	if err != nil {
		APIError(w, err)
		return
	}
	cookieUserId := m["user_id"].(string)

	w.Header().Set("Content-Type", "application.json;charset=utf-8")
	list := make(map[string][]model.WhatDoneRes)
	list["list"], _ = GetListByMonthService(r, cookieUserId)

	resJson, _ := json.Marshal(list)
	w.Write(resJson)
}

func saveWhatDone(w http.ResponseWriter, r *http.Request) {
	m, err := AuthByJwt(r)
	if err != nil {
		APIError(w, err)
		return
	}
	cookieUserId := m["user_id"].(string)

	if err := SaveWhatDoneService(r, cookieUserId); err != nil {
		APIError(w, err)
	}
}

func updateWhatDone(w http.ResponseWriter, r *http.Request) {
	m, err := AuthByJwt(r)
	if err != nil {
		APIError(w, err)
		return
	}
	cookieUserId := m["user_id"].(string)

	if err := UpdateWhatDoneService(r, cookieUserId); err != nil {
		APIError(w, err)
	}
}

func deleteWhatDone(w http.ResponseWriter, r *http.Request) {
	m, err := AuthByJwt(r)
	if err != nil {
		APIError(w, err)
		return
	}
	cookieUserId := m["user_id"].(string)

	if err := DeleteWhatDoneService(r, cookieUserId); err != nil {
		APIError(w, err)
	}

}

func getWhatDoneDetail(w http.ResponseWriter, r *http.Request) {
	detailObj, err := GetWhatDoneDetailService(r)
	if err != nil {
		APIError(w, err)
		return
	}
	resJson, _ := json.Marshal(detailObj)
	w.Write(resJson)
}

func getToDoListByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	todoAllRes, err := GetToDoListService(r)
	if err != nil {
		APIError(w, err)
		return
	}
	todoAllResJson, _ := json.Marshal(todoAllRes)
	w.Write(todoAllResJson)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	todoRes, err := GetTodo(r)
	if err != nil {
		APIError(w, err)
		return
	}
	todoResJson, _ := json.Marshal(todoRes)
	w.Write(todoResJson)
}

func insertToDo(w http.ResponseWriter, r *http.Request) {
	if err := InsertTodoService(r); err != nil {
		APIError(w, err)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	if err := UpdateToDoService(r); err != nil {
		APIError(w, err)
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if err := DeleteTodoService(r); err != nil {
		APIError(w, err)
	}
}

type JSONError struct {
	error
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func apiMakeHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func APIError(w http.ResponseWriter, jsonErr *JSONError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(jsonErr.Code)
	jsonErrStr, err := json.Marshal(jsonErr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	w.Write(jsonErrStr)
}

func WebServer() {
	http.Handle("/", initialHandle())
	http.Handle("/signUp", apiMakeHandler(userSignUpHandle))

	http.Handle("/getList", apiMakeHandler(getList))

	http.Handle("/saveWhatDone", apiMakeHandler(saveWhatDone))
	http.Handle("/updateWhatDone", apiMakeHandler(updateWhatDone))
	http.Handle("/getWhatDoneDetail", apiMakeHandler(getWhatDoneDetail))
	http.Handle("/deleteWhatDoneDetail", apiMakeHandler(deleteWhatDone))

	http.Handle("/saveTodo", apiMakeHandler(insertToDo))
	http.Handle("/updateTodo", apiMakeHandler(updateTodo))
	http.Handle("/deleteTodo", apiMakeHandler(deleteTodo))
	http.Handle("/getTodoList", apiMakeHandler(getToDoListByUserId))
	http.Handle("/getTodo", apiMakeHandler(getTodo))

	http.Handle("/checkIsOnSession", apiMakeHandler(isOnSession))
	http.HandleFunc("/login", apiMakeHandler(userLoginHandle))
	log.Fatalln(http.ListenAndServe(":5080", nil))
}
