package multicurrency

import "reflect"

// MoneyInterface interface
type MoneyInterface interface {
	Amount() int
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
	amount int
}

// NewDollar method to create new Dollar instance
func NewDollar(amount int) *Dollar {
	d := &Dollar{
		amount: amount,
	}
	return d
}

// Times multiply dollar amount and return new instance
func (d *Dollar) Times(multiplier int) *Dollar {
	p := &Dollar{
		amount: d.amount * multiplier,
	}
	return p
}

// Amount returns amount
func (d *Dollar) Amount() int {
	return d.amount
}

// Franc structure
type Franc struct {
	amount int
}

// NewFranc method to create new instance
func NewFranc(amount int) *Franc {
	f := &Franc{
		amount: amount,
	}
	return f
}

// Times multiply and return a new instance
func (f *Franc) Times(multiplier int) *Franc {
	n := &Franc{
		amount: f.amount * multiplier,
	}
	return n
}

// Amount returns amount
func (f *Franc) Amount() int {
	return f.amount
}
