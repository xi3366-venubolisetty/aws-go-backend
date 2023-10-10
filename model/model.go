package model

type BlogPosts struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	TagLine     string `json:"tagline"`
	Description string `json:"description"`
}
