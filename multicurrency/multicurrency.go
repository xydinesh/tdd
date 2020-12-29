package multicurrency

// Money structure to capture common functions
type Money struct {
	amount int
}

// Equals compare two objects
func (m *Money) Equals(n *Money) bool {
	return m.amount == n.amount
}

// Dollar structure for the book
type Dollar struct {
	Money
}

// NewDollar method to create new Dollar instance
func NewDollar(amount int) *Dollar {
	d := &Dollar{
		Money{amount: amount},
	}
	return d
}

// Times multiply dollar amount and return new instance
func (d *Dollar) Times(multiplier int) *Dollar {
	p := &Dollar{
		Money{amount: d.amount * multiplier},
	}
	return p
}

// Franc structure
type Franc struct {
	Money
}

// NewFranc method to create new instance
func NewFranc(amount int) *Franc {
	f := &Franc{
		Money{amount: amount},
	}
	return f
}

// Times multiply and return a new instance
func (f *Franc) Times(multiplier int) *Franc {
	n := &Franc{
		Money{amount: f.amount * multiplier},
	}
	return n
}
