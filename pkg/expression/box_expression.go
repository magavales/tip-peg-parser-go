package expression

import "fmt"

type BoxExpression struct {
	Expression Expression
}

func (e BoxExpression) String() string {
	return fmt.Sprintf("( %s )", e.Expression)
}

func (e BoxExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpBox}
}
