package test

import (
	"errors"
	"github.com/jahs/clinic-backend/adapter/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pass_Is_Ok(t *testing.T) {
	pass := password.NewService()

	err := pass.Compare("$2a$10$FPl13deAraS0lzRd8R0IduCYbmH4Dv3nIyN7baW7PJwir0K6ktCUi", "admin")
	assert.Nil(t, err)

}

func Test_Pass_Is_Not_Ok(t *testing.T) {
	pass := password.NewService()

	err := pass.Compare("$2a$10$FPl13deAraS0lzRd8R0IduCYbmH4Dv3nIyN7baW7PJwir0K6ktCUy", "admin")
	assert.Equal(t, err, errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"))
}
