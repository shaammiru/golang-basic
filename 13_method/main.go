package main

import "fmt"

/* Method
Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya).
Method bisa diakses lewat variabel objek.
*/

type car struct {
	brand, model string
}

// Method Declaration
func (c car) startEngine() {
	fmt.Println(c.brand, "- Engine started")
}

/* Method Pointer Declaration
Method pointer digunakan untuk mengubah nilai variabel objek.
Jika tidak menggunakan pointer, maka nilai variabel objek tidak akan berubah.
Hanya akan berubah di dalam method saja.
*/
func (c *car) changeModel(newModel string) {
	c.model = newModel
}

func main() {
	// Method Usage
	{
		fmt.Println(">> Method Declaration <<")

		car1 := car{"Porsche", "911 GT3 RS"}
		car1.startEngine()
		fmt.Println()
	}

	// Method Pointer Usage
	{
		fmt.Println(">> Method Pointer Declaration <<")

		car1 := car{"Porsche", "911 GT3 RS"}
		fmt.Println("Before change:", car1.brand, car1.model)
		car1.changeModel("Taycan Turbo S")
		fmt.Println("After change:", car1.brand, car1.model)
	}
}
