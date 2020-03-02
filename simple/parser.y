%{
package simple
import (
	"github.com/lovewin99/goyacc/simple/ast"
)
%}

%union {
	offset int // offset
	item interface{}
	ident string
	statement ast.StmtNode
}

%token	<ident>
	/*yy:token "%c"     */	identifier      "identifier"
	/*yy:token "_%c"    */  underscoreCS	"UNDERSCORE_CHARSET"
	/*yy:token "\"%c\"" */	stringLit       "string literal"
	selectKwd		"SELECT"
	from			"FROM"

%type   <item>
	StatementList		"statement list"
	Field			"field expression"
	TableRefsClause		"Table references clause"

%type	<statement>
	SelectStmt		"SELECT statement"
	Statement		"statement"

%start	Start

%%
Start:
	StatementList

StatementList:
	Statement
	{
		if $1 != nil {
			s := $1
//			if lexer, ok := yylex.(stmtTexter); ok {
//				s.SetText(lexer.stmtText())
//			}
			parser.result = append(parser.result, s)
		}
	}
//|	StatementList ';' Statement
//	{
//		if $3 != nil {
//			s := $3
//			if lexer, ok := yylex.(stmtTexter); ok {
//				s.SetText(lexer.stmtText())
//			}
//			parser.result = append(parser.result, s)
//		}
//	}

Field:
	'*'
	{

	}

TableRefsClause:
	identifier
	{
		$$ = $1
	}

SelectStmt:
	"SELECT" Field "FROM" TableRefsClause
	{
		$$ = &ast.SelectStmt{From: $4.(string)}
	}

Statement:
	SelectStmt
	{
		$$ = $1.(*ast.SelectStmt)
	}

%%