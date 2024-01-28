package main

import (
	"14_property_modifier/lib"
	"fmt"
)

/* Property Modifier
Property modifier digunakan untuk mengubah hak akses property.
Di Go, hak akses property ditentukan dari huruf pertama nama property.
Jika huruf pertama nama property adalah huruf besar, maka property tersebut bisa diakses dari luar package (exported).
Jika huruf pertama nama property adalah huruf kecil, maka property tersebut hanya bisa diakses dari dalam package (unexported).
*/

func main() {
	// Membuat instance dari struct Car
	car := lib.NewCar("Porsche", "Cayman", 2015)

	// Property Brand dan Model bisa diakses karena diset public
	fmt.Println(car.Brand, car.Model)

	// Property year tidak bisa diakses karena diset private
	// fmt.Println(car.year) // Error

	// Method PrintCarDetail bisa diakses karena diset public
	// Method ini bisa mengakses property/method private karena ada dalam satu package
	car.PrintCarDetail()

	// Method setCarYear tidak bisa diakses karena diset private
	// car.setCarYear(2016) // Error
	car.ChangeCarYear(2016)
	car.PrintCarDetail()
}
