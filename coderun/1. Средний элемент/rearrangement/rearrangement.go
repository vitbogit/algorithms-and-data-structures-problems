package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a > b {
		a, b = b, a
	}
	// at this point, a <= b
	if a > c {
		a, c = c, a
	}
	// at this point, a <= b and a <= c
	if b > c {
		b, c = c, b
	}
	// at this point, a <= b <= c
	fmt.Print(b)
}
