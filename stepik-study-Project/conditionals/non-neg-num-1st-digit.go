package main

import "fmt"

func main() {
	var n uint16
	fmt.Scan(&n)

	switch {
	case n == 10000:
		fmt.Println(n / 10000)
	case n >= 1000:
		fmt.Println(n / 1000)
	case n >= 100:
		fmt.Println(n / 100)
	case n >= 10:
		fmt.Println(n / 10)
	case n < 10:
		fmt.Println(n / 1)
	}
}
