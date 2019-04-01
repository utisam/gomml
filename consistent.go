package gomml

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

type consistentMatcher struct {
	m       gomock.Matcher
	decided bool
	val     interface{}
}

// Consistent is a Matcher that matches with consistent.
//
// The value must be the time between the constructor was called and
// it is validated.
//
// Example usage:
//
// 		consistentNow := gomml.Consistent(gomml.Now())
// 		dbMock.EXPECT().
// 			Insert(consistentNow, consistentNow).
//
func Consistent(o interface{}) gomock.Matcher {
	m, ok := o.(gomock.Matcher)
	if !ok {
		m = gomock.Eq(o)
	}
	return &consistentMatcher{m: m}
}

func (m *consistentMatcher) Matches(x interface{}) bool {
	if m.decided {
		return reflect.DeepEqual(m.val, x)
	}

	m.decided = true
	m.val = x
	return m.m.Matches(x)
}

func (m *consistentMatcher) String() string {
	if m.decided {
		return fmt.Sprintf("%v and must be consistent with %v", m.m.String(), m.val)
	}
	return fmt.Sprintf("%v and must be consistent", m.m.String())
}
