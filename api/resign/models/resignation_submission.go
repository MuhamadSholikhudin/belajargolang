package models

type Resignation_submission struct {
	Number_of_employees          string `json:"number_of_employees"`
	Name                         string `json:"name"`
	Position                     string `json:"position"`
	Department                   string `json:"department"`
	Building                     string `json:"building"`
	Hire_date                    string `json:"hire_date"`
	Date_out                     string `json:"date_out"`
	Date_resignation_submissions string `json:"date_resignation_submissions"`
	Type                         string `json:"type"`
	Reason                       string `json:"reason"`
	Detail_reason                string `json:"detail_reason"`
	Suggestion                   string `json:"suggestion"`
	Periode_of_service           int    `json:"periode_of_service"`
	Status_resignsubmisssion     string `json:"status_resignsubmisssion"`
	Age                          int    `json:"age"`
	Using_media                  string `json:"using_media"`
	Classification               string `json:"classification"`
	Created_at                   string `json:"created_at"`
	Updated_at                   string `json:"updated_at"`
}
