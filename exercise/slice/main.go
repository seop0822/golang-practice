package main

import "fmt"

func main() {
	var s []int
	var t []int
	var u []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	t = append(s, 400)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))

	fmt.Println("///////////////")
	s = append(s, 500, 600, 700, 800)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(u, len(u), cap(u))

}
