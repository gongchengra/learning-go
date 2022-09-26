package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gongchengra/learning-go/05.advanced-programs/GraphQL-CRUD/database"
	"github.com/gongchengra/learning-go/05.advanced-programs/GraphQL-CRUD/graph"
	"github.com/gongchengra/learning-go/05.advanced-programs/GraphQL-CRUD/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDatabase()
	createTable()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDatabase() {
	var connectionString string

	// Get environment variable for connection string
	host := os.Getenv("HOST")
	if "" == host {
		connectionString = "host=localhost port=5432 user=postgres dbname=pq-demo password=example sslmode=disable"
	} else {
		connectionString = "host=" + host + " port=5432 user=postgres dbname=pq-demo password=example sslmode=disable"
	}

	var err error
	database.DBConn, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func createTable() {
	db := database.DBConn

	// Delete table if you changed/updated the schema
	deleteQuery := `DROP TABLE IF EXISTS items`
	db.Exec(deleteQuery)

	query := `
	CREATE TABLE IF NOT EXISTS items
	(
	 id serial NOT NULL,
	 Title character varying NOT NULL,
	 Owner character varying,
	 Rating integer,
	 created_at date,
	 updated_at date,
	 deleted_at date,
	 CONSTRAINT pk_books PRIMARY KEY (id )
	)
	WITH (
	 OIDS=FALSE
	);
	ALTER TABLE items
	 OWNER TO postgres;`

	db.Exec(query)
}
