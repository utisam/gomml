package gomml

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

type cmpMatcher struct {
	t    *testing.T
	want interface{}
	opts []cmp.Option
}

func Cmp(t *testing.T, want interface{}, opts ...cmp.Option) gomock.Matcher {
	return &cmpMatcher{t, want, opts}
}

func (m *cmpMatcher) Matches(x interface{}) bool {
	diff := cmp.Diff(m.want, x, m.opts...)
	if diff == "" {
		return true
	}

	m.t.Logf("(-want +got):\n%s", diff)
	return false
}

func (m *cmpMatcher) String() string {
	return fmt.Sprintf("is equal to %v with %s", m.want, cmp.Options(m.opts).String())
}
