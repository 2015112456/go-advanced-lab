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
