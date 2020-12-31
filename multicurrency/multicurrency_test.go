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
	five := franc(5)
	prd := five.Times(2)
	if !Equals(prd, franc(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !Equals(prd, franc(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestMultiplication(t *testing.T) {
	five := dollar(5)
	prd := five.Times(2)
	if !Equals(prd, dollar(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !Equals(prd, dollar(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestEquality(t *testing.T) {
	five := dollar(5)
	if !Equals(five, dollar(5)) {
		t.Errorf("Expected true but got false")
	}
	if Equals(five, dollar(6)) {
		t.Errorf("Expected false but got true")
	}

	six := franc(6)
	if !Equals(six, franc(6)) {
		t.Errorf("Expected true but got false")
	}
	if Equals(six, franc(5)) {
		t.Errorf("Expected false but got true")
	}

	if Equals(six, dollar(6)) {
		t.Errorf("Expected false but got true")
	}
}

func TestCurrency(t *testing.T) {
	one := dollar(1)
	if one.Currency() != "USD" {
		t.Errorf("Expected USD got %s", one.Currency())
	}
	two := franc(2)
	if two.Currency() != "CHF" {
		t.Errorf("Expected CHF got %s", two.Currency())
	}
}

func TestSimpleAddition(t *testing.T) {
	five := dollar(5)
	sum := five.Plus(five)
	bank := &Bank{}
	reduced := bank.Reduce(sum, "USD")
	if !Equals(reduced, dollar(10)) {
		t.Errorf("Expected 10 got %d", reduced.Amount())
	}
}
