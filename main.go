package main

import "fmt"

func main () {
	//hello world
	fmt.Println("hello, world!")

	//var const string
	var nama string = "example"
	nama = "example2"
	fmt.Println(nama)

	const nomor = 12
	//nomor = 13 // cannot assign to nomor (declared const)
	fmt.Println(nomor)

	//boolean
	var pelajar bool = true
	fmt.Println("apakah kamu pelajar?", pelajar)

	//conversion
	var word = "your"
	var w = word[0]
	var wWord = string(w)
	fmt.Println(word)
	fmt.Println(wWord)

	//type data number
	var nilai8 int8 = 127
	var nilai16 int16 = 32767
	var nilai32 int32 = 2147483647
	var nilai64 int64 = 9223372036854775807
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//int8	    -128 – 127
	//init16	  -32768 – 32767 // 32.767
	//int32	    -2147483648 – 2	147483647 //214.7483647
	//int64	    -9223372036854775808 – 9223372036854775807

	//type declaration
	type str string
	type numb int

	var produk str = "susu"
	var harga numb = 100000
	fmt.Println(produk)
	fmt.Println(harga)

	//perbandingan

	var dimas = "boy"
	var venny = "girl"
	fmt.Println(dimas == venny)

	//operasi matematika
	var a = 10
	a++
	a += 9
	fmt.Println(a)

	var jumlah = 10 + 10
	fmt.Println(jumlah)

	//operasi boolean

	var ujian = 80
	var magang = 85

	var kkmUjian = 80
	var kkmMagang = 80

	var hasiUjian = ujian >= kkmUjian
	var hasilMagang = magang >= kkmMagang

	var lulus = hasiUjian && hasilMagang
	fmt.Println(lulus)

	//array
	var jurusan [3]string
	jurusan[0] = "multimedia"
	jurusan[1] = "tkj"
	jurusan[2] = "pkm"
	fmt.Println(jurusan[0])
	fmt.Println(jurusan[1])
	fmt.Println(jurusan[2])

	//array langsung

	var id = [3]int {
		12,
		22,
		13,
	}
	fmt.Println(id)

}
