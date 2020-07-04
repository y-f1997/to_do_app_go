package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId       string    `db:"user_id, primarykey"`
	Password     string    `db:"password"`
	Name         string    `db:"name"`
	CrtTimestamp time.Time `db:"crt_timestamp"`
}

var UserCookieCache []string

func NewUser(userId, password, name string) *User {
	return &User{
		userId,
		password,
		name,
		time.Now(),
	}
}

func (u *User) UserSignUp() error {
	log.Println(u.Password)
	u.Password = hash(u.Password)
	u.CrtTimestamp = time.Now()

	if err := DbMap.Insert(u); err != nil {
		return err
	}
	return nil
}

func (isUser *User) UserLogin() (*User, error) {
	var dbUser User
	DbMap.SelectOne(&dbUser, `select * from users where user_id=$1`, isUser.UserId)
	log.Println(dbUser.Password)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(isUser.Password))
	if err == nil {
		return &dbUser, nil
	} else {
		log.Println(err)
		return nil, err
	}
}

func (u *User) UserWithdraw() {

}

func hash(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}
