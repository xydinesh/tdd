package multicurrency

import "testing"

/*
* $5 + 10 CHF = $10 if rate is 2:1
* Money rounding
* equals()
* hashCode()
* Equal null
* Equal object
 */

func TestMultiplication(t *testing.T) {
	five := New(5)
	prd := five.Times(2)
	if !prd.Equals(New(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !prd.Equals(New(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestEquality(t *testing.T) {
	five := New(5)
	if !five.Equals(New(5)) {
		t.Errorf("Expected true but got false")
	}
	if five.Equals(New(6)) {
		t.Errorf("Expected false but got true")
	}
}
