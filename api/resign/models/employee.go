package models

type Employee struct {
	//penamaan Camel Cas untuk Import Package supaya bisa di pakai dari luar
	Number_of_employees string `json:"number_of_employees"` // `` cara membuat penamaan ulang pada golang pada saat di GET
	National_id         string `json:"national_id"`
	Name                string `json:"name"`
	Job_id              string `json:"job_id"`
	Department_id       string `json:"department_id"`
	Date_of_birth       string `json:"date_of_birth"`
	Hire_date           string `json:"hire_date"`
	Date_out            string `json:"date_out"`
	Place_of_birth      string `json:"place_of_birth"`
	Address_jalan       string `json:"address_jalan"`
	Address_rt          string `json:"address_rt"`
	Address_rw          string `json:"address_rw"`
	Address_village     string `json:"address_village"`
	Address_district    string `json:"address_district"`
	Address_city        string `json:"address_city"`
	Address_province    string `json:"address_province"`
	Status_employee     string `json:"status_employee"`
}
