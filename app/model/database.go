package model

import "github.com/leesper/couchdb-golang"

var btDB *couchdb.Database

func init(){
	var err error
	btDB, err = couchdb.NewDatabase("http://localhost:5984/sportsmap")
	if err!=nil {
		panic(err)
	}
}