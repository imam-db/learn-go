/*
================================================================================
PELAJARAN 5: PERULANGAN (LOOPS)
================================================================================

APA ITU PERULANGAN?
-------------------
Perulangan memungkinkan kita menjalankan kode berulang kali
tanpa menulis kode yang sama banyak kali.

DI GO HANYA ADA "FOR"
---------------------
Go HANYA memiliki satu kata kunci untuk perulangan: "for"
Tapi "for" di Go sangat fleksibel dan bisa digunakan untuk berbagai kebutuhan:
1. For standar (mirip C/Java)
2. For sebagai while
3. For tanpa kondisi (infinite loop)
4. For dengan range (untuk array/slice/map)

================================================================================
1. FOR STANDAR
================================================================================

Struktur:
┌─────────────────────────────────────────┐
│  for init; kondisi; post {              │
│      // kode yang diulang               │
│  }                                      │
└─────────────────────────────────────────┘

Penjelasan:
- init     : Dieksekusi sekali di awal (biasanya inisialisasi variabel)
- kondisi  : Dicek sebelum setiap iterasi, loop berjalan jika true
- post     : Dieksekusi setelah setiap iterasi (biasanya increment)

================================================================================
2. FOR SEBAGAI WHILE
================================================================================

Go tidak punya "while", tapi bisa pakai for:
┌─────────────────────────────────────────┐
│  for kondisi {                          │
│      // kode yang diulang               │
│  }                                      │
└─────────────────────────────────────────┘

================================================================================
3. FOR TANPA KONDISI (INFINITE LOOP)
================================================================================

┌─────────────────────────────────────────┐
│  for {                                  │
│      // kode yang diulang terus         │
│      // gunakan break untuk keluar      │
│  }                                      │
└─────────────────────────────────────────┘

================================================================================
4. FOR DENGAN RANGE
================================================================================

Digunakan untuk mengiterasi array, slice, map, atau string:
┌─────────────────────────────────────────┐
│  for index, value := range koleksi {    │
│      // index = posisi                  │
│      // value = nilai                   │
│  }                                      │
└─────────────────────────────────────────┘

Gunakan _ untuk mengabaikan index atau value:
- for _, value := range koleksi  // ignore index
- for index, _ := range koleksi  // ignore value
- for range koleksi              // ignore both (Go 1.4+)

================================================================================
KONTROL LOOP
================================================================================

- break    : Menghentikan loop sepenuhnya
- continue : Melewatkan iterasi saat ini, lanjut ke iterasi berikutnya
*/

package main

import "fmt"

