package tokeniser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokeniser(t *testing.T) {
	tokens := new(Tokeniser).Tokenise("abcd-efgh ijkl'mnop")
	s := make([]string, 0, len(tokens))
	for _, t := range tokens {
		s = append(s, t.String())
	}

	assert := assert.New(t)
	assert.Equal(s, []string{"abcd/WRD", "-/SYM", "efgh/WRD", "ijkl/WRD", "'/SYM", "mnop/WRD", "/EOF"})
}
