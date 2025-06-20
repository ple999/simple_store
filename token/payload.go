package token

import(
	"time"
	"github.com/google/uuid"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
)

var ErrorExpiredToken=fmt.Errorf("Token Expired");

type Payload struct{
	ID uuid.UUID  `json:"id"`
	Username string `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func NewPayload(username string,duration time.Duration) (*Payload,error){
	tokenId,err:=uuid.NewRandom();
	if err!=nil{
		return nil,err;
	}

	payload:=&Payload{
		ID:tokenId,
		Username:username,
		IssuedAt:time.Now(),
		ExpiredAt:time.Now().Add(duration),
	}
	return payload,nil;
}

func (this Payload) Valid() bool{
	if this.ExpiredAt.Unix()!=0 && time.Now().Unix()>this.ExpiredAt.Unix() {
		return true;
	}
	return false;
}

