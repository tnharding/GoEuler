// Problem: Solution 1
// Find the sum of all the multiples of 3 or 5 below 1000
package main

import "fmt"

func main() {

	total1 := Solution1()
	total2 := Solution2()
	total3 := Solution3()

	fmt.Println("The sum of all multiples of 3 or 5 below 1000.")
	fmt.Println("Solution 1:", total1)
	fmt.Println("Solution 2:", total2)
	fmt.Println("Solution 3:", total3)
}

//Solution1 sum of all the multiples of 3 or 5 below 1000 using for loop
func Solution1() int {
	var total int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

//Solution2 sum of all the multiples of 3 or 5 below 1000
func Solution2() int {
	total := sumDivisibleBy(3, 999)
	total += sumDivisibleBy(5, 999)
	total -= sumDivisibleBy(15, 999)
	return total
}

func sumDivisibleBy(n, target int) int {
	var total int

	for i := n; i <= target; i += n {
		total += i
	}
	return total
}

//Solution3 sum of all the multiples of 3 or 5 below 1000 using a for loop
func Solution3() int {
	total := sumDivisibleBy2(3, 999)
	total += sumDivisibleBy2(5, 999)
	total -= sumDivisibleBy2(15, 999)
	return total
}

func sumDivisibleBy2(n, target int) int {
	p := target / n
	return n * (p * (p + 1)) / 2
}
