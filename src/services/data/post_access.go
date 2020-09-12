package data

import (
	"errors"

	"github.com/camolezi/MicroservicesGolang/src/model"
	"github.com/camolezi/MicroservicesGolang/src/utils"
)

//PostAccess - Maybe create separated interfaces latter
type PostAccess interface {
}

//Mock database for now
var dbMock = map[int64]model.Post{
	0: {ID: 0, Title: "First post ever, you are a lucky guy for seeing this"},
	1: {ID: 1, Title: "Second post ever, you are a lucky guy for seeing this"},
	2: {ID: 2, Title: "third post ever, you are a lucky guy for seeing this"},
	3: {ID: 3, Title: "forth post ever, you are a lucky guy for seeing this"},
	4: {ID: 4, Title: "fifth post ever, you are a lucky guy for seeing this"},
}

//GetPost retrieves a post from the database
func GetPost(id int64) (model.Post, error) {
	post, contain := dbMock[id]
	if !contain {
		return model.Post{}, &utils.ResourceError{ErrorMessage: "Post not Found"}
	}
	return post, nil
}

//NewPost creates a new post
func NewPost(id int64, post model.Post) error {
	_, contain := dbMock[id]
	if contain {
		return errors.New("Post on this ID already exist")
	}

	dbMock[id] = post
	return nil
}

//GetLatestPosts for now will return all posts in the map
func GetLatestPosts(size uint) ([]model.Post, error) {
	postSlice := make([]model.Post, 0)
	for _, post := range dbMock {
		postSlice = append(postSlice, post)
	}

	return postSlice, nil
}
