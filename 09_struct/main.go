/*
================================================================================
PELAJARAN 9: STRUCT DAN METHOD
================================================================================

STRUCT (Structure)
------------------
Struct adalah tipe data yang digunakan untuk mengelompokkan beberapa field/property
yang terkait menjadi satu kesatuan. Struct memungkinkan kita membuat tipe data
kustom (user-defined type) sesuai kebutuhan.

Bayangkan struct seperti formulir yang memiliki beberapa kolom. Setiap kolom
(field) bisa menyimpan tipe data yang berbeda-beda.

ANATOMI STRUCT:
---------------
type NamaStruct struct {
    Field1 TipeData1
    Field2 TipeData2
    ...
}

CONTOH REAL WORLD:
------------------
- Person: Nama, Umur, Alamat
- Product: Nama, Harga, Stok
- Car: Merk, Model, Tahun, Warna

METHOD PADA STRUCT
------------------
Method adalah fungsi yang "melekat" pada struct. Method memiliki akses ke field
struct melalui receiver.

Ada 2 jenis receiver:
1. Value Receiver   : (s Struct)    - Tidak bisa mengubah data asli (pass by value)
2. Pointer Receiver : (s *Struct)   - Bisa mengubah data asli (pass by reference)

NESTED STRUCT
-------------
Struct bisa berisi struct lain di dalamnya. Berguna untuk merepresentasikan
hubungan kompleks (misal: Employee punya Address).

ANONYMOUS STRUCT
----------------
Struct yang didefinisikan tanpa nama type. Berguna untuk one-time use.
*/

package main

import "fmt"

// =============================================================================
// DEFINISI STRUCT
// =============================================================================

// Person merepresentasikan data seseorang dengan nama, umur, dan alamat.
// Struct ini mendemonstrasikan struct dasar dengan field sederhana.
type Person struct {
	Nama   string // Field untuk menyimpan nama (tipe data string)
	Umur   int    // Field untuk menyimpan umur dalam tahun
	Alamat string // Field untuk menyimpan alamat tempat tinggal
}

// Product merepresentasikan barang dagangan.
// Mendemonstrasikan penggunaan berbagai tipe data dalam satu struct.
type Product struct {
	Nama     string  // Nama produk
	Harga    float64 // Harga dengan desimal (mata uang)
	Stok     int     // Jumlah unit tersedia
	Tersedia bool    // Status ketersediaan (true/false)
}

// =============================================================================
// NESTED STRUCT
// =============================================================================

// Address merepresentasikan alamat lengkap.
// Struct ini akan digunakan sebagai field di struct lain (nested struct).
type Address struct {
	Jalan   string // Nama jalan dan nomor rumah
	Kota    string // Nama kota
	KodePos string // Kode pos (disimpan sebagai string karena tidak untuk kalkulasi)
}

// Employee merepresentasikan karyawan.
// Mendemonstrasikan nested struct: Employee memiliki field Address yang juga struct.
type Employee struct {
	Nama    string  // Nama karyawan
	Umur    int     // Umur karyawan
	Address Address // Nested struct: field Address bertipe struct Address
}

// =============================================================================
// METHOD DENGAN VALUE RECEIVER
// =============================================================================

// Perkenalan adalah method dari Person yang mencetak perkenalan diri.
// (p Person) adalah VALUE RECEIVER - artinya method ini menerima COPY dari Person.
// Perubahan pada 'p' di dalam method ini TIDAK akan memengaruhi instance asli.
func (p Person) Perkenalan() {
	fmt.Printf("Halo, nama saya %s, umur %d tahun\n", p.Nama, p.Umur)
}

// IsAdult adalah method yang mengembalikan boolean apakah seseorang sudah dewasa.
// Method ini mendemonstrasikan method dengan return value.
// Di Indonesia, umur dewasa adalah 18 tahun ke atas.
func (p Person) IsAdult() bool {
	return p.Umur >= 18
}

// GetInfo mengembalikan string informasi lengkap tentang produk.
// Method ini mendemonstrasikan penggunaan kondisi di dalam method.
func (prod Product) GetInfo() string {
	status := "Tersedia"
	if !prod.Tersedia {
		status = "Habis"
	}
	return fmt.Sprintf("%s - Rp%.2f (%s)", prod.Nama, prod.Harga, status)
}

