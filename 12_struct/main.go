package main

import "fmt"

func main() {
	/* Struct
	Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus dengan nama tertentu.
	Struct biasanya digunakan sebagai data record.
	Sebuah struct bisa menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan.
	Struct dalam Go mirip seperti class dalam konsep pemrograman lainnya, namun bedanya di Go tidak ada konsep pewarisan.
	*/

	// Struct Declaration
	type car struct {
		brand string
		model string
	}

	// Struct Usage
	{
		fmt.Println(">> Struct Usage <<")

		car1 := car{
			brand: "Porsche",
			model: "911 GT3 RS",
		}

		car2 := car{"Mercedes", "AMG GT63 S"}

		fmt.Println(car1.brand, "-", car1.model)
		fmt.Println(car2.brand, "-", car2.model)
		fmt.Println()
	}

	// Pointer Struct
	{
		fmt.Println(">> Pointer Struct <<")

		car1 := car{
			brand: "Porsche",
			model: "911 GT3 RS",
		}

		car2 := &car1

		fmt.Println(car1.brand, "-", car1.model)
		fmt.Println(car2.brand, "-", car2.model)

		car2.brand = "Mercedes"
		car2.model = "AMG GT63 S"

		fmt.Println(car1.brand, "-", car1.model)
		fmt.Println(car2.brand, "-", car2.model)
		fmt.Println()
	}

	// Embedded Struct
	{
		fmt.Println(">> Embedded Struct <<")

		type company struct {
			name string
			car
		}

		company1 := company{
			name: "Volkswagen AG",
			car: car{
				brand: "Porsche",
				model: "911 GT3 RS",
			},
		}

		fmt.Println(company1.name)
		fmt.Println(company1.brand, "-", company1.model)
		fmt.Println()
	}

	// Anonymous Struct
	{
		fmt.Println(">> Anonymous Struct <<")

		car1 := struct {
			brand string
			model string
		}{}
		// {
		// 	brand: "Toyota",
		// 	model: "Supra MK4",
		// }

		// Or use var
		// var car1 struct {
		// 	brand string
		// 	model string
		// }

		car1.brand = "Toyota"
		car1.model = "Supra MK4"

		fmt.Println(car1.brand, "-", car1.model)
		fmt.Println()
	}

	// Slice Struct
	{
		fmt.Println(">> Slice Struct <<")

		car1 := []car{
			{brand: "Porsche", model: "911 GT3 RS"},
			{brand: "Mercedes", model: "AMG GT63 S"},
			{brand: "Toyota", model: "Supra MK4"},
		}

		for _, car := range car1 {
			fmt.Println(car.brand, "-", car.model)
		}

		fmt.Println()
	}

	// Nested Struct
	{
		fmt.Println(">> Nested Struct <<")

		type engine struct {
			name  string
			specs struct {
				horsepower int
				torque     int
			}
		}

		engine1 := engine{
			name: "M139",
			specs: struct {
				horsepower int
				torque     int
			}{
				horsepower: 639,
				torque:     900,
			},
		}

		fmt.Println("Engine name:", engine1.name)
		fmt.Println("Horsepower:", engine1.specs.horsepower, "HP")
		fmt.Println("Torque:", engine1.specs.torque, "Nm")
	}
}
