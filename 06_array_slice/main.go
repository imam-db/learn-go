/*
================================================================================
PELAJARAN 6: ARRAY DAN SLICE
================================================================================

APA PERBEDAAN ARRAY DAN SLICE?
==============================

┌─────────────┬──────────────────────┬──────────────────────────┐
│             │ ARRAY                │ SLICE                    │
├─────────────┼──────────────────────┼──────────────────────────┤
│ UKURAN      │ Tetap (fixed)        │ Dinamis (bisa berubah)   │
│ DEKLARASI   │ [5]int               │ []int                    │
│ KAPASITAS   │ Sama dengan ukuran   │ Bisa lebih besar dari    │
│             │                      │ panjang yang terpakai    │
│ PENGGUNAAN  │ Jarang dipakai       │ Sangat sering dipakai    │
└─────────────┴──────────────────────┴──────────────────────────┘

================================================================================
ARRAY
================================================================================

Array adalah koleksi data dengan ukuran TETAP.
Setelah dibuat, ukuran array tidak bisa diubah.

Deklarasi:
┌─────────────────────────────────────────┐
│  var nama [ukuran]tipe_data             │
│  nama := [ukuran]tipe_data{nilai}       │
│  nama := [...]tipe_data{nilai}          │ ← ukuran otomatis
└─────────────────────────────────────────┘

================================================================================
SLICE
================================================================================

Slice adalah koleksi data dengan ukuran DINAMIS.
Slice lebih fleksibel dan lebih sering digunakan di Go.

Deklarasi:
┌─────────────────────────────────────────┐
│  var nama []tipe_data                   │ ← slice kosong           │
│  nama := []tipe_data{nilai}             │ ← slice dengan nilai     │
│  nama := make([]tipe_data, len, cap)    │ ← slice dengan make      │
└─────────────────────────────────────────┘

Konsep penting slice:
- len(length): jumlah elemen yang ada
- cap(capacity): kapasitas maksimum sebelum perlu alokasi ulang

================================================================================
SLICING (MENGAMBIL BAGIAN)
================================================================================

Dari array atau slice, kita bisa ambil bagian tertentu:
┌─────────────────────────────────────────┐
│  slice[start:end]   → dari start sampai │
│                     sebelum end         │
│  slice[start:]      → dari start sampai │
│                     akhir               │
│  slice[:end]        → dari awal sampai  │
│                     sebelum end         │
│  slice[:]           → semua elemen      │
└─────────────────────────────────────────┘
*/

package main

import "fmt"

