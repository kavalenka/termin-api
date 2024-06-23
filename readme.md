## Readme

### It's an API for scheduling and managing appointments. 
It's built with Go and PostgreSQL.

### For local development, please copy database.env.example to database.env and fill in the values.
```shell
cp config/database.env.example config/database.env
```
### After it install migrate tool (https://github.com/golang-migrate/migrate/tree/v4.17.1/cmd/migrate) and run migrations
```shell
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations up