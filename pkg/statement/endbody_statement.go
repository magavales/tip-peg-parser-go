package statement

type EndBodyStatement struct {
}

func (s EndBodyStatement) GetToken() StmToken {
	return StmToken{Pattern: EndStm}
}

func (s EndBodyStatement) String() string {
	return " }"
}
