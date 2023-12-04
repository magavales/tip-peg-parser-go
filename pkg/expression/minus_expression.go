package expression

import "fmt"

type MinusExpression struct {
	BiExpression
}

func (e MinusExpression) String() string {
	return fmt.Sprintf("%s - %s", e.Left, e.Right)
}

func (e MinusExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpMinusExp}
}
