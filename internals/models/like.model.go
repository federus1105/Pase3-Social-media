package models

import "time"

type Like struct {
	ID        int       `json:"-"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"-"`
}

type LikeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Like   `json:"data"`
}
