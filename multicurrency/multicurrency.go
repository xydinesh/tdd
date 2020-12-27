package multicurrency

type Dollar struct {
	amount int
}

func New(amount int) *Dollar {
	d := &Dollar{amount: amount}
	return d
}

func (d *Dollar) Times(multiplier int) *Dollar {
	p := &Dollar{amount: d.amount * multiplier}
	return p
}

func (d *Dollar) Equals(n *Dollar) bool {
	return d.amount == n.amount
}
