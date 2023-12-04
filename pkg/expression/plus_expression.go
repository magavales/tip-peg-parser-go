package expression

import "fmt"

type PlusExpression struct {
	BiExpression
}

func (e PlusExpression) String() string {
	return fmt.Sprintf("%s + %s", e.Left, e.Right)
}

func (e PlusExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpPlusExp}
}
