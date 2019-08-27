package main

import "fmt"

func fact(n int) int{
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func sumTo(a int) int {
	if a == 1 {
		return a
	}
	return a + sumTo(a - 1)
}

func main(){
	fmt.Println(fact(7))
	fmt.Println(sumTo(10))
}