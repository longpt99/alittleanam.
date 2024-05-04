package jwt

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	id := uuid.NewString()
	token, err := SignToken(id)
	fmt.Println(token)
	assert.NoError(t, err)

	claims, err := VerifyToken(token)
	assert.NoError(t, err)

	tokenSub, err := claims.GetSubject()
	assert.NoError(t, err)
	assert.Equal(t, id, tokenSub)

	// err = ComparePassword(wrongPassword, hashedPassword1)
	// assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// hashedPassword2, err := HashPassword(password)
	// assert.NoError(t, err)
	// assert.NotEmpty(t, hashedPassword2)
	// assert.NotEqual(t, hashedPassword1, hashedPassword2)
}
