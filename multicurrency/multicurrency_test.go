package multicurrency

import "testing"

/*
* $5 + 10 CHF = $10 if rate is 2:1
* Make amount private
* Money rounding
 */

func TestMultiplication(t *testing.T) {
	five := New(5)
	prd := five.Times(2)
	if prd.amount != 10 {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if prd.amount != 15 {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}
