package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceJoin(t *testing.T) {
	cases := []struct {
		values   []interface{}
		sep      string
		expected string
	}{
		{[]interface{}{}, " ", ""},
		{[]interface{}{"a", "b", "c"}, ",", "[a,b,c]"},
		{[]interface{}{"a", "b", "c"}, " ", "[a b c]"},
		{[]interface{}{"a", "b", "c"}, "", "[abc]"},
		{[]interface{}{"a", 12, "c"}, ", ", "[a, 12, c]"},
		{[]interface{}{"15", 12, "c"}, ", ", "[15, 12, c]"},
		{[]interface{}{15, "12", 133}, ", ", "[15, 12, 133]"},
	}

	for _, c := range cases {
		actual := InterfaceJoin(c.values, c.sep)
		assert.Equal(t, c.expected, actual)
	}
}
