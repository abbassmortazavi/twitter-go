migrate:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@172.23.0.1:5432/twitter_dev?sslmode=disable up

rollback:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@172.23.0.1:5432/twitter_dev?sslmode=disable down

drop:
	migrate -source file://postgres/migrations \
			-database postgres://postgres:postgres@172.23.0.1:5432/twitter_dev?sslmode=disable drop
force:
	migrate -source file://postgres/migrations \
 			-database postgres://postgres:postgres@127.0.0.1:5432/twitter_dev?sslmode=disable force 20261005061032

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migrations $$name

run:
	go run cmd/graphqlserver/main.go
