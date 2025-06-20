package utils

import(
	"github.com/stretchr/testify/require"
	"testing"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func TestPassword(t *testing.T){
	password:=RandomString(20);
	salt:=RandomString(10);

	var combinedpassword strings.Builder;
	combinedpassword.WriteString(password);
	combinedpassword.WriteString(salt);

	hashedPassword,err:=HashPassword(combinedpassword.String());
	require.NoError(t,err);
	require.NotEmpty(t,hashedPassword);
	err=CheckPassword(hashedPassword,combinedpassword.String());
	require.NoError(t,err);
	
	wrongPassword:=RandomString(6);
	err=CheckPassword(hashedPassword,wrongPassword);
	require.EqualError(t,err,bcrypt.ErrMismatchedHashAndPassword.Error());
}