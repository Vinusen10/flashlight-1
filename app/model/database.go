package model

import "github.com/leesper/couchdb-golang"

var (
	userFlashlightDB, flashlightDB *couchdb.Database
)

func init() {
	var err error
	userFlashlightDB, err = couchdb.NewDatabase("http://127.0.0.1:5984/flashlight_db")
	flashlightDB, err = couchdb.NewDatabase("http://127.0.0.1:5984/flashlight_user_db")
	if err != nil {
		panic(err)
	}
}
