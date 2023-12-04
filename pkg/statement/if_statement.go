package statement

import (
	"fmt"
	"tip-peg-parser-go/pkg/expression"
)

type IfStatement struct {
	Condition expression.Expression
	Body      BodyStatement
}

func (s *IfStatement) GetToken() StmToken {
	return StmToken{Pattern: IfExp}
}

func (s *IfStatement) String() string {
	return fmt.Sprintf("if ( %s ) %s", s.Condition, s.Body)
}

func (s *IfStatement) Put(statement Statement) {
	s.Body.Put(statement)
}

func (s *IfStatement) GetBodyStatement() []*Statement {
	return s.Body.GetBodyStatement()
}
