package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)
	hashPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword)
}

func TestCheckPassword(t *testing.T) {
	password := RandomString(6)
	hashPassword, err := HashPassword(password)
	require.NoError(t, err)

	err = CheckPassword(password, hashPassword)
	require.NoError(t, err)
}

func TestCheckWrongPassword(t *testing.T) {
	password := RandomString(6)
	hashPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword1)

	err = CheckPassword(RandomString(6), hashPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)

	require.NotEmpty(t, hashPassword1, hashPassword2)
}
