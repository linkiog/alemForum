package main

import (
	"fmt"
	"forum/internal/handler"
	"forum/internal/service"
	"forum/internal/storage"
	"log"
	"net/http"
)

func main() {
	confDb := storage.ConfDb()
	db, err := storage.CreateDb(confDb)
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.CreateTab(db); err != nil {
		log.Fatal(err)
	}
	storage := storage.InitStorage(db)
	services := service.InitService(storage)
	handler := handler.InitHandler(services)
	handler.Routers()
	fmt.Println("http://localhost:8081")
	if err = http.ListenAndServe(":8081", handler.Mux); err != nil {
		log.Fatal(err)
	}
}
