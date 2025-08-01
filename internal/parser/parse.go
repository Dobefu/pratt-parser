package parser

// Parse parses the expression string supplied in the struct.
func (p *Parser) Parse() error {
	tokens, err := p.tokenizer.Tokenize()

	if err != nil {
		return err
	}

	p.tokens = tokens
	p.tokenLen = len(tokens)

	for !p.isEOF {
		_, err := p.GetNextToken()

		if err != nil {
			return err
		}
	}

	return nil
}
