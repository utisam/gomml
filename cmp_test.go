package gomml_test

import (
	"testing"

	"github.com/utisam/gomml"
)

func TestCmp(t *testing.T) {
	tests := []struct {
		name  string
		x     interface{}
		y     interface{}
		match bool
	}{
		{
			x: map[int]int{
				1: 1,
				2: 2,
			},
			y: map[int]int{
				1: 1,
				2: 2,
			},
			match: true,
		},
		{
			x: map[int]int{
				1: 1,
				2: 2,
			},
			y: map[int]int{
				1: 1,
			},
			match: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gomml.Cmp(t, tt.x); !got.Matches(tt.y) == tt.match {
				t.Errorf("gomml.Cmp(%v).Matches(%v)", tt.x, tt.y)
			}
		})
	}
}
