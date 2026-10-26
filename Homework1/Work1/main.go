package main

import "fmt"

const N = 8

func main() {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if (row+col)%2 == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
