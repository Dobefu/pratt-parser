package evaluator

import "math"

var identifierRegistry = map[string]identifierInfo{
	"PI": {
		handler: func() (float64, error) {
			return math.Pi, nil
		},
	},
	"TAU": {
		handler: func() (float64, error) {
			return math.Pi * 2, nil
		},
	},
	"E": {
		handler: func() (float64, error) {
			return math.E, nil
		},
	},
	"PHI": {
		handler: func() (float64, error) {
			return math.Phi, nil
		},
	},
}
