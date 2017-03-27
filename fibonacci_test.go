package fibonacci

import (
	"testing"
)

// the number we use for testing the fib functions
var n = uint64(1000)

// expected value for fib(1000) as a string
var expectedValue = "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875"

// test the easy implementation using bottom-up dynamic programming
func TestEasyFib(t *testing.T) {
	fn := EasyFibTest(n, true)
	if expectedValue != fn.Text(10) {
		t.Fatal("Value for fib(", n, ") is not the expected one")
	}
}

func TestMemoFib(t *testing.T) {
	memo, fn := MemoFibTest(n, true)
	fncopy := getMemoFibNumber(n, memo)

	if (expectedValue != fn.Text(10)) && (expectedValue != fncopy.Text(10)) {
		t.Fatal("Value for fib(", n, ") is not the expected one")
	}
}

func TestMemoFibGoroutine(t *testing.T) {
	MemoFibGoroutineTest(uint64(1000), true)
}

func TestMemoFibService(t *testing.T) {
	fn := MemoFibServiceTest(n)
	if expectedValue != fn {
		t.Fatal("Value for fib(", n, ") is not the expected one")
	}
}
