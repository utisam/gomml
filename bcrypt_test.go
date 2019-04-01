package gomml_test

import (
	"testing"

	"github.com/utisam/gomml"
	"golang.org/x/crypto/bcrypt"
)

func TestBCrypt(t *testing.T) {
	password := []byte("password")
	hashedPasswprd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	wrongHashedPasswprd, err := bcrypt.GenerateFromPassword([]byte("aaa"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		password []byte
		hash     []byte
		match    bool
	}{
		{
			password: password,
			hash:     hashedPasswprd,
			match:    true,
		},
		{
			password: password,
			hash:     wrongHashedPasswprd,
			match:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gomml.BCrypt(tt.password); !got.Matches(tt.hash) == tt.match {
				t.Errorf("gomml.BCrypt(%s).Matches(%s)", tt.password, tt.hash)
			}
		})
	}
}
