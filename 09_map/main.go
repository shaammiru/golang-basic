package main

import "fmt"

func main() {
	/* Map
	Map adalah tipe data asosiatif yang berisi kumpulan pasangan key-value.
	Map mirip seperti array, hanya saja indexnya bisa ditentukan sendiri.
	*/
	{
		capitals := map[string]string{
			"Indonesia": "Jakarta",
			"Malaysia":  "Kuala Lumpur",
			"Japan":     "Tokyo",
		}

		// Get length
		fmt.Println(len(capitals))

		// Looping map
		for key, value := range capitals {
			fmt.Println(key, ":", value)
		}

		// Add new item
		capitals["South Korea"] = "Seoul"


		// Delete item
		delete(capitals, "Japan")

		// Get item
		japan, ok := capitals["Japan"]
		if ok {
			fmt.Println(japan)
		} else {
			fmt.Println("Japan not found")
		}

		// Get item with one line
		fmt.Println(capitals["South Korea"])

		// Get item with one line and check if item exist
		if japan, ok := capitals["Japan"]; ok {
			fmt.Println(japan)
		} else {
			fmt.Println("Japan not found")
		}
	}
}
