package main

import "fmt"

func main() {
	/* Pointer
	Pointer adalah sebuah variabel yang menyimpan alamat memori dari sebuah variabel lain.
	Pointer digunakan untuk mengakses variabel lain secara tidak langsung.
	*/
	{
		numberA := 4
		numberB := &numberA
		fmt.Println("numberA (value) :", numberA)
		fmt.Println("numberB (value) :", *numberB)
	}

	// Parameter Pointer
	{
		number := 4
		fmt.Println("before :", number)
		changeNumber(&number, 10)
		fmt.Println("after  :", number)
	}
}

// Parameter Pointer
func changeNumber(origin *int, value int) {
	*origin = value
}
