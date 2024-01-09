package main

import "fmt"

func main() {
	/* Array
	Array di golang adalah kumpulan data bertipe sama dan bersifat fixed length
	*/
	{
		// var names = [...]string{"Monkey", "D", "Luffy"}
		// var names = [3]string{"Monkey", "D", "Luffy"}
		// var names = [3]string{
		// 	"Monkey",
		// 	"D",
		// 	"Luffy",
		// }
		// var names = make([]string, 3)
		var names [3]string
		names[0] = "Monkey"
		names[1] = "D"
		names[2] = "Luffy"
		fmt.Println(names)
	}

	/* Two dimension array
	Array 2 dimensi adalah array yang memiliki elemen array di dalamnya
	*/
	{
		// var numbers = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
		var numbers = [2][3]int{{3, 2, 3}, {3, 4, 5}}
		fmt.Println(numbers)
	}

	// For usage
	{
		var fruits = [4]string{"apple", "grape", "banana", "melon"}

		// for i := 0; i < len(fruits); i++ {
		// 	fmt.Println(fruits[i])
		// }
		for _, fruit := range fruits {
			fmt.Println(fruit)
		}
	}
}
