package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"project-name/db"
	"project-name/web"
)

func main() {
	conn, err := pgx.Connect(context.Background(), connString())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(db.NewDB(conn), cors)
	err = app.Serve()
	log.Println("Error", err)
}

func connString() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "postgresql://" + host + ":5432/goxygen?user=goxygen&password=" + pass
}
