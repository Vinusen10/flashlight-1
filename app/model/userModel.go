package model

import "github.com/leesper/couchdb-golang"

type User struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Type     string `json:"type"`
	Username string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`

	couchdb.Document
}

//TODO
func AddUser() {

}

func DeleteUser() {

}

func GetUserbyUsername() {

}

func UserExist() {

}

func checkPassword() {

}
