package model

import "github.com/leesper/couchdb-golang"

type Post struct {
	ID        string                 `json:"_id"`
	Rev       string                 `json:"_rev"`
	Type      string                 `json:"type"`
	Author    string                 `json:"author"`
	Timestamp string                 `json:"timestamp"`
	Likes     int                    `json:"likes"`
	Caption   string                 `json:"caption"`
	Comments  map[string]interface{} `json:"comments"` //
	couchdb.Document
}

//TODO
func AddPicture() {

}

func DeletePicture() {

}
func GetPostByPostId() {

}

func GetComments() {

}

func AppendComment() {

}
func AppendLike() {

}
func DisappendLike() {

}
