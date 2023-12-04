package expression

import "fmt"

type GreaterExpression struct {
	BiExpression
}

func (e GreaterExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpGrtExp}
}

func (e GreaterExpression) String() string {
	return fmt.Sprintf("%s > %s", e.Left, e.Right)
}
