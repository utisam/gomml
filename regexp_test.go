package gomml_test

import (
	"testing"

	"github.com/utisam/gomml"
)

func TestRegexp(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		value   string
		match   bool
	}{
		{
			pattern: "\\d+",
			value:   "1234",
			match:   true,
		},
		{
			pattern: "\\d+",
			value:   "abc",
			match:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gomml.Regexp(tt.pattern); !got.Matches(tt.value) == tt.match {
				t.Errorf("gomml.Regexp(%v).Matches(%v)", tt.pattern, tt.value)
			}
		})
	}
}

func Test_regexpMatcher_String(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		want    string
	}{
		{
			pattern: "^prefix-\\d+$",
			want:    "matches to the pattern of ^prefix-\\d+$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gomml.Regexp(tt.pattern).String(); got != tt.want {
				t.Errorf("regexpMatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
