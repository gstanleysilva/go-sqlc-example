# Go SQLC/UOW example

### Installing Pre-requisites:

Installing Wire for DI:
```
go install github.com/google/wire/cmd/wire@latest
```
Installing SQLC for generating query helpers:
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
Installing Go Migrate:
```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

### Pre-Run
Generate the Dependency with: make wire
Start the database using: make db-up 
Run the migrations with: make migrate-up

### Running
Execute with: make run

