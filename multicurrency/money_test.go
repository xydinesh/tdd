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
	if !prd.Equals(franc(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !prd.Equals(franc(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestMultiplication(t *testing.T) {
	five := dollar(5)
	prd := five.Times(2)
	if !prd.Equals(dollar(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3)
	if !prd.Equals(dollar(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestEquality(t *testing.T) {
	five := dollar(5)
	if !five.Equals(dollar(5)) {
		t.Errorf("Expected true but got false")
	}
	if five.Equals(dollar(6)) {
		t.Errorf("Expected false but got true")
	}

	six := franc(6)
	if !six.Equals(franc(6)) {
		t.Errorf("Expected true but got false")
	}
	if six.Equals(franc(5)) {
		t.Errorf("Expected false but got true")
	}

	if six.Equals(dollar(6)) {
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
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if !reduced.Equals(dollar(10)) {
		t.Errorf("Expected 10 got %d", reduced.Amount())
	}
}

func TestReduceSum(t *testing.T) {
	sum := &Sum{
		augend: dollar(3),
		addend: dollar(4),
	}
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	if !result.Equals(dollar(7)) {
		t.Errorf("Expected 7 got %d", result.Amount())
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(franc(2), "USD")
	if !result.Equals(dollar(1)) {
		t.Errorf("Expected 1 got %d", result.Amount())
	}
}

func TestMixedAddition(t *testing.T) {
	fiveDollars := dollar(5)
	tenFrancs := franc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveDollars.Plus(tenFrancs), "USD")
	if !result.Equals(dollar(10)) {
		t.Errorf("Expected 10 got %d", result.Amount())
	}
}
