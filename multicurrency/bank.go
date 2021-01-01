package multicurrency

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
