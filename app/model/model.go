package model

import "github.com/leesper/couchdb-golang"

//Models
type (
	User struct {
		ID       string `json:"_id"`
		Rev      string `json:"_rev"`
		Type     string `json:"type"`
		Username string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`

		couchdb.Document
	}

	Picture struct {
		ID        string   `json:"_id"`
		Rev       string   `json:"_rev"`
		Type      string   `json:"type"`
		Creator   string   `json:"creator"`
		Comments  []string `json:"comment"`
		Timestamp int      `json:"timestamp"`
		//TODO: Image Path?
		Image []byte `json:"img"`
		couchdb.Document
	}
)
