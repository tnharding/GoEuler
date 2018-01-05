// Problem:	Solution 3
// Find the sum of all the multiples of 3 or 5 below 1000
package main

import "fmt"

func main() {

	total := sumDivisibleBy(3, 999)
	total += sumDivisibleBy(5, 999)
	total -= sumDivisibleBy(15, 999)

	fmt.Println("The sum of all multiples of 3 or 5 below 1000 is", total)
}

func sumDivisibleBy(n, target int) int {
	p := target / n
	return n * (p * (p + 1)) / 2
}
