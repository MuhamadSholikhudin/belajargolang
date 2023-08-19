package helper

import (
	"belajargolang/api/resign/models"
	"fmt"
	"strconv"
	"time"

	"github.com/theTardigrade/age"
)

const (
	LINKFRONTEND   string = "http://10.10.40.190:8080"
	DDMMYYYYhhmmss        = "2006-01-02 15:04:05"
	DDMMYYYY              = "2006-01-02"
)

var Dept = map[string]string{
	"ACCOUNTING":                 "ACCOUNTING",
	"ASSEMBLY":                   "ASSEMBLY",
	"ASSEMBLY-A":                 "ASSEMBLY",
	"ASSEMBLY-B":                 "ASSEMBLY",
	"ASSEMBLY-C":                 "ASSEMBLY",
	"ASSEMBLY-D":                 "ASSEMBLY",
	"ASSEMBLY-D1":                "ASSEMBLY",
	"ASSEMBLY-D2":                "ASSEMBLY",
	"ASSEMBLY-E1":                "ASSEMBLY",
	"ASSEMBLY-E2":                "ASSEMBLY",
	"ASSEMBLY-H":                 "ASSEMBLY",
	"CHEMICAL ENGINEERING":       "CHEMICAL ENGINEERING",
	"CUTTING":                    "CUTTING",
	"CUTTING-A":                  "CUTTING",
	"CUTTING-B":                  "CUTTING",
	"CUTTING-C":                  "CUTTING",
	"CUTTING-D1":                 "CUTTING",
	"CUTTING-D2":                 "CUTTING",
	"CUTTING-E1":                 "CUTTING",
	"CUTTING-E2":                 "CUTTING",
	"CUTTING-H":                  "CUTTING",
	"DEVELOPMENT":                "DEVELOPMENT",
	"EMBOSS":                     "EMBOSS",
	"ENGINEERING":                "ENGINEERING",
	"EPTE":                       "EPTE",
	"EXIM":                       "EXIM",
	"FACTORY MGR A":              "FACTORY MANAGER",
	"FACTORY MGR C":              "FACTORY MANAGER",
	"FACTORY MGR D":              "FACTORY MANAGER",
	"FACTORY MGR E":              "FACTORY MANAGER",
	"FINISH GOOD":                "FINISH GOOD",
	"FINISH GOOD A":              "FINISH GOOD",
	"FINISH GOOD C":              "FINISH GOOD",
	"FINISH GOOD D":              "FINISH GOOD",
	"FINISH GOOD E":              "FINISH GOOD",
	"FINISH GOOD H":              "FINISH GOOD",
	"FINISH GOOD O":              "FINISH GOOD",
	"GA":                         "GA",
	"GA (DRIVER)":                "DRIVER",
	"GA (KEBERSIHAN)":            "GA",
	"GA (MESS)":                  "GA",
	"GA (SECURITY)":              "SECURITY",
	"GA (SIPIL)":                 "GA",
	"GA (WWTP)":                  "GA",
	"GUDANG MATERIAL":            "GUDANG MATERIAL",
	"HRD":                        "HRD",
	"IE":                         "IE",
	"IT":                         "IT",
	"LABORAT":                    "LABORAT",
	"LAMINATING":                 "LAMINATING",
	"MAGANG":                     "MAGANG",
	"MARKETING":                  "MARKETING",
	"ME":                         "ME",
	"MT":                         "MT",
	"PPIC":                       "PPIC",
	"PRODUCTION DIRECTOR":        "PRODUCTION DIRECTOR",
	"PURCHASING":                 "PURCHASING",
	"QIP":                        "QIP",
	"QIP-A":                      "QIP",
	"QIP-B":                      "QIP",
	"QIP-C":                      "QIP",
	"QIP-D":                      "QIP",
	"QIP-E":                      "QIP",
	"QIP-F":                      "QIP",
	"QIP-G":                      "QIP",
	"QIP-H":                      "QIP",
	"QIP-M":                      "QIP",
	"QIP-S":                      "QIP",
	"QSM":                        "QSM",
	"SABLON":                     "SABLON",
	"SABLON EMBOSS":              "SABLON EMBOSS",
	"SEA":                        "SEA",
	"SERIKAT NON-JOB":            "SERIKAT NON-JOB",
	"SEWING COMP":                "SEWING COMPUTER",
	"SEWING COMP A":              "SEWING COMPUTER",
	"SEWING COMP B":              "SEWING COMPUTER",
	"SEWING COMP C":              "SEWING COMPUTER",
	"SEWING COMP D":              "SEWING COMPUTER",
	"SEWING COMP D1":             "SEWING COMPUTER",
	"SEWING COMP D2":             "SEWING COMPUTER",
	"SEWING COMP E1":             "SEWING COMPUTER",
	"SEWING COMP E":              "SEWING COMPUTER",
	"SEWING COMP E2":             "SEWING COMPUTER",
	"SEWING COMP H":              "SEWING COMPUTER",
	"SEWING MEKANIK":             "MEKANIK SEWING",
	"SEWING MEKANIK A":           "MEKANIK SEWING",
	"SEWING MEKANIK B":           "MEKANIK SEWING",
	"SEWING MEKANIK C":           "MEKANIK SEWING",
	"SEWING MEKANIK D":           "MEKANIK SEWING",
	"SEWING MEKANIK D1":          "MEKANIK SEWING",
	"SEWING MEKANIK D2":          "MEKANIK SEWING",
	"SEWING MEKANIK E":           "MEKANIK SEWING",
	"SEWING MEKANIK E1":          "MEKANIK SEWING",
	"SEWING MEKANIK E2":          "MEKANIK SEWING",
	"SEWING MEKANIK H":           "MEKANIK SEWING",
	"SEWING":                     "SEWING",
	"SEWING-A":                   "SEWING",
	"SEWING-B":                   "SEWING",
	"SEWING-C":                   "SEWING",
	"SEWING-D":                   "SEWING",
	"SEWING-D1":                  "SEWING",
	"SEWING-D2":                  "SEWING",
	"SEWING-E":                   "SEWING",
	"SEWING-E1":                  "SEWING",
	"SEWING-E2":                  "SEWING",
	"SEWING-H":                   "SEWING",
	"SMART":                      "SMART",
	"STOCKFIT":                   "STOCKFIT",
	"TECHNICAL":                  "TECHNICAL",
	"TECHNICAL HOTPRESS":         "TECHNICAL HOTPRESS",
	"TECHNICAL LAB":              "TECHNICAL LAB",
	"TECHNICAL ROLLING COMPOUND": "TECHNICAL ROLLING COMPOUND",
	"TECHNICAL SUPERMARKET":      "TECHNICAL SUPERMARKET",
	"TRAINING":                   "TRAINING",
	"TRAINING CENTER":            "TRAINING CENTER",
	"NONE":                       "NONE",
	"QIP-OFFICE":                 "QIP",
	"PRESIDENT DIRECTOR":         "PRESIDENT DIRECTOR",
	"SENIOR PRODUCTION DIRECTOR": "SENIOR PRODUCTION DIRECTOR",
	"BUSINESS MATERIAL":          "BUSINESS MATERIAL",
	"QIP LAB CE":                 "QIP LAB CE",
	"PUBLIC RELATION":            "PUBLIC RELATION",
	"HR GA SEA":                  "HR GA SEA",
	"PLANT E":                    "PLANT",
	"MATL PUR":                   "MATL PUR",
	"CHANGE":                     "CHANGE",
	"PLANT C":                    "PLANT",
	"PLANT H":                    "PLANT",
	"PLANT D":                    "PLANT",
	"BUSINESS":                   "BUSINESS",
	"PLANT A AND B":              "PLANT A AND B",
	"CHEMICAL ADVISOR":           "CHEMICAL ADVISOR",
	"EXECUTIVE SR. DIR HWI 2":    "EXECUTIVE SR. DIR HWI 2",
	"CUTTING-D":                  "CUTTING",
	"ASSEMBLY-E":                 "ASSEMBLY",
	"CUTTING-E":                  "CUTTING",
	"AMT":                        "AMT",
}

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

