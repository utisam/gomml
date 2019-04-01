package gomml

import (
	"fmt"
	"time"

	"github.com/golang/mock/gomock"
)

type nowMatcher struct {
	start time.Time
}

// Now is a Matcher that matches with now.
//
// The value must be the time between the constructor was called and
// it is validated.
//
// Example usage:
//
// 		dbMock.EXPECT().
// 			Insert(gomml.Now()).
//
func Now() gomock.Matcher {
	return &nowMatcher{time.Now()}
}

func (m *nowMatcher) Matches(x interface{}) bool {
	switch t := x.(type) {
	case time.Time:
		return m.start.Before(t) && t.Before(time.Now())
	case *time.Time:
		return m.start.Before(*t) && t.Before(time.Now())
	}
	return false
}

func (m *nowMatcher) String() string {
	return fmt.Sprintf("is in range between %v and now", m.start)
}
