# Pratt Parser

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_pratt-parser&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Dobefu_pratt-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dobefu/pratt-parser)](https://goreportcard.com/report/github.com/Dobefu/pratt-parser)

## Usage

- Run the application with the expression you wish to parse as the first argument, e.g.:

  ```bash
  go run main.go "1 + 1"
  ```

## Supported constants

- `PI` - π
- `TAU` - τ (2π)
- `E` - Euler's number
- `PHI` - Golden ratio

## Supported functions

- `abs(x)` - Absolute value of `x`
- `sin(x)` - Sine value of `x`
- `cos(x)` - Cosine value of `x`
- `tan(x)` - Tangent value of `x`
- `sqrt(x)` - Square root
- `round(x)` - Round `x` to the nearest integer value
- `floor(x)` - Round `x` down to the nearest integer value
- `ceil(x)` - Round `x` up to the nearest integer value
- `min(x, y)` - Get the smallest of the values provided
- `max(x, y)` - Get the largest of the values provided
