// CAUTION: Generated file - DO NOT EDIT.

package simple

import __yyfmt__ "fmt"

import (
	"github.com/lovewin99/goyacc/simple/ast"
)

type yySymType struct {
	yys       int
	offset    int // offset
	item      interface{}
	ident     string
	statement ast.StmtNode
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault    = 57351
	yyEOFCode    = 57344
	yyErrCode    = 57345
	from         = 57350
	identifier   = 57346
	selectKwd    = 57349
	stringLit    = 57348
	underscoreCS = 57347

	yyMaxDepth = 200
	yyTabOfs   = -7
)

var (
	yyXLAT = map[int]int{
		57344: 0,  // $end (6x)
		57350: 1,  // from (2x)
		42:    2,  // '*' (1x)
		57352: 3,  // Field (1x)
		57346: 4,  // identifier (1x)
		57349: 5,  // selectKwd (1x)
		57353: 6,  // SelectStmt (1x)
		57354: 7,  // Start (1x)
		57355: 8,  // Statement (1x)
		57356: 9,  // StatementList (1x)
		57357: 10, // TableRefsClause (1x)
		57351: 11, // $default (0x)
		57345: 12, // error (0x)
		57348: 13, // stringLit (0x)
		57347: 14, // underscoreCS (0x)
	}

	yySymNames = []string{
		"$end",
		"from",
		"'*'",
		"Field",
		"identifier",
		"selectKwd",
		"SelectStmt",
		"Start",
		"Statement",
		"StatementList",
		"TableRefsClause",
		"$default",
		"error",
		"stringLit",
		"underscoreCS",
	}

	yyReductions = []struct{ xsym, components int }{
		{0, 1},
		{7, 1},
		{9, 1},
		{3, 1},
		{10, 1},
		{6, 4},
		{8, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [11][]uint8{
		// 0
		{5: 11, 12, 8, 10, 9},
		{7},
		{6},
		{5},
		{2: 13, 14},
		// 5
		{1},
		{1: 4},
		{1: 15},
		{4: 16, 10: 17},
		{3},
		// 10
		{2},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Errorf(format string, a ...interface{}) error
	AppendError(err error)
	Errors() (warns []error, errs []error)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEOFCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer, parser *Parser) int {
	const yyError = 12

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	parser.yylval = yySymType{}
	parser.yyVAL = yySymType{}
	yyS := parser.cache

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
		parser.cache = yyS
	}
	yyS[yyp] = parser.yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &parser.yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		parser.yyVAL = parser.yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if !ok || msg == "" {
				msg = "syntax error"
			}
			// ignore goyacc error message
			yylex.AppendError(yylex.Errorf(""))
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEOFCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
		parser.cache = yyS
	}
	parser.yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 2:
		{
			if yyS[yypt-0].statement != nil {
				s := yyS[yypt-0].statement
				//			if lexer, ok := yylex.(stmtTexter); ok {
				//				s.SetText(lexer.stmtText())
				//			}
				parser.result = append(parser.result, s)
			}
		}
	case 4:
		{
			parser.yyVAL.item = yyS[yypt-0].ident
		}
	case 5:
		{
			parser.yyVAL.statement = &ast.SelectStmt{From: yyS[yypt-0].item.(string)}
		}
	case 6:
		{
			parser.yyVAL.statement = yyS[yypt-0].statement.(*ast.SelectStmt)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &parser.yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
