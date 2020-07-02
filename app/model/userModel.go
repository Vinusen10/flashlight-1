package model

import (
	"encoding/json"
	"github.com/leesper/couchdb-golang"
)

type User struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Type     string `json:"type"`
	Password string `json:"password"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	couchdb.Document
}

//TODO
func AddUser() {

}

func GetUserbyMail(mailUsername string) (user User, err error) {

	return
}

func UserExist(mailUsername string) (status bool, err error) {
	return
}

/* Converts Mail to Usernamr : E.G Bylat93 -> [Bylat93]@examplemail.com
 */
func MailToUsername(mail string) {

}

func CheckPassword(pw string) (status bool, err error) {
	return
}

/*-------Helper Function--------*/

// Convert from map[string]interface{} to User Struct as required by golang-couchdb method
func map2User(user map[string]interface{}) (u User, err error) {
	uJSON, err := json.Marshal(user)
	json.Unmarshal(uJSON, &u)
	return u, err
}

// Convert from User Struct to map[string]interface{}  as required by golang-couchdb method
func user2Map(u User) (user map[string]interface{}, err error) {
	uJSON, err := json.Marshal(u)
	json.Unmarshal(uJSON, &user)
	return user, err
}
