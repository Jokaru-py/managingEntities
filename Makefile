setupdocker: ## add postgres image to docker
	docker-compose up --build

migrateup:  ## Dropdown db
	migrate -path db/migration -database "postgresql://root:WEBdeveloepr1452@localhost:12012/contractors_acounts?sslmode=disable" -verbose up

migratedown:  ## Dropdown db
	migrate -path db/migration -database "postgresql://root:WEBdeveloepr1452@localhost:12012/contractors_acounts?sslmode=disable" -verbose down

sqlc: ## sqlc generates type-safe code from SQL
	sqlc generate

.PHONY: setupdocker migrateup migratedown sqlc