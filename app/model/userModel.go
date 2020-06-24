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

// At the moment, there is no use case for this function
func DeleteUser() {

}

func GetUserbyUsername(username string) (user User, err error) {

	return
}

func UserExist(username string) (status bool, err error) {
	return
}

func checkPassword() {

}
