// Problem:	Solution 2
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

	var total int

	for i := n; i <= target; i += n {
		total += i
	}
	return total
}
