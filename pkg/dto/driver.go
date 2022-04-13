package dto

type Driver struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Driving  string `json:"driving"`
}
