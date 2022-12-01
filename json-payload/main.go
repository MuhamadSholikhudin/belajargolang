package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/save", handleSave)
	http.HandleFunc("/", handle)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	resp, err := json.Marshal("OKE")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write([]byte(resp))
	return
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Data                        interface{} `json:"data"`
			Date_resignation_submission string      `json:"date_resignation_submission"`
			Selectdatesubmission        string      `json:"selectdatesubmission"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var isitr string = ""
		for i := 1; i < 5; i++ {
			isitr = fmt.Sprintf(`
				%s<tr>
					<td>%d.</td>
					<td>2203051857</td>
					<td>Muhammad Sholikhudin</td>
					<td>2022-03-04</td>
					<td class="text-center">
						<span>
							<input class="form-check-input checkboxsubmission" id="1" type="checkbox"
								checked="checked">
						</span>
					</td>
					</tr>
					`, isitr, i)
		}

		textisi := fmt.Sprintf(`<div class="card">
				<div class="card-header">
				<div class="custom-control custom-checkbox">
					<input class="custom-control-input" type="checkbox" id="checklistallsubmission" checked="checked" value="checkall" onclick="CheckboxSubmission();">
					<label for="checklistallsubmission" class="custom-control-label"> Checklist All</label>
				</div>
				</div>
				<div class="card-body p-0">
					<table class="table table-sm">
						<thead>
						<tr>
							<th>NO</th>
							<th>NIK</th>
							<th>Nama</th>
							<th>Tanggal</th>
							<th style="width: 10px;">Check</th>
						</tr>
						</thead>
						<tbody> %s %s`, isitr, `</tbody>
				</table>
			</div>
		</div>`)

		resp, err := json.Marshal(textisi)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Write([]byte(resp))
		return
	}

	// http.Error(w, "Only accept POST request", http.StatusBadRequest)
	message := http.StatusBadRequest

	resp, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Write([]byte(resp))
	return
}
