package main

import "fmt"

const N = 8

func main() {
	for a := 0; a < N; a++ {
		for b := 0; b < N; b++ {
			if (a+b)%2 == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
