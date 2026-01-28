package main

import "fmt"

// Definisi struct - blueprint untuk data
// Struct mengelompokkan beberapa field/property terkait
type Person struct {
	Nama   string
	Umur   int
	Alamat string
}

// Struct dengan field berbeda tipe data
type Product struct {
	Nama     string
	Harga    float64
	Stok     int
	Tersedia bool
}

// Nested struct - struct di dalam struct
type Address struct {
	Jalan  string
	Kota   string
	KodePos string
}

type Employee struct {
	Nama    string
	Umur    int
	Address Address // nested struct
}

// Method - fungsi yang "melekat" pada struct
// (p Person) disebut receiver - struct yang menerima method ini
func (p Person) Perkenalan() {
	fmt.Printf("Halo, nama saya %s, umur %d tahun\n", p.Nama, p.Umur)
}

// Method dengan parameter
func (p Person) IsAdult() bool {
	return p.Umur >= 18
}

// Method dengan pointer receiver - bisa mengubah data asli
func (p *Person) Birthday() {
	p.Umur++ // umur bertambah 1
}

// Method untuk Product
func (prod Product) Info() string {
	status := "Tersedia"
	if !prod.Tersedia {
		status = "Habis"
	}
	return fmt.Sprintf("%s - Rp%.2f (%s)", prod.Nama, prod.Harga, status)
}

func main() {
	fmt.Println("=== STRUCT DI GO ===\n")

	// ========================================
	// 1. MEMBUAT INSTANCE STRUCT
	// ========================================
	fmt.Println("--- 1. Membuat Instance Struct ---")

	// Cara 1: Inisialisasi lengkap dengan field names (REKOMENDASI)
	person1 := Person{
		Nama:   "Budi",
		Umur:   25,
		Alamat: "Jakarta",
	}
	fmt.Printf("Person 1: %+v\n", person1)

	// Cara 2: Tanpa menyebutkan field names (urutan harus benar!)
	person2 := Person{"Ani", 22, "Bandung"}
	fmt.Printf("Person 2: %+v\n", person2)

	// Cara 3: Struct kosong (zero value)
	var person3 Person
	fmt.Printf("Person 3 (kosong): %+v\n", person3)

	// Cara 4: Menggunakan new() - mengembalikan pointer
	person4 := new(Person)
	person4.Nama = "Caca"
	person4.Umur = 30
	fmt.Printf("Person 4 (pointer): %+v\n", *person4)

	fmt.Println()

	// ========================================
	// 2. MENGAKSES FIELD STRUCT
	// ========================================
	fmt.Println("--- 2. Mengakses Field ---")
	fmt.Printf("Nama person1: %s\n", person1.Nama)
	fmt.Printf("Umur person1: %d\n", person1.Umur)

	// Mengubah nilai field
	person1.Umur = 26
	fmt.Printf("Umur person1 setelah update: %d\n", person1.Umur)

	fmt.Println()

	// ========================================
	// 3. NESTED STRUCT
	// ========================================
	fmt.Println("--- 3. Nested Struct ---")
	employee1 := Employee{
		Nama: "Doni",
		Umur: 28,
		Address: Address{
			Jalan:   "Jl. Sudirman No. 1",
			Kota:    "Jakarta",
			KodePos: "10110",
		},
	}
	fmt.Printf("Employee: %+v\n", employee1)
	fmt.Printf("Kota: %s\n", employee1.Address.Kota)

	fmt.Println()

	// ========================================
	// 4. METHOD PADA STRUCT
	// ========================================
	fmt.Println("--- 4. Method ---")
	person1.Perkenalan()

	if person1.IsAdult() {
		fmt.Println("Sudah dewasa")
	} else {
		fmt.Println("Belum dewasa")
	}

	fmt.Println()

	// ========================================
	// 5. POINTER DAN METHOD
	// ========================================
	fmt.Println("--- 5. Pointer Receiver ---")
	fmt.Printf("Umur sebelum birthday: %d\n", person1.Umur)
	person1.Birthday() // otomatis dipass sebagai pointer
	fmt.Printf("Umur setelah birthday: %d\n", person1.Umur)

	fmt.Println()

	// ========================================
	// 6. SLICE DAN MAP DENGAN STRUCT
	// ========================================
	fmt.Println("--- 6. Slice dan Map dengan Struct ---")

	// Slice of struct
	students := []Person{
		{Nama: "Eka", Umur: 20},
		{Nama: "Fani", Umur: 21},
		{Nama: "Gilang", Umur: 19},
	}

	fmt.Println("Daftar Mahasiswa:")
	for _, student := range students {
		fmt.Printf("  - %s (%d tahun)\n", student.Nama, student.Umur)
	}

	// Map dengan struct sebagai value
	products := map[string]Product{
		"P001": {Nama: "Laptop", Harga: 10000000, Stok: 10, Tersedia: true},
		"P002": {Nama: "Mouse", Harga: 150000, Stok: 0, Tersedia: false},
	}

	fmt.Println("\nDaftar Produk:")
	for code, product := range products {
		fmt.Printf("  %s: %s\n", code, product.Info())
	}

	fmt.Println()

	// ========================================
	// 7. ANONYMOUS STRUCT
	// ========================================
	fmt.Println("--- 7. Anonymous Struct ---")
	// Struct sekali pakai, tidak perlu definisi type terpisah
	user := struct {
		Username string
		Email    string
	}{
		Username: "john_doe",
		Email:    "john@example.com",
	}
	fmt.Printf("Anonymous struct: %+v\n", user)

	fmt.Println()

	// ========================================
	// 8. PERBANDINGAN STRUCT
	// ========================================
	fmt.Println("--- 8. Perbandingan Struct ---")
	a := Person{Nama: "Budi", Umur: 25}
	b := Person{Nama: "Budi", Umur: 25}
	c := Person{Nama: "Ani", Umur: 25}

	fmt.Printf("a == b: %v\n", a == b) // true
	fmt.Printf("a == c: %v\n", a == c) // false

	fmt.Println("\n=== SELESAI ===")
}
