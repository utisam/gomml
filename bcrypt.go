package gomml

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

type bcryptMatcher struct {
	password []byte
}

// BCrypt is a Matcher that matches if the hashed parameter to the mock
// function is equivalent to the password.
//
// Example usage:
//
// 		dbMock.EXPECT().
// 			Insert(gomml.BCrypt("password").
// 			Return(errors.New("DB error"))
//
func BCrypt(password []byte) gomock.Matcher {
	return &bcryptMatcher{password}
}

func (m *bcryptMatcher) Matches(x interface{}) bool {
	if hpw, ok := x.([]byte); ok {
		return bcrypt.CompareHashAndPassword(hpw, m.password) == nil
	}
	return false
}

func (m *bcryptMatcher) String() string {
	return fmt.Sprintf("matches to %s with bcrypt", m.password)
}
