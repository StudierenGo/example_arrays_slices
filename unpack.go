package main

import "fmt"

func unpack() {
	var sl1 = []int{1, 2, 3}
	var sl2 = []int{4, 5, 6}

	var sl3 = append(sl1, sl2...)
	fmt.Println(sl3)

	for index, value := range sl3 {
		fmt.Println(index, value)
	}
}
