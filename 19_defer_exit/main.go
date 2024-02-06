package main

import "fmt"

/* Defer
- Defer digunakan untuk menunda eksekusi sebuah statement sampai blok fungsi selesai dieksekusi
- Defer akan selalu dieksekusi walaupun terjadi error di blok fungsi
*/

/* Exit
- Exit digunakan untuk menghentikan program
- Exit akan menghentikan program dan tidak akan mengeksekusi defer setelahnya
*/

func main() {
	// Defer Example
	fmt.Println(">> Defer Example <<")
	deferExample()

	fmt.Println()

	// IIFE Defer Example
	fmt.Println(">> IIFE Defer Example <<")
	iifeDeferExample()

	fmt.Println()

	// Exit Example
	fmt.Println(">> Exit Example <<")
	exitExample()

	fmt.Println("This line also will not be executed")
}
