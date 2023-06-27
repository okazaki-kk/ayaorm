package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullString(t *testing.T) {
	var i NullString

	assert.Equal(t, "", i.String())
	assert.Equal(t, false, i.Valid())

	i.Set("foo")

	assert.Equal(t, "foo", i.String())
	assert.Equal(t, true, i.Valid())
}
