package ast

type StmtNode interface {
	Text() string
}

type SelectStmt struct {
	From string
}

func (s *SelectStmt) Text() string {
	return s.From
}
