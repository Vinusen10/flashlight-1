package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/leesper/couchdb-golang"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func (u User) AddUser() (err error) {

	//Checks if username already exits

	//Generate a hashed Password
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	b64HashedPW := base64.StdEncoding.EncodeToString(hashedPW)

	u.Password = b64HashedPW
	u.Type = "user"

	// Convert User struct to map[string]interface as required by Save() method
	newUser, err := user2Map(u)

	if err != nil {
		log.Print(err)
	}
	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(newUser, "_id")
	delete(newUser, "_rev")

	// Add newUser to DB
	_, _, errr := flashlightDB.Save(newUser, nil)

	if errr != nil {
		fmt.Printf("[Add] error: %s", errr)
	}

	return err

}

func GetUserByMail(mailUsername string) (user User, err error) {

	return user, err
}

func UserExist(mailUsername string) bool {
	query := `{
		"selector": {
			 "type": "User",
			 "email": "%s"
		}
	}`
	u, _ := flashlightDB.QueryJSON(fmt.Sprintf(query, mailUsername))
	if len(u) < 1 {
		return false
	}
	return true
}

func CheckPassword(username, pw string) bool {
	user, err := GetUserByMail(username)
	passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordDB, []byte(pw))
	if err == nil {
		return true
	} else {
		return false
	}
}

/* Converts Mail to Usernamr : E.G Bylat93 -> [Bylat93]@examplemail.com
 */
func MailToUsername(mail string) {

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
