package main

import (
	"fmt"
	"strconv"
)

// Datatype Conversion

func main() {
	// String to Int
	{
		fmt.Println(">> String to Int <<")

		str := "123"
		num, err := strconv.Atoi(str)
		if err == nil {
			fmt.Println(num)
		}

		fmt.Println()
	}

	// Int to String
	{
		fmt.Println(">> Int to String <<")

		num := 123
		str := strconv.Itoa(num)
		fmt.Println(str)

		fmt.Println()
	}

	// String to Int (Base 2, 8, 16)
	{
		fmt.Println(">> String to Int (Base 2, 8, 16) <<")

		str := "1010"
		num, err := strconv.ParseInt(str, 2, 64)
		if err == nil {
			fmt.Println(num)
		}

		fmt.Println()
	}

	// Int to String (Base 2, 8, 16)
	{
		fmt.Println(">> Int to String (Base 2, 8, 16) <<")

		num := 10
		str := strconv.FormatInt(int64(num), 8)
		fmt.Println(str)

		fmt.Println()
	}

	// String to Float
	{
		fmt.Println(">> String to Float <<")

		str := "123.45"
		num, err := strconv.ParseFloat(str, 64)
		if err == nil {
			fmt.Println(num)
		}

		fmt.Println()
	}

	// Float to String
	{
		fmt.Println(">> Float to String <<")

		num := 123.45
		str := strconv.FormatFloat(num, 'f', 2, 64)
		fmt.Println(str)

		fmt.Println()
	}

	// String to Bool
	{
		fmt.Println(">> String to Bool <<")

		str := "true"
		b, err := strconv.ParseBool(str)
		if err == nil {
			fmt.Println(b)
		}

		fmt.Println()
	}

	// Bool to String
	{
		fmt.Println(">> Bool to String <<")

		b := true
		str := strconv.FormatBool(b)
		fmt.Println(str)

		fmt.Println()
	}

	// Convert with Casting
	{
		fmt.Println(">> Convert with Casting <<")

		num1 := float64(123)
		num2 := int(12.00)
		fmt.Println(num1)
		fmt.Println(num2)

		fmt.Println()
	}

	// Casting String <-> Byte
	{
		fmt.Println(">> Casting String <-> Byte <<")

		str := "Hello"
		strByte := []byte(str)
		fmt.Println(strByte)

		str = string(strByte)
		fmt.Println(str)

		fmt.Println(int64('H'))
		fmt.Println(string(rune(72)))

		fmt.Println()
	}

	// Type Assertion on Interface{}
	{
		fmt.Println(">> Type Assertion on Interface{} <<")

		car := map[string]interface{}{
			"brand": "Toyota",
			"year":  2020,
			"isNew": true,
		}

		fmt.Println(car["brand"].(string))
		fmt.Println(car["year"].(int))
		fmt.Println(car["isNew"].(bool))

		// Pastikan tipe data yang di-assert ada dan sesuai
		// Gunakan error handling jika tidak yakin, seperti ini:
		// if value, ok := car["brand"].(string); ok {
		// 	fmt.Println(value)
		// }

		fmt.Println()
	}
}
