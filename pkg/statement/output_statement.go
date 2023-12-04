package statement

import (
	"fmt"
	"tip-peg-parser-go/pkg/expression"
)

type OutputStatement struct {
	Operation  string
	Expression expression.Expression
}

func (s OutputStatement) GetToken() StmToken {
	return StmToken{Pattern: OutputExp}
}

func (s OutputStatement) String() string {
	return fmt.Sprintf("%s %s; ", s.Operation, s.Expression)
}
