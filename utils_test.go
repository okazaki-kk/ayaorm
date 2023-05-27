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

func TestIsZero(t *testing.T) {
	cases := []struct {
		input interface{}
		want  bool
	}{
		{0, true},
		{1, false},
		{0.0, true},
		{0.1, false},
		{false, true},
		{true, false},
		{"", true},
		{"a", false},
		{struct{}{}, true},
		{struct{ A int }{A: 1}, false},
	}

	for _, tt := range cases {
		got := IsZero(tt.input)
		assert.Equal(t, tt.want, got)
	}
}
