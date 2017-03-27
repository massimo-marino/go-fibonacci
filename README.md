# Fibonacci

**Files:** *fibonacci.go, memoFib.go, fibService.go, fibonacci_test.go*

This is an implementation of computing Fibonacci numbers using big.Int's in go.

#### Fibonacci Numbers: The Easy Way

**File:** *fibonacci.go*

This is the easy way of implementing a bottom-up dynamic programming algorithm.

#### Fibonacci Numbers: A memoization implementation

**File:** *memoFib.go*

Compute Fibonacci numbers using memoization.
The only limits are your the memory, and your time if you dare to ask a huge `fib(n)` for a very high value of `n` :-)

#### Fibonacci Numbers: A fib service using memoization

**File:** *fibService.go*

A back-end service that runs forever using memoization.

#### The Tests

**File:** *fibonacci_test.go*

Tests for the fib functions.
