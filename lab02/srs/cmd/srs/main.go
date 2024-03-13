package main

import (
	"srs/internal/app"
	"srs/internal/transport/rest"
)

func main() {
	app := app.New()

	app.Engine.GET("/products", rest.GetProducts)
	app.Engine.POST("/product", rest.PostProduct)
	app.Engine.GET("/product/:id", rest.GetProduct)
	app.Engine.PUT("/product/:id", rest.PutProduct)
	app.Engine.DELETE("/product/:id", rest.DeleteProduct)

	app.Engine.Run(":8080")
}
