package models

type Letter struct {
	Id                  int    `json:"id"`
	Resign_id           int    `json:"resign_id"`
	Number_of_employees string `json:"number_of_employees"`
	Date                string `json:"date"`
	No                  string `json:"no"`
	Rom                 string `json:"rom"`
	Status              string `json:"status"`
	Action              string `json:"action"`
	Created_at          string `json:"created_at"`
	Updated_at          string `json:"updated_at"`
}
