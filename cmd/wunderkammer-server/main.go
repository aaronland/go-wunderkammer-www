package main

import (
	_ "github.com/mattn/go-sqlite3"
)

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-wunderkammer-www/templates"
	"github.com/aaronland/go-wunderkammer-www/www"
	"github.com/aaronland/go-wunderkammer/oembed"
	"log"
	"net/http"
)

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	dsn := flag.String("database-dsn", "sql://sqlite3/oembed.db", "A valid wunderkammer database DSN string.")

	path_templates := flag.String("path-templates", "static/templates/html/*", "The path to valid wunderkammer-www HTML templates.")

	flag.Parse()

	ctx := context.Background()

	db, err := oembed.NewSQLOEmbedDatabase(ctx, *dsn)

	if err != nil {
		log.Fatalf("Failed to create database, %v", err)
	}

	defer db.Close()

	t, err := templates.LoadHTMLTemplates(ctx, *path_templates)

	if err != nil {
		log.Fatalf("Failed to load HTML templates, %v", err)
	}

	mux := http.NewServeMux()

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}

	random_handler, err := www.NewRandomObjectHandler(db)

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/random", random_handler)

	image_handler, err := www.NewImageHandler(db, t)

	if err != nil {
		log.Fatal(err)
	}

	image_handler = bootstrap.AppendResourcesHandler(image_handler, bootstrap_opts)
	mux.Handle("/image", image_handler)

	object_handler, err := www.NewObjectHandler(db, t)

	if err != nil {
		log.Fatal(err)
	}

	object_handler = bootstrap.AppendResourcesHandler(object_handler, bootstrap_opts)
	mux.Handle("/object", object_handler)

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatal(err)
	}

	oembed_handler, err := www.NewOEmbedHandler(db)

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/oembed", oembed_handler)

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
