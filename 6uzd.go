package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Ievadi stundas `(1-12)`")
	fmt.Scan(&a)
	fmt.Print("Ievadi minutes `(1-60)`")
	fmt.Scan(&b)

	c := a * 30
	d := b * 6

	if c > d {
		//fmt.Print("Lenkis: ", e)
		fmt.Printf("Lenkis: %d", c-d)
	} else {
		//fmt.Print("Lenkis: ", f)
		fmt.Printf("Lenkis: %d", d-c)
	}
}
