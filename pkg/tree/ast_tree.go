package tree

type ASTree struct {
	Str string
}

func (t *ASTree) String() string {
	return t.Str
}
