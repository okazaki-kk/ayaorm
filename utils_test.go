package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"Id", "id"},
		{"CreatedAt", "created_at"},
		{"UpdatedAt", "updated_at"},
		{"Content", "content"},
		{"Author", "author"},
		{"PostId", "post_id"},
		{"GreatManyThings", "great_many_things"},
		{"AboutGreatManyThings", "about_great_many_things"},
	}

	for _, tt := range cases {
		got := ToSnakeCase(tt.input)
		assert.Equal(t, tt.want, got)
	}
}

func TestToCamelCase(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"id", "Id"},
		{"created_at", "CreatedAt"},
		{"updated_at", "UpdatedAt"},
		{"content", "Content"},
		{"author", "Author"},
		{"post_id", "PostId"},
		{"great_many_things", "GreatManyThings"},
		{"about_great_many_things", "AboutGreatManyThings"},
	}

	for _, tt := range cases {
		got := ToCamelCase(tt.input)
		assert.Equal(t, tt.want, got)
	}
}

type TestIsZeroStruct struct {
	A int
	B string
	C float64
	D bool
	E int64
}

func TestIsZero(t *testing.T) {
	ti := TestIsZeroStruct{}
	assert.True(t, IsZero(ti.A))
	assert.True(t, IsZero(ti.B))
	assert.True(t, IsZero(ti.C))
	assert.True(t, IsZero(ti.D))
	assert.True(t, IsZero(ti.E))

	ti.A = 1
	ti.B = "a"
	ti.C = 1.0
	ti.D = true
	ti.E = 1
	assert.False(t, IsZero(ti.A))
	assert.False(t, IsZero(ti.B))
	assert.False(t, IsZero(ti.C))
	assert.False(t, IsZero(ti.D))
	assert.False(t, IsZero(ti.E))

	ti.A = 0
	ti.B = ""
	ti.C = 0.0
	ti.D = false
	ti.E = 0
	assert.True(t, IsZero(ti.A))
	assert.True(t, IsZero(ti.B))
	assert.True(t, IsZero(ti.C))
	assert.True(t, IsZero(ti.D))
	assert.True(t, IsZero(ti.E))
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	assert.True(t, Contains(s, 1))
	assert.True(t, Contains(s, 2))
	assert.True(t, Contains(s, 3))
	assert.True(t, Contains(s, 4))
	assert.True(t, Contains(s, 5))
	assert.False(t, Contains(s, 6))
	assert.False(t, Contains(s, 7))
	assert.False(t, Contains(s, 8))
	assert.False(t, Contains(s, 9))
	assert.False(t, Contains(s, 10))
}
