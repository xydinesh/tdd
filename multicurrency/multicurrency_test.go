package multicurrency

import "testing"

/*
* $5 + 10 CHF = $10 if rate is 2:1
* $5 * 2 = $10
* Make amount private
* Dollar side effects
* Money rounding
 */

func TestMultiplication(t *testing.T) {
	five := New(5)
	five.Times(2)
	if five.amount != 10 {
		t.Errorf("Expected 10 but got %d", five.amount)
	}
}
