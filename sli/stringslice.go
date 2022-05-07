package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (x StringSlice) Len() int {
	return len(x)
}

func (x StringSlice) Sort() {
	// do something to sort my slice
}

func main() {
	stringArray := []string{"one", "two", "three"}
	
	stringSlice1 := []string(stringArray)
	fmt.Printf("String Slice 1 length is %d", len(stringSlice1))
	
	stringSlice2 := StringSlice(stringArray)
	fmt.Printf("String Slice 2 length is %d", stringSlice2.Len())

	stringSlice2.Sort()

	stringSlice3 := sort.StringSlice(stringArray)
	fmt.Printf("String slice 3 length is %d", stringSlice3.Len())
	stringSlice3.Sort()
}