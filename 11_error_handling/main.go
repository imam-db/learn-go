/*
================================================================================
PELAJARAN 11: ERROR HANDLING
================================================================================

ERROR HANDLING DI GO
--------------------
Go memiliki pendekatan error handling yang berbeda dari bahasa lain:
- TIDAK ADA try-catch-finally seperti Java, Python, JavaScript
- Error diperlakukan sebagai VALUE (bukan exception)
- Fungsi mengembalikan error sebagai return value terakhir
- Konvensi: return type adalah (result, error)

MENGAPA GO TIDAK PAKAI TRY-CATCH?
---------------------------------
1. EXPLICIT - Programmer harus explicitly menangani error
2. SIMPLE - Tidak ada stack unwinding yang kompleks
3. CONTROL - Programmer punya full control alur program
4. PERFORMANCE - Tidak ada overhead exception handling

INTERFACE ERROR
---------------
type error interface {
    Error() string
}

Semua tipe yang mengimplementasikan method Error() string adalah error.

MEMBUAT ERROR BARU
------------------
1. errors.New("pesan error")     - Error sederhana
2. fmt.Errorf("format", args...) - Error dengan formatting
3. Custom error type             - Error dengan informasi tambahan

ERROR HANDLING PATTERNS
-----------------------
1. Check and Return    - if err != nil { return err }
2. Check and Handle    - if err != nil { // handle error }
3. Check and Wrap      - if err != nil { return fmt.Errorf("context: %w", err) }

PANIC DAN RECOVER
-----------------
- panic()  - Menghentikan program secara abnormal (seperti exception)
- recover() - Menangkap panic, mencegah program crash
- defer()   - Menunda eksekusi hingga fungsi selesai, berguna untuk cleanup

KAPAN MENGGUNAKAN PANIC?
------------------------
- Hanya untuk kondisi yang benar-benar fatal dan tidak bisa dipulihkan
- Jangan gunakan panic untuk error biasa yang bisa ditangani
*/

package main

import (
	"errors"
	"fmt"
)

// =============================================================================
// CUSTOM ERROR TYPE
// =============================================================================

// ValidationError adalah custom error untuk validasi input
// Custom error memungkinkan kita menambahkan informasi kontekstual
type ValidationError struct {
	Field   string // Field yang gagal validasi
	Message string // Pesan error detail
}

// Error() mengimplementasikan interface error
// Method ini wajib ada agar ValidationError bisa diperlakukan sebagai error
func (e ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal pada field '%s': %s", e.Field, e.Message)
}

// NotFoundError adalah custom error untuk data tidak ditemukan
type NotFoundError struct {
	Resource string // Jenis resource (user, product, etc)
	ID       string // Identifier yang dicari
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s dengan ID '%s' tidak ditemukan", e.Resource, e.ID)
}

// =============================================================================
// FUNGSI DENGAN ERROR RETURN
// =============================================================================

// Divide membagi dua angka dengan integer division
// Mengembalikan (hasil, error) - konvensi Go untuk error handling
// Jika pembagi 0, kembalikan error (tidak bisa dibagi nol)
func Divide(a, b int) (int, error) {
	// Cek kondisi error
	if b == 0 {
		// Kembalikan zero value untuk hasil dan error
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	// Jika sukses, kembalikan hasil dan nil (tidak ada error)
	return a / b, nil
}

// Sqrt menghitung akar kuadrat
// Mengembalikan error untuk input negatif (tidak ada akar real)
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("tidak bisa menghitung akar dari bilangan negatif: %f", x)
	}
	// Simulasi perhitungan (sederhana)
	return x * 0.5, nil // Ini bukan sqrt beneran, hanya contoh
}

// FindUser mensimulasikan pencarian user di database
// Mengembalikan custom error NotFoundError jika user tidak ada
func FindUser(id string) (string, error) {
	// Simulasi database
	users := map[string]string{
		"001": "Budi",
		"002": "Ani",
		"003": "Caca",
	}

	name, exists := users[id]
	if !exists {
		// Kembalikan custom error dengan konteks
		return "", NotFoundError{
			Resource: "User",
			ID:       id,
		}
	}
	return name, nil
}

// ValidateAge memvalidasi umur dengan custom error
func ValidateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Message: "umur tidak boleh negatif",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Message: "umur terlalu tinggi (maksimal 150)",
		}
	}
	return nil
}

