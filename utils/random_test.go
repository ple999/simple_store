package utils

import(
	"github.com/stretchr/testify/require"
	"testing"
)


func TestRandomInt(t *testing.T){
	var randomval int64=RandomInt(0,10);
	require.NotEmpty(t,randomval);
}

func TestRandomString(t *testing.T){
	var string1 string=RandomString(5)
	var string2 string=RandomString(5)
	require.NotEmpty(t,string1);
	require.NotEmpty(t,string2);
	require.NotEqual(t,string1,string2);

}