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
	age := 24
	reflectValue := reflect.ValueOf(age)
	reflectType := reflect.TypeOf(age)

	fmt.Println(reflectValue)
	fmt.Println(reflectType)	
}
