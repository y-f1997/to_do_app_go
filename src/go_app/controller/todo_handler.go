package controller

import (
	"encoding/json"
	"fmt"
	"go_app/model"
	"io/ioutil"
	"net/http"
	"time"
)

func GetTodo(r *http.Request) (*model.TodoRes, *JSONError) {
	m, err401 := AuthByJwt(r)
	if err401 != nil {
		return nil, err401
	}
	userId := m["user_id"].(string)

	query := r.URL.Query()["keyTime"]
	if query == nil {
		return nil, &JSONError{nil, "検索キーがありません", http.StatusBadRequest}
	}
	keyTime, _ := time.Parse(time.RFC3339, query[0])
	todoRes, err := model.FindToDoById(keyTime, userId)
	if err != nil {
		return nil, &JSONError{err, "データが見つかりません", http.StatusBadRequest}
	}
	return todoRes, nil
}

func GetToDoListService(r *http.Request) ([]*model.TodoRes, *JSONError) {
	m, err401 := AuthByJwt(r)
	if err401 != nil {
		return nil, err401
	}
	requestUserId := m["user_id"].(string)

	toDoListRes, err := model.FindToDoByUserId(requestUserId)
	if err != nil {
		return nil, &JSONError{err, "検索に失敗しました", http.StatusBadRequest}
	}
	return toDoListRes, nil
}

func UpdateToDoService(r *http.Request) *JSONError {
	m, err401 := AuthByJwt(r)
	if err401 != nil {
		return err401
	}
	requestUserId := m["user_id"].(string)
	// リクエストボディを読み取る
	reqJsonByte, _ := ioutil.ReadAll(r.Body)
	// リクエストをstructにパースする ...1
	var updateTodoReq *model.TodoReq
	json.Unmarshal(reqJsonByte, &updateTodoReq)
	updateTodoReq.UserId = requestUserId

	// パースしたstructにはuserIdが格納されていないのでcookieから格納する
	// CRUD操作を実行するstructにパースする
	updateTodoReqDao := updateTodoReq.TodoReqToTodoDaoConverter()
	if err := updateTodoReqDao.Update(); err != nil {
		return &JSONError{err, fmt.Sprintf("更新に失敗しました。エラーメッセージ:%v", err), http.StatusBadRequest}
	}
	return nil
}

func InsertTodoService(r *http.Request) *JSONError {
	m, err401 := AuthByJwt(r)
	if err401 != nil {
		return err401
	}
	requestUserId := m["user_id"].(string)
	// read request body
	reqJsonByte, _ := ioutil.ReadAll(r.Body)
	// parse to ReqTodo struct
	var reqTodo *model.TodoReq
	// emb userid from cookie
	json.Unmarshal(reqJsonByte, &reqTodo)
	reqTodo.UserId = requestUserId
	reqTodo.CrtTimestamp = time.Now()

	// parse struct for dao
	reqTodoDao := reqTodo.TodoReqToTodoDaoConverter()
	if err := reqTodoDao.Insert(); err != nil {
		return &JSONError{err, fmt.Sprintf("登録に失敗しました。エラーメッセージ:%v", err), http.StatusBadRequest}
	}
	return nil
}

func DeleteTodoService(r *http.Request) *JSONError {
	m, err401 := AuthByJwt(r)
	if err401 != nil {
		return err401
	}
	requestUserId := m["user_id"].(string)
	// read request body
	reqJsonByte, _ := ioutil.ReadAll(r.Body)
	// parse request struct
	var reqTodo *model.TodoReq
	json.Unmarshal(reqJsonByte, &reqTodo)
	reqTodo.UserId = requestUserId
	// parse dao struct
	reqTodoDao := reqTodo.TodoReqToTodoDaoConverter()

	if err := reqTodoDao.Delete(); err != nil {
		return &JSONError{err, fmt.Sprintf("削除に失敗しました。エラーメッセージ:%v", err), http.StatusBadRequest}
	}
	return nil
}
