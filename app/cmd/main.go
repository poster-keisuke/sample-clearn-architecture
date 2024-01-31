package main

import (
	"context"
	"github.com/poster-keisuke/sample-clearn-architecture/app/infra/sqlite3/db"
	"github.com/poster-keisuke/sample-clearn-architecture/app/server"
	"log"
	"os"
)

const (
	exitOK int = iota
	exitError
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := db.NewDB()
	if err != nil {
		log.Printf("%q\n", err)
		os.Exit(exitError)
	}

	dbClient := db.GetDBConnection()
	defer dbClient.Close()

	server.Run(ctx)
}
