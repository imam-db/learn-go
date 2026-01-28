/*
================================================================================
PELAJARAN 10: POINTER
================================================================================

POINTER (Penunjuk)
------------------
Pointer adalah variabel yang menyimpan ALAMAT MEMORI dari variabel lain,
bukan nilai langsung. Bayangkan seperti alamat rumah: pointer menyimpan
alamatnya, bukan isi rumahnya.

KENAPA PERLU POINTER?
---------------------
1. EFISIENSI MEMORY - Passing pointer lebih murah daripada copy data besar
2. MENGUBAH NILAI ASLI - Fungsi bisa mengubah variabel di luar scope-nya
3. KONKURENSI - Dibutuhkan untuk goroutine dan channel
4. DATA STRUKTUR - Linked list, tree, graph memerlukan pointer

ANATOMI POINTER:
----------------
var namaPointer *TipeData

- Tanda * (asterisk) menandakan ini adalah pointer
- TipeData adalah tipe data yang ditunjuk oleh pointer

DUA OPERATOR PENTING:
---------------------
1. & (Address-of) - Mendapatkan alamat memori suatu variabel
   Contoh: &variabel → mengembalikan alamat memori

2. * (Dereference) - Mengakses nilai di alamat yang ditunjuk pointer
   Contoh: *pointer → mengembalikan nilai asli

POINTER ZERO VALUE:
-------------------
Nilai default pointer adalah nil (bukan null seperti di bahasa lain)
nil berarti pointer tidak menunjuk ke mana-mana (tidak valid)

PERBANDINGAN: PASS BY VALUE vs PASS BY REFERENCE
-------------------------------------------------
PASS BY VALUE (default):
- Mengirim COPY dari data
- Aman, data asli tidak berubah
- Boros memory untuk data besar

PASS BY REFERENCE (pointer):
- Mengirim ALAMAT data
- Efisien, bisa mengubah data asli
- Perlu hati-hati agar tidak merusak data
*/

package main

import "fmt"

// Person adalah struct yang akan digunakan untuk demo pointer to struct
type Person struct {
	Nama string
	Umur int
}

// =============================================================================
// FUNGSI DENGAN PASS BY VALUE (TANPA POINTER)
// =============================================================================

// UpdateUmurValue menggunakan pass by value - tidak mengubah data asli
// Parameter p adalah COPY dari Person yang dikirim
func UpdateUmurValue(p Person, umurBaru int) {
	p.Umur = umurBaru
	fmt.Printf("  [Di dalam fungsi] Umur: %d\n", p.Umur)
	// Perubahan di sini HANYA berlaku di copy, data asli tidak berubah
}

// =============================================================================
// FUNGSI DENGAN PASS BY REFERENCE (DENGAN POINTER)
// =============================================================================

// UpdateUmurPointer menggunakan pointer - MENGUBAH data asli
// Parameter p adalah pointer (*Person), menunjuk ke data asli di memori
func UpdateUmurPointer(p *Person, umurBaru int) {
	p.Umur = umurBaru // Otomatis dereference, sama dengan (*p).Umur = umurBaru
	fmt.Printf("  [Di dalam fungsi] Umur: %d\n", p.Umur)
	// Perubahan di sini BERLAKU untuk data asli karena kita mengubah melalui alamat memori
}

// =============================================================================
// SWAP DENGAN POINTER
// =============================================================================

// Swap menukar nilai dua variabel menggunakan pointer
// Hanya bisa dilakukan dengan pointer, tidak bisa dengan pass by value
func Swap(a, b *int) {
	fmt.Printf("  [Swap] Sebelum: a=%d, b=%d\n", *a, *b)
	temp := *a // Simpan nilai yang ditunjuk a
	*a = *b    // Ubah nilai yang ditunjuk a menjadi nilai yang ditunjuk b
	*b = temp  // Ubah nilai yang ditunjuk b menjadi temp (nilai awal a)
	fmt.Printf("  [Swap] Sesudah: a=%d, b=%d\n", *a, *b)
}

