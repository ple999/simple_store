package simple_store_sqlc;

import(
	"context"
	"testing"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestInsertCountry(t *testing.T){
	var par InsertCountryParams=InsertCountryParams{
		CountryCode:pgtype.Text{String:"NO",Valid:true},
		CountryName:pgtype.Text{String:"Norway",Valid:true},
		ContinentName:pgtype.Text{String:"Europe",Valid:true},
	}
	result,err:=testQuery.InsertCountry(context.Background(),par);
	if(err!=nil){
		t.Errorf("Error Occured:%v",err);
	}
	t.Logf("Country Inserted:%v",result)
}

func TestGetAllCountries(t *testing.T){
	result,err:=testQuery.GetAllCountries(context.Background());
	if err!=nil{
		t.Errorf("Function Call Return Error:%v",err);
	}

	for i,v:= range result{
		t.Logf("Index:%v",i);
		t.Logf("Country ID:%v",v.CountryID);
		t.Logf("Country Name:%v",v.CountryName);
		t.Logf("Country Code:%v",v.CountryCode);
		t.Logf("Continent Name:%v",v.ContinentName);
	}
}

func TestGetAllOrders(t *testing.T){
	result,err:=testQuery.GetAllOrders(context.Background());
	if err!=nil{
		t.Errorf("Function Call Failed:%v",err);
	}

	for i,v:= range result{
		t.Logf("index:%v",i);
		t.Logf("Order ID:%v",v.OrderHeaderID);
		t.Logf("Order Date:%v",v.OrderDate);		
	}
}