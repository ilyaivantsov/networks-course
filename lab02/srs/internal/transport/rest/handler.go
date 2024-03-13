package rest

import (
	"net/http"
	"srs/internal/models"
	"srs/internal/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostProduct(ctx *gin.Context) {
	var product models.ProductCreate

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	p, err := storage.CreateProduct(product.Name, product.Description)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, p)
}

func PutProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	var product models.ProductUpdate
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	_, err = storage.UpdateProduct(id, product.Name, product.Description)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	product, err := storage.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func GetProducts(ctx *gin.Context) {
	products := storage.GetProducts()
	ctx.JSON(http.StatusOK, products)
}

func GetProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	product, err := storage.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
