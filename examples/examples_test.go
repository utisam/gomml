package examples_test

import (
	"testing"
	"time"

	"github.com/utisam/gomml"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang/mock/gomock"
	"github.com/utisam/gomml/examples"
)

func TestExampleInterface_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := examples.NewMockExampleInterface(ctrl)

	consistentNow := gomml.Consistent(gomml.Now())
	mock.EXPECT().Create(
		gomml.Regexp("^prefix-\\d+$"),
		gomml.BCrypt([]byte("passwd")),
		consistentNow, consistentNow,
	)

	testedMethod(mock)
}

func testedMethod(i examples.ExampleInterface) {
	now := time.Now()
	randomNumber := "34567"
	hashedPasswrod, _ := bcrypt.GenerateFromPassword([]byte("passwd"), bcrypt.DefaultCost)
	i.Create("prefix-"+randomNumber, hashedPasswrod, now, now)
}
