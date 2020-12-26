package multicurrency

type Dollar struct {
	amount int
}

func New(amount int) *Dollar {
	d := &Dollar{amount: amount}
	return d
}

func (d *Dollar) Times(multiplier int) {
	d.amount = d.amount * multiplier
}
