package main

import (
	"fmt"
	"regexp"
)

/* Regular Expression
- Package regex menyediakan fungsi untuk memanipulasi string menggunakan regular expression.
- Regular expression adalah sebuah string yang digunakan untuk mencocokkan teks berdasarkan pola tertentu.
- Regular expression sangat berguna untuk mencari, mengganti, atau memvalidasi teks.
- Go mengadopsi standar regex RE2 yang dikembangkan oleh Google.
*/

func main() {
	// Regexp Usage
	{
		fmt.Println(">> Regexp Usage <<")

		text := "The quick brown fox jumps over the lazy dog"
		regex, err := regexp.Compile(`[a-zA-Z]+`)

		if err != nil {
			fmt.Println(err.Error())
		}

		res1 := regex.FindAllString(text, 2)
		res2 := regex.FindAllString(text, -1)
		fmt.Println(res1)
		fmt.Println(res2)

		fmt.Println()
	}

	// Match String
	{
		fmt.Println(">> Match String <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[a-z]+`)
		isMatch := regex.MatchString(text)
		fmt.Println(isMatch)

		fmt.Println()
	}

	// Find String
	{
		fmt.Println(">> Find String <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[a-z]+`)
		res := regex.FindString(text)
		fmt.Println(res)

		fmt.Println()
	}

	// Find String Index
	{
		fmt.Println(">> Find String Index <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[a-z]+`)
		res := regex.FindStringIndex(text)
		fmt.Println(res)

		fmt.Println()
	}

	// Find All String
	{
		fmt.Println(">> Find All String <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[a-z]+`)
		res := regex.FindAllString(text, -1)
		fmt.Println(res)

		fmt.Println()
	}

	// Replace All String
	{
		fmt.Println(">> Replace All String <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[a-z]+`)
		res := regex.ReplaceAllString(text, "(car brand)")
		fmt.Println(res)

		fmt.Println()
	}

	// Replace All String Func
	{
		fmt.Println(">> Replace All String Func <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[a-z]+`)
		res := regex.ReplaceAllStringFunc(text, func(s string) string {
			if s == "nissan" {
				return "(car brand)"
			}
			return s
		})
		fmt.Println(res)

		fmt.Println()
	}

	// Split String
	{
		fmt.Println(">> Split String <<")

		text := "toyota nissan honda"
		regex, _ := regexp.Compile(`[ao]+`)
		res := regex.Split(text, -1)
		fmt.Println(res)

		fmt.Println()
	}
}
