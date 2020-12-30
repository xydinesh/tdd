package multicurrency

import "fmt"

// MoneyInterface interface
type MoneyInterface interface {
	Amount() int
	Currency() string
}

// Equals function compares two objects for the value
func Equals(a MoneyInterface, b MoneyInterface) bool {
	return a.Amount() == b.Amount() && a.Currency() == b.Currency()
}

// String method prints the currency value and type
func String(a MoneyInterface) string {
	return fmt.Sprintf("%d %s", a.Amount(), a.Currency())
}

// Money is a super class for Dollar and Franc
type Money struct {
	amount   int
	currency string
}

// Amount returns the amount
func (m *Money) Amount() int {
	return m.amount
}

// Currency retuns the currenecy
func (m *Money) Currency() string {
	return m.currency
}

// Times returns the multiplications
func (m *Money) Times(multiplier int) *Money {
	return &Money{
		amount:   m.Amount() * multiplier,
		currency: m.Currency(),
	}
}

// Dollar function to create new dollar
func dollar(d int) *Money {
	return NewDollar(d)
}

// franc function to create new franc
func franc(f int) *Money {
	return NewFranc(f)
}

// Dollar structure for the book
type Dollar struct {
	Money
}

// NewDollar method to create new Dollar instance
func NewDollar(amount int) *Money {
	d := &Dollar{
		Money{
			amount:   amount,
			currency: "USD",
		},
	}
	return &d.Money
}

// Franc structure
type Franc struct {
	Money
}

// NewFranc method to create new instance
func NewFranc(amount int) *Money {
	f := &Franc{
		Money{
			amount:   amount,
			currency: "CHF",
		},
	}
	return &f.Money
}