// =============================================================================
// METHOD DENGAN POINTER RECEIVER
// =============================================================================

// Birthday adalah method yang menambah umur sebanyak 1 tahun.
// (p *Person) adalah POINTER RECEIVER - artinya method ini menerima ALAMAT MEMORI.
// Perubahan pada 'p' di dalam method ini AKAN memengaruhi instance asli.
// Tanda * menunjukkan ini adalah pointer (referensi ke data asli).
func (p *Person) Birthday() {
	p.Umur++ // Umur bertambah 1 karena mengakses data asli melalui pointer
}

// UpdateAlamat mengubah alamat person.
// Menggunakan pointer receiver agar perubahan tersimpan.
func (p *Person) UpdateAlamat(alamatBaru string) {
	p.Alamat = alamatBaru
}

// KurangiStok mengurangi stok produk sebanyak jumlah yang diberikan.
// Menggunakan pointer receiver karena stok berubah.
func (prod *Product) KurangiStok(jumlah int) {
	if jumlah > prod.Stok {
		fmt.Println("ERROR: Stok tidak mencukupi!")
		return
	}
	prod.Stok -= jumlah
	fmt.Printf("Stok %s berkurang %d unit. Sisa: %d\n", prod.Nama, jumlah, prod.Stok)

	// Jika stok habis, update status
	if prod.Stok == 0 {
		prod.Tersedia = false
	}
}

