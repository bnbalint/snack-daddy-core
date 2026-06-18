package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
	_ "github.com/mattes/migrate/source/file"
	pgmodule "github.com/testcontainers/testcontainers-go/modules/postgres"
)

var (
	DbClient SnackDaddyDatabaseClient
	ctx      = context.Background()
)

// Sets up all tests in this package
// Starts the postgres Docker container
// Runs the migrations
// Sets the dbClient variable
func TestMain(testingFramework *testing.M) {

	// Set the connection parameters
	dbName := "testdb"
	dbUser := "postgres"
	dbPassword := "password"

	// Start the container
	postgresContainer, err := pgmodule.Run(ctx,
		"postgres:18-alpine",
		pgmodule.WithDatabase(dbName),
		pgmodule.WithUsername(dbUser),
		pgmodule.WithPassword(dbPassword),
		pgmodule.BasicWaitStrategies(),
	)

	// Ensure the container started
	if err != nil {
		fmt.Printf("failed to start container: %s", err)
		os.Exit(1)
	}

	// Get the dynamic connection string from the running container
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		fmt.Printf("failed to get connection string: %s", err)
		os.Exit(1)
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("failed to connect to database: %s", err)
		os.Exit(1)
	}
	defer db.Close()

	// ------------------------------------------------
	// Run Migrations using golang-migrate

	// create the migrator
	migrator, err := migrate.New("file://../../db/migrations", connStr)
	if err != nil {
		fmt.Printf("failed to create migrator instance: %s", err)
		os.Exit(1)
	}

	// run the up migrations
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("failed to run migrations up: %s", err)
		os.Exit(1)
	}

	// Get the dynamic connection host (no port) from the running container (needed to create client)
	dbHost, err := postgresContainer.Host(ctx)
	if err != nil {
		fmt.Printf("failed to get connection host string: %s", err)
		os.Exit(1)
	}

	// Get the dynamic connection port from the running container (needed to create client)
	dbPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		fmt.Printf("failed to get connection port string: %s", err)
		os.Exit(1)
	}

	// Create the Database Client, using the credentials & info from the testcontainer
	fmt.Printf("host=%s user=%s password=%s dbname=%s  dbPort=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	dbClient, err := NewDatabaseClient(dbHost, dbUser, dbPassword, dbName, int32(dbPort.Num()), "disable")
	if err != nil {
		fmt.Printf("failed to create DatabaseClient: %s", err)
		os.Exit(1)
	}
	DbClient = dbClient // set the global variable

	// Run all tests in this package
	exitCode := testingFramework.Run()

	// Terminate the container when ALL tests finish
	if err := postgresContainer.Terminate(ctx); err != nil {
		fmt.Printf("failed to terminate container: %s", err)
	}

	// exit with the exitCode from the tests
	os.Exit(exitCode)

}