// =============================================================================
// ERROR WRAPPING (Go 1.13+)
// =============================================================================

// DatabaseError untuk error layer database
type DatabaseError struct {
	Op  string // Operasi yang gagal
	Err error  // Error asli dari database
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error saat %s: %v", e.Op, e.Err)
}

// Unwrap memungkinkan error chain inspection
func (e DatabaseError) Unwrap() error {
	return e.Err
}

// GetUserFromDB mensimulasikan pengambilan data dari database
func GetUserFromDB(id string) (string, error) {
	// Simulasi error dari database
	dbErr := errors.New("connection timeout")

	// Wrap error dengan konteks tambahan
	return "", DatabaseError{
		Op:  "get user by id",
		Err: dbErr,
	}
}

// =============================================================================
// PANIC DAN RECOVER
// =============================================================================

// SafeDivide membagi dengan recover dari panic
// recover() hanya bekerja di dalam defer function
func SafeDivide(a, b int) (result int, err error) {
	// defer akan dieksekusi saat fungsi selesai, meski ada panic
	defer func() {
		// recover() menangkap panic dan mengembalikan nilai yang dipanic
		if r := recover(); r != nil {
			// Konversi panic message ke error
			err = fmt.Errorf("panic recovered: %v", r)
			result = 0
		}
	}()

	// Jika b == 0, kita panic (meski ini tidak direkomendasikan untuk case sederhana)
	// Dalam praktik nyata, gunakan error return untuk kasus ini
	if b == 0 {
		panic("tidak bisa membagi dengan nol") // ❌ Jangan gunakan panic untuk ini!
	}

	return a / b, nil
}

// ProcessFile mensimulasikan pemrosesan file dengan defer
// defer digunakan untuk cleanup resources
func ProcessFile(filename string) error {
	// Simulasi: buka file
	fmt.Printf("Membuka file: %s\n", filename)

	// defer akan dieksekusi saat fungsi return, meski ada error
	// Urutan defer: LIFO (Last In First Out)
	defer fmt.Println("Defer 1: Cleanup resource A")
	defer fmt.Println("Defer 2: Cleanup resource B")
	defer fmt.Println("Defer 3: Close file")

	// Simulasi error
	if filename == "" {
		return errors.New("filename tidak boleh kosong")
	}

	fmt.Println("Memproses file...")
	return nil
}

