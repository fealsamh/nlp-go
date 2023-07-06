package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	assert := assert.New(t)

	t.Run("string", func(t *testing.T) {
		tokens, err := Tokenise(`"abcd"`)
		assert.Nil(err)

		expr, err := ParseExpr(tokens)
		assert.Nil(err)

		ctx := new(EvalContext)

		v, err := expr.Eval(ctx)
		assert.Nil(err)
		assert.Equal("abcd", v.Interface())
	})

	t.Run("int", func(t *testing.T) {
		tokens, err := Tokenise(`1234`)
		assert.Nil(err)

		expr, err := ParseExpr(tokens)
		assert.Nil(err)

		ctx := new(EvalContext)

		v, err := expr.Eval(ctx)
		assert.Nil(err)
		assert.Equal(1234, v.Interface())
	})

	t.Run("float", func(t *testing.T) {
		tokens, err := Tokenise(`12.34`)
		assert.Nil(err)

		expr, err := ParseExpr(tokens)
		assert.Nil(err)

		ctx := new(EvalContext)

		v, err := expr.Eval(ctx)
		assert.Nil(err)
		assert.Equal(12.34, v.Interface())
	})

	t.Run("equality", func(t *testing.T) {
		tokens, err := Tokenise(`x = "abcd"`)
		assert.Nil(err)

		expr, err := ParseExpr(tokens)
		assert.Nil(err)

		ctx := new(EvalContext)
		ctx.Set("x", &String{Value: "abcd"})

		v, err := expr.Eval(ctx)
		assert.Nil(err)
		assert.Equal(true, v.Interface())
	})

	t.Run("inequality", func(t *testing.T) {
		tokens, err := Tokenise(`x /= "abcd"`)
		assert.Nil(err)

		expr, err := ParseExpr(tokens)
		assert.Nil(err)

		ctx := new(EvalContext)
		ctx.Set("x", &String{Value: "abcd"})

		v, err := expr.Eval(ctx)
		assert.Nil(err)
		assert.Equal(false, v.Interface())
	})
}
