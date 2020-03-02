package simple

import (
	"fmt"
	"github.com/lovewin99/goyacc/simple/ast"
)

// Parser represents a parser instance. Some temporary objects are stored in it to reduce object allocation during Parse function.
type Parser struct {
	charset   string
	collation string
	result    []ast.StmtNode
	src       string
	//lexer     Scanner

	// the following fields are used by yyParse to reduce allocation.
	cache  []yySymType
	yylval yySymType
	yyVAL  yySymType
}

func Parse() {
	fmt.Println("haha1")
	sql := "select * from haha"
	p := &Parser{cache: make([]yySymType, 200)}
	var l yyLexer
	l = &Scanner{}
	p.src = sql
	yyParse(l, p)
	t := p.result[0].Text()
	fmt.Printf("t == %v\n", t)
	fmt.Println(p)
}
