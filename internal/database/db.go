package database

import(
	"time"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectToDB() *pgxpool.Pool{
	host := "db"
	port := "5432"
	database := "ctf"
	user := "postgres"
	password := "ctf_task"

	connString := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"

	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil{
		log.Fatalf("Unable to connect to the db: %v\n", err)
	}

	dbpool.Config().MaxConns = 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Ping database was failed: %v\n", err)
	}

	return dbpool
}