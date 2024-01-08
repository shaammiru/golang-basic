package main

import "fmt"

func main() {
	// /* For Loop
	// For loop di go hanya memiliki satu keyword yaitu for
	// */
	// {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(i)
	// 	}
	// }

	// /* For Range
	// For range digunakan untuk mengiterasi array, slice, map, string, dan channel
	// */
	// {
	// 	names := []string{"John", "Wick", "Ethan", "Hunt"}

	// 	for index, name := range names {
	// 		fmt.Println("Index:", index, "Name:", name)
	// 	}
	// }

	// /* While Loop
	// While loop di go tidak memiliki keyword while,
	// untuk membuat while loop gunakan keyword for dan kondisi
	// */
	// {
	// 	i := 0

	// 	for i < 3 {
	// 		fmt.Println(i)
	// 		i++
	// 	}
	// }

	// // For Loop Tanpa Kondisi
	// {
	// 	j := 0

	// 	for {
	// 		if j == 1 {
	// 			j++
	// 			continue
	// 		}

	// 		fmt.Println(j)
	// 		j++

	// 		if j == 3 {
	// 			break
	// 		}
	// 	}
	// }

	// Label
	{
	outerLoop:
		for i := 0; i < 5; i++ {
		innerLoop:
			for j := 0; j < 5; j++ {
				if i == 3 {
					break outerLoop
				}

				if i == 1 {
					break innerLoop
				}

				fmt.Println("Matriks [", i, "][", j, "]")
			}
		}
	}
}
