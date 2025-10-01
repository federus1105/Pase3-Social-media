package models

import "time"

type Comment struct {
	Id          int       `json:"id,omitempty"`
	PostinganId int       `json:"-"`
	Teks        string    `json:"teks"`
	UserID      int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

