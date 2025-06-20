package api

import(
	gin "github.com/gin-gonic/gin"
	db "simple_store/simple_store_sqlc"
	"simple_store/token"
	"fmt"
	"simple_store/utils"
)

type Server struct{
	store *db.Store
	tokenMaker token.Maker
	router *gin.Engine
	config utils.Config
}


func NewServer(config utils.Config,store *db.Store) (*Server,error){
	maker,err:=token.NewPasetoMaker(config.PasetoSymetricKey);
	if err!=nil{
		return nil,fmt.Errorf("Token Creation Failed:%v",err);
	}
	var server *Server=&Server{
		store:store,
		tokenMaker:maker,
		router:gin.Default(),
		config:config,
	};
	authRoutes:=server.router.Group("/").Use(AuthMiddleware(server.tokenMaker));
	server.router.POST("/insertCountry",server.createCountry)
	server.router.POST("/insertUser",server.insertUser)
	server.router.POST("/login",server.UserLogin)
	authRoutes.GET("/getCountryByPage",server.getAllCountriesByPage)
	authRoutes.GET("/getCountryById/:id",server.getCountryById)
	authRoutes.GET("/getAllCountries",server.getAllCountries)
	
	return server,nil;
}

func (this *Server) StartServer(address string) error{
	return this.router.Run(address);
}

func errorRequest(err error) gin.H{
	return gin.H{"error":err.Error()};
}

