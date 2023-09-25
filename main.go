package main

import "fmt"

func main() {
	var xx [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&xx[i])
	}

	for i := 0; i < 10; i++ {
		if xx[i] <= 0 {
			fmt.Println("1")
		} else {
			fmt.Println(&xx[i])
		}
	}
}
