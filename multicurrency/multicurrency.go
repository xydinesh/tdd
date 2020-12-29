package multicurrency

// Dollar structure for the book
type Dollar struct {
	amount int
}

// NewDollar method to create new Dollar instance
func NewDollar(amount int) *Dollar {
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

// Franc structure
type Franc struct {
	amount int
}

// NewFranc method to create new instance
func NewFranc(amount int) *Franc {
	f := &Franc{amount: amount}
	return f
}

// Times multiply and return a new instance
func (f *Franc) Times(multiplier int) *Franc {
	n := &Franc{amount: f.amount * multiplier}
	return n
}

// Equals compare two objects
func (f *Franc) Equals(n *Franc) bool {
	return f.amount == n.amount
}
