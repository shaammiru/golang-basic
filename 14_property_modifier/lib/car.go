package lib

import "fmt"

// Init function
// Fungsi ini akan dipanggil sebelum fungsi main dipanggil
// Atau ketika package ini diimport
func init() {
	fmt.Println("Library initialized")
}

type Car struct {
	// Public Field
	Brand, Model string

	// Private Field
	year int
}

// Public Method
func NewCar(brand, model string, year int) Car {
	return Car{brand, model, year}
}

// Dalam pacakage ini method public bisa mengakses property/method private
func (c *Car) PrintCarDetail() {
	// Akses property private c.year
	fmt.Printf("Car: %s %s (%d)\n", c.Brand, c.Model, c.year)
}

func (c *Car) ChangeCarYear(newYear int) {
	c.setCarYear(newYear)
}

// Private Method
func (c *Car) setCarYear(year int) {
	c.year = year
}
