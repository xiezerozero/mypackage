package main

import "fmt"

func main() {
	for pos, char := range "日本語" {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
}
