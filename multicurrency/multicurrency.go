package multicurrency

import "reflect"

// MoneyInterface interface
type MoneyInterface interface {
	Amount() int
	Currency() string
}

// Equals function compares two objects for the value
func Equals(a MoneyInterface, b MoneyInterface) bool {
	return a.Amount() == b.Amount() && reflect.TypeOf(a) == reflect.TypeOf(b)
}

// Dollar function to create new dollar
func dollar(d int) *Dollar {
	return NewDollar(d)
}

// franc function to create new franc
func franc(f int) *Franc {
	return NewFranc(f)
}

// Dollar structure for the book
type Dollar struct {
	amount   int
	currency string
}

// NewDollar method to create new Dollar instance
func NewDollar(amount int) *Dollar {
	d := &Dollar{
		amount:   amount,
		currency: "USD",
	}
	return d
}

// Times multiply dollar amount and return new instance
func (d *Dollar) Times(multiplier int) *Dollar {
	return dollar(d.amount * multiplier)
}

// Currency return currency format
func (d *Dollar) Currency() string {
	return d.currency
}

// Amount returns amount
func (d *Dollar) Amount() int {
	return d.amount
}

// Franc structure
type Franc struct {
	amount   int
	currency string
}

// NewFranc method to create new instance
func NewFranc(amount int) *Franc {
	f := &Franc{
		amount:   amount,
		currency: "CHF",
	}
	return f
}

// Times multiply and return a new instance
func (f *Franc) Times(multiplier int) *Franc {
	return franc(f.amount * multiplier)
}

// Amount returns amount
func (f *Franc) Amount() int {
	return f.amount
}

// Currency return currency format
func (f *Franc) Currency() string {
	return f.currency
}
