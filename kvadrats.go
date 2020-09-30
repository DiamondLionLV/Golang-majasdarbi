package main
import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	c := a * b
	d := (2 * a) + (2 * b)

	fmt.Print("Perimetrs: ", d, "Laukums: ", c)
}
