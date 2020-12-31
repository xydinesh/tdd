package multicurrency

import "fmt"

// MoneyInterface interface
type MoneyInterface interface {
	Amount() int
	Currency() string
}

// ExpressionInterface interface
type ExpressionInterface interface {
	Reduce(bank *Bank, to string) *Money
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
func (m *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(m.Currency(), to)
	return &Money{
		amount:   m.amount / rate,
		currency: to,
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
	rates map[Pair]int
}

// NewBank creates an instance of a Bank structure
func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]int),
	}
}

// Reduce use the expression
func (b *Bank) Reduce(ex ExpressionInterface, to string) *Money {
	return ex.Reduce(b, to)
}

// AddRate to add different rates
func (b *Bank) AddRate(from string, to string, rate int) {
	b.rates[Pair{
		from: from,
		to:   to,
	}] = rate
}

// Rate return exchange rate for from -> to
func (b *Bank) Rate(from string, to string) int {
	if v, ok := b.rates[Pair{
		from: from,
		to:   to,
	}]; ok {
		return v
	}
	return 1
}

// Sum structure to return expression
type Sum struct {
	augend *Money
	addend *Money
}

// Reduce function implementation for ExpressionInterface
func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.Amount() + s.addend.Amount()
	return &Money{
		amount:   amount,
		currency: to,
	}
}

// Pair data structure to store currency exchanges
type Pair struct {
	from string
	to   string
}

// Equals compare two pairs
func (p *Pair) Equals(n *Pair) bool {
	return p.from == n.from && p.to == n.to
}
