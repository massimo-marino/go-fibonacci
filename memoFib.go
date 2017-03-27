package fibonacci

import (
	"fmt"
	"math/big"
	"time"
)

// the memo table
type Memo map[uint64]big.Int

// the memo function type
type MemoFib func(n big.Int, memo Memo) (Memo, big.Int)

func getMemoFibNumber(n uint64, memo Memo) big.Int {
	return memo[n]
}

// compute the n-th fibonacci number using memoization
func memofib(n big.Int, memo Memo) (Memo, big.Int) {
	// check if already memoized, and if so return it
	if v, vok := memo[n.Uint64()]; vok == true {
		// cached
		return memo, v
	}

	// Initialize two big ints with the first two numbers in the sequence
	if _, vok := memo[0]; vok == false {
		// not yet cached
		memo[0] = *big.NewInt(0)
	}

	if _, vok := memo[1]; vok == false {
		// not yet cached
		memo[1] = *big.NewInt(1)
	}

	// Initialize values we need for the computation
	one := big.NewInt(1)
	v1 := big.NewInt(0)
	v2 := big.NewInt(0)

	i := big.NewInt(2)
	for i.Cmp(&n) <= 0 {
		if _, vok := memo[i.Uint64()]; vok == false {
			// not cached
			*v1 = memo[i.Uint64()-1]
			*v2 = memo[i.Uint64()-2]
			f := big.NewInt(0)
			f.Add(v1, v2)
			memo[i.Uint64()] = *f
		}
		i.Add(i, one)
	}

	return memo, memo[i.Uint64()-1]
}

func memoFibCall(memofib MemoFib, n big.Int, memo Memo, printFibN bool, prompt string) (Memo, big.Int) {
	start := time.Now()
	fn := big.NewInt(0)
	memo, *fn = memofib(n, memo)
	end := time.Now()
	delta := end.Sub(start)

	if printFibN {
		fmt.Println("memoFibCall:", prompt, "memofib(", n.Uint64(), ") = ", fn)
		fmt.Println("memoFibCall: Computed in: ", delta)
	} else {
		fmt.Println("memoFibCall:", prompt, "memofib(", n.Uint64(), ") Computed in: ", delta)
	}

	return memo, *fn
}

///////////////////////////////////////////////////////////////////////////////
// Test function

func MemoFibTest(narg uint64, printFibN bool) (Memo, big.Int) {
	var n big.Int
	fn := big.NewInt(0)
	n.SetUint64(narg)

	memo := make(Memo, 10)

	// 1st call
	memo, *fn = memoFibCall(memofib, n, memo, printFibN, "memoFibTest: First call to")

	// 2nd call
	memo, *fn = memoFibCall(memofib, n, memo, printFibN, "memoFibTest: Second call to")

	return memo, *fn
}
