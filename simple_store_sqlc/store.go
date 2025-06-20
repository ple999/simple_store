package simple_store_sqlc

import(
		pgx "github.com/jackc/pgx/v5"
)

type Store struct{
	Queries *Queries
	Db *pgx.Conn
}

func NewStore(conn *pgx.Conn) *Store{
	return &Store{
		Queries:New(conn),
		Db:conn,
	};
}