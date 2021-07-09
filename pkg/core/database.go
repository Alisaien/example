package core

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

// DB reader endpoint
var DBR *pgxpool.Pool

// DB writer endpoint
var DBW *pgxpool.Pool

func Connect(ctx context.Context) {
	var err error
	DBR, err = pgxpool.Connect(ctx, os.Getenv("PGR_DSN"))
	if err != nil {
		panic(err)
	}

	DBW, err = pgxpool.Connect(ctx, os.Getenv("PGW_DSN"))
	if err != nil {
		panic(err)
	}
}

func Close() {
	DBR.Close()
	DBW.Close()
}
