package main

import "fmt"

/*
- Package time berisi fungsi-fungsi untuk mengatur waktu dan tanggal.
- Package time juga berisi fungsi-fungsi untuk mengkonversi string ke time dan sebaliknya.
- Package time juga berisi fungsi-fungsi untuk mengatur zona waktu, durasi waktu, timer, ticker, sleep, dsb.
*/

func main() {
	// Time Usage
	fmt.Println(">> Time Usage <<")
	timeUsage()
	fmt.Println()

	// Parsing String to Time
	fmt.Println(">> Parsing String to Time <<")
	parsingStringToTime()
	fmt.Println()

	// Predefined Layout Time
	fmt.Println(">> Predefined Layout Time <<")
	predefinedLayout()
	fmt.Println()

	// Format Time to String
	fmt.Println(">> Format Time to String <<")
	formatTimeToString()
	fmt.Println()

	// Handle Parsing Error
	fmt.Println(">> Handle Parsing Error <<")
	handleParsingError()
	fmt.Println()
}
