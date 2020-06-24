package model

import "github.com/leesper/couchdb-golang"

var (
	flashlightDB *couchdb.Database
)

func init() {
	var err error
	flashlightDB, err = couchdb.NewDatabase("http://127.0.0.1:5984/flashlight_db")
	if err != nil {
		panic(err)
	}
}
