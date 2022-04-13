package forum

import "time"

type comment struct {
	CommentID      int
	Username       string // author
	Content        string
	CommentTime    time.Time
	CommentTimeStr string
	Likes          int
	Dislikes       int
}

type post struct {
	PostID      int
	Username    string // author
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
	Posts         []post
	Comments      []comment
	LikedPost     []post
	LikedComments []comment
}