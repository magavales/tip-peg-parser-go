package expression

import "fmt"

type MultiExpression struct {
	BiExpression
}

func (e MultiExpression) String() string {
	return fmt.Sprintf("%s * %s", e.Left, e.Right)
}

func (e MultiExpression) GetToken() ExpToken {
	return ExpToken{Pattern: ExpMultiExp}
}
