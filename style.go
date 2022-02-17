package structenum

import (
	"regexp"
	"strings"
)

var Styles struct {
	CamelCase string "CamelCase"
	SnakeCase string "snake_case"
	LowerCase string "lowercase"
}

func toSnakeCase(str string) string {
	snakeCaseStr := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1}_${2}")
	snakeCaseStr = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snakeCaseStr, "${1}_${2}")
	return strings.ToLower(snakeCaseStr)
}

func convert(str, style string) string {
	switch style {
	case Styles.CamelCase:
		return str
	case Styles.SnakeCase:
		return toSnakeCase(str)
	case Styles.LowerCase:
		return strings.ToLower(str)
	default:
		return str
	}
}
