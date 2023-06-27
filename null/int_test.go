package null

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullInt64Stringer(t *testing.T) {
	var i NullInt

	assert.Equal(t, "", fmt.Sprint(i))
	assert.Equal(t, false, i.Valid())

	i.Set(3)

	assert.Equal(t, "3", fmt.Sprint(i))
	assert.Equal(t, true, i.Valid())
}
