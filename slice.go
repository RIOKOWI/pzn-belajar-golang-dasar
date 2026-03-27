package main

import	"fmt"

func main(){
	names := [...]string{
		"udin",
		"dudun",
		"abel",
		"yogo",
		"adly",
		"dono",
		"zee",
		"antonio",
		"faris",
		"coki",
	}

	slice := names[4:6]
	// pointer = index 4
	// length = 2
	// capacity = 6

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])


	slice1 := names[:3]

	fmt.Println(slice1)


	slice2 := names[3:]

	fmt.Println(slice2)

	slice3 := names[:]

	fmt.Println(slice3)

	fmt.Println("")
	fmt.Println("FUNCTION SLICE")
	fmt.Println("")

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)
	// [Sabtu Minggu]

	daysSlice1[0] = "Sabtu Cuti"
	daysSlice1[1] = "Minggu Cuti"

	fmt.Println(days)
	//[Senin Selasa Rabu Kamis Jumat Sabtu Cuti Minggu Cuti]
	
	daysSlice2 := append(daysSlice1, "Ruka")
	daysSlice2[0] = "Adoooh"
	fmt.Println(daysSlice2)
	// [Adoooh Minggu Cuti Ruka]

	fmt.Println(days)
	// [Senin Selasa Rabu Kamis Jumat Sabtu Cuti Minggu Cuti]



	// make

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rio"
	newSlice[1] = "Achyar"
	

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSlice1 := append(newSlice, "Ganteng")
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	fromSlice := days
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	
	copy(toSlice, fromSlice[:])

	fmt.Print(fromSlice)
	fmt.Print(toSlice)



	// perbedaan array dan slice
	iniArray := [...]int{1,2,3}
	iniSlice := []int{1,2,3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}