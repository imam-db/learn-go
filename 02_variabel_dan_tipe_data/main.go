/*
================================================================================
PELAJARAN 2: VARIABEL DAN TIPE DATA
================================================================================

VARIABEL (Variable)
-------------------
Variabel adalah tempat menyimpan data yang nilainya bisa berubah.
Di Go, variabel harus dideklarasikan sebelum digunakan.

Cara mendeklarasikan variabel di Go:
1. var nama tipe_data = nilai  (cara panjang)
2. var nama = nilai             (tipe data otomatis terdeteksi)
3. nama := nilai                (cara singkat, paling sering digunakan)

KONSEP := (SHORT VARIABLE DECLARATION)
--------------------------------------
- := adalah cara singkat mendeklarasikan variabel
- Hanya bisa digunakan DI DALAM fungsi (seperti main())
- Go akan otomatis mendeteksi tipe data dari nilainya
- Tidak perlu menulis kata kunci "var" dan tipe data

TIPE DATA DASAR GO
------------------
| Tipe Data | Deskripsi           | Contoh         |
|-----------|---------------------|----------------|
| string    | Teks/karakter       | "Hello"        |
| int       | Bilangan bulat      | 42, -10        |
| float64   | Bilangan desimal    | 3.14, 15000.50 |
| bool      | True/False          | true, false    |

KONSTANTA (const)
-----------------
Konstanta adalah variabel yang nilainya TIDAK BOLEH berubah.
Dideklarasikan dengan kata kunci "const".
Biasanya ditulis dengan HURUF BESAR untuk membedakan dengan variabel.
*/

package main

import "fmt"

func main() {
	// =============================================================================
	// CARA 1: var nama tipe_data (lalu assign nilai terpisah)
	// =============================================================================
	// Kata kunci "var" digunakan untuk mendeklarasikan variabel
	// "nama" adalah nama variabel
	// "string" adalah tipe data (untuk menyimpan teks)
	var nama string

	// Assignment (memberi nilai ke variabel)
	nama = "Budi"

	// fmt.Println menerima multiple argument yang dipisahkan koma
	// Setiap argument akan dicetak dengan spasi di antaranya
	fmt.Println("Cara 1 - Nama:", nama)

	// =============================================================================
	// CARA 2: var nama = nilai (tipe data otomatis dikenali)
	// =============================================================================
	// Go akan otomatis mendeteksi bahwa 20 adalah int (integer/bilangan bulat)
	var umur = 20
	fmt.Println("Cara 2 - Umur:", umur)

	// Kita bisa cek tipe data dengan fmt.Printf dan %T
	fmt.Printf("Tipe data umur: %T\n", umur)

	// =============================================================================
	// CARA 3: nama := nilai (cara singkat, PALING SERING DIGUNAKAN)
	// =============================================================================
	// := artinya "deklarasikan dan assign"
	// Tipe data akan otomatis terdeteksi dari nilai "Jakarta" (string)
	alamat := "Jakarta"
	fmt.Println("Cara 3 - Alamat:", alamat)

	// =============================================================================
	// TIPE DATA STRING
	// =============================================================================
	// String adalah tipe data untuk teks
	// Harus diapit dengan tanda petik ganda ("")
	hello := "Halo, Go!"
	selamat := "Selamat Belajar!"

	fmt.Println("\n=== TIPE DATA STRING ===")
	fmt.Println("hello:", hello)
	fmt.Println("selamat:", selamat)

	// String concatenation (menggabungkan string)
	gabungan := hello + " " + selamat
	fmt.Println("Gabungan:", gabungan)

	// =============================================================================
	// TIPE DATA INTEGER (int)
	// =============================================================================
	// int adalah bilangan bulat (tidak ada desimal)
	// Ada beberapa varian: int8, int16, int32, int64 (beda di ukuran memori)
	// Untuk umum, gunakan "int" saja
	angka := 42
	nilaiNegatif := -10

	fmt.Println("\n=== TIPE DATA INTEGER ===")
	fmt.Println("angka:", angka)
	fmt.Println("nilaiNegatif:", nilaiNegatif)

	// =============================================================================
	// TIPE DATA FLOAT (float64)
	// =============================================================================
	// float64 adalah bilangan desimal (ada koma)
	// 64 menunjukkan ukuran di memori (semakin besar semakin presisi)
	harga := 15000.50
	berat := 65.5

	fmt.Println("\n=== TIPE DATA FLOAT ===")
	fmt.Println("harga:", harga)
	fmt.Println("berat:", berat)

	// =============================================================================
	// TIPE DATA BOOLEAN (bool)
	// =============================================================================
	// bool hanya punya 2 nilai: true atau false
	// Sangat berguna untuk kondisi dan logika
	isActive := true
	isDone := false

	fmt.Println("\n=== TIPE DATA BOOLEAN ===")
	fmt.Println("isActive:", isActive)
	fmt.Println("isDone:", isDone)

	// =============================================================================
	// MULTIPLE VARIABLE DECLARATION
	// =============================================================================
	// Bisa mendeklarasikan banyak variabel sekaligus
	var a, b, c int = 1, 2, 3

	fmt.Println("\n=== MULTIPLE VARIABLE ===")
	fmt.Println("a:", a, "b:", b, "c:", c)

	// Dengan cara singkat
	x, y, z := 10, 20, 30
	fmt.Println("x:", x, "y:", y, "z:", z)

	// =============================================================================
	// KONSTANTA (const)
	// =============================================================================
	// Konstanta nilainya TIDAK BISA diubah setelah dideklarasikan
	// Biasa digunakan untuk nilai tetap seperti PI, gravitasi, dll
	const PI = 3.14159
	const NEGARA = "Indonesia"

	fmt.Println("\n=== KONSTANTA ===")
	fmt.Println("PI:", PI)
	fmt.Println("NEGARA:", NEGARA)

	// Jika kita coba mengubah konstanta, akan ERROR:
	// PI = 3.14  // ❌ ERROR: cannot assign to PI

	// =============================================================================
	// ZERO VALUE (Nilai Default)
	// =============================================================================
	// Di Go, variabel yang dideklarasikan tapi belum diisi nilai
	// akan otomatis punya nilai default (zero value)

	var defaultString string  // default: "" (string kosong)
	var defaultInt int        // default: 0
	var defaultFloat float64  // default: 0
	var defaultBool bool      // default: false

	fmt.Println("\n=== ZERO VALUE ===")
	fmt.Printf("string kosong: \"%s\"\n", defaultString)
	fmt.Println("int kosong:", defaultInt)
	fmt.Println("float kosong:", defaultFloat)
	fmt.Println("bool kosong:", defaultBool)
}
