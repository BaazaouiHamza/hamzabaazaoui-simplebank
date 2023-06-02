package token

import (
	"testing"
	"time"

	"github.com/hamzabaazaoui/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	Payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, Payload)

	require.NotZero(t, Payload.ID)
	require.Equal(t, username, Payload.Username)
	require.WithinDuration(t, issuedAt, Payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, Payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
