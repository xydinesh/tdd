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
* Common times
* Compare Franc's with Dollars
 */

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	prd := five.Times(2)
	if !Equals(prd, NewFranc(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !Equals(prd, NewFranc(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	prd := five.Times(2)
	if !Equals(prd, NewDollar(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !Equals(prd, NewDollar(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)
	if !Equals(five, NewDollar(5)) {
		t.Errorf("Expected true but got false")
	}
	if Equals(five, NewDollar(6)) {
		t.Errorf("Expected false but got true")
	}

	six := NewFranc(6)
	if !Equals(six, NewFranc(6)) {
		t.Errorf("Expected true but got false")
	}
	if Equals(six, NewFranc(5)) {
		t.Errorf("Expected false but got true")
	}

	if Equals(six, NewDollar(6)) {
		t.Errorf("Expected false but got true")
	}
}
