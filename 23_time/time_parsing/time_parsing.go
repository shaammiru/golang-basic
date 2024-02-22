package time_parsing

import (
	"fmt"
	"time"
)

/*
- Package time berisi fungsi-fungsi untuk mengatur waktu dan tanggal.
- Package time juga berisi fungsi-fungsi untuk mengkonversi string ke time dan sebaliknya.
- Package time juga berisi fungsi-fungsi untuk mengatur zona waktu, durasi waktu, timer, ticker, sleep, dsb.
*/

func init() {
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

func timeUsage() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	definedTime := time.Date(2024, 2, 21, 20, 47, 0, 0, time.UTC)
	fmt.Println("Defined Time:", definedTime)
}

/* time.Now() method list
- Year() - Tahun(int)
- YearDay() - Hari dalam tahun(int)
- Month() - Bulan(int)
- Weekday() - Hari dalam minggu(time.Weekday)
- ISOWeek() - Tahun, minggu, hari(int, int, int)
- Day() - Hari(int)
- Hour() - Jam(int)
- Minute() - Menit(int)
- Second() - Detik(int)
- Nanosecond() - Nanosekon(int)
- Local() - Waktu lokal(time.Time)
- Location() - Lokasi(time.Location)
- Zone() - Zona waktu(string)
- IsZero() - Apakah waktu nol(bool)
- UTC() - Waktu UTC(time.Time)
- Unix() - Detik sejak 1 Januari 1970(int64)
- UnixNano() - Nanodetik sejak 1 Januari 1970(int64)
- String() - Waktu dalam string(string)
*/

func parsingStringToTime() {
	format := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	date, _ := time.Parse(format, value)
	fmt.Println("Value:", value, "=> Time:", date)
}

/* Layout Time
- 2006 = Tahun 4 digit
- 006 = Tahun 3 digit
- 06 = Tahun 2 digit
- 01 = Bulan 2 digit
- 1 = Bulan 1 digit jika di bawah 10
- January = Nama bulan
- Jan = Nama bulan 3 huruf
- 02 = Tanggal 2 digit
- 2 = Tanggal 1 digit jika di bawah 10
- Monday = Nama hari
- Mon = Nama hari 3 huruf
- 15 = Jam 24 jam
- 03 = Jam 12 jam 2 digit
- 3 = Jam 12 jam 1 digit jika di bawah 10
- PM = AM/PM
- 04 = Menit 2 digit
- 4 = Menit 1 digit jika di bawah 10
- 05 = Detik 2 digit
- 5 = Detik 1 digit jika di bawah 10
- 999999 = Nanosekon
- MST = Zona waktu
- Z0700 = Offset zona waktu
*/

func predefinedLayout() {
	date, _ := time.Parse(time.RFC822, "02 Jan 24 15:04 WIB")
	fmt.Println("Value:", "02 Jan 06 15:04 WIB", "=> Time:", date)
}

func formatTimeToString() {
	date, _ := time.Parse(time.RFC822, "02 Jan 24 15:04 WIB")
	dateS1 := date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("DateS1:", dateS1)

	dateS2 := date.Format(time.RFC3339)
	fmt.Println("DateS2:", dateS2)
}

func handleParsingError() {
	date, err := time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(date)
}
