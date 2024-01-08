package main

import "fmt"

func main() {
	point := 10

	// If, Else If, Else
	{
		if point > 10 {
			fmt.Println("Point is greater than 10")
		} else if point == 10 {
			fmt.Println("Point is equal to 10")
		} else {
			fmt.Println("Point is less than 10")
		}
	}

	// Temporary Variable
	{
		if point := 10; point > 10 {
			fmt.Println("Point is greater than 10")
		} else {
			fmt.Println("Point is less than 10")
		}
	}

	/* Switch Case
	Switch case di go tidak memerlukan keyword break,
	untuk melanjutkan ke case selanjutnya gunakan keyword fallthrough
	*/
	{
		switch point {
		case 10:
			fmt.Println("Point is equal to 10")
		case 20, 30:
			fmt.Println("Point is equal to 20 or 30")
			fallthrough
		default:
			fmt.Println("Point is not equal to 10 or 20")
		}
	}
}
