package statement

import (
	"fmt"
	"tip-peg-parser-go/pkg/expression"
)

type EqualStatement struct {
	Id         string
	Expression expression.Expression
}

func (s EqualStatement) GetToken() StmToken {
	return StmToken{Pattern: IdEqExp}
}

func (s EqualStatement) String() string {
	return fmt.Sprintf("%s = %s ; ", s.Id, s.Expression)
}
