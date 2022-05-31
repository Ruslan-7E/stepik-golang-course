package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var firDig int = a / 100000
	var secDig int = (a % 100000) / 10000
	var thirdDig int = (a % 10000) / 1000
	var fourthDig int = (a % 1000) / 100
	var fifDig int = (a % 100) / 10
	var sixDig int = a % 10

	if firDig+secDig+thirdDig == fourthDig+fifDig+sixDig {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
