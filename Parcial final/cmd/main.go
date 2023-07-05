package main

import (
	"database/sql"
	"parcialfinal/cmd/server/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"parcialfinal/pkg/store"
	"parcialfinal/internal/domain/dentista"
)

func main(){
	// Set up database connection
	datasource := "root:password@tcp(localhost:3306)/parcialfinalback"
	storageDB, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	defer storageDB.Close()

	// Set up storage
	storage := store.NewSqlStore(storageDB)

	// Set up services
	dentistaService := dentista.NewService(storage)

	// Set up handlers
	dentistaHandler := handler.NewDentistaHandler(dentistaService)

	// Set up router
	router := gin.Default()
	v1 := router.Group("/api/v1")
	dentistaHandler.RegisterRoutes(v1.Group("/"))

	// Start server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}