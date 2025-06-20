package api

import(
	gin "github.com/gin-gonic/gin"
	"net/http"
	db "simple_store/simple_store_sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"database/sql"
	"fmt"
	pgconn 	"github.com/jackc/pgx/v5/pgconn"
)

type createCountryRequest struct{
	CountryCode string `json:"country_code" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"` 
	ContinentName string `json:"continent_name" binding:"required"`
}

func (this Server) createCountry(ctx *gin.Context){
	var request createCountryRequest;
	err:= ctx.ShouldBindJSON(&request);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}
	
	arg:=db.InsertCountryParams{
		CountryCode:pgtype.Text{String:request.CountryCode,Valid:true},
		CountryName:pgtype.Text{String:request.CountryName,Valid:true},
		ContinentName:pgtype.Text{String:request.ContinentName,Valid:true},
	}
	
	country,err:= this.store.Queries.InsertCountry(ctx,arg);
	if err!=nil{
		pgErr,ok:=err.(*pgconn.PgError);
		if ok {
			switch pgErr.Code{
				case "23505":
				ctx.JSON(http.StatusForbidden,errorRequest(err));
				return
			}
		}
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;		
	}

	ctx.JSON(http.StatusOK,country);

}

type GetAllCountriesByPageRequest struct{
	PageSize int32 `form:"page_size" binding:"required"`
	PageID int32 `form:"page_id" binding:"required"`
}

func (this Server) getAllCountriesByPage(ctx *gin.Context){
	var request GetAllCountriesByPageRequest;
	err:= ctx.ShouldBindQuery(&request);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}

	
	arg:=db.GetAllCountriesByPageParams{
		Offset:(request.PageID-1) * request.PageSize,
		Limit:request.PageSize,
	}
	
	countrybypage,err:= this.store.Queries.GetAllCountriesByPage(ctx,arg);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;		
	}

	ctx.JSON(http.StatusOK,countrybypage);

}

type GetCountriesByIDRequest struct{
	CountryID int64 `uri:"id" binding:"required,min=1"`
}

func (this Server) getCountryById(ctx *gin.Context){
	var request GetCountriesByIDRequest;
	err:= ctx.ShouldBindUri(&request);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}
	fmt.Println(request);
	countrybyid,err:= this.store.Queries.GetCountryByID(ctx,request.CountryID);
	if err!=nil{

		if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,errorRequest(err));
			return;				
		}

		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;		
	}

	ctx.JSON(http.StatusOK,countrybyid);
}

func (this Server) getAllCountries(ctx *gin.Context){
	allcountry,err:=this.store.Queries.GetAllCountries(ctx);
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,errorRequest(err));
		return;
	}
	ctx.JSON(http.StatusOK,allcountry);
}