func main() {
	fmt.Println("================================================================================")
	fmt.Println("POINTER")
	fmt.Println("================================================================================\n")

	// =============================================================================
	// 1. DEKLARASI POINTER DASAR
	// =============================================================================
	// Pointer dideklarasikan dengan tanda * sebelum tipe data
	// Zero value pointer adalah nil (tidak menunjuk ke mana-mana)
	fmt.Println("--- 1. Deklarasi Pointer Dasar ---")

	var ptr *int // Pointer ke int, nilai awal: nil
	fmt.Printf("Pointer ptr: %v (nil)\n", ptr)

	// Inisialisasi variabel biasa
	nilai := 42

	// & (address-of) mendapatkan alamat memori variabel
	ptr = &nilai

	fmt.Printf("Variabel nilai: %d\n", nilai)
	fmt.Printf("Alamat nilai (&nilai): %p\n", &nilai)
	fmt.Printf("Pointer ptr: %p\n", ptr)
	fmt.Printf("Nilai yang ditunjuk ptr (*ptr): %d\n", *ptr)

	// =============================================================================
	// 2. DEREFERENCE (MENGAKSES NILAI MELALUI POINTER)
	// =============================================================================
	// *pointer mengakses nilai di alamat yang ditunjuk pointer
	// Ini disebut "dereferencing"
	fmt.Println("\n--- 2. Dereference Pointer ---")

	x := 100
	p := &x // p menunjuk ke x

	fmt.Printf("x awal: %d\n", x)
	fmt.Printf("*p (dereference): %d\n", *p)

	// Mengubah nilai melalui pointer akan mengubah nilai asli!
	*p = 200 // x juga berubah menjadi 200
	fmt.Printf("Setelah *p = 200:\n")
	fmt.Printf("  x: %d\n", x)
	fmt.Printf("  *p: %d\n", *p)

	// =============================================================================
	// 3. PERBEDAAN: PASS BY VALUE vs PASS BY REFERENCE
	// =============================================================================
	// Demonstrasi pentingnya pointer untuk mengubah data di luar fungsi
	fmt.Println("\n--- 3. Pass By Value vs Pass By Reference ---")

	person := Person{Nama: "Budi", Umur: 25}
	fmt.Printf("Person awal: %+v\n", person)

	// PASS BY VALUE - tidak mengubah data asli
	fmt.Println("\nPass By Value (UpdateUmurValue):")
	UpdateUmurValue(person, 30)
	fmt.Printf("  [Setelah fungsi] Person: %+v (TIDAK BERUBAH!)\n", person)

	// PASS BY REFERENCE - mengubah data asli
	fmt.Println("\nPass By Reference (UpdateUmurPointer):")
	UpdateUmurPointer(&person, 30) // Kirim alamat dengan &
	fmt.Printf("  [Setelah fungsi] Person: %+v (BERUBAH!)\n", person)

	// =============================================================================
	// 4. POINTER SEBAGAI PARAMETER FUNGSI
	// =============================================================================
	// Pointer memungkinkan fungsi mengembalikan multiple result melalui parameter
	fmt.Println("\n--- 4. Swap dengan Pointer ---")

	a, b := 10, 20
	fmt.Printf("Sebelum Swap: a=%d, b=%d\n", a, b)
	Swap(&a, &b) // Kirim alamat a dan b
	fmt.Printf("Setelah Swap: a=%d, b=%d\n", a, b)

	// =============================================================================
	// 5. POINTER KE ARRAY
	// =============================================================================
	// Pointer bisa menunjuk ke array, tapi Go lebih sering menggunakan slice
	fmt.Println("\n--- 5. Pointer ke Array ---")

	arr := [3]int{10, 20, 30}
	arrPtr := &arr // Pointer ke array

	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Pointer ke array: %p\n", arrPtr)

	// Akses elemen array melalui pointer
	fmt.Printf("(*arrPtr)[0]: %d\n", (*arrPtr)[0])

	// Mengubah elemen melalui pointer
	(*arrPtr)[1] = 200
	fmt.Printf("Array setelah diubah via pointer: %v\n", arr)

	// Cara alternatif (Go mengizinkan sintaks yang lebih bersih)
	arrPtr[2] = 300 // Sama dengan (*arrPtr)[2] = 300
	fmt.Printf("Array setelah diubah lagi: %v\n", arr)

	// =============================================================================
	// 6. POINTER DAN SLICE
	// =============================================================================
	// Slice sudah merupakan reference type (mirip pointer ke array)
	// Mengirim slice ke fungsi sudah otomatis pass by reference
	fmt.Println("\n--- 6. Slice (Sudah Reference Type) ---")

	slice := []int{1, 2, 3}
	fmt.Printf("Slice awal: %v\n", slice)

	// Fungsi yang menerima slice bisa mengubah data asli (tanpa pointer!)
	modifikasiSlice := func(s []int) {
		s[0] = 999 // Ini akan mengubah slice asli!
		fmt.Printf("  [Di dalam fungsi] Slice: %v\n", s)
	}

	modifikasiSlice(slice)
	fmt.Printf("Slice setelah modifikasi: %v\n", slice)

	// Tapi append tidak mengubah slice asli karena mungkin alokasi memori baru
	appendKeSlice := func(s []int) {
		s = append(s, 4) // Ini TIDAK mengubah slice asli
		fmt.Printf("  [Di dalam fungsi] Slice: %v (len=%d)\n", s, len(s))
	}

	fmt.Printf("\nSlice sebelum append: %v (len=%d)\n", slice, len(slice))
	appendKeSlice(slice)
	fmt.Printf("Slice setelah append: %v (len=%d) - TIDAK BERUBAH!\n", slice, len(slice))

	// =============================================================================
	// 7. NIL POINTER
	// =============================================================================
	// Nil pointer adalah pointer yang tidak menunjuk ke mana-mana
	// Mengakses nil pointer akan menyebabkan RUNTIME PANIC!
	fmt.Println("\n--- 7. Nil Pointer ---")

	var nilPtr *int // Default: nil
	fmt.Printf("Nil pointer: %v\n", nilPtr)

	// WAJIB cek nil sebelum dereference!
	if nilPtr != nil {
		fmt.Printf("Nilai: %d\n", *nilPtr)
	} else {
		fmt.Println("Pointer adalah nil, tidak bisa di-dereference!")
	}

	// =============================================================================
	// 8. NEW() FUNCTION
	// =============================================================================
	// new(T) mengalokasikan memori untuk tipe T dan mengembalikan pointer ke T
	// Nilai diinisialisasi dengan zero value dari tipe tersebut
	fmt.Println("\n--- 8. new() Function ---")

	// Cara 1: new() - mengembalikan pointer
	ptrInt := new(int)    // Alokasi memori untuk int, return *int
	ptrStr := new(string) // Alokasi memori untuk string, return *string

	fmt.Printf("ptrInt: %p, nilai: %d\n", ptrInt, *ptrInt) // *ptrInt = 0 (zero value)
	fmt.Printf("ptrStr: %p, nilai: \"%s\"\n", ptrStr, *ptrStr) // *ptrStr = "" (zero value)

	// Mengisi nilai
	*ptrInt = 42
	*ptrStr = "Hello"
	fmt.Printf("Setelah diisi - ptrInt: %d, ptrStr: \"%s\"\n", *ptrInt, *ptrStr)

	// Cara 2: &struct{} - lebih umum untuk struct
	// Person baru dengan new
	personNew := new(Person)
	personNew.Nama = "Ani"
	personNew.Umur = 22
	fmt.Printf("Person dengan new(): %+v\n", *personNew)

	// Cara 3: &struct{} literal - lebih idiomatic
	personLiteral := &Person{
		Nama: "Budi",
		Umur: 25,
	}
	fmt.Printf("Person dengan &struct{{}}: %+v\n", *personLiteral)

	// =============================================================================
	// 9. POINTER TO POINTER (DOUBLE POINTER)
	// =============================================================================
	// Pointer bisa menunjuk ke pointer lain
	// Berguna untuk mengubah nilai pointer itu sendiri (misal: realloc)
	fmt.Println("\n--- 9. Pointer to Pointer ---")

	xVal := 10
	p1 := &xVal   // p1 menunjuk ke xVal
	p2 := &p1     // p2 menunjuk ke p1 (pointer to pointer)

	fmt.Printf("xVal: %d\n", xVal)
	fmt.Printf("p1 (alamat xVal): %p, *p1: %d\n", p1, *p1)
	fmt.Printf("p2 (alamat p1): %p, *p2: %p, **p2: %d\n", p2, *p2, **p2)
	// **p2 = dereference 2 kali: p2 → p1 → xVal

	// =============================================================================
	// 10. POINTER RECEIVER PADA METHOD (REVIEW DARI MATERI STRUCT)
	// =============================================================================
	// Method dengan pointer receiver bisa mengubah data struct asli
	fmt.Println("\n--- 10. Pointer Receiver pada Method ---")

	pReceiver := Person{Nama: "Caca", Umur: 30}
	fmt.Printf("Sebelum Birthday: %+v\n", pReceiver)

	// Method dengan pointer receiver
	pReceiver.Birthday() // Go otomatis konversi (&pReceiver).Birthday()
	fmt.Printf("Setelah Birthday: %+v\n", pReceiver)

	// =============================================================================
	// 11. BEST PRACTICES DAN PITFALLS
	// =============================================================================
	fmt.Println("\n--- 11. Best Practices dan Pitfalls ---")

	// ✅ DO: Gunakan pointer untuk struct besar agar efisien
	tipe BigStruct struct {
		Data [1000]int
	}
	big := BigStruct{}
	bigPtr := &big // Hanya menyimpan alamat (8 byte), bukan copy 1000 int
	fmt.Printf("Ukuran big: sekitar %d bytes\n", len(big.Data)*8)
	fmt.Printf("Ukuran bigPtr: 8 bytes (hanya alamat memori)\n")

	// ✅ DO: Selalu cek nil pointer sebelum dereference
	var maybeNil *int
	if maybeNil != nil {
		fmt.Println(*maybeNil)
	} else {
		fmt.Println("Aman: pointer dicek sebelum digunakan")
	}

	// ❌ DON'T: Return pointer ke variabel lokal (stack) - bisa jadi dangling pointer
	// Di Go ini sebenarnya aman karena compiler escape analysis,
	// tapi di bahasa lain sangat berbahaya!

	// ❌ DON'T: Mengakses nil pointer - akan PANIC
	// var pPanic *int
	// fmt.Println(*pPanic) // PANIC: runtime error: invalid memory address

	// =============================================================================
	// 12. POINTER VS VALUE: KAPAN MENGGUNAKAN?
	// =============================================================================
	fmt.Println("\n--- 12. Pointer vs Value: Panduan Pemilihan ---")

	fmt.Println("Gunakan POINTER ketika:")
	fmt.Println("  - Method perlu mengubah field struct")
	fmt.Println("  - Struct sangat besar (hemat memory copy)")
	fmt.Println("  - Fungsi perlu mengubah parameter asli (Swap, dll)")
	fmt.Println("  - Data bisa bernilai nil (opsional)")

	fmt.Println("\nGunakan VALUE ketika:")
	fmt.Println("  - Data kecil (int, bool, small struct)")
	fmt.Println("  - Tidak perlu mengubah data asli")
	fmt.Println("  - Inin menghindari side effect)")
	fmt.Println("  - Data harus immutable (tidak berubah)")

	fmt.Println("\n================================================================================")
	fmt.Println("SELESAI - Pointer adalah konsep fundamental untuk efisiensi dan flexibilitas")
	fmt.Println("================================================================================")
}

// Birthday adalah method dengan pointer receiver
// Method ini menambah umur 1 tahun dan mengubah data asli
func (p *Person) Birthday() {
	p.Umur++ // Otomatis dereference, sama dengan (*p).Umur++
	fmt.Printf("  [Birthday] Selamat ulang tahun! Umur sekarang: %d\n", p.Umur)
}
