package main

import(
	"github.com/jackc/pgx/v5"
	"log"
	"context"
	"simple_store/simple_store_sqlc"
	"simple_store/api"
	"simple_store/utils"
)

func main(){
	config,err:=utils.LoadConfig(".");
	if err!=nil{
		log.Fatal("Fail Reading Config:",err);
	}

	var myctx context.Context =context.Background();
	conn,err:=pgx.Connect(myctx,config.DBConnection)
	if err!=nil{
		log.Fatal("Connection to DB Failed:",err);
	}
	defer conn.Close(myctx);
	var store = simple_store_sqlc.NewStore(conn);
	server,err := api.NewServer(config,store);
	if err!=nil{
		log.Fatal("Failed To Create Server:",err);
	}	
	err= server.StartServer(config.ServerAddress);
	if err!=nil{
		log.Fatal("Failed to Start Server:",err);	
	}

}