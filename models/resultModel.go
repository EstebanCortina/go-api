package models

type User struct {
	Id        string
	Name      string
	User_name string
	Followers int64
}

type Post struct {
	Id      string
	Content string
	Likes   int64
	Author  string
}
