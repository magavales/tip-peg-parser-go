package statement

type Pattern string

type Statement interface {
	GetToken() StmToken
}
