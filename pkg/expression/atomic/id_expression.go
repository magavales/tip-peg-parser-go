package atomic

import "tip-peg-parser-go/pkg/expression"

type IdExpression struct {
	Id string
}

func (e IdExpression) GetToken() expression.ExpToken {
	return expression.ExpToken{Pattern: expression.Id}
}

func (e IdExpression) String() string {
	return e.Id
}
