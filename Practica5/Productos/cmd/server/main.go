package main

import (
	"database/sql"

	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/cmd/server/handler"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/product"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	databaseConfig := mysql.Config{
		User:   "user1",
		Passwd: "secret_password",
		Addr:   "127.0.0.1:3306",
		DBName: "my_db",
	}

	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		panic(err)
	}

	storage := store.NewRDBStore(*database)
	//storage := store.NewJsonStore("./products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
