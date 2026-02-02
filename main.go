// Roger Zheng
// CMPS2242 Lab #2 - Advanced Go
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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

// part 4
func ExploreProcess() {
	fmt.Printf("Current Process ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Memory address of slice: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])
	fmt.Println("Other processes cannot access these memory addresses directly due to process isolation.")
}

//A process ID is a unique identifier assigned to each running program by the OS to track and manage all running programs.
//Process isolation is important because it prevents one process from interfering with another resulting in enhanced security and stability.
//Without process isolation, if one process crashes, all other running processes could crash as well.
//The difference between the slice header address and the element address is that the slice header points to the memory location of the slice pointer
//while the element address points to the actual data stored in the slice.

// part 5
func DoubleValue(x int) { //Question: Will this modify the original variable? Why or why not?
	x = x * 2 //Answer: No, this will not modify the original variable because Go uses pass-by-value for function arguments.
}

func DoublePointer(x *int) { //Question: Will this modify the original variable? Why or why not?
	*x = *x * 2 //Answer: Yes, this will modify the original variable because we are passing a pointer to the variable which allows us to change its value directly.
}

func CreateOnStack() int {
	stackX := 123
	return stackX //This variable stays on the stack
}

func CreateOnHeap() *int {
	heapX := 456
	return &heapX //This variable escapes to the heap
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

func AnalyzeEscape() {
	CreateOnStack() //CreateOnStack() variable stays on the stack
	CreateOnHeap()  //CreateOnHeap() variable escapes to the heap
} //heapX variable escapes to the heap to the heap because a reference to it's value was returned
//"Escapes to heap" means that the variable lifetime exceeds beyond the function it was created in thereby requiring heap allocation

func main() {
	// part 1
	fmt.Println("Part 1: ")
	fmt.Println(Factorial(5))
	fmt.Println(IsPrime(7))
	fmt.Println(Power(2, 3))

	// part 2
	fmt.Println("\nPart 2: ")
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
	fmt.Println("\nPart 3: ")
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

	// part 4
	fmt.Println("\nPart 4: ")
	ExploreProcess()

	//part 5
	fmt.Println("\nPart 5: ")
	a, b := 5, 10
	fmt.Printf("Before SwapValues: a = %d, b = %d\n", a, b)
	a, b = SwapValues(a, b)
	fmt.Printf("After SwapValues: a = %d, b = %d\n", a, b)
	SwapPointers(&a, &b)
	fmt.Printf("After SwapPointers: a = %d, b = %d\n", a, b)
	AnalyzeEscape()
}
