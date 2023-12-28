package main

import (
	"fmt"
	"slices"
)

func main() {

	var arr1 []string
	fmt.Printf("Type: %T Value: %v\n", arr1, arr1)
	fmt.Println("uninit:", arr1, arr1 == nil, len(arr1) == 0)

	var slc1 = make([]string, 3)
	fmt.Printf("Type: %T Value: %v\n", slc1, slc1)
	fmt.Println("emp:", slc1, "len:", len(slc1), "cap:", cap(slc1))

	slc1[0] = "a"
	slc1[1] = "b"
	slc1[2] = "c"
	fmt.Println("set:", slc1)
	fmt.Println("get:", slc1[2])

	fmt.Println("len:", len(slc1))

	slc1 = append(slc1, "d")
	slc1 = append(slc1, "e", "f")
	fmt.Println("Append:", slc1)

	c := make([]string, len(slc1))
	copy(c, slc1)
	fmt.Println("Copy:", c)

	l := slc1[2:5]
	fmt.Println("Slice 1:", l)

	l = slc1[:5]
	fmt.Println("Slice 2:", l)

	l = slc1[2:]
	fmt.Println("Slice 3:", l)

	t := []string{"g", "h", "i"}
	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Println("Declare and Initialize:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	tdSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		tdSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			tdSlice[i][j] = i + j
		}
	}
	fmt.Println("2d Slice: ", tdSlice)

	ts := make([][]int, 3, 5)
	fmt.Println("ts: ", ts)
}
