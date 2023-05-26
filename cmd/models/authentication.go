package models

type LoginUser struct {
	Usercd   string `json:"usercd"`
	Password string `json:"password"`
}
