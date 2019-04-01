package gomml_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/utisam/gomml"
)

func TestNow(t *testing.T) {
	m := gomml.Now()
	now := time.Now()
	if !m.Matches(now) {
		t.Errorf("gomml.Now().Matches(%v)", now)
	}
}

func TestNowPtr(t *testing.T) {
	m := gomml.Now()
	now := time.Now()
	if !m.Matches(&now) {
		t.Errorf("gomml.Now().Matches(%v)", now)
	}
}

func TestNow_Before(t *testing.T) {
	now := time.Now()
	m := gomml.Now()
	if m.Matches(now) {
		t.Errorf("gomml.Now().Matches(%v)", now)
	}
}

func TestNow_After(t *testing.T) {
	now := time.Now().Add(time.Second)
	m := gomml.Now()
	if m.Matches(now) {
		t.Errorf("gomml.Now().Matches(%v)", now)
	}
}

func Test_nowMatcher_String(t *testing.T) {
	got := gomml.Now().String()
	match, err := regexp.MatchString("^is in range between .* and now$", got)
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Errorf("nowMatcher.String() = %v", got)
	}
}
