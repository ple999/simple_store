package token

import(
	"github.com/stretchr/testify/require"
	"simple_store/utils"
	"time"
	"testing"
	_ "fmt"
)

func TestMakePasetoToken(t *testing.T){
	var symetricKey string = utils.RandomString(32);
	maker,err:=NewPasetoMaker(symetricKey);
	require.NoError(t,err);
	require.NotEmpty(t,maker);

	token,err:=maker.CreateToken(utils.RandomString(6),time.Hour);
	require.NoError(t,err);
	require.NotEmpty(t,token);

	payload,err:=maker.VerifyToken(token);
	require.NoError(t,err);
	require.NotEmpty(t,payload);
}

func TestMakeExpiredPasetoToken(t *testing.T){
	var symetricKey string=utils.RandomString(32);
	maker,err:=NewPasetoMaker(symetricKey);
	require.NoError(t,err);
	require.NotEmpty(t,maker);

	token,err:=maker.CreateToken(utils.RandomString(6),-time.Hour);
	require.NoError(t,err);
	require.NotEmpty(t,maker);

	payload,err:=maker.VerifyToken(token);
	require.Error(t,err);
	require.Empty(t,payload);
}