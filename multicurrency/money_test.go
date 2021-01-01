package multicurrency

import (
	"reflect"
	"testing"
)

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
	// Casting interface back to struct pointer using
	// IBar.(*foo) type casting
	prd := five.Times(2).(*Money)
	if !prd.Equals(franc(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3).(*Money)
	if !prd.Equals(franc(15)) {
		t.Errorf("Expected 15 but got %d", prd.amount)
	}
}

func TestMultiplication(t *testing.T) {
	five := dollar(5)
	prd := five.Times(2).(*Money)
	if !prd.Equals(dollar(10)) {
		t.Errorf("Expected 10 but got %d", prd.amount)
	}
	prd = five.Times(3).(*Money)
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

func TestSumPlusMoney(t *testing.T) {
	fiveDollars := dollar(5)
	tenFrancs := franc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := &Sum{
		augend: fiveDollars,
		addend: tenFrancs,
	}
	result := bank.Reduce(sum.Plus(fiveDollars), "USD")
	if !result.Equals(dollar(15)) {
		t.Errorf("Expected 15 got %d", result.Amount())
	}

}

func TestSumTimes(t *testing.T) {
	fiveDollars := dollar(5)
	tenFrancs := franc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := &Sum{
		augend: fiveDollars,
		addend: tenFrancs,
	}
	result := bank.Reduce(sum.Times(2), "USD")
	if !result.Equals(dollar(20)) {
		t.Errorf("Expected 20 got %d", result.Amount())
	}
}

func TestPlusSameCurrencyReturnsMonesy(t *testing.T) {
	sum := dollar(1).Plus(dollar(1))
	castVar := reflect.ValueOf(sum)
	if castVar.Type() != reflect.TypeOf(&Money{}) {
		t.Errorf("Expected %v type got %v", reflect.TypeOf(&Money{}), castVar.Type())
	}
}
