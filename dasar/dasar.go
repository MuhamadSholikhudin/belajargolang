package main

// HEllO word
import "fmt"

func main() {
	fmt.Println("Hello world")

	// komentar kode
	// menampilkan pesan hello world

	/*
	   komentar kode
	   menampilkan pesan hello world
	*/

	// Variabel
	/*
	   var <nama-variabel> <tipe-data>
	   var <nama-variabel> <tipe-data> = <nilai>
	*/
	var firstName string = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"
	// var lastName string
	// lastName = "wick"

	// "halo %s %s!\n" => %s di ganti nilai variabel
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// Deklarasi Multi Variabel

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Printf("halo %s %s %s!\n", first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Printf("halo %s %s %s!\n", fourth, fifth, sixth)
	fmt.Printf("halo %s %s %s!\n", seventh, eight, ninth)

	// fmt.Printf("bilangan positif: %d %t %d %s\n", one, isFriday, twoPointTwo, say)

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	// name, _ := "john", "wick"

	// Deklarasi Variabel Menggunakan Keyword new
	// name := new(string)

	// fmt.Println(name)   // 0x20818a220
	// fmt.Println(*name)  // ""

	// TIPE DATA
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	// Template %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal.
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// Tipe Data Numerik Desimal
	var decimalNumber = 2.62
	// Template %f digunakan untuk memformat data numerik desimal menjadi string
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// Tipe Data bool (Boolean)
	var exist bool = true
	// Gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf().
	fmt.Printf("exist? %t \n", exist)

	// Tipe Data string
	// var message string = "Halo"

	var message = `Nama saya "John Wick".
    Salam kenal.
    Mari belajar "Golang".`

	fmt.Println(message)
	fmt.Printf("message: %s \n", message)

	// Konstanta
	const firstNameconst string = "john"
	fmt.Print("halo ", firstNameconst, "!\n")

	// Operator
	var value1 = (((2+6)%3)*4 - 2) / 3
	fmt.Println(value1)

	// Operator Perbandingan
	var value2 = (((2+6)%3)*4 - 2) / 3
	var isEqual2 = (value2 == 2)

	fmt.Printf("nilai %d (%t) \n", value2, isEqual2)

	// Operator Logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	// Seleksi Kondisi

	// Seleksi Kondisi Menggunakan Keyword if, else if, & else

	var point1 = 8

	if point1 == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point1 > 5 {
		fmt.Println("lulus")
	} else if point1 == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point1)

	}

	// Variabel Temporary Pada if - else

	var point2 = 8840.0

	if percent2 := point2 / 100; percent2 >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent2, "%")
	} else if percent2 >= 70 {
		fmt.Printf("%.1f%s good\n", percent2, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent2, "%")
	}

	// Seleksi Kondisi Menggunakan Keyword switch - case

	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Pemanfaatan case Untuk Banyak Kondisi

	var point4 = 6

	switch point4 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Kurung Kurawal Pada Keyword case & default

	var point5 = 6

	switch point5 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Switch Dengan Gaya if - else

	var point6 = 6

	switch {
	case point6 == 8:
		fmt.Println("perfect")
	case (point6 < 8) && (point6 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch

	var point7 = 6

	switch {
	case point7 == 8:
		fmt.Println("perfect")
	case (point7 < 8) && (point7 > 3):
		fmt.Println("awesome")
		fallthrough
	case point7 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Seleksi Kondisi Bersarang
	var point8 = 10

	if point8 > 7 {
		switch point8 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point8 == 5 {
			fmt.Println("not bad")
		} else if point8 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point8 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	// Perulangan

	// Perulangan Menggunakan Keyword for
	for i1 := 0; i1 < 5; i1++ {
		fmt.Println("Angka", i1)
	}

	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var i2 = 0

	for i2 < 5 {
		fmt.Println("Angka", i2)
		i2++
	}

	// Penggunaan Keyword for Tanpa Argumen
	var i3 = 0

	for {
		fmt.Println("Angka", i3)

		i3++
		if i3 == 5 {
			break
		}
	}

	// Penggunaan Keyword break & continue
	for i4 := 1; i4 <= 10; i4++ {
		if i4%2 == 1 {
			continue
		}

		if i4 > 8 {
			break
		}

		fmt.Println("Angka", i4)
	}

	// Perulangan Bersarang
	for i5 := 0; i5 < 5; i5++ {
		for j5 := i5; j5 < 5; j5++ {
			fmt.Print(j5, " ")
		}

		fmt.Println()
	}

	// Pemanfaatan Label Dalam Perulangan

outerLoop:
	for i6 := 0; i6 < 5; i6++ {
		for j6 := 0; j6 < 5; j6++ {
			if i6 == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i6, "][", j6, "]", "\n")
		}
	}

	// Array
	var names1 [4]string
	names1[0] = "trafalgar"
	names1[1] = "d"
	names1[2] = "water"
	names1[3] = "law"

	fmt.Println(names1[0], names1[1], names1[2], names1[3])

	// Pengisian Elemen Array yang Melebihi Alokasi Awal
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"
	// names[4] = "ez" // baris kode ini menghasilkan error

	// Inisialisasi Nilai Awal Array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Inisialisasi Nilai Array Dengan Gaya Vertikal
	var fruits2 [4]string

	// cara horizontal
	// fruits  = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits2 = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Isi semua element \t", fruits2)

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen

	var numbersarray = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbersarray)
	fmt.Println("jumlah elemen \t:", len(numbersarray))

	// Array Multidimensi
	var numbersarray1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbersarray2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbersarray1)
	fmt.Println("numbers2", numbersarray2)

	// Perulangan Elemen Array Menggunakan Keyword for

	var fruitsfor = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruitsfor); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruitsfor[i])
	}

	// Perulangan Elemen Array Menggunakan Keyword for-range

	var fruitsforrange = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruitsforrange {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// Penggunaan Variabel Underscore _ Dalam for - range

	var fruitsonderscore = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruitsonderscore {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// Alokasi Elemen Array Menggunakan Keyword make
	var fruitsmake = make([]string, 2)
	fruitsmake[0] = "apple"
	fruitsmake[1] = "manggo"

	fmt.Println(fruitsmake) // [apple manggo]

	// Inisialisasi Slice
	var fruitsSlice = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruitsSlice[0]) // "apple"

	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array

	fmt.Println(fruitsA[0])
	fmt.Println("fruitsB :", fruitsB)
	fmt.Println("fruitsC :", fruitsC)

	// Hubungan Slice Dengan Array & Operasi Slice

	var fruitss1 = []string{"apple", "grape", "banana", "melon"}
	var newFruitss1 = fruitss1[0:2]

	fmt.Println(newFruitss1) // ["apple", "grape"]

	// Slice Merupakan Tipe Data Reference

	var fruitss2 = []string{"apple", "grape", "banana", "melon"}

	var afruitss2 = fruitss2[0:3]
	var bfruitss2 = fruitss2[1:4]

	var aafruitss2 = afruitss2[1:2]
	var bafruitss2 = bfruitss2[0:1]

	fmt.Println(fruitss2)   // [apple grape banana melon]
	fmt.Println(afruitss2)  // [apple grape banana]
	fmt.Println(bfruitss2)  // [grape banana melon]
	fmt.Println(aafruitss2) // [grape]
	fmt.Println(bafruitss2) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	bafruitss2[0] = "pinnaple"

	fmt.Println(fruitss2)   // [apple pinnaple banana melon]
	fmt.Println(afruitss2)  // [apple pinnaple banana]
	fmt.Println(bfruitss2)  // [pinnaple banana melon]
	fmt.Println(aafruitss2) // [pinnaple]
	fmt.Println(bafruitss2) // [pinnaple]

	// Fungsi len() pada slice

	var fruitslen = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruitslen)) // 4

	//Fungsi cap() pada slice

	var fruitscap = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruitscap)) // len: 4
	fmt.Println(cap(fruitscap)) // cap: 4

	var afruitscap = fruitscap[0:3]
	fmt.Println(len(afruitscap)) // len: 3
	fmt.Println(cap(afruitscap)) // cap: 4

	var bfruitscap = fruitscap[1:4]
	fmt.Println(len(bfruitscap)) // len: 3
	fmt.Println(cap(bfruitscap)) // cap: 3

	// Fungsi append() pada slice

	var fruitsappend = []string{"apple", "grape", "banana"}
	var cfruitsappend = append(fruitsappend, "papaya")

	fmt.Println(fruitsappend)  // ["apple", "grape", "banana"]
	fmt.Println(cfruitsappend) // ["apple", "grape", "banana", "papaya"]

	var fruitsappend2 = []string{"apple", "grape", "banana"}
	var bfruitsappend2 = fruitsappend2[0:2]

	fmt.Println(cap(bfruitsappend2)) // 3
	fmt.Println(len(bfruitsappend2)) // 2

	fmt.Println(fruitsappend2)  // ["apple", "grape", "banana"]
	fmt.Println(bfruitsappend2) // ["apple", "grape"]

	var cfruitsappend2 = append(bfruitsappend2, "papaya")

	fmt.Println(fruitsappend2)  // ["apple", "grape", "papaya"]
	fmt.Println(bfruitsappend2) // ["apple", "grape"]
	fmt.Println(cfruitsappend2) // ["apple", "grape", "papaya"]

	//Fungsi copy() pada slice

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pinnaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2) // watermelon pinnaple potato
	fmt.Println(src2) // watermelon pinnaple
	fmt.Println(n2)   // 2

	// Pengaksesan Elemen Slice Dengan 3 Indeks

	var fruitsSlice3 = []string{"apple", "grape", "banana"}
	var afruitsSlice3 = fruitsSlice3[0:2]
	var bfruitsSlice3 = fruitsSlice3[0:2:2]

	fmt.Println(fruitsSlice3)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruitsSlice3)) // len: 3
	fmt.Println(cap(fruitsSlice3)) // cap: 3

	fmt.Println(afruitsSlice3)      // ["apple", "grape"]
	fmt.Println(len(afruitsSlice3)) // len: 2
	fmt.Println(cap(afruitsSlice3)) // cap: 3

	fmt.Println(bfruitsSlice3)      // ["apple", "grape"]
	fmt.Println(len(bfruitsSlice3)) // len: 2
	fmt.Println(cap(bfruitsSlice3)) // cap: 2

	// . Map

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Inisialisasi Nilai Map
	// var datamap map[string]int
	// datamap["one"] = 1
	// akan muncul error!

	// fmt.Println("one", datamap["one"])

	// data = map[string]int{}
	// data["one"] = 1
	// tidak ada error

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	// Pendeklarasian map lainnya
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	fmt.Println("chicken1", chicken1["januari"])
	fmt.Println("chicken2", chicken2["januari"])

	// Iterasi Item Map Menggunakan for - range

	var chickenfor = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chickenfor {
		fmt.Println(key, "  \t:", val)
	}

	// Deteksi Keberadaan Item Dengan Key Tertentu
	var chickenKeberadaan = map[string]int{"januari": 50, "februari": 40}
	var valueKeberadaan, isExist = chickenKeberadaan["mei"]

	if isExist {
		fmt.Println(valueKeberadaan)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi Slice & Map

	var chickensSliceMap = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickensSliceMap {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// Kombinasi Slice & Map versi baru
	var chickensSliceMap2 = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickensSliceMap2 {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// tiap elemen bisa saja memiliki key yang berbeda-beda, sebagai contoh seperti kode berikut.
	var keyberbeda = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	for _, chicken := range keyberbeda {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
