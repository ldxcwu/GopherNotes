package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a
	c := &a
	a[0] = 0
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	// m := []int{1, 2, 3}
	m := make([]int, 0, 4)
	m = append(m, 1, 2, 3)
	fmt.Println(len(m), cap(m))
	n := m
	n[0] = 0
	n = append(n, 4)
	fmt.Println(n)
	fmt.Println(m)
	m = append(m, -1)
	fmt.Println(n)
	fmt.Println(m)
}
