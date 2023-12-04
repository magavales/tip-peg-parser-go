package statement

import (
	"regexp"
)

const (
	IfExp     Pattern = "if \\( (.*) \\)"
	IfElse            = "else "
	WhileExp          = "while \\( (.*) \\)"
	EndStm            = "\\{ \\}"
	IdEqExp           = "([a-zA-Z]+) \\= (.*) ;"
	OutputExp         = "(output) (.*) ;"
)

type StmToken struct {
	Pattern Pattern
}

func (token StmToken) GetPattern() (*regexp.Regexp, error) {
	return regexp.Compile(string(token.Pattern))
}

func GetAllStatementTokens() []StmToken {
	var output []StmToken
	output = append(output, StmToken{Pattern: IfExp})
	output = append(output, StmToken{Pattern: IfElse})
	output = append(output, StmToken{Pattern: WhileExp})
	output = append(output, StmToken{Pattern: EndStm})
	output = append(output, StmToken{Pattern: IdEqExp})
	output = append(output, StmToken{Pattern: OutputExp})
	return output
}
