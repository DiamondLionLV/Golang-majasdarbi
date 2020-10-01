package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Print("Ievadi stundas `(1-12)`")
	fmt.Scan(&a)
	fmt.Print("Ievadi minutes `(1-60)`")
	fmt.Scan(&b)

	c = 360 / b
	d = 60 / a
	e = c - d
	f = c + d

	if c > d {
		fmt.Print("Lenkis: ", e)
	} else {
		fmt.Print("Lenkis: ", f)
	}
}
