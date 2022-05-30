package main

import "fmt"

func main() {
	var a int
	var b int
	sum := 0
	fmt.Scan(&a)
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}
