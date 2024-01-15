package main

import "fmt"

func main() {
	/* Slice
	Slice singkatnya adalah reference elemen dari data array.
	Slice bersifat dinamis dan panjangnya bisa berubah-ubah.
	*/
	{
		var names = []string{"Monkey", "D", "Luffy"}
		/* Slice operator
		[2:3] = ["Luffy"]
		[:] = ["Monkey", "D", "Luffy"]
		[1:] = ["D", "Luffy"]
		[:2] = ["Monkey", "D"]
		*/
		var nickname = names[2:3]
		var nickname2 = names[2:]
		names[2] = "Dragon" // Slice akan berubah jika array berubah
		fmt.Println(names, nickname, nickname2)

		nickname[0] = "Luffy" // Array juga akan berubah jika slice berubah
		fmt.Println(names, nickname, nickname2)

		// Basic function for slice
		fruits := []string{"apple", "grape", "banana"}
		aFruits := fruits[0:2] // [apple, grape, ...] (cap = 3)
		bFruits := fruits[1:3] // ..., [grape, banana] (cap = 2)

		// len() cap()
		fmt.Println(len(aFruits), cap(aFruits))
		fmt.Println(len(bFruits), cap(bFruits))

		// append()
		cFruits := append(fruits, "papaya")
		fmt.Println(cFruits)

		// copy()
		dFruits := make([]string, 2)
		copy(dFruits, fruits)
		fmt.Println(dFruits)

		// 3 Index 
		eFruits := fruits[0:2]
		fFruits := fruits[0:2:2] // Last :2 is size of slice

		fmt.Println(len(eFruits), cap(eFruits))
		fmt.Println(len(fFruits), cap(fFruits))
	}
}