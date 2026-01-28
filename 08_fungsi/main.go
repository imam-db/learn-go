/*
================================================================================
PELAJARAN 8: FUNGSI (FUNCTION)
================================================================================

APA ITU FUNGSI?
---------------
Fungsi adalah blok kode yang dirancang untuk melakukan tugas tertentu.
Fungsi memungkinkan kita:
1. Mengorganisir kode dengan lebih baik
2. Menghindari pengulangan kode (DRY - Don't Repeat Yourself)
3. Membuat kode lebih mudah dibaca dan dipelihara

================================================================================
STRUKTUR FUNGSI DI GO
================================================================================

┌─────────────────────────────────────────────────────────────┐
│  func namaFungsi(parameter1 tipe1, parameter2 tipe2) tipeReturn {  │
│      // kode fungsi                                         │
│      return nilaiReturn                                     │
│  }                                                          │
└─────────────────────────────────────────────────────────────┘

Penjelasan:
- func         : kata kunci untuk mendefinisikan fungsi
- namaFungsi   : nama fungsi (camelCase untuk private, PascalCase untuk public)
- parameter    : input untuk fungsi (opsional)
- tipeReturn   : tipe data yang dikembalikan (opsional)
- return       : mengembalikan nilai ke pemanggil

================================================================================
JENIS-JENIS FUNGSI
================================================================================

1. Fungsi tanpa parameter, tanpa return
2. Fungsi dengan parameter
3. Fungsi dengan return
4. Fungsi dengan multiple return
5. Fungsi dengan named return
6. Variadic function (parameter tak terbatas)
7. Function as value (fungsi sebagai nilai)
8. Anonymous function (fungsi tanpa nama)
9. Closure (fungsi yang mengakses variabel di luar scope-nya)
10. Recursive function (fungsi yang memanggil dirinya sendiri)

================================================================================
NAMING CONVENTION DI GO
================================================================================

- Exported (public)   : Huruf awal BESAR - bisa diakses dari package lain
  Contoh: fmt.Println(), strings.ToUpper()

- Unexported (private): Huruf awal kecil - hanya bisa diakses dalam package
  Contoh: fungsi main() selalu private karena hanya untuk package main
*/

package main

import "fmt"

// =============================================================================
// 1. FUNGSI TANPA PARAMETER, TANPA RETURN
// =============================================================================
// Fungsi paling sederhana

// sapa adalah fungsi yang tidak menerima input dan tidak mengembalikan output
// Fungsi ini hanya mencetak pesan ke layar
func sapa() {
	fmt.Println("Halo! Selamat datang di Go!")
	fmt.Println("Semoga harimu menyenangkan!")
}

// =============================================================================
// 2. FUNGSI DENGAN PARAMETER
// =============================================================================
// Parameter adalah input yang diterima fungsi

// sapaNama menerima satu parameter bertipe string
// nama adalah nama parameter yang bisa digunakan dalam fungsi
func sapaNama(nama string) {
	fmt.Printf("Halo, %s! Selamat datang!\n", nama)
}

// Fungsi dengan multiple parameter
// Parameter dipisahkan dengan koma
// Setiap parameter harus dideklarasikan tipenya (Go tidak bisa infer tipe parameter)
func hitungLuasPersegi(panjang int, lebar int) {
	luas := panjang * lebar
	fmt.Printf("Luas persegi %d x %d = %d\n", panjang, lebar, luas)
}

// Jika parameter bertipe sama, bisa ditulis sekali di akhir
// a dan b keduanya bertipe int
func tambah(a, b int) int {
	hasil := a + b
	return hasil
}

// =============================================================================
// 3. FUNGSI DENGAN RETURN (MENGEMBALIKAN NILAI)
// =============================================================================
// return digunakan untuk mengembalikan nilai ke pemanggil fungsi

// kalikan menerima 2 int dan mengembalikan hasil perkalian (int)
func kalikan(a, b int) int {
	return a * b
}

// =============================================================================
// 4. FUNGSI DENGAN MULTIPLE RETURN
// =============================================================================
// Go memungkinkan fungsi mengembalikan lebih dari satu nilai
// Ini sangat berguna untuk mengembalikan hasil + error/status

// hitung mengembalikan dua nilai: hasil tambah dan hasil kurang
func hitung(a, b int) (int, int) {
	penjumlahan := a + b
	pengurangan := a - b
	return penjumlahan, pengurangan
}

// Contoh: mengembalikan hasil dan status error
func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa membagi dengan nol")
	}
	hasil := a / b
	return hasil, nil
}

// =============================================================================
// 5. FUNGSI DENGAN NAMED RETURN
// =============================================================================
// Memberi nama pada return value membuat kode lebih jelas
// Dengan named return, kita bisa return tanpa menyebut variabelnya

// Named return: luas dan keliling sudah dideklarasikan di signature
func hitungLuasKeliling(panjang, lebar int) (luas int, keliling int) {
	// Variabel luas dan keliling sudah tersedia
	luas = panjang * lebar           // tidak perlu := karena sudah dideklarasikan
	keliling = 2 * (panjang + lebar) // tidak perlu :=

	// return tanpa variabel - Go otomatis mengembalikan luas dan keliling
	return
	// atau bisa juga: return luas, keliling
}

// =============================================================================
// 6. VARIADIC FUNCTION
// =============================================================================
// Fungsi yang bisa menerima jumlah parameter tak terbatas
// Ditandai dengan ... sebelum tipe parameter

