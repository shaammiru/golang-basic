package main

import (
	"fmt"
	"strings"
)

/* Interface
- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung.
- Sebuah interface berisikan definisi-definisi method.
- Biasanya interface digunakan sebagai kontrak.
- Sebuah tipe data bisa mengimplementasi banyak interface.
- Interface juga bisa disebut sebagai tipe data polimorfik.
- Nilai kosong dari interface adalah nil.
*/

// Interface Declaration
type calculate interface {
	area() float32
	perimeter() float32
}

func main() {
	// var c calculate = rectangle{10, 5}
	var c calculate

	c = rectangle{10, 5}
	fmt.Println("Rectangle")
	fmt.Println("Area      :", c.area())
	fmt.Println("Perimeter :", c.perimeter())

	c = square{10.5}
	fmt.Println("Square")
	fmt.Println("Area      :", c.area())
	fmt.Println("Perimeter :", c.perimeter())

	// Empty Interface
	// Empty interface tidak memiliki deklarasi method satupun.
	// Sehingga semua tipe data bisa disimpan sebagai nilai dari empty interface.
	fmt.Println("Empty Interface")
	// var secret any
	var secret interface{}

	secret = 2
	numVar := secret.(int) + 10
	fmt.Println("Int    :", secret)
	fmt.Println("numVar :", numVar)

	secret = "John Wick"
	strVar := secret.(string) + " Wick"
	fmt.Println("String :", secret)
	fmt.Println("strVar :", strVar)

	secret = []string{"apple", "manggo", "banana"}
	sliceVar := strings.Join(secret.([]string), ", ")
	fmt.Println("Slice    :", secret)
	fmt.Println("sliceVar :", sliceVar)
}