func main() {
	// =============================================================================
	// ARRAY
	// =============================================================================

	fmt.Println("=== ARRAY ===")

	// Deklarasi array dengan ukuran 5, tipe int
	// Semua elemen diinisialisasi dengan zero value (0 untuk int)
	var angka [5]int

	fmt.Printf("Array kosong: %v\n", angka)
	fmt.Printf("Panjang array: %d\n", len(angka))

	// Mengisi array dengan index
	// Index array dimulai dari 0
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50

	fmt.Printf("Array setelah diisi: %v\n", angka)
	fmt.Printf("Elemen ke-0: %d\n", angka[0])
	fmt.Printf("Elemen ke-2: %d\n", angka[2])

	// Deklarasi array dengan nilai awal
	nama := [3]string{"Budi", "Ani", "Caca"}
	fmt.Printf("\nArray nama: %v\n", nama)

	// Array dengan ukuran otomatis [...]
	// Go akan menghitung ukuran dari nilai yang diberikan
	kota := [...]string{"Jakarta", "Bandung", "Surabaya", "Medan"}
	fmt.Printf("Array kota (auto size): %v\n", kota)
	fmt.Printf("Panjang array kota: %d\n", len(kota))

	// Array multidimensi (array di dalam array)
	fmt.Println("\n=== ARRAY MULTIDIMENSI ===")
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("Matrix: %v\n", matrix)
	fmt.Printf("Element [0][1]: %d\n", matrix[0][1]) // 2

	// =============================================================================
	// SLICE
	// =============================================================================

	fmt.Println("\n=== SLICE ===")

	// Deklarasi slice kosong
	// Bedanya dengan array: tidak ada ukuran di dalam []
	var buah []string
	fmt.Printf("Slice kosong: %v\n", buah)
	fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(buah), cap(buah))

	// append() - menambah elemen ke slice
	// append mengembalikan slice baru, harus ditangkap
	buah = append(buah, "Apel")
	buah = append(buah, "Mangga")
	buah = append(buah, "Jeruk")

	fmt.Printf("\nSetelah append: %v\n", buah)
	fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(buah), cap(buah))

	// Append multiple elemen sekaligus
	buah = append(buah, "Pisang", "Durian", "Anggur")
	fmt.Printf("\nSetelah append multiple: %v\n", buah)
	fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(buah), cap(buah))
	// Perhatikan: kapasitas berubah ketika penuh

	// =============================================================================
	// MAKE - MEMBUAT SLICE DENGAN KAPASITAS TERTENTU
	// =============================================================================

	fmt.Println("\n=== MAKE ===")

	// make(tipe, length, capacity)
	// Membuat slice dengan panjang dan kapasitas yang bisa dikontrol
	nilai := make([]int, 3, 5)

	fmt.Printf("Slice dari make: %v\n", nilai)
	fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(nilai), cap(nilai))

	// Elemen sudah terisi zero value (0)
	nilai[0] = 80
	nilai[1] = 90
	nilai[2] = 85

	fmt.Printf("Setelah diisi: %v\n", nilai)

	// Bisa append karena masih ada capacity
	nilai = append(nilai, 95)
	fmt.Printf("Setelah append: %v\n", nilai)
	fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(nilai), cap(nilai))

	// =============================================================================
	// SLICING - MENGAMBIL BAGIAN DARI ARRAY/SLICE
	// =============================================================================

	fmt.Println("\n=== SLICING ===")

	hewan := []string{"Kucing", "Anjing", "Kelinci", "Burung", "Ikan", "Ular"}
	fmt.Printf("Original: %v\n", hewan)
	fmt.Printf("Panjang: %d\n\n", len(hewan))

	// slice[start:end] → elemen start sampai sebelum end
	// Index: 0    1      2       3       4     5
	//       Kucing Anjing Kelinci Burung Ikan Ular

	fmt.Printf("hewan[1:4]  → %v\n", hewan[1:4])
	// Index 1, 2, 3 → Anjing, Kelinci, Burung

	fmt.Printf("hewan[:3]   → %v\n", hewan[:3])
	// Dari 0 sampai sebelum 3 → Kucing, Anjing, Kelinci

	fmt.Printf("hewan[2:]   → %v\n", hewan[2:])
	// Dari 2 sampai akhir → Kelinci, Burung, Ikan, Ular

	fmt.Printf("hewan[:]    → %v\n", hewan[:])
	// Semua elemen

	// =============================================================================
	// SLICE DAN ARRAY Saling Terhubung!
	// =============================================================================

	fmt.Println("\n=== HUBUNGAN SLICE DAN ARRAY ===")

	// Slice adalah "view" ke array di belakangnya
	// Mengubah slice bisa mengubah array asal!

	srcArray := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array asal: %v\n", srcArray)

	// Buat slice dari array
	slice1 := srcArray[1:4]
	fmt.Printf("Slice [1:4]: %v\n", slice1)

	// Ubah slice
	slice1[0] = 999
	fmt.Printf("\nSetelah ubah slice[0]:\n")
	fmt.Printf("Slice: %v\n", slice1)
	fmt.Printf("Array asal: %v\n", srcArray)
	// Array asal juga berubah!

	// =============================================================================
	// COPY SLICE
	// =============================================================================

	fmt.Println("\n=== COPY SLICE ===")

	src := []int{1, 2, 3}
	dst := make([]int, len(src))

	// copy(destination, source)
	// Mengembalikan jumlah elemen yang tercopy
	n := copy(dst, src)

	fmt.Printf("Source:      %v\n", src)
	fmt.Printf("Destination: %v\n", dst)
	fmt.Printf("Tercopy: %d elemen\n", n)

	// Ubah destination, source tidak berubah (independent)
	dst[0] = 100
	fmt.Printf("\nSetelah ubah destination:\n")
	fmt.Printf("Source:      %v\n", src)
	fmt.Printf("Destination: %v\n", dst)

	// =============================================================================
	// DELETE ELEMEN SLICE
	// =============================================================================

	fmt.Println("\n=== DELETE ELEMEN SLICE ===")

	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	fmt.Printf("Original: %v\n", names)

	// Delete index 2 (Charlie)
	// Cara: gabungkan slice sebelum index dan setelah index
	indexToDelete := 2
	names = append(names[:indexToDelete], names[indexToDelete+1:]...)
	fmt.Printf("After delete index %d: %v\n", indexToDelete, names)

	// Penjelasan:
	// names[:2]     → [Alice, Bob]
	// names[3:]     → [David, Eve]
	// append(..., ...) → menggabungkan keduanya
	// ...           → spread operator untuk melepas elemen slice
}
