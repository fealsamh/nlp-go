package misc

import (
	"fmt"
	"reflect"
	"strings"
)

func CheckStructTags(tagAttr string, typ reflect.Type, terms ...string) error {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		tag := f.Tag.Get(tagAttr)
		if tag == "" {
			return fmt.Errorf("nil '%s' struct tag for %s.%s", tagAttr, typ, f.Name)
		}
		if !f.Anonymous {
			tagName := strings.Split(tag, ",")[0]
			if tagName != "-" {
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
