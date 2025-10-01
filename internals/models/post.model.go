package models

import (
	"mime/multipart"
	"time"
)

type PostBody struct {
	Id       int                   `form:"id,omitempty"`
	User     int                   `form:"user"`
	Image    *multipart.FileHeader `form:"image,omitempty" json:"-"`
	ImageStr string                `form:"imagestr,omitempty"`
	Caption  string                `form:"caption,omitempty"`
}

type Post struct {
	Id        int       `json:"id"`
	User      int       `json:"user_id"`
	Image     string    `json:"content"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUserPostsResponse struct {
	Success bool   `json:"success"`
	Data    []Post `json:"data"`
}
