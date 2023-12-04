package statement

import "fmt"

type RootStatement struct {
	Body BodyStatement
}

func (s *RootStatement) String() string {
	return fmt.Sprintf("%s", s.Body.String())
}

func (s *RootStatement) Put(statement Statement) {
	s.Body.Put(statement)
}

func (s *RootStatement) GetBodyStatement() []*Statement {
	return s.Body.GetBodyStatement()
}

func (s *RootStatement) GetToken() StmToken {
	return StmToken{}
}
