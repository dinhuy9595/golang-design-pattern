package main

import "fmt"

func main() {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()

	expect := 1
	fmt.Printf("expect %d got %d \n", expect, res)
}
