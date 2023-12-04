package statement

type Container interface {
	Put(statement Statement)
	GetBodyStatement() []*Statement
}
