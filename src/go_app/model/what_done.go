package model

import (
	"log"
	"time"

	_ "github.com/lib/pq"
)

var CategoryMap = map[int]string{1: "食事", 2: "勉強", 3: "仕事", 4: "遊び"}

type WhatDoneDB struct {
	UserId         string     `db:"user_id" json:"userId"`
	CrtTimestamp   time.Time  `db:"crt_timestamp, primarykey,notnull" json:"crtTimestamp"`
	Category       int        `db:"category" json:"category,string"`
	WhatDone       string     `db:"what_done" json:"whatDone"`
	AdditionalInfo string     `db:"additional_info" json:"additionalInfo"`
	StartDateTime  *time.Time `db:"start_date_time,notnull" json:"startDateTime"`
	EndDateTime    *time.Time `db:"end_date_time,notnull" json:"endDateTime"`
}

type WhatDoneRes struct {
	UserId         string     `json:"userId"`
	CrtTimestamp   time.Time  `json:"crtTimestamp"`
	Category       int        `json:"category,string"`
	WhatDone       string     `json:"whatDone"`
	AdditionalInfo string     `json:"additionalInfo"`
	CategoryRes    string     `json:"categoryRes"`
	StartDateTime  *time.Time `json:"startDateTime"`
	EndDateTime    *time.Time `json:"endDateTime"`
}

func NewWhatDone(userId string, category int, whatDone string, additionalInfo string) WhatDoneDB {
	return WhatDoneDB{
		UserId:         userId,
		CrtTimestamp:   time.Now(),
		Category:       category,
		WhatDone:       whatDone,
		AdditionalInfo: additionalInfo,
	}
}

func NewWhatDoneRes(userId string, crtTimestamp time.Time, category int, whatDone string, additionalInfo string, categoryRes string, startDateTime *time.Time, endDateTime *time.Time) WhatDoneRes {
	return WhatDoneRes{
		UserId:         userId,
		CrtTimestamp:   crtTimestamp,
		Category:       category,
		WhatDone:       whatDone,
		AdditionalInfo: additionalInfo,
		CategoryRes:    categoryRes,
		StartDateTime:  startDateTime,
		EndDateTime:    endDateTime,
	}
}

func (o *WhatDoneDB) Insert() error {
	if err := DbMap.Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *WhatDoneDB) Update() error {
	log.Println(o.StartDateTime)
	_, err := DbMap.Update(o)
	if err != nil {
		return err
	}
	return nil
}

func (o *WhatDoneDB) Delete() error {
	if _, err := DbMap.Delete(o); err != nil {
		return err
	}
	return nil
}

func (o *WhatDoneDB) FindById(keyTime time.Time, cookieKey string) (*WhatDoneDB, error) {
	log.Printf("time=%v\n", keyTime)
	if err := DbMap.SelectOne(&o, "select * from public.what_done where crt_timestamp=$1 and user_id=$2", keyTime, cookieKey); err != nil {
		return nil, err
	}
	o.StartDateTime = ResTimeHandler(o.StartDateTime)
	o.EndDateTime = ResTimeHandler(o.EndDateTime)

	return o, nil
}

func FindListByMonth(spanStart time.Time, spanEnd time.Time, cookieUserId string) ([]WhatDoneRes, error) {
	var lists []WhatDoneDB
	var resLists []WhatDoneRes
	var err error
	log.Printf("spanStart:%v", spanStart.Format(time.RFC3339))
	log.Printf("spanStart:%v", spanEnd.Format(time.RFC3339))
	if spanStart.Equal(spanEnd) {
		log.Println(spanStart.Format(time.RFC3339))
		where := map[string]interface{}{"spanStart": spanStart.Format(time.RFC3339), "UserId": cookieUserId}
		// _, err = DbMap.Select(&lists, `SELECT * from what_done where date_trunc('day', start_date_time) = date_trunc('day', (timestamp :spanStart)) and user_id=:UserId; `, where)
		_, err = DbMap.Select(&lists, `SELECT * from what_done where date(start_date_time) = date(:spanStart) and user_id=:UserId ORDER BY start_date_time; `, where)
	} else {
		_, err = DbMap.Select(&lists, `select * from public.what_done where start_date_time BETWEEN date($1) and (date($2)+1) and user_id=$3 ORDER BY end_date_time;`, spanStart.Format(time.RFC3339), spanEnd.Format(time.RFC3339), cookieUserId)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, v := range lists {
		v.StartDateTime = ResTimeHandler(v.StartDateTime)
		v.EndDateTime = ResTimeHandler(v.EndDateTime)
		convertedRes := NewWhatDoneRes(v.UserId, v.CrtTimestamp, v.Category, v.WhatDone, v.AdditionalInfo, CategoryMap[v.Category], v.StartDateTime, v.EndDateTime)
		resLists = append(resLists, convertedRes)
	}
	log.Println(resLists)
	return resLists, nil
}
