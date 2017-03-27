package fibonacci

import (
	"fmt"
	"math/big"
	"time"
)

// compute the n-th fibonacci number: the easy way
func Fib(n big.Int) big.Int {
	// Initialize values we need for the computation
	one := big.NewInt(1)

	// Initialize two big ints with the first two numbers in the sequence
	a := big.NewInt(0)
	b := big.NewInt(1)

	i := big.NewInt(1)
	for i.Cmp(&n) <= 0 {
		// Compute the next Fibonacci number, storing it in a
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence
		a, b = b, a

		i.Add(i, one)
	}

	return *a
}

///////////////////////////////////////////////////////////////////////////////
// Test function

func EasyFibTest(narg uint64, printFibN bool) big.Int {
	var n big.Int
	n.SetUint64(narg)

	start := time.Now()
	fn := Fib(n)
	end := time.Now()
	delta1 := end.Sub(start)

	if printFibN {
		fmt.Println("easyFibTest: fib(", n.Uint64(), ") = ", &fn)
		fmt.Println("easyFibTest: Computed in: ", delta1)
	} else {
		fmt.Println("easyFibTest: fib(", n.Uint64(), ") Computed in: ", delta1)
	}
	return fn
}
