package helper

import (
	"belajargolang/api/resign/models"
	"fmt"
	"strconv"
	"time"

	"github.com/theTardigrade/age"
)

const (
	LINKFRONTEND string = "http://10.10.40.190:8080"

	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	DDMMYYYY       = "2006-01-02"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DMYhms() string {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
	}
	datetimenow := t.In(location).Format(DDMMYYYYhhmmss)
	return datetimenow
}

func DMY() string {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
	}
	datetimenow := t.In(location).Format(DDMMYYYY)
	return datetimenow
}

func StringMonth() string {
	var now = time.Now()
	var stringmonth string
	stringmonth = strconv.Itoa(int(now.Month()))
	var length = len([]rune(stringmonth))
	var value string = stringmonth
	if length == 1 {
		value = fmt.Sprintf("0%s", stringmonth)
	}
	return value
}

func Rom(stringmonth string) string {
	var Rom string
	if stringmonth == "01" {
		Rom = "I"
	}
	if stringmonth == "02" {
		Rom = "II"
	}
	if stringmonth == "03" {
		Rom = "III"
	}
	if stringmonth == "04" {
		Rom = "IV"
	}
	if stringmonth == "05" {
		Rom = "V"
	}
	if stringmonth == "06" {
		Rom = "VI"
	}
	if stringmonth == "07" {
		Rom = "VII"
	}
	if stringmonth == "08" {
		Rom = "VIII"
	}
	if stringmonth == "09" {
		Rom = "IX"
	}
	if stringmonth == "10" {
		Rom = "X"
	}
	if stringmonth == "11" {
		Rom = "XI"
	}
	if stringmonth == "12" {
		Rom = "XII"
	}
	return Rom
}

func NomorLetter(No string, SK string, Rom string, DateString string) string {
	var s string
	s = DateString
	yearDate, _ := strconv.Atoi(string(s[0:4]))
	output := fmt.Sprintf("%s%s%s/%d", No, SK, Rom, yearDate)
	return output
}

func Age(DateString string) int {
	var s string
	s = DateString
	yearDate, _ := strconv.Atoi(string(s[0:4]))
	monthDate, _ := strconv.Atoi(string(s[5:7]))
	dayDate, _ := strconv.Atoi(string(s[8:10]))

	month := time.Month(monthDate)

	date := time.Date(yearDate, month, dayDate, 0, 0, 0, 0, time.UTC)
	dateAge := age.Calculate(date)

	return dateAge
}

func Periode_of_serve(DateString string, DateString2 string) int {
	var s string
	s = DateString
	yearDate, _ := strconv.Atoi(string(s[0:4]))
	monthDate, _ := strconv.Atoi(string(s[5:7]))
	dayDate, _ := strconv.Atoi(string(s[8:10]))

	var s2 string
	s2 = DateString2
	yearDate2, _ := strconv.Atoi(string(s2[0:4]))
	monthDate2, _ := strconv.Atoi(string(s2[5:7]))
	dayDate2, _ := strconv.Atoi(string(s2[8:10]))

	t1 := Date(yearDate, monthDate, dayDate)
	t2 := Date(yearDate2, monthDate2, dayDate2)
	days := t2.Sub(t1).Hours() / 24
	return int(days)
}

func CekDateSubmission(Number_of_employees string) string {
	var dbresign, err = models.ConnResign()
	defer dbresign.Close()

	var Count_id int
	var Submission models.Resignation_submission
	var output string

	err = dbresign.QueryRow("SELECT COUNT(id) as id, COALESCE(date_resignation_submissions, '0000-00-00') FROM resignation_submissions WHERE number_of_employees = ? AND status_resignsubmisssion = 'wait'  ", Number_of_employees).
		Scan(&Count_id, &Submission.Date_resignation_submissions)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch Count_id {
	case 1:
		var s string
		s = Submission.Date_resignation_submissions
		yearDate, _ := strconv.Atoi(string(s[0:4]))
		monthDate, _ := strconv.Atoi(string(s[5:7]))
		dayDate, _ := strconv.Atoi(string(s[8:10]))

		month := time.Month(monthDate)

		theTime := time.Date(yearDate, month, dayDate, 0, 0, 0, 0, time.Local)

		after := theTime.AddDate(0, 0, 14)

		var stringafter string
		stringafter = after.Format("2006-01-02")

		var currentTime string
		currentTime = time.Now().Format("2006-01-02")

		status_type := Periode_of_serve(currentTime, stringafter)

		if status_type <= 0 {
			output = "Sudah mengajukan resign dan resign sesuai procedure"
		} else {
			output = "Mengajukan resign tetapi resign sebelum waktunya"
		}
	default:
		output = "Resign dahulu sebelum mengajukan resign"
	}

	return output
}

func TypeResign(Number_of_employees string, Date_out string) map[string]interface{} {
	var db, _ = models.ConnHrd()
	defer db.Close()

	var dbhwi, err = models.ConnResign()
	defer dbhwi.Close()

	var Count_id, Count_id_submission int
	var Type string = "false"
	var Status string = "wait"
	var classification string = "Resign dahulu sebelum mengajukan resign"

	err = dbhwi.QueryRow("SELECT COUNT(id) as id FROM resignation_submissions WHERE number_of_employees = ? AND status_resignsubmisssion = 'wait' ", Number_of_employees).
		Scan(&Count_id_submission)
	if err != nil {
		fmt.Println(err.Error())
	}
	if Count_id_submission > 0 {
		classification = "Mengajukan resign tetapi resign sebelum waktunya"
	}

	err = dbhwi.QueryRow("SELECT COUNT(id) as id FROM resignation_submissions WHERE number_of_employees = ? AND date_resignation_submissions <= ?", Number_of_employees, Date_out).
		Scan(&Count_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	if Count_id > 0 {
		Type = "true"
		classification = "Sudah mengajukan resign dan resign sesuai procedure"
	}
	output := map[string]interface{}{
		"type":           Type,
		"classification": classification,
		"status":         Status,
	}
	return output
}
