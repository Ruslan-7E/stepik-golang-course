package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a int = n / 100
	var b int = (n / 10) % 10
	var c int = n % 10

	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
