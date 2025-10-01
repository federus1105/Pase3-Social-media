package models

import "time"

type Follow struct {
	ID          int       `json:"-"`
	FollowerID  int       `json:"follower_id"`
	FollowingID int       `json:"following_id"`
	CreatedAt   time.Time `json:"-"`
}

type FollowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Follow `json:"data"`
}
