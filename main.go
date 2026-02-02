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

func main() {
	Factorial(5)
	IsPrime(7)
	Power(2, 3)

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
}
