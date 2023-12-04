package expression

import "fmt"

type EqualExpression struct {
	BiExpression
}

func (e EqualExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpEqExp}
}

func (e EqualExpression) String() string {
	return fmt.Sprintf("%s == %s", e.Left, e.Right)
}