// sum menerima banyak angka int dan menjumlahkannya
// angka adalah slice ([]int) yang berisi semua argument
func sum(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

// Variadic parameter harus di posisi terakhir
func greet(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// =============================================================================
// 7. FUNCTION AS VALUE
// =============================================================================
// Di Go, fungsi adalah first-class citizen
// Fungsi bisa disimpan dalam variabel, dipassing sebagai parameter, dll

// Deklarasi tipe fungsi
// type Operasi func(int, int) int

type Operasi func(int, int) int

// jalankanOperasi menerima fungsi sebagai parameter
func jalankanOperasi(a, b int, op Operasi) int {
	return op(a, b)
}

// =============================================================================
// 8. ANONYMOUS FUNCTION (FUNGSI TANPA NAMA)
// =============================================================================
// Fungsi yang dideklarasikan tanpa nama
// Biasa digunakan untuk callback atau IIFE (Immediately Invoked Function Expression)

// =============================================================================
// 9. CLOSURE
// =============================================================================
// Closure adalah fungsi yang mengakses variabel di luar scope-nya
// Variabel tersebut akan "tertangkap" dalam closure

func counter() func() int {
	n := 0 // variabel di luar scope fungsi anonim
	return func() int {
		n++ // mengakses variabel n
		return n
	}
}

// =============================================================================
// 10. RECURSIVE FUNCTION (REKURSIF)
// =============================================================================
// Fungsi yang memanggil dirinya sendiri
// Harus memiliki base case untuk menghentikan rekursi

func factorial(n int) int {
	// Base case
	if n <= 1 {
		return 1
	}
	// Recursive case
	return n * factorial(n-1)
}

// =============================================================================
// FUNGSI MAIN
// =============================================================================

func main() {
	fmt.Println("=== 1. FUNGSI TANPA PARAMETER ===")
	sapa() // Memanggil fungsi sapa

	fmt.Println("\n=== 2. FUNGSI DENGAN PARAMETER ===")
	sapaNama("Budi")
	sapaNama("Ani")
	hitungLuasPersegi(5, 3)

	fmt.Println("\n=== 3. FUNGSI DENGAN RETURN ===")
	hasilTambah := tambah(5, 3)
	fmt.Printf("5 + 3 = %d\n", hasilTambah)

	hasilKali := kalikan(4, 7)
	fmt.Printf("4 x 7 = %d\n", hasilKali)

	fmt.Println("\n=== 4. MULTIPLE RETURN ===")
	jumlah, kurang := hitung(10, 4)
	fmt.Printf("10 + 4 = %d\n", jumlah)
	fmt.Printf("10 - 4 = %d\n", kurang)

	// Ignore salah satu return value dengan _
	hanyaJumlah, _ := hitung(7, 2)
	fmt.Printf("Hanya jumlah: %d\n", hanyaJumlah)

	fmt.Println("\n=== 5. NAMED RETURN ===")
	l, k := hitungLuasKeliling(5, 3)
	fmt.Printf("Luas: %d, Keliling: %d\n", l, k)

	fmt.Println("\n=== 6. VARIADIC FUNCTION ===")
	fmt.Printf("Sum(1,2,3): %d\n", sum(1, 2, 3))
	fmt.Printf("Sum(10,20): %d\n", sum(10, 20))
	fmt.Printf("Sum(): %d\n", sum())

	// Pass slice ke variadic function dengan ...
	angka := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum(slice...): %d\n", sum(angka...))

	greet("Selamat pagi", "Budi", "Ani", "Caca")

	fmt.Println("\n=== 7. FUNCTION AS VALUE ===")
	// Menyimpan fungsi dalam variabel
	operasiTambah := tambah
	hasil := operasiTambah(10, 20)
	fmt.Printf("Hasil operasiTambah(10,20): %d\n", hasil)

	// Passing fungsi sebagai parameter
	hasilOperasi := jalankanOperasi(5, 3, tambah)
	fmt.Printf("jalankanOperasi(5,3,tambah): %d\n", hasilOperasi)

	fmt.Println("\n=== 8. ANONYMOUS FUNCTION ===")
	// Fungsi tanpa nama yang langsung disimpan dalam variabel
	kali := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Anonymous func kali(4,5): %d\n", kali(4, 5))

	// IIFE - Immediately Invoked Function Expression
	hasilIIFE := func(a, b int) int {
		return a*a + b*b
	}(3, 4) // langsung dipanggil dengan argument (3, 4)
	fmt.Printf("IIFE 3^2 + 4^2: %d\n", hasilIIFE)

	fmt.Println("\n=== 9. CLOSURE ===")
	hitung1 := counter()
	hitung2 := counter()

	fmt.Printf("Hitung1: %d\n", hitung1()) // 1
	fmt.Printf("Hitung1: %d\n", hitung1()) // 2
	fmt.Printf("Hitung1: %d\n", hitung1()) // 3

	fmt.Printf("Hitung2: %d\n", hitung2()) // 1 (independen dari hitung1)
	fmt.Printf("Hitung2: %d\n", hitung2()) // 2

	fmt.Println("\n=== 10. RECURSIVE FUNCTION ===")
	fmt.Printf("Faktorial 5: %d\n", factorial(5)) // 120
	fmt.Printf("Faktorial 0: %d\n", factorial(0)) // 1
}
