package multicurrency

import "fmt"

// MoneyInterface interface
type MoneyInterface interface {
	Amount() int
	Currency() string
}

// ExpressionInterface interface
type ExpressionInterface interface {
	Reduce(to string) *Money
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

// Plus adds two and return a new one
func (m *Money) Plus(n *Money) *Money {
	return &Money{
		amount:   m.Amount() + n.Amount(),
		currency: m.Currency(),
	}
}

// Reduce return value of the operation already done
func (m *Money) Reduce(to string) *Money {
	return &Money{
		amount:   m.amount,
		currency: m.currency,
	}
}

// Dollar function to create new dollar
func dollar(d int) *Money {
	return &Money{
		amount:   d,
		currency: "USD",
	}
}

// franc function to create new franc
func franc(f int) *Money {
	return &Money{
		amount:   f,
		currency: "CHF",
	}
}

// Bank structure
type Bank struct {
}

// Reduce use the expression
func (b *Bank) Reduce(ex ExpressionInterface, to string) *Money {
	return ex.Reduce(to)
}

// Sum structure to return expression
type Sum struct {
	augend *Money
	addend *Money
}

// Reduce function implementation for ExpressionInterface
func (s *Sum) Reduce(to string) *Money {
	amount := s.augend.Amount() + s.addend.Amount()
	return &Money{
		amount:   amount,
		currency: to,
	}
}
