# Running the migrations

Store the connection information: <br>
`export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/snackdaddy?sslmode=disable'`


Run the migrations: <br>
`migrate -database ${POSTGRESQL_URL} -path db/migrations up`



# Resolving issues

Seeing  `Dirty database version 1. Fix and force version.`? <br>

Fix things and reset the version with `migrate -database ${POSTGRESQL_URL} -path db/migrations force 1`