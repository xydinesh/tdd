package multicurrency

// Pair data structure to store currency exchanges
type Pair struct {
	from string
	to   string
}

// Equals compare two pairs
func (p *Pair) Equals(n *Pair) bool {
	return p.from == n.from && p.to == n.to
}
