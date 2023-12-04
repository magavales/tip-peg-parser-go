package expression

import "regexp"

type Pattern string

const (
	ExpPlusExp  Pattern = "(.*) \\+ (.*)"
	ExpMinusExp         = "(.*) \\- (.*)"
	ExpMultiExp         = "(.*) \\* (.*)"
	ExpDivExp           = "(.*) \\/ (.*)"
	ExpGrtExp           = "(.*) \\> (.*)"
	ExpEqExp            = "(.*) == (.*)"
	ExpBox              = "\\( (.*) \\)"
)

const (
	Id    Pattern = "([a-zA-Z]+)"
	Int           = "([\\-0-9]+)"
	Input         = "input"
)

type ExpToken struct {
	Pattern Pattern
}

func (token ExpToken) GetPattern() (*regexp.Regexp, error) {
	return regexp.Compile(string(token.Pattern))
}

func GetAllExpressionTokens() []ExpToken {
	var output []ExpToken
	output = append(output, ExpToken{Pattern: ExpPlusExp})
	output = append(output, ExpToken{Pattern: ExpMinusExp})
	output = append(output, ExpToken{Pattern: ExpMultiExp})
	output = append(output, ExpToken{Pattern: ExpDivExp})
	output = append(output, ExpToken{Pattern: ExpGrtExp})
	output = append(output, ExpToken{Pattern: ExpEqExp})
	output = append(output, ExpToken{Pattern: ExpBox})
	return output
}

func GetAllAtomicTokens() []ExpToken {
	var output []ExpToken
	output = append(output, ExpToken{Pattern: Id})
	output = append(output, ExpToken{Pattern: Int})
	output = append(output, ExpToken{Pattern: Input})
	return output
}
