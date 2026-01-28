/*
================================================================================
PELAJARAN 3: OPERASI (OPERATOR)
================================================================================

APA ITU OPERATOR?
-----------------
Operator adalah simbol yang digunakan untuk melakukan operasi pada data/variabel.
Go memiliki beberapa jenis operator:

1. OPERATOR ARITMATIKA
   Untuk operasi matematika dasar

2. OPERATOR PERBANDINGAN
   Untuk membandingkan dua nilai, hasilnya bool (true/false)

3. OPERATOR LOGIKA
   Untuk menggabungkan kondisi boolean

================================================================================
*/

package main

import "fmt"

func main() {
	// Deklarasi variabel untuk operasi
	a := 10
	b := 3

	// =============================================================================
	// 1. OPERATOR ARITMATIKA
	// =============================================================================
	// Digunakan untuk perhitungan matematika

	fmt.Println("=== OPERATOR ARITMATIKA ===")
	fmt.Printf("a = %d, b = %d\n\n", a, b)

	// Penjumlahan (+)
	// Menjumlahkan dua bilangan
	hasilTambah := a + b
	fmt.Printf("Penjumlahan: %d + %d = %d\n", a, b, hasilTambah)

	// Pengurangan (-)
	// Mengurangi bilangan kedua dari bilangan pertama
	hasilKurang := a - b
	fmt.Printf("Pengurangan: %d - %d = %d\n", a, b, hasilKurang)

	// Perkalian (*)
	// Mengalikan dua bilangan
	hasilKali := a * b
	fmt.Printf("Perkalian:   %d * %d = %d\n", a, b, hasilKali)

	// Pembagian (/)
	// Membagi bilangan pertama dengan bilangan kedua
	// CATATAN: Jika kedua operand adalah int, hasilnya int (dibulatkan ke bawah)
	hasilBagi := a / b
	fmt.Printf("Pembagian:   %d / %d = %d\n", a, b, hasilBagi)
	// 10 / 3 = 3 (bukan 3.33) karena a dan b bertipe int

	// Untuk hasil desimal, minimal salah satu harus float
	hasilBagiFloat := float64(a) / float64(b)
	fmt.Printf("Pembagian (float): %.2f\n", hasilBagiFloat)

	// Modulo/Modulus (%)
	// Menghasilkan sisa bagi
	// 10 % 3 = 1 (karena 10 = 3*3 + 1)
	hasilModulo := a % b
	fmt.Printf("Modulo:      %d %% %d = %d\n", a, b, hasilModulo)

	// =============================================================================
	// OPERATOR PENUGASAN (ASSIGNMENT) dengan OPERASI
	// =============================================================================
	// Cara singkat melakukan operasi dan assignment sekaligus

	fmt.Println("\n=== OPERATOR PENUGASAN ===")
	x := 5
	fmt.Printf("Nilai awal x = %d\n", x)

	x += 3 // Sama dengan: x = x + 3
	fmt.Printf("Setelah x += 3: x = %d\n", x)

	x -= 2 // Sama dengan: x = x - 2
	fmt.Printf("Setelah x -= 2: x = %d\n", x)

	x *= 2 // Sama dengan: x = x * 2
	fmt.Printf("Setelah x *= 2: x = %d\n", x)

	x /= 3 // Sama dengan: x = x / 3
	fmt.Printf("Setelah x /= 3: x = %d\n", x)

	// =============================================================================
	// 2. OPERATOR PERBANDINGAN (COMPARISON)
	// =============================================================================
	// Membandingkan dua nilai, hasilnya selalu boolean (true/false)

	fmt.Println("\n=== OPERATOR PERBANDINGAN ===")
	fmt.Printf("a = %d, b = %d\n\n", a, b)

	// Sama dengan (==)
	// true jika nilai sama, false jika berbeda
	fmt.Printf("a == b: %v (apakah %d sama dengan %d?)\n", a == b, a, b)

	// Tidak sama dengan (!=)
	// true jika nilai berbeda
	fmt.Printf("a != b: %v (apakah %d tidak sama dengan %d?)\n", a != b, a, b)

	// Lebih besar (>)
	fmt.Printf("a > b:  %v (apakah %d lebih besar dari %d?)\n", a > b, a, b)

	// Lebih kecil (<)
	fmt.Printf("a < b:  %v (apakah %d lebih kecil dari %d?)\n", a < b, a, b)

	// Lebih besar sama dengan (>=)
	fmt.Printf("a >= b: %v (apakah %d >= %d?)\n", a >= b, a, b)

	// Lebih kecil sama dengan (<=)
	fmt.Printf("a <= b: %v (apakah %d <= %d?)\n", a <= b, a, b)

	// =============================================================================
	// 3. OPERATOR LOGIKA (LOGICAL)
	// =============================================================================
	// Digunakan untuk menggabungkan kondisi boolean
	// Sangat penting untuk if statement dan loop

	fmt.Println("\n=== OPERATOR LOGIKA ===")

	// AND (&&)
	// Hasil true HANYA jika KEDUANYA true
	// Tabel kebenaran:
	// true && true  = true
	// true && false = false
	// false && true = false
	// false && false = false

	fmt.Println("\n--- Operator AND (&&) ---")
	fmt.Println("true && true  =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", false && false)

	// Contoh praktis AND
	umur := 25
	hasilUjian := 80
	bisaKerja := umur >= 18 && hasilUjian >= 75
	fmt.Printf("Umur %d dan nilai %d -> Bisa kerja? %v\n", umur, hasilUjian, bisaKerja)

	// OR (||)
	// Hasil true jika SALAH SATU atau KEDUANYA true
	// Tabel kebenaran:
	// true || true  = true
	// true || false = true
	// false || true = true
	// false || false = false

	fmt.Println("\n--- Operator OR (||) ---")
	fmt.Println("true || true  =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || true =", false || true)
	fmt.Println("false || false =", false || false)

	// Contoh praktis OR
	punyaKTP := true
	punyaSIM := false
	bisaMasuk := punyaKTP || punyaSIM
	fmt.Printf("Punya KTP: %v atau SIM: %v -> Bisa masuk? %v\n", punyaKTP, punyaSIM, bisaMasuk)

	// NOT (!)
	// Membalikkan nilai boolean
	// !true  = false
	// !false = true

	fmt.Println("\n--- Operator NOT (!) ---")
	fmt.Println("!true  =", !true)
	fmt.Println("!false =", !false)

	// Contoh kombinasi operator logika
	fmt.Println("\n=== KOMBINASI OPERATOR LOGIKA ===")
	sudahMakan := true
	sudahMinum := false
	uangCukup := true

	bisaBelanja := sudahMakan && (sudahMinum || uangCukup)
	fmt.Printf("Makan: %v, Minum: %v, Uang: %v -> Bisa belanja? %v\n",
		sudahMakan, sudahMinum, uangCukup, bisaBelanja)
}
