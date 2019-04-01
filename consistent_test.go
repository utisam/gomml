package gomml_test

import (
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/utisam/gomml"
)

func TestConsistent(t *testing.T) {
	consistentNow := gomml.Consistent(gomml.Now())
	now := time.Now()
	if !consistentNow.Matches(now) {
		t.Errorf("consistentNow.Matches(%v)", now)
	}
	if !consistentNow.Matches(now) {
		t.Errorf("consistentNow.Matches(%v)", now)
	}
}

func TestConsistent__inconsistent(t *testing.T) {
	consistentNow := gomml.Consistent(gomml.Now())
	if now := time.Now(); !consistentNow.Matches(now) {
		t.Errorf("consistentNow.Matches(%v)", now)
	}
	if now := time.Now(); consistentNow.Matches(now) {
		t.Errorf("consistentNow.Matches(%v)", now)
	}
}

func Test_consistentMatcher_String(t *testing.T) {
	m := gomml.Consistent(gomock.Any())

	if got := m.String(); !strings.HasSuffix(got, " and must be consistent") {
		t.Errorf("consistentAny.String() = %v", got)
	}

	m.Matches(nil)

	if got := m.String(); !strings.Contains(got, " and must be consistent with ") {
		t.Errorf("consistentAny.String() = %v", got)
	}
}
