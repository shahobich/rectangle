package main

import "fmt"

type rectangle struct {
	a int
	b int

}

func S(rec rectangle) (s int)	 {
	s = rec.a * rec.b
	return s

}

func P(rec rectangle) (p int) {
	p = 2 * (rec.a * rec.b)
	return P

}
func main() {
	fmt.Println(rectangle{})

}
