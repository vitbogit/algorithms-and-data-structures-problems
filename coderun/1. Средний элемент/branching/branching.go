package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a < b {
		if b <= c {
			fmt.Print(b)
		} else if a <= c { // b > c
			fmt.Print(c)
		} else {
			fmt.Print(a)
		}
	} else if a == b {
		fmt.Print(a)
	} else { // a > b
		if a <= c {
			fmt.Print(a)
		} else if b <= c {
			fmt.Print(c)
		} else {
			fmt.Print(b)
		}
	}
}