func main() {
	fmt.Println("================================================================================")
	fmt.Println("ERROR HANDLING")
	fmt.Println("================================================================================\n")

	// =============================================================================
	// 1. ERROR HANDLING DASAR
	// =============================================================================
	// Pattern: result, err := Function(); if err != nil { // handle }
	fmt.Println("--- 1. Error Handling Dasar ---")

	// Contoh sukses
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d (sukses)\n", result)
	}

	// Contoh error
	result, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("10 / 0 = Error: %v\n", err)
	} else {
		fmt.Printf("Hasil: %d\n", result)
	}

	// =============================================================================
	// 2. ERROR DENGAN FORMATTING
	// =============================================================================
	// fmt.Errorf untuk error dengan informasi dinamis
	fmt.Println("\n--- 2. Error dengan Formatting ---")

	_, err = Sqrt(-4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// =============================================================================
	// 3. CUSTOM ERROR TYPE
	// =============================================================================
	// Custom error memberikan informasi lebih kontekstual
	fmt.Println("\n--- 3. Custom Error Type ---")

	// NotFoundError
	name, err := FindUser("999")
	if err != nil {
		// Cek tipe error dengan type assertion
		if notFound, ok := err.(NotFoundError); ok {
			fmt.Printf("NotFoundError terdeteksi: %v\n", notFound)
			fmt.Printf("  Resource: %s\n", notFound.Resource)
			fmt.Printf("  ID: %s\n", notFound.ID)
		} else {
			fmt.Printf("Error lain: %v\n", err)
		}
	} else {
		fmt.Printf("User ditemukan: %s\n", name)
	}

	// ValidationError
	err = ValidateAge(-5)
	if err != nil {
		if validationErr, ok := err.(ValidationError); ok {
			fmt.Printf("ValidationError: %v\n", validationErr)
			fmt.Printf("  Field: %s\n", validationErr.Field)
		}
	}

	// =============================================================================
	// 4. ERROR WRAPPING DAN UNWRAPPING
	// =============================================================================
	// errors.Is dan errors.As untuk memeriksa error chain
	fmt.Println("\n--- 4. Error Wrapping ---")

	_, err = GetUserFromDB("001")
	if err != nil {
		fmt.Printf("Wrapped error: %v\n", err)

		// Unwrap untuk mendapatkan error asli
		if dbErr, ok := err.(DatabaseError); ok {
			fmt.Printf("Operasi yang gagal: %s\n", dbErr.Op)
			fmt.Printf("Error asli: %v\n", dbErr.Unwrap())
		}
	}

	// =============================================================================
	// 5. DEFER (DEFERRED EXECUTION)
	// =============================================================================
	// defer menunda eksekusi hingga fungsi selesai
	// Berguna untuk cleanup: close file, close database, unlock mutex, dll
	fmt.Println("\n--- 5. defer Statement ---")

	err = ProcessFile("data.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\nPanggil ProcessFile dengan error:")
	err = ProcessFile("")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// Perhatikan: defer tetap dieksekusi meski ada error!

	// =============================================================================
	// 6. PANIC
	// =============================================================================
	// panic menghentikan program secara abnormal
	// Gunakan hanya untuk kondisi yang benar-benar fatal!
	fmt.Println("\n--- 6. panic ---")

	fmt.Println("Contoh panic (akan di-recover):")
	result, err = SafeDivide(10, 0)
	if err != nil {
		fmt.Printf("Recovered from panic: %v\n", err)
	} else {
		fmt.Printf("Hasil: %d\n", result)
	}

	// =============================================================================
	// 7. MULTIPLE ERROR CHECKING
	// =============================================================================
	// Pattern untuk beberapa operasi yang masing-masing bisa error
	fmt.Println("\n--- 7. Multiple Error Checking ---")

	// Cara 1: Sequential dengan check tiap step
	func() {
		result, err := Divide(100, 5)
		if err != nil {
			fmt.Printf("Step 1 error: %v\n", err)
			return
		}

		sqrt, err := Sqrt(float64(result))
		if err != nil {
			fmt.Printf("Step 2 error: %v\n", err)
			return
		}

		fmt.Printf("Chain result: %f\n", sqrt)
	}()

	// =============================================================================
	// 8. BEST PRACTICES
	// =============================================================================
	fmt.Println("\n--- 8. Best Practices Error Handling ---")

	fmt.Println("\n✅ DO:")
	fmt.Println("   - Selalu check error: if err != nil { ... }")
	fmt.Println("   - Return error sebagai value terakhir: (result, error)")
	fmt.Println("   - Gunakan custom error untuk konteks spesifik")
	fmt.Println("   - Wrap error dengan konteks tambahan")
	fmt.Println("   - Gunakan defer untuk cleanup resources")

	fmt.Println("\n❌ DON'T:")
	fmt.Println("   - Abaikan error dengan _ (kecuali memang sengaja)")
	fmt.Println("   - Gunakan panic untuk error yang bisa ditangani")
	fmt.Println("   - Return nil untuk error jika sukses")
	fmt.Println("   - Buat error message yang terlalu generic")

	// =============================================================================
	// 9. ERROR CHECKING PATTERN
	// =============================================================================
	fmt.Println("\n--- 9. Error Checking Pattern ---")

	// Pattern 1: Check and return
	checkAndReturn := func(x int) (int, error) {
		if x < 0 {
			return 0, errors.New("x harus positif")
		}
		return x * 2, nil
	}

	// Pattern 2: Check and handle
	value := -5
	result, err = checkAndReturn(value)
	if err != nil {
		// Handle dengan default value
		fmt.Printf("Error, menggunakan default: %v\n", err)
		result = 0
	}
	fmt.Printf("Result: %d\n", result)

	// Pattern 3: Check and wrap (Go 1.13+)
	wrapError := func() error {
		_, err := Divide(10, 0)
		if err != nil {
			// Wrap error dengan konteks tambahan
			return fmt.Errorf("operasi matematika gagal: %w", err)
		}
		return nil
	}

	if err := wrapError(); err != nil {
		fmt.Printf("Wrapped error: %v\n", err)
	}

	fmt.Println("\n================================================================================")
	fmt.Println("SELESAI - Error handling di Go: explicit, simple, dan full control")
	fmt.Println("================================================================================")
}
