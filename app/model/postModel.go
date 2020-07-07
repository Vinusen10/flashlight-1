package model

import (
	"encoding/json"
	"fmt"
	"github.com/leesper/couchdb-golang"
)

type Post struct {
	ID        string   `json:"_id"`
	Rev       string   `json:"_rev"`
	Type      string   `json:"type"`
	Author    string   `json:"author"`
	Timestamp string   `json:"timestamp"`
	Likes     []string `json:"likes"`
	Caption   string   `json:"caption"`
	Comments  []string `json:"comments"`
	ImagPath  string   `json:"imgpath"`

	couchdb.Document
}

//TODO
func AddPost() {

}

func GetAllPosts() ([]map[string]interface{}, error) {
	query := `{
	"selector": {
	"type": "post"
}
}
`
	allPosts, err := flashlightDB.QueryJSON(query)
	if err != nil {
		return nil, err
	} else {
		return allPosts, nil
	}
}

func GetPostsbyUser(user string) ([]map[string]interface{}, error) {

	query := `{
   "selector": {
      "type": "post",
      "author": "%s"
   }
}`
	userPost, err := flashlightDB.QueryJSON(fmt.Sprintf(query, user))
	return userPost, err
}

func AppendComment(id, comment, user string) {
}
func AppendLike() {

}
func AppendDislike(user string) {

}

/*-------Helper Function--------*/

// Convert from  Post Struct to map[string]interface{} as required by golang-couchdb method

func post2map(p Post) (post map[string]interface{}, err error) {
	uJSON, err := json.Marshal(p)
	json.Unmarshal(uJSON, &post)
	return post, err
}

// Convert from map[string]interface{} to Post Struct as required by golang-couchdb method
func map2post(post map[string]interface{}) (p Post, err error) {
	uJSON, err := json.Marshal(post)
	json.Unmarshal(uJSON, &p)
	return p, err
}
