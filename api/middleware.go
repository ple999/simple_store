package api

import(
	gin "github.com/gin-gonic/gin"
	"simple_store/token"
	"errors"
	"net/http"
	"strings"
)


func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc{
	return func(ctx *gin.Context){
		authorizationHeader:=ctx.GetHeader("authorization");
		if len(authorizationHeader) ==0{
			err:=errors.New("Authorization Header is not provided");
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorRequest(err));
			return;
		}

		fields:=strings.Fields(authorizationHeader);
		if len(fields)<2{
			err:=errors.New("Authorization Header invalid Format");
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorRequest(err));
			return;
		}

		authorizationType:=fields[0];
		if authorizationType != "Bearer"{
			err:=errors.New("Authorization Type Invalid");
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorRequest(err));
			return;
		}

		accessToken:=fields[1];
		payload,err:=tokenMaker.VerifyToken(accessToken);
		if err!=nil{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorRequest(err));
			return;
		}

		ctx.Set("authorization_payload",payload);
		ctx.Next();
	}
}
