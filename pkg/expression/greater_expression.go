package expression

import "fmt"

type GreaterExpression struct {
	BiExpression
}

func (e GreaterExpression) String() string {
	return fmt.Sprintf("%s > %s", e.Left, e.Right)
}
