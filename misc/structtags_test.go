package misc

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructJson1 struct {
	SomeInterestingUUIDUsedInTheCloud string   `json:"some_interesting_uuid_used_in_the_cloud"`
	SomeIDs                           []string `json:"some_ids"`
}

type StructJson2 struct {
	SomeField int `json:"somefield"`
}

type StructXML1 struct {
	XMLName xml.Name `xml:"struct_xml1"`
	Name    string   `xml:"name"`
}

func TestCheckJSONTags(t *testing.T) {
	assert := assert.New(t)

	err := CheckJSONTags(reflect.TypeOf((*StructJson1)(nil)), "id")
	assert.Nil(err)

	err = CheckJSONTags(reflect.TypeOf((*StructJson2)(nil)), "id")
	assert.NotNil(err)
	assert.Equal(err.Error(), "bad 'json' struct tag for misc.StructJson2.SomeField (some_field != somefield)")
}

func TestCheckXMLTags(t *testing.T) {
	assert := assert.New(t)

	err := CheckXMLTags(reflect.TypeOf((*StructXML1)(nil)), "id")
	assert.Nil(err)
}
