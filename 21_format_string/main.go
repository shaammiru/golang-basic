package main

import "fmt"

/* Format String
- Format string adalah cara untuk memformat string yang akan ditampilkan ke layar
- Layout format string digunakan dalam konversi tipe data ke string
*/

func main() {
	// Format %b - Number to Binary
	fmt.Printf("Number to Binary: (10) %b\n", 10)

	// Format %c - Number to Character
	fmt.Printf("Number to Character: (65) %c\n", 65)

	// Format %d - Number to Decimal
	fmt.Printf("Number to Decimal: (10) %d\n", 10)

	// Format %e / %E - Number to Decimal (Float) with Exponent
	fmt.Printf("Number to Decimal (Float) with Exponent: (3.14) %e\n", 3.14)

	// Format %f - Number to Decimal (Float)
	fmt.Printf("Number to Decimal (Float): (3.14) %f\n", 3.14)

	// Format %g / %G - Number to Decimal (Float) without Trailing Zero
	fmt.Printf("Number to Decimal (Float) without Trailing Zero: (3.14) %g\n", 3.14)

	// Format %o - Number to Octal
	fmt.Printf("Number to Octal: (10) %o\n", 10)

	// Format %p - Pointer
	var a int = 10
	fmt.Printf("Pointer: (&a) %p\n", &a)

	// Format %q - Quoted String
	fmt.Printf("Quoted String: (Hello) %q\n", "Hello")

	// Format %s - String
	fmt.Printf("String: (Hello) %s\n", "Hello")

	// Format %t - Boolean
	fmt.Printf("Boolean: (true) %t\n", true)

	// Format %T - Type
	fmt.Printf("Type: (10) %T\n", 10)

	// Format %U - Unicode
	fmt.Printf("Unicode: (65) %U\n", 65)

	// Format %x / %X - Number to Hexadecimal
	fmt.Printf("Number to Hexadecimal: (10) %x\n", 10)

	// Format %v - Default Format
	fmt.Printf("Default Format: (10) %v\n", 10)
}
