package structenum

import "reflect"

func init() {
	Make(&Styles)
}

func Make(v interface{}, style ...string) {
	// Do nothing if given value is not a pointer
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return
	}

	// Dereference pointer and do nothing if it is not a struct
	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	// Get type of struct
	elemType := elem.Type()

	// Loop through struct fields
	for i := 0; i < elemType.NumField(); i++ {
		elemField := elem.Field(i)
		switch elemField.Kind() {
		case reflect.String:
			// Set string value according to tag
			structField := elemType.Field(i)
			setString(elemField, structField, style...)
		case reflect.Struct:
			// Recursively parse nested struct
			Make(elemField.Addr().Interface(), style...)
		default:
			// Ignore other types
		}
	}
}

func setString(elemField reflect.Value, structField reflect.StructField, style ...string) {
	tag := string(structField.Tag)

	// If tag is empty, default to field name
	if tag == "" {
		tag = structField.Name
	}

	// If style was given, apply it to tag
	if len(style) > 0 {
		tag = convert(tag, style[0])
	}

	// Set struct field value to tag
	elemField.Set(reflect.ValueOf(tag))
}
