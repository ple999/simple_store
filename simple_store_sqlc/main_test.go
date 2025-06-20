package simple_store_sqlc;

import(
	"testing"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"context"
)

const (
	dbDriver="postgres"
	dbConnection="postgresql://alpine12:password12@localhost:5000/simple_store?sslmode=disable"
)

var testQuery *Queries;

func TestMain(t *testing.M){
	var myctx context.Context =context.Background();
	conn,err:=pgx.Connect(myctx,dbConnection)
	if err!=nil{
		log.Fatal("Connection to DB Failed:",err);
	}
	defer conn.Close(myctx);
	testQuery=New(conn);

	os.Exit(t.Run());
}