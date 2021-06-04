package types

import (
	"errors"
	"sync"
)

type Blog struct {
	posts      map[int]Post
	lastPostID int
	sync.Mutex
}

type Post struct {
	ID    int
	Title string
	Body  string
}

func NewBlog() *Blog {
	return &Blog{
		posts: make(map[int]Post),
	}
}

func (b *Blog) AddPost(payload, reply *Post) error {
	b.Lock()
	defer b.Unlock()

	if payload.Title == "" || payload.Body == "" {
		return errors.New("Title and Body must not be empty")
	}

	b.lastPostID++

	*reply = Post{ID: b.lastPostID, Title: payload.Title, Body: payload.Body}
	b.posts[reply.ID] = *reply

	return nil
}

func (b *Blog) GetPostByID(payload int, reply *Post) error {
	b.Lock()
	defer b.Unlock()

	*reply = b.posts[payload]

	return nil
}
