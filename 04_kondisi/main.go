/*
================================================================================
PELAJARAN 4: KONDISI (CONDITIONAL STATEMENTS)
================================================================================

APA ITU KONDISI?
----------------
Kondisi memungkinkan program untuk membuat keputusan.
Program akan mengeksekusi kode berdasarkan kondisi tertentu.

Di Go ada 2 cara membuat kondisi:
1. if - else if - else
2. switch - case

================================================================================
1. IF - ELSE IF - ELSE
================================================================================

Struktur:
┌─────────────────────────────────────────┐
│  if kondisi {                           │
│      // jalankan jika kondisi true      │
│  } else if kondisi2 {                   │
│      // jalankan jika kondisi2 true     │
│  } else {                               │
│      // jalankan jika semua false       │
│  }                                      │
└─────────────────────────────────────────┘

CATATAN PENTING:
- Go TIDAK perlu tanda kurung () untuk kondisi
- Tetapi WAJIB menggunakan kurung kurawal {}
- else if dan else bersifat opsional

================================================================================
2. SWITCH - CASE
================================================================================

Struktur:
┌─────────────────────────────────────────┐
│  switch variabel {                      │
│  case nilai1:                           │
│      // jalankan jika variabel == nilai1│
│  case nilai2:                           │
│      // jalankan jika variabel == nilai2│
│  default:                               │
│      // jalankan jika tidak cocok       │
│  }                                      │
└─────────────────────────────────────────┘

Keunggulan switch:
- Lebih rapi untuk banyak kondisi
- Otomatis break (tidak perlu tulis break)
*/

package main

import "fmt"

func main() {
	// =============================================================================
	// IF - ELSE IF - ELSE
	// =============================================================================

	fmt.Println("=== CONTOH IF-ELSE IF-ELSE ===")

	nilai := 75

	fmt.Printf("Nilai: %d\n", nilai)

	// Kondisi pertama: apakah nilai >= 80?
	if nilai >= 80 {
		// Kode di dalam {} hanya dijalankan jika kondisi true
		fmt.Println("Grade: A (Sangat Baik)")
		fmt.Println("Selamat! Anda lulus dengan nilai memuaskan!")
	} else if nilai >= 70 {
		// Jalankan ini jika kondisi pertama false, tapi nilai >= 70
		fmt.Println("Grade: B (Baik)")
		fmt.Println("Anda lulus!")
	} else if nilai >= 60 {
		// Jalankan ini jika kedua kondisi di atas false
		fmt.Println("Grade: C (Cukup)")
		fmt.Println("Anda lulus, perlu belajar lebih giat!")
	} else {
		// Jalankan ini jika SEMUA kondisi di atas false
		fmt.Println("Grade: D (Kurang)")
		fmt.Println("Maaf, Anda tidak lulus.")
	}

	// =============================================================================
	// IF DENGAN DEKLARASI VARIABEL
	// =============================================================================
	// Go memungkinkan deklarasi variabel sebelum kondisi
	// Variabel hanya bisa diakses di dalam block if tersebut

	fmt.Println("\n=== IF DENGAN DEKLARASI VARIABEL ===")

	// umur dideklarasikan dan langsung digunakan untuk kondisi
	// umur hanya bisa diakses di dalam block if-else ini
	if umur := 17; umur >= 18 {
		fmt.Printf("Umur %d tahun -> Status: Dewasa\n", umur)
	} else {
		fmt.Printf("Umur %d tahun -> Status: Belum dewasa\n", umur)
	}

	// fmt.Println(umur)  // ❌ ERROR: umur tidak terdefinisi di sini

	// =============================================================================
	// IF BERSARANG (NESTED IF)
	// =============================================================================
	// If di dalam if

	fmt.Println("\n=== IF BERSARANG (NESTED IF) ===")

	umur2 := 25
	punyaSIM := true

	if umur2 >= 17 {
		fmt.Println("Umur cukup untuk SIM")

		if punyaSIM {
			fmt.Println("Anda sudah punya SIM, boleh mengemudi!")
		} else {
			fmt.Println("Anda belum punya SIM, daftar dulu ya!")
		}
	} else {
		fmt.Println("Umur belum cukup untuk SIM")
	}

	// =============================================================================
	// SWITCH - CASE
	// =============================================================================

	fmt.Println("\n=== CONTOH SWITCH-CASE ===")

	hari := 3

	fmt.Printf("Hari ke-%d adalah: ", hari)

	// switch akan memeriksa nilai hari
	switch hari {
	case 1:
		// Jalankan jika hari == 1
		fmt.Println("Senin")
	case 2:
		// Jalankan jika hari == 2
		fmt.Println("Selasa")
	case 3:
		// Jalankan jika hari == 3
		fmt.Println("Rabu")
	case 4:
		// Jalankan jika hari == 4
		fmt.Println("Kamis")
	case 5:
		// Jalankan jika hari == 5
		fmt.Println("Jumat")
	case 6, 7:
		// Multiple case: jalankan jika hari == 6 ATAU hari == 7
		fmt.Println("Weekend")
	default:
		// Jalankan jika tidak cocok dengan case manapun
		fmt.Println("Hari tidak valid")
	}

	// =============================================================================
	// SWITCH TANPA KONDISI
	// =============================================================================
	// Bisa digunakan sebagai alternatif if-else yang panjang

	fmt.Println("\n=== SWITCH TANPA KONDISI ===")

	angka := 15

	// Sama seperti serangkaian if-else
	switch {
	case angka < 0:
		fmt.Printf("%d adalah bilangan negatif\n", angka)
	case angka >= 0 && angka <= 10:
		fmt.Printf("%d ada di rentang 0-10\n", angka)
	case angka > 10 && angka <= 20:
		fmt.Printf("%d ada di rentang 11-20\n", angka)
	default:
		fmt.Printf("%d adalah bilangan besar\n", angka)
	}

	// =============================================================================
	// SWITCH DENGAN FALLTHROUGH
	// =============================================================================
	// Secara default, Go otomatis break setelah case cocok
	// fallthrough memaksa eksekusi ke case berikutnya

	fmt.Println("\n=== SWITCH DENGAN FALLTHROUGH ===")

	nilaiHuruf := "B"

	fmt.Printf("Nilai huruf %s:\n", nilaiHuruf)

	switch nilaiHuruf {
	case "A":
		fmt.Println("Sangat Baik")
		fallthrough // Lanjut ke case berikutnya
	case "B":
		fmt.Println("Baik")
		fallthrough // Lanjut ke case berikutnya
	case "C":
		fmt.Println("Cukup")
		fallthrough // Lanjut ke case berikutnya
	default:
		fmt.Println("Nilai tercatat")
	}
	// Output akan menampilkan: Baik, Cukup, Nilai tercatat

	// =============================================================================
	// CONTOH PRAKTIS: CEK TAHUN KABISAT
	// =============================================================================

	fmt.Println("\n=== CONTOH PRAKTIS: CEK TAHUN KABISAT ===")

	tahun := 2024

	// Tahun kabisat jika:
	// 1. Bisa dibagi 4 DAN
	// 2. (Tidak bisa dibagi 100 ATAU bisa dibagi 400)

	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				fmt.Printf("%d adalah tahun kabisat\n", tahun)
			} else {
				fmt.Printf("%d bukan tahun kabisat\n", tahun)
			}
		} else {
			fmt.Printf("%d adalah tahun kabisat\n", tahun)
		}
	} else {
		fmt.Printf("%d bukan tahun kabisat\n", tahun)
	}
}
