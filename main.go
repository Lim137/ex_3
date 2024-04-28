package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for divisor := 2; divisor < num/2; divisor++ {
		if num%divisor == 0 {
			return false
		}
	}
	return true
}

func getPrimeNums(from, to int) []int {
	var primeNums []int
	for num := from; num <= to; num++ {
		if isPrime(num) {
			primeNums = append(primeNums, num)
		}
	}
	return primeNums
}

func main() {
	var from, to int
	fmt.Scan(&from, &to)
	primeNums := getPrimeNums(from, to)
	fmt.Println(primeNums)
}
