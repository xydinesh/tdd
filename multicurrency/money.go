package multicurrency

import "fmt"

// ExpressionInterface interface
type ExpressionInterface interface {
	Reduce(bank *Bank, to string) *Money
	Plus(addend ExpressionInterface) ExpressionInterface
	Times(multiplier int) ExpressionInterface
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

// Equals returns a bool
func (m *Money) Equals(n *Money) bool {
	return m.Amount() == n.Amount() && m.Currency() == n.Currency()
}

// Strinig return a string value
func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount(), m.Currency())
}

// Currency retuns the currenecy
func (m *Money) Currency() string {
	return m.currency
}

// Times returns the multiplications
func (m *Money) Times(multiplier int) ExpressionInterface {
	return &Money{
		amount:   m.Amount() * multiplier,
		currency: m.Currency(),
	}
}

// Plus adds two and return a new one
func (m *Money) Plus(n ExpressionInterface) ExpressionInterface {
	f, ok := n.(*Money)
	if ok && m.Currency() == f.Currency() {
		return &Money{
			amount:   m.Amount() + f.Amount(),
			currency: m.Currency(),
		}
	}

	return &Sum{
		augend: m,
		addend: n,
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
