package main

/* Variadic Function
Variadic function adalah function yang bisa menerima banyak parameter.
Parameter yang diterima bisa nol atau lebih.
*/

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}