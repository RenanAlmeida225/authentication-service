n =
text =
migrateup:
	migrate -path infra/database/postgres/migrations -database "postgres://postgres:postgres@127.0.0.1:5432/auth-api?sslmode=disable" -verbose up $(n)

migratedown:
	migrate -path infra/database/postgres/migrations -database "postgres://postgres:postgres@127.0.0.1:5432/auth-api?sslmode=disable" -verbose down $(n)	

migratecreate:
	migrate create -ext sql -dir infra/database/postgres/migrations -seq $(text) 

swaggergenerate:
	swagger generate spec --scan-models  -o ./1-api/swagger/swagger.json

swaggervalidate: 
	swagger validate ./1-api/swagger/swagger.json

.PHONY: migrateup migratedown migratecreate	swaggergenerate swaggervalidate