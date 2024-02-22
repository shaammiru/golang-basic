package main

import (
	"fmt"
	"strings"
)

/* String Functions
- Package strings menyediakan banyak fungsi untuk memanipulasi string.
*/

func main() {
	// strings.Contains()
	{
		fmt.Println(">> string.Contains() <<")

		// Mengecek apakah string mengandung kata 'Supra' atau 'Skyline'
		result := strings.Contains("Toyota Supra", "Supra")
		fmt.Println(result)

		fmt.Println()
	}

	// strings.Count()
	{
		fmt.Println(">> string.Count() <<")

		// Menghitung jumlah huruf 'a' pada string
		count := strings.Count("Toyota Supra", "a")
		fmt.Println(count)

		fmt.Println()
	}

	// strings.HasPrefix() & strings.HasSuffix()
	{
		fmt.Println(">> string.HasPrefix() & string.HasSuffix() <<")

		// Mengecek apakah string diawali atau diakhiri dengan kata 'Toyota'
		prefix := strings.HasPrefix("Toyota Supra", "Toyota")
		suffix := strings.HasSuffix("Toyota Supra", "Supra")
		fmt.Println(prefix, suffix)

		fmt.Println()
	}

	// strings.Index() & strings.LastIndex()
	{
		fmt.Println(">> string.Index() & string.LastIndex() <<")

		// Mencari index pertama dan terakhir dari kata 'a' pada string
		index := strings.Index("Toyota Supra", "a")
		lastIndex := strings.LastIndex("Toyota Supra", "a")
		fmt.Println(index, lastIndex)

		fmt.Println()
	}

	// strings.Replace()
	{
		fmt.Println(">> string.Replace() <<")

		// Mengganti kata 'Supra' dengan 'AE86' pada string
		result := strings.Replace("Toyota Supra", "Supra", "AE86", -1)
		fmt.Println(result)

		fmt.Println()
	}

	// strings.Repeat()
	{
		fmt.Println(">> string.Repeat() <<")

		// Mengulang string sebanyak 3 kali
		result := strings.Repeat("Toyota", 3)
		fmt.Println(result)

		fmt.Println()
	}

	// strings.Split()
	{
		fmt.Println(">> string.Split() <<")

		// Memisahkan string berdasarkan spasi (bisa juga berdasarkan karakter lain)
		result := strings.Split("Toyota Supra", " ")
		fmt.Println(result)

		fmt.Println()
	}

	// strings.Join()
	{
		fmt.Println(">> string.Join() <<")

		// Menggabungkan string dengan '-' sebagai pemisah
		cars := []string{"Nissan", "Toyota", "Mitsubishi"}
		result := strings.Join(cars, "-")
		fmt.Println(result)

		fmt.Println()
	}

	// strings.ToLower() & strings.ToUpper()
	{
		fmt.Println(">> string.ToLower() & string.ToUpper() <<")

		// Mengubah string menjadi huruf kecil dan besar
		lower := strings.ToLower("Toyota Supra")
		upper := strings.ToUpper("Toyota Supra")
		fmt.Println(lower, upper)

		fmt.Println()
	}
}
