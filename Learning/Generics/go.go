package main

import (
	"fmt"
	"time"
)

// Thanks -> https://go.dev/doc/tutorial/generics

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

/*
Declare a SumIntsOrFloats function with two type parameters (inside the square brackets), K and V, and one argument that uses the type parameters, m of type map[K]V. The function returns a value of type V.
Specify for the K type parameter the type constraint comparable. Intended specifically for cases like these, the comparable constraint is predeclared in Go. It allows any type whose values may be used as an operand of the comparison operators == and !=.
Go requires that map keys be comparable. So declaring K as comparable is necessary so you can use K as the key in the map variable. It also ensures that calling code uses an allowable type for map keys.
Specify for the V type parameter a constraint that is a union of two types: int64 and float64. Using | specifies a union of the two types, meaning that this constraint allows either type.
Either type will be permitted by the compiler as an argument in the calling code. Specify that the m argument is of type map[K]V, where K and V are the types already specified for the type parameters.
Note that we know map[K]V is a valid map type because K is a comparable type. If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.
*/

type Number interface {
	int64 | float64
}

/*
Declare the Number interface type to use as a type constraint.
Declare a union of int64 and float64 inside the interface.
Essentially, you’re moving the union from the function declaration into a new type constraint. That way, when you want to constrain a type parameter to either int64 or float64, you can use this Number type constraint instead of writing out int64 | float64.
*/

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

/*
Declare a generic function with the same logic as the generic function you declared previously, but with the new interface type instead of the union as the type constraint. As before, you use the type parameters for the argument and return types.
*/

func main() {
	start := time.Now()

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
