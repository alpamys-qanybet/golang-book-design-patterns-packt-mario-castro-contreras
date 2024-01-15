package main

import (
	"fmt"
	"golang-book-design-patterns-packt-mario-castro-contreras/chapter2/p55_singleton"
)

func main() {

	a := p55_singleton.GetInstance()

	fmt.Println(a)
	a.AddOne()
	fmt.Println(a)

	b := p55_singleton.GetInstance()
	fmt.Println(b)
	b.AddOne()
	fmt.Println(b)
	fmt.Println(a)

	a.AddOne()
	fmt.Println(b)
	fmt.Println(a)
}

/*

&{0}
&{1}
&{1}
&{2}
&{2}
&{3}
&{3}

*/
