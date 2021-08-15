package calculate

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	InitOpMap()
	testCalculate(t, "+", "Test1", 2, 3, 5)
	testCalculate(t, "+", "Test2", 5, 5, 10)
	testCalculate(t, "-", "Test3", 5, 2, 3)
	testCalculate(t, "-", "Test4", 2, 6, -4)
	testCalculate(t, "*", "Test5", 3, 3, 9)
	testCalculate(t, "*", "Test6", 2, 6, 12)
	testCalculate(t, "/", "Test7", 9, 3, 3)
	testCalculate(t, "/", "Test8", 20, 4, 5)
}

func testCalculate(t *testing.T ,op, testName string, a, b, expected int) {
	t.Run(fmt.Sprintf("Calculate %s", op), func(t *testing.T) {
		s := Calculate(op, a, b)
		if s != expected {
			t.Errorf("%s is Failed! expected: %d, output: %d\n",testName, expected, s)
		}
	})
}