func main() {
	// =============================================================================
	// 1. FOR LOOP STANDAR
	// =============================================================================

	fmt.Println("=== FOR LOOP STANDAR ===")

	// Struktur: for init; kondisi; post
	// i := 1     -> inisialisasi, hanya dijalankan sekali
	// i <= 5     -> kondisi, dicek setiap iterasi
	// i++        -> increment, dijalankan setelah setiap iterasi

	for i := 1; i <= 5; i++ {
		fmt.Printf("Iterasi ke-%d\n", i)
	}

	/*
		Alur eksekusi:
		1. i = 1 (inisialisasi)
		2. Cek i <= 5? (1 <= 5 = true) -> lanjut
		3. Jalankan fmt.Printf("Iterasi ke-1")
		4. i++ (i menjadi 2)
		5. Cek i <= 5? (2 <= 5 = true) -> lanjut
		6. Jalankan fmt.Printf("Iterasi ke-2")
		7. i++ (i menjadi 3)
		...dan seterusnya sampai i = 6
		Saat i = 6, kondisi 6 <= 5 = false, loop berhenti
	*/

	// =============================================================================
	// VARIASI INCREMENT/DECREMENT
	// =============================================================================

	fmt.Println("\n=== COUNTDOWN ===")

	// Decrement (mengurangi)
	for i := 5; i >= 1; i-- {
		fmt.Printf("%d... ", i)
	}
	fmt.Println("Mulai!")

	// =============================================================================
	// 2. FOR SEBAGAI WHILE
	// =============================================================================
	// Hanya kondisi, tanpa init dan tanpa post

	fmt.Println("\n=== FOR SEBAGAI WHILE ===")

	x := 1
	for x <= 5 {
		fmt.Printf("x = %d\n", x)
		x++ // Increment manual di dalam loop
	}

	// =============================================================================
	// 3. FOR TANPA KONDISI (INFINITE LOOP)
	// =============================================================================
	// Berjalan terus sampai ada break

	fmt.Println("\n=== INFINITE LOOP DENGAN BREAK ===")

	y := 1
	for {
		// Loop ini akan berjalan selamanya tanpa break
		fmt.Printf("y = %d\n", y)

		if y >= 3 {
			fmt.Println("Mencapai batas, keluar dari loop!")
			break // Keluar dari loop
		}

		y++
	}

	// =============================================================================
	// KONTROL LOOP: BREAK dan CONTINUE
	// =============================================================================

	fmt.Println("\n=== CONTOH CONTINUE ===")
	// continue = lewati iterasi ini, lanjut ke iterasi berikutnya

	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Printf("(%d dilewati) ", i)
			continue // Lewati i = 3, lanjut ke i = 4
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// =============================================================================
	// 4. FOR DENGAN RANGE
	// =============================================================================
	// Cara paling umum untuk mengiterasi array/slice/map di Go

	fmt.Println("\n=== FOR DENGAN RANGE (ARRAY/SLICE) ===")

	namaBuah := []string{"Apel", "Mangga", "Jeruk", "Pisang"}

	// range mengembalikan dua nilai: index dan value
	fmt.Println("Dengan index dan value:")
	for index, value := range namaBuah {
		fmt.Printf("Index %d: %s\n", index, value)
	}

	// Mengabaikan index dengan _ (underscore)
	// Gunakan ini jika hanya butuh value, tidak butuh index
	fmt.Println("\nHanya value (index diabaikan):")
	for _, buah := range namaBuah {
		fmt.Printf("Buah: %s\n", buah)
	}

	// Mengabaikan value, hanya ambil index
	// Jarang digunakan, tapi bisa dilakukan
	fmt.Println("\nHanya index (value diabaikan):")
	for index, _ := range namaBuah {
		fmt.Printf("Index: %d\n", index)
	}

	// =============================================================================
	// RANGE DENGAN MAP
	// =============================================================================

	fmt.Println("\n=== RANGE DENGAN MAP ===")

	nilai := map[string]int{
		"Matematika": 90,
		"Fisika":     85,
		"Kimia":      88,
	}

	// Iterasi map dengan range
	for pelajaran, nilaiPelajaran := range nilai {
		fmt.Printf("%s: %d\n", pelajaran, nilaiPelajaran)
	}

	// =============================================================================
	// NESTED LOOP (LOOP BERSARANG)
	// =============================================================================

	fmt.Println("\n=== NESTED LOOP ===")

	// Loop di dalam loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("Baris %d: ", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// =============================================================================
	// CONTOH PRAKTIS: MENJUMLAHKAN ARRAY
	// =============================================================================

	fmt.Println("\n=== CONTOH PRAKTIS: PENJUMLAHAN ===")

	angka := []int{10, 20, 30, 40, 50}
	total := 0

	for _, nilai := range angka {
		total += nilai // total = total + nilai
	}

	fmt.Printf("Angka: %v\n", angka)
	fmt.Printf("Total: %d\n", total)

	// =============================================================================
	// LABEL DAN BREAK (MENGKELUARKAN NESTED LOOP)
	// =============================================================================

	fmt.Println("\n=== LABEL DAN BREAK ===")

	// Label untuk mengidentifikasi loop luar
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
			if i == 2 && j == 2 {
				fmt.Println("\nBreak ke outer loop!")
				break outerLoop // Keluar dari kedua loop
			}
		}
		fmt.Println()
	}
}
