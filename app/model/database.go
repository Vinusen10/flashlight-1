package model

import "github.com/leesper/couchdb-golang"

var (
	userFlashlightDB, imageFlashlightDB *couchdb.Database
)

func init() {
	var err error
	userFlashlightDB, err = couchdb.NewDatabase("http://localhost:5984/sportsmap")
	imageFlashlightDB, err = couchdb.NewDatabase("")
	if err != nil {
		panic(err)
	}

}
