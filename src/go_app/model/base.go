package model

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

var DbMap *gorp.DbMap

func init() {

	DbMap = initDb()
}

func initDb() *gorp.DbMap {
	var err error
	log.Println("started: DB connection...")
	dbConnection, err := sql.Open("postgres", "user=postgres host=db dbname=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	timeZoneSettingCmd := `ALTER DATABASE postgres SET timezone TO 'Asia/Tokyo'; 
												COMMIT; SELECT pg_reload_conf(); 
												COMMIT;`
	_, err = dbConnection.Exec(timeZoneSettingCmd)
	if err != nil {
		log.Println(err)
	}

	cmd := `CREATE TABLE IF NOT EXISTS users (user_id varchar(100) primary key, password varchar(60), name varchar(100), crt_timestamp timestamp with time zone ); COMMIT;`
	cmd1 := `CREATE TABLE IF NOT EXISTS what_done (user_id varchar(100), crt_timestamp timestamp with time zone, category int, what_done varchar(400), additional_info varchar(400) , start_date_time timestamp with time zone, end_date_time timestamp with time zone, PRIMARY KEY(user_id, crt_timestamp), FOREIGN KEY(user_id) REFERENCES public.users(user_id) ); COMMIT;`
	cmd2 := `CREATE TABLE IF NOT EXISTS to_do (user_id varchar(100), crt_timestamp timestamp with time zone,importance int ,category int, to_do varchar(400), time_to_be_done timestamp with time zone, time_done timestamp with time zone, PRIMARY KEY(user_id, crt_timestamp), FOREIGN KEY(user_id) REFERENCES public.users(user_id) ); COMMIT;`

	_, err = dbConnection.Exec(cmd)
	if err != nil {
		log.Println(err)
	}
	_, err = dbConnection.Exec(cmd1)
	if err != nil {
		log.Println(err)
	}
	_, err = dbConnection.Exec(cmd2)
	if err != nil {
		log.Println(err)
	}

	dbmap := &gorp.DbMap{Db: dbConnection, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(WhatDoneDB{}, "what_done")
	dbmap.AddTableWithName(User{}, "users")
	dbmap.AddTableWithName(Todo{}, "to_do")

	return dbmap

}
