/*
================================================================================
PELAJARAN 7: MAP
================================================================================

APA ITU MAP?
------------
Map adalah tipe data koleksi yang menyimpan data dalam format key-value.
- Key: identifier unik (tidak boleh duplikat)
- Value: nilai yang disimpan

Analogi: Map seperti kamus/katalog
- Key = kata yang dicari
- Value = arti/penjelasan kata tersebut

================================================================================
KARAKTERISTIK MAP DI GO
================================================================================
1. Unordered - Tidak ada urutan tertentu
2. Key harus unik - Tidak boleh ada duplikat key
3. Key bisa dicek keberadaannya - Mengembalikan 2 nilai (value, exists)
4. Reference type - Map di-pass by reference
5. Nilai default adalah nil - Map harus diinisialisasi sebelum digunakan

================================================================================
DEKLARASI MAP
================================================================================

Cara 1: Menggunakan make
┌─────────────────────────────────────────┐
│  nama := make(map[tipe_key]tipe_value)  │
└─────────────────────────────────────────┘

Cara 2: Map literal
┌─────────────────────────────────────────┐
│  nama := map[tipe_key]tipe_value{       │
│      "key1": value1,                    │
│      "key2": value2,                    │
│  }                                      │
└─────────────────────────────────────────┘

================================================================================
OPERASI PADA MAP
================================================================================

┌────────────────┬──────────────────────────────────────┐
│ Operasi        │ Cara                                 │
├────────────────┼──────────────────────────────────────┤
│ Tambah/Ubah    │ map["key"] = value                   │
│ Ambil          │ value := map["key"]                  │
│ Cek keberadaan │ value, exists := map["key"]          │
│ Hapus          │ delete(map, "key")                   │
│ Panjang        │ len(map)                             │
└────────────────┴──────────────────────────────────────┘
*/

package main

import "fmt"

