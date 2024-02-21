package main

import "fmt"

/* Random
- Package math/rand adalah package yang berisi fungsi-fungsi untuk menghasilkan angka/data acak.
*/

func main() {
	// Random Usage
	fmt.Println(">> Random Usage <<")
	randomUsage()
	fmt.Println()

	// Unique Seed
	fmt.Println(">> Unique Seed <<")
	uniqueSeed()
	fmt.Println()

	// Random Index
	fmt.Println(">> Random Index <<")
	randomIndex()
	fmt.Println()

	// Random String
	fmt.Println(">> Random String <<")
	fmt.Println(randomString(10))
	fmt.Println(randomString(10))
	fmt.Println()
}
