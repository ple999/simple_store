generate_postgres_container:
	docker run --name postgres_alpine12  --network mynetwork -h 0.0.0.0 -e POSTGRES_USER=alpine12 -e POSTGRES_PASSWORD=password12 -e POSTGRES_DATABASE=alpine12 -d -v simple_store_postgres_db:/var/lib/postgresql/data -p 5000:5432 postgres:12-alpine 
stop_postgres_container:
	docker stop postgres_alpine12
run_postgres_container:
	docker start postgres_alpine12
createdb:
	docker exec -it postgres_alpine12 createdb --username=alpine12 --owner=alpine12 simple_store
dropdb:
	docker exec -it postgres_alpine12 dropdb simple_store  -U alpine12
migrateup:
	migrate -path ./migration_file -database "postgresql://alpine12:password1@localhost:5432/simple_store" -verbose up
migratedown:
	migrate -path ./migration_file -database "postgresql://alpine12:password1@localhost:5432/simple_store" -verbose down

.PHONY: generate_postgres_container run_postgres_container stop_postgres_container createdb dropdb migrateup migratedown