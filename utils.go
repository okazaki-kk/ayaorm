package ayaorm

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
}

func ToCamelCase(s string) string {
	var camel string
	for _, split := range strings.Split(s, "_") {
		c := []rune(split)
		c[0] = unicode.ToUpper(c[0])
		camel += string(c)
	}
	return camel
}

func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).Interface() == reflect.Zero(reflect.TypeOf(v)).Interface()
}
