package main

import "fmt"

func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n-1)
}

func a() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func sqr(peace int) int {
	return peace * peace
}

func main() {
	fmt.Print("Enter a number ")

	var i int
	fmt.Scanf("%d", &i)
	fmt.Printf("Fact(%v) = %v ", i, factorial(i))
	fmt.Printf("Sqr(Fact(%v)) = %v", i, sqr(factorial(i)))
}