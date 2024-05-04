package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := GenerateString(6)

	hashedPassword1, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword1)

	err = ComparePassword(password, hashedPassword1)
	assert.NoError(t, err)

	wrongPassword := GenerateString(6)
	err = ComparePassword(wrongPassword, hashedPassword1)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword2)
	assert.NotEqual(t, hashedPassword1, hashedPassword2)
}
