package atomic

import (
	"tip-peg-parser-go/pkg/expression"
)

type InputExpression struct {
}

func (e InputExpression) GetToken() expression.ExpToken {
	return expression.ExpToken{Pattern: expression.Input}
}

func (e InputExpression) String() string {
	return "input"
}
