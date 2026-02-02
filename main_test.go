package main

import (
	"testing"
)

// part 1
func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 12, want: 479001600, wantErr: false},
		{name: "factorial of -1", input: -1, want: 0, wantErr: true},
		{name: "factorial of -10", input: -10, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "Prime of 2", input: 2, want: true, wantErr: false},
		{name: "Prime of 3", input: 3, want: true, wantErr: false},
		{name: "Prime of 4", input: 4, want: false, wantErr: false},
		{name: "Prime of 17", input: 17, want: true, wantErr: false},
		{name: "Prime of 25", input: 25, want: false, wantErr: false},
		{name: "Prime of -2", input: -2, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name          string
		inputBase     int
		inputExponent int
		want          int
		wantErr       bool
	}{
		{name: "Power of 2 to the 4", inputBase: 2, inputExponent: 4, want: 16, wantErr: false},
		{name: "Power of 3 to the 2", inputBase: 3, inputExponent: 2, want: 9, wantErr: false},
		{name: "Power of 5 to the 0", inputBase: 5, inputExponent: 0, want: 1, wantErr: false},
		{name: "Power of 0 to the 2", inputBase: 0, inputExponent: 2, want: 0, wantErr: false},
		{name: "Power of 2 to the -3", inputBase: 2, inputExponent: -3, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.inputBase, tt.inputExponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

// part 2
func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "counter starting at 0", input: 0, want: 1},
		{name: "counter starting at 5", input: 5, want: 6},
		{name: "counter starting at -3", input: -3, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.input)
			got := counter()
			if got != tt.want {
				t.Errorf("MakeCounter() input = %v, got %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "multiplier of 2", factor: 2, input: 2, want: 4},
		{name: "multiplier of 3", factor: 6, input: 3, want: 18},
		{name: "multiplier of -1", factor: -3, input: -1, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)
			got := multiplier(tt.input)
			if got != tt.want {
				t.Errorf("MakeMultiplier() factor = %v, input = %v, got %v, want %v", tt.factor, tt.input, got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		add      int
		subtract int
		want     int
	}{
		{name: "accumulator starting at 0", initial: 0, add: 5, subtract: 0, want: 5},
		{name: "accumulator starting at 10", initial: 100, add: 50, subtract: 30, want: 120},
		{name: "accumulator starting at -5", initial: -5, add: 0, subtract: 5, want: -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addFunc, subFunc, getFunc := MakeAccumulator(tt.initial)
			addFunc(tt.add)
			subFunc(tt.subtract)
			got := getFunc()
			if got != tt.want {
				t.Errorf("MakeAccumulator() initial = %v, add = %v, subtract = %v, got %v, want %v", tt.initial, tt.add, tt.subtract, got, tt.want)
			}
		})
	}
}

// part 3
func TestApply(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		function func(int) int
		want     []int
	}{
		{name: "square function", nums: []int{1, 2, 3}, function: func(x int) int { return x * x }, want: []int{1, 4, 9}},
		{name: "double function", nums: []int{0, -1, -2}, function: func(x int) int { return x * 2 }, want: []int{0, -2, -4}},
		{name: "negate function", nums: []int{1, -2, 3}, function: func(x int) int { return -x }, want: []int{-1, 2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.function)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply() got[%v] = %v, want[%v] = %v", i, got[i], i, tt.want[i])
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{name: "even numbers", nums: []int{1, 2, 3, 4, 5}, predicate: func(x int) bool { return x%2 == 0 }, want: []int{2, 4}},
		{name: "positive numbers", nums: []int{-2, -1, 0, 1, 2, 15}, predicate: func(x int) bool { return x > 0 }, want: []int{1, 2, 15}},
		{name: "greater than ten", nums: []int{1, 2, 30, 4, 50}, predicate: func(x int) bool { return x > 10 }, want: []int{30, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter() got[%v] = %v, want[%v] = %v", i, got[i], i, tt.want[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{name: "sum", nums: []int{1, 2, 3, 4}, initial: 0, operation: func(a, b int) int { return a + b }, want: 10},
		{name: "product", nums: []int{1, 2, 3, 4}, initial: 1, operation: func(a, b int) int { return a * b }, want: 24},
		{name: "max", nums: []int{1, 5, 3, 4}, initial: 0,
			operation: func(a, b int) int {
				if a > b {
					return a
				} else {
					return b
				}
			}, want: 5},
		{name: "min", nums: []int{4, 2, 8, 1}, initial: 100,
			operation: func(a, b int) int {
				if a < b {
					return a
				} else {
					return b
				}
			}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{name: "double then add two", f: func(x int) int { return x + 2 }, g: func(x int) int { return x * 2 }, input: 5, want: 12},
		{name: "square then increment", f: func(x int) int { return x + 1 }, g: func(x int) int { return x * x }, input: 3, want: 10},
		{name: "negate then double", f: func(x int) int { return x * 2 }, g: func(x int) int { return -x }, input: 4, want: -8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composedFunc := Compose(tt.f, tt.g)
			got := composedFunc(tt.input)
			if got != tt.want {
				t.Errorf("Compose() got = %v, want %v", got, tt.want)
			}
		})
	}
}
