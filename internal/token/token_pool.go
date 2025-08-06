package token

import (
	"sync"
)

// Pool provides a thread-safe pool of tokens.
type Pool struct {
	pool     map[string]*Token
	poolSize int
	mu       sync.RWMutex
}

// NewPool creates a new token pool with some pre-allocated common tokens.
func NewPool() *Pool {
	p := &Pool{
		pool:     make(map[string]*Token),
		poolSize: 0,
		mu:       sync.RWMutex{},
	}

	commonTokens := []struct {
		atom      string
		tokenType Type
	}{
		{"+", TokenTypeOperationAdd},
		{"-", TokenTypeOperationSub},
		{"*", TokenTypeOperationMul},
		{"/", TokenTypeOperationDiv},
		{"%", TokenTypeOperationMod},
		{"**", TokenTypeOperationPow},
		{"(", TokenTypeLParen},
		{")", TokenTypeRParen},
		{",", TokenTypeComma},
	}

	for _, t := range commonTokens {
		p.pool[t.atom] = NewToken(t.atom, t.tokenType)
		p.poolSize++
	}

	return p
}

// GetToken gets an existing token if it exists, or creates a new one.
func (p *Pool) GetToken(atom string, tokenType Type) *Token {
	if token, exists := p.getFromPool(atom); exists {
		return token
	}

	return NewToken(atom, tokenType)
}

// GetPoolSize returns the current size of the pool.
func (p *Pool) GetPoolSize() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.poolSize
}

func (p *Pool) getFromPool(atom string) (*Token, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	token, exists := p.pool[atom]

	return token, exists
}
