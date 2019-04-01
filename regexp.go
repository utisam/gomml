package gomml

import (
	"fmt"
	"regexp"

	"github.com/golang/mock/gomock"
)

type regexpMatcher struct {
	re *regexp.Regexp
}

// Regexp is a Matcher that matches with regexp.
//
// Example usage:
//
// 		dbMock.EXPECT().
// 			Insert(gomml.Regexp("prefix-\\d+").
//
func Regexp(pattern string) gomock.Matcher {
	return &regexpMatcher{regexp.MustCompile(pattern)}
}

func (m *regexpMatcher) Matches(x interface{}) bool {
	if s, ok := x.(string); ok {
		return m.re.MatchString(s)
	}
	return false
}

func (m *regexpMatcher) String() string {
	return fmt.Sprintf("matches to the pattern of %s", m.re.String())
}
