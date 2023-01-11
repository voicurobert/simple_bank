package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "github.com/voicurobert/simple_bank/db/sqlc"
	"github.com/voicurobert/simple_bank/util"
	"testing"
	"time"
)

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func createRandomUser(t *testing.T) db.User {

	hashPass, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	args := db.CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashPass,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestQueries_GetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.Email, user2.Email)

	require.WithinDuration(t, user2.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user2.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}
