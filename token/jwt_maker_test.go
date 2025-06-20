package token

import(
	"github.com/stretchr/testify/require"
	"simple_store/utils"
	"time"
	"testing"
	_ "fmt"
)

func TestMakeJWTToken(t *testing.T){

	maker,err:=NewJWTMaker(utils.RandomString(32));
	require.NoError(t,err);

	userName:=utils.RandomString(6);
	duration:=time.Minute;
	issuedAt:=time.Now();
	expiredAt:=time.Now().Add(duration);

	token,err:=maker.CreateToken(userName,duration);
	require.NoError(t,err);
	require.NotEmpty(t,token);

	payload,err:=maker.VerifyToken(token);
	require.NoError(t,err);
	require.NotEmpty(t,payload);
	require.Equal(t,userName,payload.Username);
	require.WithinDuration(t,issuedAt,payload.IssuedAt,time.Second);
	require.WithinDuration(t,expiredAt,payload.ExpiredAt,time.Second)
}



func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.RandomString(5), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrorExpiredToken.Error())
	require.Nil(t, payload)
}

