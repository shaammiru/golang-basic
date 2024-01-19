package main

/* Multiple Return Function
Function bisa mengembalikan lebih dari satu nilai.
*/

// Multiple return function
func getFullName() (string, string) {
	return "John", "Doe"
}

// Predifined return value
func getCityCountry() (city, country string) {
	city = "Jakarta"
	country = "Indonesia"
	return
}