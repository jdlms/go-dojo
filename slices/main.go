package main

import "fmt"

func main() {

	// # three ways to initialize a slice
	var mySlice []int = []int{}
	var myOtherSlice []int
	yetAnotherSlice := []int{}

	fmt.Println(mySlice, myOtherSlice, yetAnotherSlice)
	//

	// # convert array into slice
	// probably the bad way, mutates:
	anArray := []int{1, 2, 3, 4}

	// convert full array
	// (v[:] is called a slice expression)
	nowSlice := anArray[:]
	// or a selection
	smallerSlice := anArray[2:]

	//mutates
	anArray[0] = 9
	fmt.Println(nowSlice, smallerSlice) // [9 2 3 4] [3 4]

	// try copy instead:

	differentArray := []int{5, 6, 7, 8}
	newSlice := make([]int, len(differentArray))
	copy(newSlice, differentArray[:])
	// doesn't mutate
	differentArray[0] = 9
	fmt.Println("newSlice", newSlice)
	//

	// # concat two arrays into a slice
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{5, 6, 7, 8}

	slice1 := arr1[:]
	slice2 := arr2[:]

	concat := append(slice1, slice2...)

	fmt.Println(concatFunc(arr1, arr2))

	fmt.Println(concat)
	//

	// # concat two slices into an array
	soonToBeArray1 := []int{5, 4, 3, 2, 1}
	soonToBeArray2 := []int{10, 9, 8, 7, 6}

	emptyArray := [10]int{}

	copy(emptyArray[:], soonToBeArray1)
	copy(emptyArray[len(soonToBeArray1):], soonToBeArray2)

	fmt.Println(emptyArray)
}

func concatFunc(x [4]int, y [4]int) (string, []int) {
	s1 := x[:]
	s2 := y[:]

	concat := append(s1, s2...)
	return "concatFunc:", concat
}
