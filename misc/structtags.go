package misc

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fealsamh/nlp-go/utils/slices"
)

func CheckStructTags(tagAttr string, typ reflect.Type, terms ...string) error {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if slices.Contains([]string{"0", "no", "false"}, f.Tag.Get("check")) {
			continue
		}
		tag := f.Tag.Get(tagAttr)
		if tag == "" {
			return fmt.Errorf("nil '%s' struct tag for %s.%s", tagAttr, typ, f.Name)
		}
		if !f.Anonymous {
			tagName := strings.Split(tag, ",")[0]
			if f.Name == "XMLName" {
				typName := SnakecasedFromCamelcased(typ.Name())
				comps := strings.Split(tagName, " ")
				tagName = comps[len(comps)-1]
				if strings.ToLower(typName) != tagName {
					return fmt.Errorf("bad '%s' struct tag for %s.XMLName (%s != %s)", tagAttr, typ, typName, tagName)
				}
			} else if tagName != "-" {
				snakeField := SnakecasedFromCamelcased(f.Name, terms...)
				if snakeField != tagName {
					return fmt.Errorf("bad '%s' struct tag for %s.%s (%s != %s)", tagAttr, typ, f.Name, snakeField, tagName)
				}
			}
		}
	}
	return nil
}

func CheckJSONTags(typ reflect.Type, terms ...string) error {
	return CheckStructTags("json", typ, terms...)
}

func CheckXMLTags(typ reflect.Type, terms ...string) error {
	return CheckStructTags("xml", typ, terms...)
}
