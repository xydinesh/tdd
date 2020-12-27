package multicurrency

// Dollar structure for the book
type Dollar struct {
	amount int
}

// New method to create new Dollar instance
func New(amount int) *Dollar {
	d := &Dollar{amount: amount}
	return d
}

// Times multiply dollar amount and return new instance
func (d *Dollar) Times(multiplier int) *Dollar {
	p := &Dollar{amount: d.amount * multiplier}
	return p
}

// Equals compare two objects for the value
func (d *Dollar) Equals(n *Dollar) bool {
	return d.amount == n.amount
}
