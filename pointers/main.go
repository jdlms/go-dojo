package main

import "fmt"

type aStruct struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)

	// Pointer to f, creation of a new pointer header, which also references f's location
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	// The value of f changes:
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)

	// The value of f does not change
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)
	fmt.Printf("Value of f: %.2f\n", f)

	// The value of f still hasn't changed
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)
	fmt.Printf("Value of f: %.2f\n", f)

	// Check for empty struct
	var k *aStruct

	// this points to nowhere:
	fmt.Println(k)
	// so you can do this:
	if k == nil {
		k = new(aStruct)
	}

	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil")
	}

}
