package multicurrency

// Sum structure to return expression
type Sum struct {
	augend ExpressionInterface
	addend ExpressionInterface
}

// Reduce function implementation for ExpressionInterface
func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return &Money{
		amount:   amount,
		currency: to,
	}
}

// Plus function implementation for ExpressionInerface
func (s *Sum) Plus(n ExpressionInterface) ExpressionInterface {
	return &Sum{
		augend: s,
		addend: n,
	}
}

// Times returns the multiplications
func (s *Sum) Times(multiplier int) ExpressionInterface {
	return &Sum{
		augend: s.augend.Times(multiplier),
		addend: s.addend.Times(multiplier),
	}
}
