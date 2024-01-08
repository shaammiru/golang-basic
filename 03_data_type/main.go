package main

import "fmt"

func main() {
	/* Numeric non-decimal
		int8, int16, int32, int64, uint8, uint16, uint32, uint64
		int(either int32 or int64), uint(either uint32 or uint64)
		byte(int8), rune(int32)
		
		int8: -128 to 127
		int16: -32768 to 32767
		int32: -2147483648 to 2147483647
		int64: -9223372036854775808 to 9223372036854775807
		uint8: 0 to 255
		uint16: 0 to 65535
		uint32: 0 to 4294967295
		uint64: 0 to 18446744073709551615
	*/
	var positiveNumber uint8 = 69
	var negativeNumber int8 = -69
	fmt.Println("Positive:", positiveNumber, "| Negative:", negativeNumber)

	/* Numeric decimal
		float32, float64
	*/
	var decimalNumber float32 = 3.14
	fmt.Println("Decimal:", decimalNumber)

	/* Boolean
		true, false
	*/
	var boolean bool = true
	fmt.Println("Boolean:", boolean)

	/* String
		"Hello World"
	*/
	var stringText string = "Hello World"
	// Grave accent (`)
	var stringText2 string = `Hello \n "John"`
	fmt.Println("String:", stringText, "| String2:", stringText2)

	/* Nil & Zero value
		Nil berarti kosong (benar-benar kosong),
		sedangkan zero value berarti nilai awal dari suatu tipe data.
		numeric = 0 / 0.0
		boolean = false
		string = ""
	*/
}