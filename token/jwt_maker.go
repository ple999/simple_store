package token


import(
	"fmt"
	"time"
	jwt "github.com/golang-jwt/jwt/v5"
)


const minSecretKeyLen=32;
var ErrorInvalidSigningMethod=fmt.Errorf("Invalid Signing Method");
var ErrorInvalidToken=fmt.Errorf("Invalid Token");
type JWTMaker struct{
	SecretKey string
}

func NewJWTMaker(secretKey string) (Maker,error){
	if len(secretKey)<minSecretKeyLen{
		return nil,fmt.Errorf("Secret Must have length of %v",minSecretKeyLen);
	}

	return &JWTMaker{SecretKey:secretKey},nil;
}

func (this *JWTMaker) CreateToken(username string,duration time.Duration) (string,error) {
	payload,err:=NewPayload(username,duration);
	if err!=nil{
		return "",fmt.Errorf("Failed Creating Payload:%v",err);
	}
	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,payload);
	return jwtToken.SignedString([]byte(this.SecretKey));
}

func (this *JWTMaker) VerifyToken(token string) (*Payload,error){
	keyFunc:=func (token *jwt.Token) (interface{},error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC);
		if !ok{
			return nil,ErrorInvalidSigningMethod
		}
		return []byte(this.SecretKey),nil;
	}

	jwtToken,err:=jwt.ParseWithClaims(token,&Payload{},keyFunc)
	if err!=nil{
		return nil,ErrorInvalidToken
	}

	payload,ok:=jwtToken.Claims.(*Payload);
	if !ok{
		return nil,ErrorInvalidToken;
	}

	ok=payload.Valid();
	if ok{
		return nil,ErrorExpiredToken
	}

	return payload,nil;
	
}