package structenum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Example struct {
	TaggedField   string "tagged_field_123"
	UntaggedField string
	NestedStruct  struct {
		TaggedField   string "tagged_field_123"
		UntaggedField string
		NestedStruct  struct {
			TaggedField   string "tagged_field_123"
			UntaggedField string
		}
	}
}

func TestStructTagged(t *testing.T) {
	assert.NotPanics(t, func() {
		Make(&Example)
	})
	tests := map[string]string{
		"struct":               Example.TaggedField,
		"nested struct":        Example.NestedStruct.TaggedField,
		"double nested struct": Example.NestedStruct.NestedStruct.TaggedField,
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc, "tagged_field_123")
		})
	}
}

func TestStructUntagged(t *testing.T) {
	assert.NotPanics(t, func() {
		Make(&Example)
	})
	tests := map[string]string{
		"struct":               Example.UntaggedField,
		"nested struct":        Example.NestedStruct.UntaggedField,
		"double nested struct": Example.NestedStruct.NestedStruct.UntaggedField,
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc, "UntaggedField")
		})
	}
}

func TestNoPanic(t *testing.T) {
	tests := map[string]interface{}{
		"nil":          nil,
		"struct":       Example,
		"empty struct": struct{}{},
		"int":          123,
		"float":        123.45,
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				Make(tc)
				Make(&tc)
			})
		})
	}
}

func TestNoTagStyles(t *testing.T) {
	tests := map[string]string{
		Styles.CamelCase: "UntaggedField",
		Styles.SnakeCase: "untagged_field",
		Styles.LowerCase: "untaggedfield",
	}
	for style, tc := range tests {
		t.Run(style, func(t *testing.T) {
			assert.NotPanics(t, func() {
				Make(&Example, style)
			})
			assert.Equal(t, tc, Example.UntaggedField)
		})
	}
}
