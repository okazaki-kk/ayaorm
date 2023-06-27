package null

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullFloat64(t *testing.T) {
	var i NullFloat

	assert.Equal(t, "", fmt.Sprint(i))
	assert.Equal(t, false, i.Valid())

	i.Set(4.23)

	assert.Equal(t, "4.23", fmt.Sprint(i))
	assert.Equal(t, true, i.Valid())
}
