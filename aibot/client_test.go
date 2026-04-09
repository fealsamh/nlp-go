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
		reflect.TypeFor[*Bot](),
		reflect.TypeFor[*Entity](),
		reflect.TypeFor[*Intent](),
		reflect.TypeFor[*State](),
		reflect.TypeFor[*Message](),
		reflect.TypeFor[*Next](),
		reflect.TypeFor[*User](),
		reflect.TypeFor[*IntentSimilarity](),
		reflect.TypeFor[*deployRequest](),
		reflect.TypeFor[*recogniseIntentRequest](),
		reflect.TypeFor[*signinRequest](),
		reflect.TypeFor[*signinRequest](),
		reflect.TypeFor[*getBotResponse](),
		reflect.TypeFor[*listBotsResponse](),
		reflect.TypeFor[*recogniseIntentResponse](),
		reflect.TypeFor[*signinResponse](),
	} {
		err := misc.CheckJSONTags(t, "id")
		assert.Nil(err)
	}
}
