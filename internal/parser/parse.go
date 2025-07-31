package parser

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() error {
	_, err := p.Tokenize()

	if err != nil {
		return err
	}

	return nil
}
