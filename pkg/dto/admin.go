package dto

type Admin struct {
	Id         string `json:"id"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Authorized string `json:"authorized"`
}
