package main

import "fmt"

type example struct {
	a example2
	b example2
}

type example2 struct {
	c string
	d string
}

func (e *example) f(arg example) {
	e.a.c = arg.a.c
	e.b.c = arg.b.c
	e.a.d = arg.a.d
	e.b.d = arg.b.d

}

func (e *example) toString() string {
	return "{a:\n\tc:" + e.a.c + "\n\td:" + e.a.d + "\nb:\n\tc:" + e.b.c + "\n\td:" + e.b.d + "\n}"
}

func main() {
	x := example{a: example2{c: "c", d: "d"}, b: example2{c: "c", d: "d"}}
	y := example{a: example2{c: "d", d: "c"}, b: example2{c: "d", d: "c"}}
	// y := example{a: "a", b: "b"}
	x.f(y)
	fmt.Println(x.toString())

}