func main() {
	fmt.Println("================================================================================")
	fmt.Println("STRUCT DAN METHOD")
	fmt.Println("================================================================================\n")

	// =============================================================================
	// 1. MEMBUAT INSTANCE STRUCT - CARA 1: MENGGUNAKAN FIELD NAMES (REKOMENDASI)
	// =============================================================================
	// Cara ini paling jelas dan tidak bergantung pada urutan field.
	// Sangat direkomendasikan karena mudah dibaca dan aman dari kesalahan urutan.
	fmt.Println("--- 1. Inisialisasi dengan Field Names ---")

	person1 := Person{
		Nama:   "Budi Santoso",
		Umur:   25,
		Alamat: "Jl. Merdeka No. 123, Jakarta",
	}

	fmt.Printf("Person 1:\n")
	fmt.Printf("  Nama  : %s\n", person1.Nama)
	fmt.Printf("  Umur  : %d tahun\n", person1.Umur)
	fmt.Printf("  Alamat: %s\n", person1.Alamat)

	// =============================================================================
	// 2. MEMBUAT INSTANCE STRUCT - CARA 2: TANPA FIELD NAMES
	// =============================================================================
	// Cara ini lebih singkat tapi bergantung pada URUTAN field di definisi struct.
	// Harus hati-hati: jika urutan salah, data akan tertukar!
	// Urutan harus sesuai definisi: Nama (string), Umur (int), Alamat (string)
	fmt.Println("\n--- 2. Inisialisasi Tanpa Field Names ---")

	person2 := Person{"Ani Wijaya", 22, "Jl. Sudirman No. 45, Bandung"}

	fmt.Printf("Person 2: %+v\n", person2)
	// %+v menampilkan field names beserta nilainya, berguna untuk debugging

	// =============================================================================
	// 3. ZERO VALUE (NILAI DEFAULT)
	// =============================================================================
	// Di Go, variabel yang dideklarasikan tanpa nilai awal akan memiliki zero value:
	// - string: "" (string kosong)
	// - int   : 0
	// - bool  : false
	// - pointer: nil
	fmt.Println("\n--- 3. Zero Value ---")

	var person3 Person // Deklarasi tanpa inisialisasi

	fmt.Printf("Zero value Person:\n")
	fmt.Printf("  Nama  : \"%s\"\n", person3.Nama)
	fmt.Printf("  Umur  : %d\n", person3.Umur)
	fmt.Printf("  Alamat: \"%s\"\n", person3.Alamat)

	// =============================================================================
	// 4. MENGAKSES DAN MENGUBAH FIELD
	// =============================================================================
	// Mengakses field menggunakan dot notation (titik): instance.Field
	fmt.Println("\n--- 4. Mengakses dan Mengubah Field ---")

	// Membaca field
	fmt.Printf("Nama person1 sebelum: %s\n", person1.Nama)

	// Mengubah nilai field
	person1.Nama = "Budi Santoso Update"
	fmt.Printf("Nama person1 sesudah: %s\n", person1.Nama)

	// =============================================================================
	// 5. NESTED STRUCT (STRUCT BERSARANG)
	// =============================================================================
	// Struct bisa berisi struct lain, memungkinkan hierarki data yang kompleks.
	// Akses field nested menggunakan double dot: instance.FieldNested.Field
	fmt.Println("\n--- 5. Nested Struct ---")

	employee1 := Employee{
		Nama: "Doni Pratama",
		Umur: 28,
		Address: Address{
			Jalan:   "Jl. Gatot Subroto Kav. 12",
			Kota:    "Jakarta Selatan",
			KodePos: "12930",
		},
	}

	fmt.Printf("Employee: %s (%d tahun)\n", employee1.Nama, employee1.Umur)
	fmt.Printf("Alamat Lengkap:\n")
	fmt.Printf("  Jalan  : %s\n", employee1.Address.Jalan)
	fmt.Printf("  Kota   : %s\n", employee1.Address.Kota)
	fmt.Printf("  Kode Pos: %s\n", employee1.Address.KodePos)

	// =============================================================================
	// 6. METHOD DENGAN VALUE RECEIVER
	// =============================================================================
	// Method ini bekerja pada COPY dari struct, tidak mengubah data asli.
	// Cocok untuk operasi yang hanya membaca atau menghitung tanpa modifikasi.
	fmt.Println("\n--- 6. Method dengan Value Receiver ---")

	person1.Perkenalan()

	if person1.IsAdult() {
		fmt.Println("Status: Sudah dewasa (≥18 tahun)")
	} else {
		fmt.Println("Status: Belum dewasa (<18 tahun)")
	}

	// =============================================================================
	// 7. METHOD DENGAN POINTER RECEIVER
	// =============================================================================
	// Method ini bekerja pada data asli (melalui referensi/pointer).
	// Perubahan di method akan tersimpan di instance asli.
	// Go otomatis mengkonversi (&instance).Method() menjadi instance.Method()
	fmt.Println("\n--- 7. Method dengan Pointer Receiver ---")

	fmt.Printf("Umur person1 sebelum birthday: %d\n", person1.Umur)
	person1.Birthday() // Go otomatis pass sebagai pointer meski kita pakai instance
	fmt.Printf("Umur person1 sesudah birthday: %d\n", person1.Umur)

	// Update alamat
	fmt.Printf("Alamat sebelum: %s\n", person1.Alamat)
	person1.UpdateAlamat("Jl. Thamrin No. 100, Jakarta Pusat")
	fmt.Printf("Alamat sesudah: %s\n", person1.Alamat)

	// =============================================================================
	// 8. SLICE DARI STRUCT
	// =============================================================================
	// Slice bisa menyimpan banyak instance struct, berguna untuk koleksi data.
	fmt.Println("\n--- 8. Slice dari Struct ---")

	students := []Person{
		{Nama: "Eka Putri", Umur: 20, Alamat: "Surabaya"},
		{Nama: "Fani Nugraha", Umur: 21, Alamat: "Yogyakarta"},
		{Nama: "Gilang Ramadhan", Umur: 19, Alamat: "Semarang"},
	}

	fmt.Println("Daftar Mahasiswa:")
	for i, student := range students {
		fmt.Printf("  %d. %s (%d tahun) - %s\n", i+1, student.Nama, student.Umur, student.Alamat)
	}

	// =============================================================================
	// 9. MAP DENGAN STRUCT SEBAGAI VALUE
	// =============================================================================
	// Map key-value di mana value-nya adalah struct.
	// Berguna untuk lookup cepat berdasarkan key (misal: kode produk).
	fmt.Println("\n--- 9. Map dengan Struct Value ---")

	products := map[string]Product{
		"P001": {Nama: "Laptop Gaming", Harga: 15000000, Stok: 10, Tersedia: true},
		"P002": {Nama: "Mouse Wireless", Harga: 250000, Stok: 50, Tersedia: true},
		"P003": {Nama: "Headset", Harga: 500000, Stok: 0, Tersedia: false},
	}

	fmt.Println("Daftar Produk:")
	for code, product := range products {
		fmt.Printf("  %s: %s\n", code, product.GetInfo())
	}

	// Demonstrasi method dengan pointer receiver pada map
	fmt.Println("\nSimulasi pembelian:")
	productRef := products["P001"] // Dapatkan copy
	productRef.KurangiStok(2)      // Method ini tidak akan mengubah map karena kita pakai copy!

	// Cara benar mengubah struct dalam map (gunakan pointer atau langsung assign)
	prod := products["P002"]
	prod.Stok -= 5
	products["P002"] = prod // Update map dengan nilai baru
	fmt.Printf("Stok P002 setelah dikurangi: %d\n", products["P002"].Stok)

	// =============================================================================
	// 10. ANONYMOUS STRUCT
	// =============================================================================
	// Anonymous struct adalah struct tanpa nama type yang didefinisikan.
	// Berguna untuk data satu kali pakai yang tidak perlu didefinisikan sebagai type.
	// Sintaks: variable := struct { fields... }{ values... }
	fmt.Println("\n--- 10. Anonymous Struct ---")

	user := struct {
		Username string
		Email    string
		IsActive bool
	}{
		Username: "john_doe",
		Email:    "john@example.com",
		IsActive: true,
	}

	fmt.Printf("Anonymous struct - User:\n")
	fmt.Printf("  Username: %s\n", user.Username)
	fmt.Printf("  Email   : %s\n", user.Email)
	fmt.Printf("  Active  : %v\n", user.IsActive)

	// =============================================================================
	// 11. PERBANDINGAN STRUCT
	// =============================================================================
	// Struct bisa dibandingkan menggunakan operator == jika semua field-nya comparable.
	// string, int, bool adalah comparable. Slice, map, function TIDAK comparable.
	fmt.Println("\n--- 11. Perbandingan Struct ---")

	a := Person{Nama: "Budi", Umur: 25, Alamat: "Jakarta"}
	b := Person{Nama: "Budi", Umur: 25, Alamat: "Jakarta"}
	c := Person{Nama: "Ani", Umur: 25, Alamat: "Jakarta"}

	fmt.Printf("a == b: %v (identik)\n", a == b)
	fmt.Printf("a == c: %v (beda nama)\n", a == c)

	// =============================================================================
	// 12. POINTER KE STRUCT
	// =============================================================================
	// Pointer menyimpan alamat memori, bukan nilai langsung.
	// Pointer ke struct lebih efisien untuk struct besar (hemat memory copy).
	fmt.Println("\n--- 12. Pointer ke Struct ---")

	// Membuat pointer dengan &
	personPtr := &Person{
		Nama:   "Caca Handika",
		Umur:   30,
		Alamat: "Jl. Asia Afrika No. 1",
	}

	fmt.Printf("Pointer: %p\n", personPtr)
	fmt.Printf("Value  : %+v\n", *personPtr)

	// Akses field dari pointer - Go otomatis dereference
	// personPtr.Nama sama dengan (*personPtr).Nama
	fmt.Printf("Nama via pointer: %s\n", personPtr.Nama)

	// =============================================================================
	// 13. EMBEDDED STRUCT (ANONYMOUS FIELD)
	// =============================================================================
	// Struct bisa di-embed (tanpa nama field) untuk komposisi.
	// Field dari struct yang di-embed langsung bisa diakses.
	fmt.Println("\n--- 13. Embedded Struct ---")

	type Manager struct {
		Person     // Embedded struct - tanpa nama field
		Department string
	}

	manager := Manager{
		Person:     Person{Nama: "Direktur", Umur: 45, Alamat: "Kantor Pusat"},
		Department: "IT",
	}

	// Akses langsung field dari Person meski Person di-embed
	fmt.Printf("Manager: %s, Dept: %s, Umur: %d\n", manager.Nama, manager.Department, manager.Umur)

	fmt.Println("\n================================================================================")
	fmt.Println("SELESAI - Silakan eksplorasi dan modifikasi kode ini untuk pemahaman lebih baik")
	fmt.Println("================================================================================")
}
