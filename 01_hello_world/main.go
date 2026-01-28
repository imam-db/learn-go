/*
================================================================================
PELAJARAN 1: HELLO WORLD - Program Go Pertama Anda
================================================================================

APA ITU PACKAGE?
----------------
Package adalah cara Go mengorganisir kode. Setiap file Go harus memiliki package.
- package main    : Package utama untuk program yang bisa dijalankan (executable)
- package lainnya : Untuk library/fungsi yang akan dipakai program lain

APA ITU IMPORT?
---------------
Import digunakan untuk memasukkan package bawaan Go atau package buatan orang lain.
Package yang diimpor berisi fungsi-fungsi siap pakai.

Package fmt (format) adalah package paling sering digunakan untuk:
- fmt.Println() : Mencetak teks ke layar + pindah baris
- fmt.Print()   : Mencetak teks ke layar (tidak pindah baris)
- fmt.Printf()  : Mencetak dengan format tertentu
*/

// package main menandakan file ini adalah program utama yang bisa dijalankan
package main

// import "fmt" memasukkan package fmt untuk keperluan input/output
import "fmt"

/*
FUNGSI main()
-------------
- Fungsi main() adalah fungsi khusus yang akan dijalankan pertama kali
- Setiap program Go yang executable HARUS memiliki fungsi main()
- Program akan berjalan dari baris pertama dalam main() sampai akhir
*/
func main() {
	// fmt.Println() digunakan untuk mencetak teks ke layar
	// String (teks) di Go harus diapit dengan tanda petik ganda (" ")
	fmt.Println("Hello, World!")

	// Bisa mencetak multiple value dalam satu Println
	fmt.Println("Selamat datang di", "Go!")

	// \n membuat baris baru (newline)
	fmt.Println("Baris pertama\nBaris kedua")

	// fmt.Print() tidak menambahkan baris baru otomatis
	fmt.Print("Teks 1 ")
	fmt.Print("Teks 2 ")
	fmt.Print("Teks 3\n")

	// fmt.Printf() untuk format output (mirip C)
	nama := "Budi"
	umur := 20
	fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
}
