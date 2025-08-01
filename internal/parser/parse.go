package parser

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() error {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return err
	}

	p.tokens = tokens

	return nil
}