func YearMysql(date string) string {
	yearint, _ := strconv.Atoi(string(date[0:4]))
	year := fmt.Sprintf("%d", yearint)
	return year
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

func NomorLetter(No int, SK string, Rom string, DateString string) string {
	var s string
	s = DateString
	yearDate, _ := strconv.Atoi(string(s[0:4]))
	output := fmt.Sprintf("%d%s%s/%d", No, SK, Rom, yearDate)
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
	var output string
	var Submission models.Resignation_submission

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
	var Status string = "resign"
	var classification string = "Resign dahulu sebelum mengajukan resign"

	err = dbhwi.QueryRow("SELECT COUNT(id) as id FROM resignation_submissions WHERE number_of_employees = ? AND (status_resignsubmisssion = 'wait' OR status_resignsubmisssion = 'acc')", Number_of_employees).
		Scan(&Count_id_submission)
	if err != nil {
		fmt.Println(err.Error())
	}

	date, _ := time.Parse("2006-01-02", Date_out)
	datesubstract := date.AddDate(0, 0, -7)

	if Count_id_submission > 0 {
		classification = "Mengajukan resign tetapi resign sebelum waktunya"
		Status = "wait"
	}

	err = dbhwi.QueryRow("SELECT COUNT(id) as id FROM resignation_submissions WHERE number_of_employees = ? AND date_resignation_submissions <= ?", Number_of_employees, datesubstract).
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

func DateSubmissionCompareRequest(date_submission, date_request string) string {
	datecreated := string(date_submission[0:10])
	submission_date, _ := time.Parse("2006-01-02", datecreated)
	d_submission := submission_date.AddDate(0, 0, 7)
	request_date, _ := time.Parse("2006-01-02", date_request)
	d1 := request_date.After(d_submission)
	var d2 string = "false"
	if d_submission.Equal(request_date) || d1 == true {
		d2 = "true"
	}
	return d2
}

func Deptout(dept string) string {
	var datadept string
	for key, val := range Dept {
		if key == dept {
			datadept = val
		}
	}
	return datadept
}
