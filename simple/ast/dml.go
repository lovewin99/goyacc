package ast

type StmtNode interface {
	Text() string
}

type Node interface {
	Text() string
}

type SelectStmt struct {
	Node
	From string
}

func (s *SelectStmt) Text() string {
	return s.From
}
