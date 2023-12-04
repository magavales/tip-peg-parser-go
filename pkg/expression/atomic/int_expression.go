package atomic

import (
	"tip-peg-parser-go/pkg/expression"
)

type IntExpression struct {
	Value string
}

func (e IntExpression) GetToken() expression.ExpToken {
	return expression.ExpToken{Pattern: expression.Int}
}

func (e IntExpression) String() string {
	return e.Value
}