func main() {
	// =============================================================================
	// CARA 1: Membuat Map dengan make()
	// =============================================================================

	fmt.Println("=== MAP DENGAN MAKE ===")

	// Deklarasi map kosong
	// map[string]string artinya: key bertipe string, value bertipe string
	var mahasiswa map[string]string

	// Map harus diinisialisasi dengan make sebelum digunakan
	// Jika tidak, akan panic (runtime error)
	mahasiswa = make(map[string]string)

	// Menambahkan data ke map
	// Format: map["key"] = "value"
	mahasiswa["nama"] = "Budi Santoso"
	mahasiswa["jurusan"] = "Teknik Informatika"
	mahasiswa["nim"] = "2023001"
	mahasiswa["email"] = "budi@email.com"

	fmt.Printf("Data mahasiswa: %v\n", mahasiswa)
	fmt.Printf("Nama: %s\n", mahasiswa["nama"])
	fmt.Printf("Jurusan: %s\n", mahasiswa["jurusan"])

	// =============================================================================
	// CARA 2: Map Literal
	// =============================================================================

	fmt.Println("\n=== MAP LITERAL ===")

	// Deklarasi langsung dengan nilai
	nilai := map[string]int{
		"matematika": 90,
		"fisika":     85,
		"kimia":      88,
		"biologi":    92,
	}
	// Catatan: koma setelah elemen terakhir WAJIB!

	fmt.Printf("Nilai: %v\n", nilai)

	// =============================================================================
	// MENGECEK KEY ADA ATAU TIDAK (OK IDIOM)
	// =============================================================================

	fmt.Println("\n=== CEK KEY ADA/TIDAK ===")

	// Mengambil value dari map mengembalikan 2 nilai:
	// 1. value - nilai dari key (jika tidak ada = zero value)
	// 2. ok/bool - true jika key ada, false jika tidak ada

	// Contoh: Key yang ADA
	if nilaiFisika, ada := nilai["fisika"]; ada {
		fmt.Printf("Key 'fisika' ADA dengan nilai: %d\n", nilaiFisika)
	} else {
		fmt.Println("Key 'fisika' tidak ditemukan")
	}

	// Contoh: Key yang TIDAK ADA
	if nilaiSejarah, ada := nilai["sejarah"]; ada {
		fmt.Printf("Key 'sejarah' ada dengan nilai: %d\n", nilaiSejarah)
	} else {
		fmt.Println("Key 'sejarah' TIDAK ditemukan")
	}

	// Cara alternatif (tanpa if)
	nilaiKimia, ok := nilai["kimia"]
	fmt.Printf("\nKimia: %d, Ada: %v\n", nilaiKimia, ok)

	nilaiSeni, ok := nilai["seni"]
	fmt.Printf("Seni: %d (zero value), Ada: %v\n", nilaiSeni, ok)

	// =============================================================================
	// OPERASI: TAMBAH, UBAH, HAPUS
	// =============================================================================

	fmt.Println("\n=== OPERASI MAP ===")

	// 1. MENAMBAH - assign ke key yang belum ada
	nilai["sejarah"] = 87
	fmt.Printf("Setelah tambah 'sejarah': %v\n", nilai)

	// 2. MENGUBAH - assign ke key yang sudah ada
	nilai["fisika"] = 95
	fmt.Printf("Setelah ubah 'fisika': %v\n", nilai)

	// 3. MENGHAPUS - menggunakan delete()
	delete(nilai, "kimia")
	fmt.Printf("Setelah hapus 'kimia': %v\n", nilai)

	// Menghapus key yang tidak ada - tidak error, hanya tidak ada efek
	delete(nilai, "tidakada")
	fmt.Println("Hapus key yang tidak ada: tidak error")

	// =============================================================================
	// ITERASI MAP (DENGAN RANGE)
	// =============================================================================

	fmt.Println("\n=== ITERASI MAP ===")

	// Map tidak punya urutan tertentu!
	// Hasil iterasi bisa berbeda-beda setiap dijalankan

	fmt.Println("Iterasi map nilai:")
	for pelajaran, nilaiPelajaran := range nilai {
		fmt.Printf("  %s: %d\n", pelajaran, nilaiPelajaran)
	}

	// Hanya mengambil key
	fmt.Println("\nHanya key:")
	for pelajaran := range nilai {
		fmt.Printf("  %s\n", pelajaran)
	}

	// Hanya mengambil value (ignore key dengan _)
	fmt.Println("\nHanya value:")
	for _, n := range nilai {
		fmt.Printf("  %d\n", n)
	}

	// =============================================================================
	// PANJANG MAP
	// =============================================================================

	fmt.Println("\n=== PANJANG MAP ===")
	fmt.Printf("Jumlah pelajaran: %d\n", len(nilai))

	// =============================================================================
	// MAP DENGAN TIPE DATA KOMPLEKS
	// =============================================================================

	fmt.Println("\n=== MAP KOMPLEKS ===")

	// Map dengan value berupa slice
	hobi := map[string][]string{
		"budi": {"membaca", "coding", "gaming"},
		"ani":  {"menari", "menyanyi", "traveling"},
	}

	fmt.Printf("Hobi: %v\n", hobi)
	fmt.Printf("Hobi Budi: %v\n", hobi["budi"])
	fmt.Printf("Hobi Budi ke-1: %s\n", hobi["budi"][1]) // coding

	// Map dengan value berupa map (nested map)
	kelas := map[string]map[string]string{
		"budi": {
			"jurusan": "Informatika",
			"kelas":   "A",
		},
		"ani": {
			"jurusan": "Sistem Informasi",
			"kelas":   "B",
		},
	}

	fmt.Printf("\nData Kelas: %v\n", kelas)
	fmt.Printf("Jurusan Budi: %s\n", kelas["budi"]["jurusan"])

	// =============================================================================
	// MAP ADALAH REFERENCE TYPE
	// =============================================================================

	fmt.Println("\n=== MAP REFERENCE TYPE ===")

	original := map[string]int{"a": 1, "b": 2}

	// reference menunjuk ke data yang sama dengan original
	reference := original

	reference["a"] = 100

	fmt.Printf("Original:  %v\n", original)
	fmt.Printf("Reference: %v\n", reference)
	// Keduanya berubah karena mereferensi ke data yang sama!

	// =============================================================================
	// NIL MAP (MAP YANG BELUM DIINISIALISASI)
	// =============================================================================

	fmt.Println("\n=== NIL MAP ===")

	var nilMap map[string]int
	// nilMap belum diinisialisasi, nilainya nil

	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)

	// Membaca nil map aman (return zero value)
	fmt.Printf("Read nilMap['key']: %d\n", nilMap["key"])

	// TAPI menulis ke nil map akan PANIC (runtime error)!
	// nilMap["key"] = 100  // ❌ PANIC: assignment to entry in nil map

	// Solusi: inisialisasi dulu dengan make
	nilMap = make(map[string]int)
	nilMap["key"] = 100 // ✅ Sekarang aman
	fmt.Printf("Setelah inisialisasi: %v\n", nilMap)
}
