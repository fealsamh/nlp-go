package main

import (
	"fmt"

	"github.com/fealsamh/nlp-go/aibot"
	"github.com/fealsamh/nlp-go/aibot/ast"
)

func astCmd(cl *aibot.Client) {
	tokens, err := ast.Tokenise(`option = "abcd"`)
	if err != nil {
		exitWithError(err)
	}
	expr, err := ast.ParseExpr(tokens)
	if err != nil {
		exitWithError(err)
	}
	v, err := expr.Eval(map[string]ast.Value{
		"option": &ast.String{Value: "abcd"},
	})
	if err != nil {
		exitWithError(err)
	}
	fmt.Printf("%T '%s' %T %v\n", expr, expr, v, v.Interface())
}
