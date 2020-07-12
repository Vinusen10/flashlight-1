package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leesper/couchdb-golang"
	"log"
)

type Post struct {
	ID        string    `json:"_id"`
	Rev       string    `json:"_rev"`
	Type      string    `json:"type"`
	Author    string    `json:"author"`
	Timestamp Tstamp    `json:"timestamp"`
	Likes     []string  `json:"likes"`
	Caption   string    `json:"caption"`
	Comments  []Comment `json:"comments"`
	ImagPath  string    `json:"imgpath"`

	couchdb.Document
}

type Comment struct {
	User    string `json:"user"`
	Comment string `json:"comment"`
}
type Tstamp struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

//TODO
func (p Post) AddPost() (err error) {

	p.Type = "post"

	newPost, err := post2map(p)
	if err != nil {
		log.Println(err)
	}
	delete(newPost, "_id")
	delete(newPost, "_rev")

	_, _, errr := flashlightDB.Save(newPost, nil)

	if errr != nil {
		log.Println("[Add] error: %s", errr)
	}
	return err
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

func GetPostByID(id string) (Post, error) {
	t, err := flashlightDB.Get(id, nil)
	if err != nil {
		return Post{}, err
	}

	userPost := Post{
		ID:   t["_id"].(string),
		Rev:  t["_rev"].(string),
		Type: t["type"].(string),
	}

	return userPost, err
}
func (p Post) Delete() error {
	err := flashlightDB.Delete(p.ID)
	return err
}

func AppendComment(id, comment, user string) {
	// Get the Post with id
	post, err := flashlightDB.Get(id, nil)
	if err != nil {
		log.Println(err)
	}
	// Append the comment to json field comments[]
	newPost, err := map2post(post)
	if err != nil {
		log.Println(err)
	}
	appendPost := Comment{
		User:    user,
		Comment: comment,
	}
	newPost.Comments = append(newPost.Comments, appendPost)
	tmp, _ := post2map(newPost)
	flashlightDB.Set(id, tmp)

}
func Like(id, username string) error {
	post, err := flashlightDB.Get(id, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	newPost, err := map2post(post)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, like := range newPost.Likes {
		if like == username {
			err = dislike(id, username, newPost)
			if err != nil {
				return err
			} else {
				return nil
			}

		}
	}
	newPost.Likes = append(newPost.Likes, username)
	tmp, _ := post2map(newPost)
	err = flashlightDB.Set(id, tmp)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func dislike(id, username string, newPost Post) error {
	for i, like := range newPost.Likes {
		if like == username {
			newPost.Likes = append(newPost.Likes[:i], newPost.Likes[i+1:]...)
			fmt.Println(newPost.Likes)
			tmp, _ := post2map(newPost)
			err := flashlightDB.Set(id, tmp)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	return errors.New("Undefined Error - dislike")
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
