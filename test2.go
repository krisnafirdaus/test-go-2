package main

import "fmt"

func main() {
	number := 49
	if number%7 == 0 {
		fmt.Printf("%d adalah kelipatan 7\n", number)
	} else {
		fmt.Printf("%d bukan kelipatan 7\n", number)
	}
}
