package main

import (
	//"context"
	"fmt"
	"log/slog"
	"os"
	"snack-daddy-core/internal/database"
	"snack-daddy-core/internal/server"
	"strconv"

	"github.com/joho/godotenv"
	//"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

const (
	zitadelDomain = "http://localhost:8080" // TODO - update when we move out of development
	projectID     = "378866658279686147"    // SnackDaddy's Zitadel Project ID
	clientId      = "378866830212661251"    // SnackDaddy's Zitadel ClientID
)

func main() {

	// create the logger
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	/*
	 * pull the database credentials
	 */
	if os.Getenv("env") == "" {
		err := godotenv.Load("db.env")
		if err != nil {
			logger.Error("Could not load environment", "error", err)
		}
	}

	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	sslMode := os.Getenv("sslmode")
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		logger.Error("failed to convert port to int", "error", err)
	}
	logger.Debug(fmt.Sprintf("host: %s, user: %s, password: %s, dbname: %s, port: %d, sslmode: %s", host, user, password, dbname, port, sslMode))

	/*
	 * create the database client
	 */
	databaseClient, err := database.NewDatabaseClient(host, user, password, dbname, int32(port), sslMode)
	if err != nil {
		logger.Error("Failed to initialize Database Client", "error", err)
	}

	/*
	 * connect to Zitadel
	 */

	/*
	 * create the service
	 */
	server := server.NewSnackDaddyEchoServer(databaseClient, *logger)

	/*
	 * start the service
	 */
	if err := server.Start(); err != nil {
		logger.Error(err.Error())
	}
}
