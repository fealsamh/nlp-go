package misc

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructJSON1 struct {
	SomeInterestingUUIDUsedInTheCloud string   `json:"some_interesting_uuid_used_in_the_cloud"`
	SomeIDs                           []string `json:"some_ids"`
}

type StructJSON2 struct {
	SomeField   int    `json:"somefield"`
	NonMatching string `json:"nonmatching" check:"no"`
}

type StructXML1 struct {
	XMLName     xml.Name `xml:"http://www.w3.org/2005/Atom struct_xml1"`
	Name        string   `xml:"name"`
	NonMatching string   `xml:"nonmatching" check:"no"`
	Description string   `check:"no"`
}

type StructXML2 struct {
	XMLName     xml.Name `xml:"struct_xml2"`
	Name        string   `xml:"name"`
	Description string   `check:"no"`
}

func TestCheckJSONTags(t *testing.T) {
	assert := assert.New(t)

	err := CheckJSONTags(reflect.TypeOf((*StructJSON1)(nil)), "id")
	assert.Nil(err)

	err = CheckJSONTags(reflect.TypeOf((*StructJSON2)(nil)), "id")
	assert.NotNil(err)
	assert.Equal(err.Error(), "bad 'json' struct tag for misc.StructJSON2.SomeField (some_field != somefield)")
}

func TestCheckXMLTags(t *testing.T) {
	assert := assert.New(t)

	err := CheckXMLTags(reflect.TypeOf((*StructXML1)(nil)))
	assert.Nil(err)

	err = CheckXMLTags(reflect.TypeOf((*StructXML2)(nil)))
	assert.Nil(err)
}
