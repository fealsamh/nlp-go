package aibot

import (
	"reflect"
	"testing"

	"github.com/fealsamh/nlp-go/misc"
	"github.com/stretchr/testify/assert"
)

func TestStructTags(t *testing.T) {
	assert := assert.New(t)

	for _, t := range []reflect.Type{
		reflect.TypeOf((*Bot)(nil)),
		reflect.TypeOf((*Entity)(nil)),
		reflect.TypeOf((*Intent)(nil)),
		reflect.TypeOf((*State)(nil)),
		reflect.TypeOf((*Message)(nil)),
		reflect.TypeOf((*Next)(nil)),
		reflect.TypeOf((*User)(nil)),
		reflect.TypeOf((*IntentSimilarity)(nil)),
		reflect.TypeOf((*deployRequest)(nil)),
		reflect.TypeOf((*recogniseIntentRequest)(nil)),
		reflect.TypeOf((*signinRequest)(nil)),
		reflect.TypeOf((*signinRequest)(nil)),
		reflect.TypeOf((*getBotResponse)(nil)),
		reflect.TypeOf((*listBotsResponse)(nil)),
		reflect.TypeOf((*recogniseIntentResponse)(nil)),
		reflect.TypeOf((*signinResponse)(nil)),
	} {
		err := misc.CheckJSONTags(t, "id")
		assert.Nil(err)
	}
}
