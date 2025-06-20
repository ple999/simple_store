package utils

import(
	"math/rand"
	"time"
	"strings"
)

const(
	ALPHABET="abcdefghijklmnopqrstuvwxyz"
	ALPHABET_LEN int64=int64(len(ALPHABET))
)

func init(){
	rand.Seed(time.Now().UnixNano());
}

func RandomInt(min,max int64) int64{
	return min+rand.Int63n(max-min+1);
}

func RandomString(n int) string{
	var mystr strings.Builder;
	for i :=0;i<n;i++{
		c:=ALPHABET[RandomInt(0,ALPHABET_LEN-1)]
		mystr.WriteByte(c);
	}
	return mystr.String()
}