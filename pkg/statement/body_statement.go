package statement

import (
	"fmt"
	"strings"
)

type BodyStatement struct {
	Statements []*Statement
}

func (s *BodyStatement) String() string {
	var sb strings.Builder
	sb.WriteString("{ ")
	for _, t := range s.Statements {
		sb.WriteString(fmt.Sprintf("%s", *t))
	}
	sb.WriteString(" }")
	return sb.String()
}

func (s *BodyStatement) Put(statement Statement) {
	s.Statements = append(s.Statements, &statement)
}

func (s *BodyStatement) GetBodyStatement() []*Statement {
	return s.Statements
}
