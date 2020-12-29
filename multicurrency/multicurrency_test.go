package multicurrency

import "testing"

/*
* $5 + 10 CHF = $10 if rate is 2:1
* Money rounding
* equals()
* hashCode()
* Equal null
* Equal object
* Dollar/Franc duplication
* Common equals
* Common times
 */

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	prd := five.Times(2)
	if !prd.Equals(&(NewFranc(10).Money)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !prd.Equals(&(NewFranc(15).Money)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	prd := five.Times(2)
	if !prd.Equals(&(NewDollar(10).Money)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !prd.Equals(&(NewDollar(15).Money)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)
	if !five.Equals(&(NewDollar(5).Money)) {
		t.Errorf("Expected true but got false")
	}
	if five.Equals(&(NewDollar(6).Money)) {
		t.Errorf("Expected false but got true")
	}

	six := NewFranc(6)
	if !six.Equals(&(NewFranc(6).Money)) {
		t.Errorf("Expected true but got false")
	}
	if six.Equals(&(NewFranc(5).Money)) {
		t.Errorf("Expected false but got true")
	}
}
