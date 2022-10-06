package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-graqhql-postgres/graph"
	"github.com/go-graqhql-postgres/graph/generated"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	// DBへの接続
	// host.docker.internalはコンテナをホストしてるipを指定してます
	// docker-composeでネットワークを作った方が良いとは思います。
	// なんか分からないけどうまく行かなかったのでとりあえずこうしてます。
	db, err := sql.Open("postgres", "host=host.docker.internal port=5432 user=postgres dbname=db sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

　　　　// Resolverに接続したdbを渡してResolverでdbにアクセスできるようにします。（Resolver.goに構造体を追加します）
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
