package evaluator

import "math"

var identifierRegistry = map[string]identifierInfo{
	"pi": {
		handler: func() (float64, error) {
			return math.Pi, nil
		},
	},
}
