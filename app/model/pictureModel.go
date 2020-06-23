package model

import "github.com/leesper/couchdb-golang"

type Picture struct {
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

//TODO
func AddPicture() {

}

func DeletePicture() {

}
func GetPictureByPictureId() {

}

func GetComments() {

}

func AppendComment() {

}
func AppendLike() {

}
func DisappendLike() {

}
