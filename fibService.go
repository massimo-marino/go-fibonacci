package fibonacci

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

// goroutine: runs forever waiting n on channel inch and sending back the result
// as a string on channel outch
func grMemoFib(inch <-chan big.Int, outch chan<- string, printFibN bool) {
	defer close(outch)
	defer wg.Done()

	// initialize the memo
	memo := make(Memo, 10)

	var callCounter uint64 = 1
	var timeout bool = false

	fn := big.NewInt(0)

	for {
		if timeout == false {
			fmt.Println("GrMemoFib: waiting a number on channel...")
		}
		select {
		case n, ok := <-inch:
			if !ok {
				fmt.Println("GrMemoFib. Channel was closed... ending")
				return
			} else {
				memo, *fn = memoFibCall(memofib, n, memo, printFibN, "GrMemoFib: [goroutine grMemoFib] [call #"+strconv.FormatUint(callCounter, 10)+"] Call to")
				// func (x *Int) Text(base int) string
				outch <- fn.Text(10)
				callCounter++
				timeout = false
			}
		default:
			timeout = true
		}
	}
}

// channel for sending the number
var fibNch = make(chan big.Int)

// channel for receiving the result
var fibFch = make(chan string)

// start the fib service as a background goroutine that waits for numbers for
// computing a fibonacci number
func StartFibService() {

	printFibN := true

	wg.Add(1)
	go grMemoFib(fibNch, fibFch, printFibN)

	fmt.Printf("StartFibService: Waiting for goroutines to finish\n")
	wg.Wait()
}

// a client of the service should call this function
func QueryFibService(n big.Int) string {
	fibNch <- n
	return <-fibFch
}

///////////////////////////////////////////////////////////////////////////////
// Test functions

func MemoFibGoroutineTest(narg uint64, printFibN bool) {

	var n big.Int
	n.SetUint64(narg)

	// goroutine test
	ch := make(chan big.Int)
	fch := make(chan string)

	wg.Add(1)
	go grMemoFib(ch, fch, printFibN)

	for i := 1; i <= 20; i++ {
		// increment n
		n.SetUint64(n.Uint64() + uint64(1))

		// first call should take longer
		ch <- n
		fn1 := <-fch
		if printFibN {
			fmt.Println("memoFibServiceTest: First Result: ", fn1)
			fmt.Println()
		}
		// second call should take less time than the first one
		ch <- n
		fn2 := <-fch
		if printFibN {
			fmt.Println("memoFibServiceTest: Second Result: ", fn2)
			fmt.Println()
		}
	}
}

func MemoFibServiceTest(narg uint64) string {
	// start the fib service as a background goroutine
	go StartFibService()

	var n big.Int
	n.SetUint64(narg)

	// query the service and receive the result
	fn := QueryFibService(n)

	return fn
}
