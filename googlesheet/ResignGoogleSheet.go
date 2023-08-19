package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	resp, err := http.Get("https://script.googleusercontent.com/macros/echo?user_content_key=en9LPCpWxHIi6ypowj516ylITkOMcRm0eaorXRA-GMyKo-ng7B_4N3uJdiSRP8a2Q7gPMCZq1dbxs47rN-kyK-ATKocFrm8vm5_BxDlH2jW0nuo2oDemN9CCS2h10ox_1xSncGQajx_ryfhECjZEnF7XA4IZnfrhIsfLjiFuVgnfUn6GRRnJuLRSaK1nuhAc5gBTl6uWHFKtr26VJn7A0dwf6k5YrcAy1T9U3EAstbBKoSOdad6tfw&lib=MP2gm8nYgyptxZX7bL1YrGUbkDnjr7vlQ")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	// sb := string(body)
	// log.Printf(sb)
	// fmt.Println(reflect.TypeOf(sb))

	var data map[string][][]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Jumlah :", len(data["content"]))
	/*
		0	Timestamp
		1	NIK ID Card
		2	Nama
		3	Alasan Pengunduran Diri
		4	Tanggal Permohonan Keluar
		5	Jabatan
		6	Keterangan
		7	Saran untuk perusahaan
		8	Alamat
		9	Department
		10	1. Saya terampil menyelesaikan target pekerjaan
		11	2. Atasan menggunakan kata-kata/sikap yang wajar dalam bekerja
		12	3. Rekan kerja saya membantu kesulitan saya dalam menyelesaikan pekerjaan
		13	4. Jarak perusahaan dengan tempat tinggal tidak menjadi masalah bagi saya
		14	5. Jam kerja (termasuk shift malam) tidak masalah bagi saya
		15	6. Saya berkeinginan kembali ke perusahaan (PT HWI) suatu saat nanti
		16	7. Keluarga (termasuk menikah, mengurus keluarga) bukanlah alasan bagi saya untuk meninggalkan perusahaan ini
	*/

	// Sudah Bisa Convert
	for i := 0; i < len(data["content"]); i++ {
		var get_num_str, rep_nik, number_of_employees string
		get_num_str = fmt.Sprintf("%f", data["content"][i][1])
		rep_nik = strings.Replace(get_num_str, ".", "", -1)
		number_of_employees = rep_nik[0:10]
		// fmt.Println(reflect.TypeOf(data["content"][i][0]), number_of_employees, reflect.TypeOf(data["content"][i][4]))
		fmt.Println(reflect.TypeOf(data["content"][i][4]), number_of_employees, reflect.TypeOf(cellfloattoint(data["content"][i][16])), data["content"][i][16])
	}

}

func cellfloattoint(data interface{}) int {
	var get_num_str, rep_str, get_str string
	get_num_str = fmt.Sprintf("%f", data)
	rep_str = strings.Replace(get_num_str, ".", "", -1)
	get_str = rep_str[0:1]
	var num, _ = strconv.ParseInt(get_str, 10, 64)
	return int(num)
}
