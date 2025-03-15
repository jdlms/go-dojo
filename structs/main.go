package main

import (
	"fmt"
	"strconv"
)

type Entry struct {
	Name     string
	LastName string
	Year     int
}

type Record struct {
	Field1 int
	Field2 string
}

func zeroStruct() Entry {
	return Entry{}
}

func initStruct(N, L string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, LastName: L, Year: 2000}
	}
	return Entry{Name: N, LastName: L, Year: Y}
}

func zeroPointerToStruct() *Entry {
	t := &Entry{}
	return t
}

func initPointerToString(N, L string, Y int) *Entry {
	if len(L) == 0 {
		return &Entry{Name: N, LastName: "unknown", Year: Y}
	}
	return &Entry{Name: N, LastName: L, Year: Y}
}

func main() {
	s1 := zeroStruct()
	p1 := zeroPointerToStruct()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := initStruct("Joshua", "Smith", 2025)
	p2 := initPointerToString("Joshua", "Smith", 2025)
	fmt.Println("s2:", s2, "p2:", *p2)

	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)

	pS := new(Entry)
	fmt.Println("pS:", pS)

	// slice of structs

	s := []Record{}
	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := Record{Field1: i, Field2: text}
		s = append(s, temp)
	}

	fmt.Println(s)

	fmt.Println("Index 0:", s[0].Field1, s[0].Field2)
	fmt.Println("Number of structures:", len(s))
	sum := 0
	for _, k := range s {
		sum += k.Field1
	}
	fmt.Println("Sum:", sum)
}
