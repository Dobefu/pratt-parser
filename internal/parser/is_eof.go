package parser

// IsEOF checks if the expression has ended.
func (p *Parser) IsEOF() bool {
	return p.isEOF
}
