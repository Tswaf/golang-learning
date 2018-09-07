package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)

	a[0] = 5
	a[1] = 2
	a[2] = 6
	fmt.Println(a)

	b := [2]int{3, 8}
	fmt.Println(b)

	c := [3]int{8}
	fmt.Println(c)

	d := [...]int{3, 4}
	fmt.Println(d)

	e := [2]string{"cat", "dog"}
	f := e
	fmt.Println(e, f)

	f[0] = "pig"
	fmt.Println(e, f)

	letters := [2][3]string{
		{"a", "b", "c"},
		{"A", "B", "C"},
	}
	fmt.Println(letters)
}
