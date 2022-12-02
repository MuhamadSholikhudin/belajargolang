package main

import (
	"belajargolang/api/resign/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Dashboard
	r.HandleFunc("/api.v1", controllers.Index).Methods("GET")
	r.HandleFunc("/api.v1/dashboard", controllers.Dashboard).Methods("GET")
	r.HandleFunc("/api.v1/resignjobs", controllers.GetJobs).Methods("GET")
	r.HandleFunc("/api.v1/resigndepartments", controllers.GetDepartments).Methods("GET")
	r.HandleFunc("/api.v1/resignbuildings", controllers.GetGedungs).Methods("GET")
	r.HandleFunc("/api.v1/resignalamat/{number_of_employees}", controllers.GetAlamat).Methods("GET")

	//Employee
	r.HandleFunc("/api.v1/employeeaction", controllers.EmployeeAction).Methods("POST")
	r.HandleFunc("/api.v1/employees/{number_of_employees}", controllers.Get).Methods("GET")
	r.HandleFunc("/api.v1/resign/{number_of_employees}/{national_id}", controllers.GetKaryawan).Methods("GET")
	r.HandleFunc("/api.v1/resigndate/{number_of_employees}/{national_id}", controllers.GetResign).Methods("GET")

	//Submissions
	r.HandleFunc("/api.v1/resignsubmissions", controllers.Submissions).Methods("GET")
	r.HandleFunc("/api.v1/resignsubmissions/{search}", controllers.GetResignSubmissionSearch).Methods("GET")
	r.HandleFunc("/api.v1/resignsubmissions/{number_of_employees}/{status_resign}", controllers.GetResignSubmissionStatus).Methods("GET")
	r.HandleFunc("/api.v1/resignsubmission_upload", controllers.UploadSubmission).Methods("POST")
	r.HandleFunc("/api.v1/resignsubmission_edit/{number_of_employees}", controllers.GetEditSubmission).Methods("GET")
	r.HandleFunc("/api.v1/resignsubmission_update", controllers.GetUpdateSubmission).Methods("POST")
	r.HandleFunc("/api.v1/resignsubmission_status", controllers.PostStatus).Methods("POST")
	r.HandleFunc("/api.v1/ExportSubmission", controllers.ExportSubmission).Methods("GET")
	r.HandleFunc("/api.v1/searchsubmission", controllers.SearchSubmission).Methods(http.MethodPost, http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/api.v1/processaccsubmission", controllers.ProcessAccSubmission).Methods(http.MethodPost, http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)

	//Resigns
	r.HandleFunc("/api.v1/resigns", controllers.Resigns).Methods("GET")
	r.HandleFunc("/api.v1/resigns/upload", controllers.UploadResigns).Methods("POST")
	r.HandleFunc("/api.v1/resigns_edit/{number_of_employees}", controllers.GetEditResign).Methods("GET")
	r.HandleFunc("/api.v1/resigns_update", controllers.PutResign).Methods("POST")
	r.HandleFunc("/api.v1/resigns/makecertificate", controllers.PostCertifcate).Methods("POST")
	r.HandleFunc("/api.v1/resigns/makeexperience", controllers.PostExperience).Methods("POST")
	r.HandleFunc("/api.v1/ExportResign", controllers.ExportResign).Methods("GET")

	//Parklaring
	r.HandleFunc("/api.v1/parklarings_certificate", controllers.GetParklaringCertificates).Methods("GET")
	r.HandleFunc("/api.v1/parklarings_certificateedit/{number_of_employees}", controllers.GetEditParklaringCertificate).Methods("GET")
	r.HandleFunc("/api.v1/parklarings_certificateupdate", controllers.GetUpdateParklaringCertificate).Methods("POST")

	r.HandleFunc("/api.v1/parklarings_experience", controllers.GetParklaringExperiences).Methods("GET")
	r.HandleFunc("/api.v1/parklarings_experienceedit/{number_of_employees}", controllers.GetEditParklaringExperience).Methods("GET")
	r.HandleFunc("/api.v1/parklarings_experienceupdate", controllers.GetUpdateParklaringExperience).Methods("POST")

	//Letter
	r.HandleFunc("/api.v1/ExportLetter/{dataletter}", controllers.ExportLetter).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))

	fmt.Println("LIsten on Port 10.10.42.6:8880")
	http.ListenAndServe(":8880", r)
}
