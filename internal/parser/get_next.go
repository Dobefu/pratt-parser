package parser

import "errors"

// GetNext gets the next byte in the expression.
func (p *Parser) GetNext() (byte, error) {
	if p.isEOF {
		return 0, errors.New("cannot get next byte after EOF")
	}

	next := p.exp[p.expIdx]
	p.expIdx++

	if p.expIdx >= p.expLen {
		p.isEOF = true
	}

	return next, nil
}
