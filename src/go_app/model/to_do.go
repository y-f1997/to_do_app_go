package model

import (
	"log"
	"time"
)

var CategoryMapTodo = map[int]string{1: "仕事", 2: "勉強", 3: "生活", 4: "目標"}
var ImportanceMapTodo = map[int]string{1: "最重要", 2: "重要", 3: "不急", 4: "余力"}

type Todo struct {
	UserId       string     `db:"user_id, primarykey"`
	CrtTimestamp time.Time  `db:"crt_timestamp, primarykey,notnull"`
	Importance   int        `db:"importance"`
	Category     int        `db:"category"`
	Todo         string     `db:"to_do"`
	TimeToBeDone *time.Time `db:"time_to_be_done,notnull"`
	TimeDone     *time.Time `db:"time_done,notnull"`
}

type TodoReq struct {
	UserId       string     `json:"userId"`
	CrtTimestamp time.Time  `json:"crtTimestamp"`
	Importance   int        `json:"importance,string"`
	Category     int        `json:"category,string"`
	Todo         string     `json:"toDo"`
	TimeToBeDone *time.Time `json:"timeToBeDone"`
	TimeDone     *time.Time `json:"timeDone"`
}

type TodoRes struct {
	UserId            string     `json:"userId"`
	CrtTimestamp      time.Time  `json:"crtTimestamp"`
	Importance        int        `json:"importance,string"`
	Category          int        `json:"category,string"`
	Todo              string     `json:"toDo"`
	TimeToBeDone      *time.Time `json:"timeToBeDone"`
	TimeDone          *time.Time `json:"timeDone"`
	TodoImportanceRes string     `json:"todoImportanceRes"`
	TodoCategoryRes   string     `json:"todoCategoryRes"`
}

//insert用
func NewTodo(userId string, category int, importance int, todo string, timeToBeDone *time.Time, timeDone *time.Time) *Todo {
	return &Todo{
		UserId:       userId,
		CrtTimestamp: time.Now(),
		Importance:   importance,
		Category:     category,
		Todo:         todo,
		TimeToBeDone: timeToBeDone,
		TimeDone:     timeDone,
	}
}

// リクエストstructからDao structにconvert
func (o *TodoReq) TodoReqToTodoDaoConverter() *Todo {
	return &Todo{o.UserId, o.CrtTimestamp, o.Importance, o.Category, o.Todo, o.TimeToBeDone, o.TimeDone}
}

// コンストラクタはメソッドで返す
func NewTodoRes(userId string, crtTimestamp time.Time, importance int, category int, todo string, timeToBeDone *time.Time, timeDone *time.Time) *TodoRes {
	return &TodoRes{
		UserId:            userId,
		CrtTimestamp:      crtTimestamp,
		Importance:        importance,
		Category:          category,
		Todo:              todo,
		TimeToBeDone:      timeToBeDone,
		TimeDone:          timeDone,
		TodoImportanceRes: ImportanceMapTodo[importance],
		TodoCategoryRes:   CategoryMapTodo[category],
	}
}

//レスポンスをカスタマイズしたいときはNewTodoResのみいじる
func (o *Todo) TodoDaoToToDoResConverter() *TodoRes {
	return NewTodoRes(o.UserId, o.CrtTimestamp, o.Importance, o.Category, o.Todo, o.TimeToBeDone, o.TimeDone)
}

func (o *Todo) Insert() error {
	if err := DbMap.Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *Todo) Update() error {
	if _, err := DbMap.Update(o); err != nil {
		return err
	}
	return nil
}

func (o *Todo) Delete() error {
	if _, err := DbMap.Delete(o); err != nil {
		return err
	}
	return nil
}

func FindToDoById(keyTime time.Time, userId string) (*TodoRes, error) {
	var todo *Todo
	err := DbMap.SelectOne(&todo, `SELECT * FROM to_do WHERE crt_timestamp=$1 and user_id=$2`, keyTime, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	todoRes := todo.TodoDaoToToDoResConverter()
	//デフォルトのtimeの値をnullに変換(デフォルトでない場合はそのまま)
	todoRes.TimeToBeDone = ResTimeHandler(todoRes.TimeToBeDone)
	todoRes.TimeDone = ResTimeHandler(todoRes.TimeDone)

	return todoRes, nil
}

func FindToDoByUserId(userId string) ([]*TodoRes, error) {
	var todoAll []Todo
	var todoAllRes []*TodoRes
	_, err := DbMap.Select(&todoAll, `SELECT * FROM to_do WHERE user_id=$1 ORDER BY time_to_be_done`, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, todo := range todoAll {
		todoRes := todo.TodoDaoToToDoResConverter()
		//デフォルトのtimeの値をnullに変換(デフォルトでない場合はそのまま)
		todoRes.TimeToBeDone = ResTimeHandler(todoRes.TimeToBeDone)
		todoRes.TimeDone = ResTimeHandler(todoRes.TimeDone)
		todoAllRes = append(todoAllRes, todoRes)
	}
	return todoAllRes, nil
}
