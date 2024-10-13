package main

import "fmt"

func main() {
	var myMap = map[string]uint8{ "Adam": 10, "Cole": 8 }

	for name := range myMap {
		fmt.Println(name)
	}
}
