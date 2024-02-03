package main

import (
	"fmt"
	"reflect"
)

/* Reflect
- Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
- Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.
*/

func main() {
	{
		fmt.Println(">> Reflect Usage <<")
		age := 24
		reflectValue := reflect.ValueOf(age)
		reflectType := reflect.TypeOf(age)

		fmt.Println(reflectValue)
		fmt.Println(reflectType)
		fmt.Println()
	}

	{
		fmt.Println(">> Reflect IsZero <<")
		var age int
		var year int = 2021
		reflectValue := reflect.ValueOf(age)
		reflectValue2 := reflect.ValueOf(year)

		fmt.Println(reflectValue.IsZero())
		fmt.Println(reflectValue2.IsZero())
		fmt.Println()
	}
}
