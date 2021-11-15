package main

import "fmt"

type Interface interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i Interface

	var t *T
	i = t
	fmt.Printf("%T\n", i)
	i.M()

	i = &T{"hello"}
	fmt.Printf("%T\n", i)
	i.M()
}
