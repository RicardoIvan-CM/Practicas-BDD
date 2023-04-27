package main

import (
	"database/sql"

	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/cmd/server/handler"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/product"
	"github.com/RicardoIvan-CM/Practicas-BDD/Practica5/Productos/internal/warehouse"
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

	//storage := store.NewJsonStore("./products.json")
	productStorage := store.NewRDBStore(database)
	productRepository := product.NewRepository(productStorage)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	warehouseRepository := warehouse.NewRepository(database)
	warehouseService := warehouse.NewService(warehouseRepository)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	warehouses := r.Group("/warehouses")
	{
		warehouses.GET("", warehouseHandler.GetAll())
		warehouses.GET(":id", warehouseHandler.GetByID())
		warehouses.POST("", warehouseHandler.Post())
		/*
			warehouses.DELETE(":id", warehouseHandler.Delete())
			warehouses.PATCH(":id", warehouseHandler.Patch())
			warehouses.PUT(":id", warehouseHandler.Put())
		*/
		warehouses.GET("reportProducts", warehouseHandler.GetProductsByID())
	}

	r.Run(":8080")
}
