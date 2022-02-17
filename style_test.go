package structenum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestStyleInit(t *testing.T) {
	assert.Equal(t, "snake_case", Styles.SnakeCase)
	assert.Equal(t, "CamelCase", Styles.CamelCase)
	assert.Equal(t, "lowercase", Styles.LowerCase)
}

func TestToSnakeCase(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"empty":              {input: "", want: ""},
		"only numbers":       {input: "123", want: "123"},
		"one word":           {input: "Foo", want: "foo"},
		"one word lowercase": {input: "foo", want: "foo"},
		"two words":          {input: "FooBar", want: "foo_bar"},
		"begins lowercase":   {input: "fooBar", want: "foo_bar"},
		"with numbers":       {input: "foo123Bar", want: "foo123_bar"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := toSnakeCase(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
