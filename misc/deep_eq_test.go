package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DeepEqStruct struct {
	S string
	N int
	I interface{}
}

type BadDeepEqStruct struct {
	Ch chan string
}

func TestDeepEqualitySuccess(t *testing.T) {
	assert := assert.New(t)

	eq, reason := IsDeeplyEqual(&DeepEqStruct{S: "abcd", N: 1234, I: 12.34}, &DeepEqStruct{S: "abcd", N: 1234, I: 12.34})
	assert.Equal(true, eq)
	assert.Nil(reason)
}

func TestDeepEqualityFailure_DifferentValues(t *testing.T) {
	assert := assert.New(t)

	eq, reason := IsDeeplyEqual(&DeepEqStruct{S: "abcd", N: 1234, I: 12.34}, &DeepEqStruct{S: "abcd", N: 1234, I: 12.35})
	assert.Equal(false, eq)
	assert.NotNil(reason)
	assert.Equal(false, reason.IsError)
	assert.Equal("float64s not equal (12.34, 12.35)", reason.String())
}

func TestDeepEqualityFailure_DifferentTypes(t *testing.T) {
	assert := assert.New(t)

	eq, reason := IsDeeplyEqual(&DeepEqStruct{S: "abcd", N: 1234, I: 12.34}, &DeepEqStruct{S: "abcd", N: 1234, I: true})
	assert.Equal(false, eq)
	assert.NotNil(reason)
	assert.Equal(false, reason.IsError)
	assert.Equal("different types (float64, bool)", reason.String())
}

func TestDeepEqualityFailure_UnknownType(t *testing.T) {
	assert := assert.New(t)

	eq, reason := IsDeeplyEqual(&BadDeepEqStruct{}, &BadDeepEqStruct{})
	assert.Equal(false, eq)
	assert.NotNil(reason)
	assert.Equal(true, reason.IsError)
	assert.Equal("failed to deeply compare values of type 'chan string'", reason.String())
}
