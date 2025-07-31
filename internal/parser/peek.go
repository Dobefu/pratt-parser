package parser

import "errors"

// Peek gets the byte of the expression at the current index without advancing it.
func (p *Parser) Peek() (byte, error) {
	if p.isEOF {
		return 0, errors.New("cannot get next byte after EOF")
	}

	return p.exp[p.expIdx], nil
}
