package main

import (
	"fmt"
)

func main()  {
	var a int
	a = 5
	var b string = "golang"
	fmt.Println(a,b)

	var c = 7
	var d = "golang"
	fmt.Println(c,d)

	e,f := 8,"golang"
	fmt.Println(e,f)

	var g int = 9
	var h float32 = 12.9
	var i int
	i = g + int(h)
	fmt.Println(g,h,i)

	const(
		j = 8
		k = "golang"
	)
	fmt.Println(j,k)

}
