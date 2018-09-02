package main

import (
	"fmt"
)

func main()  {
	a := 9
	b := 9.8

	c := a + int(b)

	fmt.Println(a,b,c)

	const (
		c = 9
		e = "month"
	)

}
