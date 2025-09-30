package models

type UserRegister struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserAuth struct {
	Id       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email" binding:"required,email"`
	Password string `db:"password" json:"password" binding:"required,min=8"`
}