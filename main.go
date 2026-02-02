package main

import (
	"errors"
	"fmt"
	"math"
)

// part 1
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	sqrtVal := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtVal; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return -1, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

// part 2
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial
	add = func(x int) {
		total += x
	}
	subtract = func(x int) {
		total -= x
	}
	get = func() int {
		return total
	}
	return add, subtract, get
}

// part 3
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = operation(num)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(accumulator int, current int) int) int {
	acc := initial
	for _, num := range nums {
		acc = operation(acc, num)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	// part 1
	Factorial(5)
	IsPrime(7)
	Power(2, 3)

	// part 2
	counter1 := MakeCounter(0)
	fmt.Println(counter1()) //1
	fmt.Println(counter1()) //2

	counter2 := MakeCounter(10)
	fmt.Println(counter2()) //11
	fmt.Println(counter1()) //3 (independent from counter2)

	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)
	fmt.Println(double(5)) //10
	fmt.Println(triple(5)) //15

	add, sub, get := MakeAccumulator(100)
	add(50)
	fmt.Println(get()) //150
	sub(30)
	fmt.Println(get()) //120

	// part 3
	nums1 := []int{1, 2, 3, 4}
	squared := Apply(nums1, func(x int) int { return x * x })
	fmt.Println(squared) // [1 4 9 16]

	nums2 := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums2, func(x int) bool { return x%2 == 0 })
	fmt.Println(evens) // [2 4 6]

	sum := Reduce(nums1, 0, func(acc, curr int) int { return acc + curr })
	fmt.Println(sum) // 10

	addTwoFunc := func(x int) int { return x + 2 }
	doubleFunc := func(x int) int { return x * 2 }
	doubleThenAddTwo := Compose(addTwoFunc, doubleFunc)
	result := doubleThenAddTwo(5) // (5 * 2) + 2 = 12
	fmt.Println(result)
}
