package models

type Category struct {
	CategoryId int    `json:"id"`
	Name       string `json:"name"`
	UserId     int    `json:"user_id"`
}
