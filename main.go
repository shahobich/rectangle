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
func main() {
	fmt.Println(rectangle{})

}
