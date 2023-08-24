package main

import "fmt"

// An anonymous struct will only be used where it is declared
var data = []struct {
	string // An anonymous field will have the name of its type
	bool
	number int32
}{{"This is a string", true, 32678},
	{"This is a another string", false, -2},
}

func main() {
	for _, datum := range data {
		// An anonymous field will have the name of its type
		fmt.Printf("%s, %v, %d\n", datum.string, datum.bool, datum.number)
	}
}
