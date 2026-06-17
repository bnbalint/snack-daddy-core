package database

import (
	"context"
	"database/sql"
	"testing"

	"snack-daddy-core/internal/models"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
	_ "github.com/mattes/migrate/source/file"
	pgmodule "github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestIngredientRepository(testingFramework *testing.T) {
	ctx := context.Background()

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
		testingFramework.Fatalf("failed to start container: %s", err)
	}

	// Ensure container cleanup when all tests finish
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			testingFramework.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Get the dynamic connection string from the running container
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		testingFramework.Fatalf("failed to get connection string: %s", err)
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		testingFramework.Fatalf("failed to connect to database: %s", err)
	}
	defer db.Close()

	// ------------------------------------------------
	// Run Migrations using golang-migrate

	migrator, err := migrate.New("file://../../db/migrations", connStr)
	if err != nil {
		testingFramework.Fatalf("failed to create migrator instance: %s", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		testingFramework.Fatalf("failed to run migrations up: %s", err)
	}

	// Get the dynamic connection host (no port) from the running container (needed to create client)
	dbHost, err := postgresContainer.Host(ctx)
	if err != nil {
		testingFramework.Fatalf("failed to get connecdb host string: %s", err)
	}

	// Create the Database Client, using the credentials & info from the testcontainer
	repo, err := NewDatabaseClient(dbHost, dbUser, dbPassword, dbName, 5432, "disable")
	if err != nil {
		testingFramework.Fatalf("failed to create DatabaseClient: %s", err)
	}

	//---------------------------------------------
	//  TESTS
	//

	// --- Subtest: Create Ingredient ---
	testingFramework.Run("Add Ingredient", func(t *testing.T) {
		ingredient := models.Ingredient{
			Name: "Pecan",
		}

		savedIngredient, err := repo.AddIngredient(ctx, &ingredient)
		if err != nil {
			t.Errorf("unexpected error creating ingredient: %v", err)
		}

		if savedIngredient.ID == 0 {
			t.Error("expected ingredient ID to be populated, got 0")
		}
	})

	// --- Subtest: Get All Ingredients ---
	testingFramework.Run("Get All Ingredients", func(t *testing.T) {
		ingredients, err := repo.GetAllIngredients(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching ingredients: %v", err)
		}

		if len(ingredients) == 0 {
			t.Errorf("expected some ingredients, got '%d'", len(ingredients))
		}
	})
}
