package main

import "fmt"

func main() {
	// Array
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)

	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	// Slice
	var z []float64

	// Create slice x with length 5 (max 10) using built-in function make
	s := make([]float64, 5, 10)
	sl := [5]float64{1, 2, 3, 4, 5}
	xs := sl[3:5] // print [4 5]

	fmt.Println(z, s)
	fmt.Println(sl, xs)

	// Slice Append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// Slice Copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	fmt.Println(slice3, slice4)
	copy(slice3, slice4)
	fmt.Println(slice3, slice4)

	// Map
	m1 := make(map[string]int)
	m1["key"] = 10
	fmt.Println(m1["key"])

	m2 := make(map[int]string)
	m2[10] = "key"
	fmt.Println(m2[10])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"]) // Prints string of empty strings

	// Validating if elements["Un"] exist or not
	name, ok := elements["Un"]
	fmt.Println(name, ok) // Prints false

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Un does not exist")
	}

	// Short map declaration
	elements2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	if name, ok := elements2["C"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Does not exist")
	}
}
