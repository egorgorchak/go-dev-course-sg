package main

import "fmt"

func main() {
	ints := getSlice(10)

	for _, el := range ints {
		if el%2 == 0 {
			fmt.Println(el, "even")
		} else {
			fmt.Println(el, "odd")
		}
	}

}

func getSlice(elements int) []int {
	slice := []int{}
	for i := 0; i < elements; i++ {
		slice = append(slice, i)
	}
	return slice
}
