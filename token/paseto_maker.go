package token

import(
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"fmt"
	"time"
)

type PasetoMaker struct{
	paseto *paseto.V2
	symetricKey []byte
}



func NewPasetoMaker(symetricKey string) (Maker,error){
	if len(symetricKey)!= chacha20poly1305.KeySize{
		return nil,fmt.Errorf("Invalid Key Size : Must be %v Character",chacha20poly1305.KeySize);
	}

	maker:=&PasetoMaker{
		paseto:paseto.NewV2(),
		symetricKey:[]byte(symetricKey),
	}

	return maker,nil
}

func (this *PasetoMaker) CreateToken(username string,duration time.Duration) (string,error){
	payload,err:=NewPayload(username,duration);
	if err!=nil{
		return "",fmt.Errorf("Failed to Create Payload:%v",payload);
	}

	return this.paseto.Encrypt(this.symetricKey,payload,nil);
}

func (this *PasetoMaker) VerifyToken(token string) (*Payload,error){
	payload:=&Payload{}
	err:=this.paseto.Decrypt(token,this.symetricKey,payload,nil);
	if err!=nil{
		return nil,ErrorInvalidToken;
	}

	if payload.Valid(){
		return nil,ErrorExpiredToken;
	}

	return payload,nil;
}

