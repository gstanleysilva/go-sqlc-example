run:
	@go run cmd/main.go

#Start DB
db-up:
	@docker-compose up -d

#Start DB
db-down:
	@docker-compose down

#Used to add the mysql driver
prepare-mysql:
	@go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#Create a new migration
migrate-create:
	@migrate create -ext=sql -dir=sql/migrations -seq init

#run migrations
migrate-up:
	@migrate -path=infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

#run migrations
migrate-down:
	@migrate -path=infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

#open mysql
#after executing this command you have to authenticate using: mysql -uroot -p courses
open-mysql:
	@docker-compose exec mysql bash

sqlc:
	@sqlc generate

.PHONY: db-up prepare-mysql migrate-create migrate-up migrate-down open-mysql sqlc