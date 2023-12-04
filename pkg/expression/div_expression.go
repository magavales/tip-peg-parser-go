package expression

import "fmt"

type DivExpression struct {
	BiExpression
}

func (e DivExpression) String() string {
	return fmt.Sprintf("%s / %s", e.Left, e.Right)
}

func (e DivExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpDivExp}
}
