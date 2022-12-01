package models

type Resign struct {
	Id                     string `json:"id"`
	Number_of_employees    string `json:"number_of_employees"`
	Name                   string `json:"name"`
	Position               string `json:"position"`
	Department             string `json:"department"`
	Hire_date              string `json:"hire_date"`
	Classification         string `json:"classification"`
	Date_out               string `json:"date_out"`
	Date_resignsubmissions string `json:"date_resignsubmissions"`
	Periode_of_service     int    `json:"periode_of_service"`
	Type                   string `json:"type"`
	Age                    int    `json:"age"`
	Status_resign          string `json:"status_resign"`
	Printed                string `json:"printed"`
	Created_at             string `json:"created_at"`
	Updated_at             string `json:"updated_at"`
}
