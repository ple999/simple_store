package api

import(
	gin "github.com/gin-gonic/gin"
	utils "simple_store/utils"
	"net/http"
	db "simple_store/simple_store_sqlc"
	pgtype "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
	"database/sql"
)

type InsertUserRequest struct{
	FirstName string `json:"first_name" binding:"required"`
	LastName string	`json:"last_name" binding:"required"`
	CountryID int64 `json:"country_id" binding:"required,min=1"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (this Server) insertUser(ctx *gin.Context){
	var request InsertUserRequest;
	err:=ctx.ShouldBindJSON(&request);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}

	var salt=utils.RandomString(8);
	var combinedPassword strings.Builder;
	combinedPassword.WriteString(request.Password);
	combinedPassword.WriteString(salt);
	
	var newPassword=combinedPassword.String();
	hashedPassword,err:=utils.HashPassword(newPassword);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return		
	}


	body:=db.InsertUserParams{
		FirstName:pgtype.Text{String:request.FirstName,Valid:true},
		LastName:pgtype.Text{String:request.LastName,Valid:true},
		Password:pgtype.Text{String:hashedPassword,Valid:true},
		Salt:pgtype.Text{String:salt,Valid:true},
		CountryID:pgtype.Int8{Int64:request.CountryID, Valid:true},
		Email:pgtype.Text{String:request.Email,Valid:true},
	}

	user,err:=this.store.Queries.InsertUser(ctx,body);
	if err!=nil{
		pgErr,ok:=err.(*pgconn.PgError);
		if ok{
			switch pgErr.Code{
				case "23505","23503":
					ctx.JSON(http.StatusForbidden,errorRequest(err));
					return
			}
		}
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return
	}
	ctx.JSON(http.StatusOK,user);
}


type LoginRequest struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct{
	Email string	`json:"email"`
	AccessToken string `json:"access_token"`
}

func (this Server) UserLogin(ctx *gin.Context){
	var req LoginRequest;
	err:=ctx.ShouldBindJSON(&req);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}

	userData,err:=this.store.Queries.GetUserByEmail(ctx,pgtype.Text{String:req.Email,Valid:true});
	if err!=nil{
		if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,errorRequest(err));
			return;			
		}
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}

	var combinedPasswordBuilder strings.Builder;
	combinedPasswordBuilder.WriteString(req.Password);
	combinedPasswordBuilder.WriteString(userData.Salt.String);
	var combinedPassword string=combinedPasswordBuilder.String();

	err=utils.CheckPassword(userData.Password.String,combinedPassword);
	if err!=nil{
		ctx.JSON(http.StatusUnauthorized,errorRequest(err));
		return;
	}
	
	token,err:= this.tokenMaker.CreateToken(utils.RandomString(6),this.config.TokenExpiredTime);
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,errorRequest(err));
		return;
	}

	response:=LoginResponse{
		Email:req.Email,
		AccessToken:token,
	}

	ctx.JSON(http.StatusOK,response);
}