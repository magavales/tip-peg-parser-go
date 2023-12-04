package statement

import "fmt"

type ElseStatement struct {
	Body BodyStatement
}

func (s *ElseStatement) GetToken() StmToken {
	return StmToken{Pattern: IfElse}
}

func (s *ElseStatement) String() string {
	return fmt.Sprintf("else %s", s.Body)
}

func (s *ElseStatement) Put(statement Statement) {
	s.Body.Put(statement)
}

func (s *ElseStatement) GetBodyStatement() []*Statement {
	return s.Body.GetBodyStatement()
}
