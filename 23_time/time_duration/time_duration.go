package time_duration

import (
	"fmt"
	"time"
)

/* Time Duration
- Package time juga memiliki tipe data Duration yang digunakan untuk mengatur durasi waktu
- Tipe data Duration memiliki konstanta seperti Nanosecond, Microsecond, Millisecond, Second, Minute, Hour
*/

func init() {
	// Duration Usage
	fmt.Println(">> Duration Usage <<")
	durationUsage()
	fmt.Println()

	// Calculate Duration
	fmt.Println(">> Calculate Duration <<")
	calculateDuration()
	fmt.Println()

	// Duration Conversion
	fmt.Println(">> Duration Conversion <<")
	durationConversion()
	fmt.Println()
}

func durationUsage() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	duration := time.Since(start)

	fmt.Println("Time elapsed in second:", duration.Seconds())
	fmt.Println("Time elapsed in milisecond:", duration.Milliseconds())
	fmt.Println("Time elapsed in microsecond:", duration.Microseconds())
}

func calculateDuration() {
	time1 := time.Now()
	time.Sleep(2 * time.Second)
	time2 := time.Now()

	duration := time2.Sub(time1)

	fmt.Println("Time elapsed in second:", duration.Seconds())
	fmt.Println("Time elapsed in milisecond:", duration.Milliseconds())
	fmt.Println("Time elapsed in microsecond:", duration.Microseconds())
}

func durationConversion() {
	fmt.Println(">> Duration Conversion <<")
	fmt.Println("Duration(2) in second:", time.Second*2)
	fmt.Println("Duration(2) in minute:", time.Minute*2)
	fmt.Println("Duration(2) in hour:", time.Hour*2)
}
