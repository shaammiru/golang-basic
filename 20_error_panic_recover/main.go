package main

import (
	"fmt"
)

/* Error, Panic, Recover
- Error digunakan untuk menangani kesalahan yang terjadi saat program berjalan
- Panic digunakan untuk menghentikan program jika terjadi kesalahan yang tidak bisa diatasi
- Recover digunakan untuk menangkap data panic
*/

func main() {
	// Error Usage
	fmt.Println(">> Error Usage <<")
	errorUsage()
	fmt.Println()

	// Custom Error
	fmt.Println(">> Custom Error <<")
	customError()
	fmt.Println()

	// Panic Recover
	fmt.Println(">> Panic Recover <<")
	panicRecover()
	fmt.Println()
}
