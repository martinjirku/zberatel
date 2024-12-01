package db

import (
	"fmt"
	"reflect"
	"strings"
)

func GetDbColumnByJsonField(input interface{}, jsonTag string) (string, any, error) {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("expected a struct or a pointer to struct, got %s", t.Kind())
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}

		// The json tag can have options, e.g., `json:"name,omitempty"`
		tagParts := strings.Split(tag, ",")
		currentJSONTag := tagParts[0]

		if currentJSONTag == "-" {
			continue
		}

		if currentJSONTag == jsonTag {
			dbTag := field.Tag.Get("db")
			if dbTag == "" {
				return "", nil, fmt.Errorf("db tag not found for field with json tag '%s'", jsonTag)
			}
			field := reflect.ValueOf(input).Field(i)
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			if !field.IsValid() {
				return dbTag, nil, nil
			}
			var fieldValue = field.Interface()
			return dbTag, fieldValue, nil
		}
	}

	return "", nil, fmt.Errorf("no field with json tag '%s' found", jsonTag)

}
