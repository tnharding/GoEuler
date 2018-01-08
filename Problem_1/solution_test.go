package main

import "testing"

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1()
	}
}
func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2()
	}
}
func BenchmarkSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution3()
	}
}
