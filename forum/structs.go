package forum

import "time"

type comment struct {
	CommentID      int
	Author         string
	PostID         int
	Content        string
	CommentTime    time.Time
	CommentTimeStr string
	Likes          int
	Dislikes       int
}

type post struct {
	PostID      int
	Author      string // author
	Image       string
	Title       string
	Content     string
	Category    string
	PostTime    time.Time
	PostTimeStr string
	Likes       int
	Dislikes    int
	Comments    []comment
}

type user struct {
	Username      string
	Access        int // 0 means no access, not logged in
	LoggedIn      bool
	Image         string
	Posts         []post
	Comments      []comment
	LikedPost     []post
	LikedComments []comment
}
