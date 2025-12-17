package controllers

import (
	"fmt"
	"net/http"
	"sayurgo/config"
	"sayurgo/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// func UpdateProduct(c *gin.Context) {
// 	id := c.Param("id")

// 	var product models.Product
// 	if err := config.DB.First(&product, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Product not found",
// 		})
// 		return
// 	}

// 	var input models.Product
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	product.Name = input.Name
// 	product.Price = input.Price
// 	product.Stock = input.Stock
// 	product.Image = input.Image

// 	config.DB.Save(&product)

// 	c.JSON(http.StatusOK, product)
// }

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// ⬇️ SEMUA DARI FORM-DATA
	if name := c.PostForm("name"); name != "" {
		product.Name = name
	}
	if price := c.PostForm("price"); price != "" {
		product.Price, _ = strconv.Atoi(price)
	}
	if stock := c.PostForm("stock"); stock != "" {
		product.Stock, _ = strconv.Atoi(stock)
	}

	// image OPTIONAL
	file, err := c.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		path := "uploads/products/" + filename
		_ = c.SaveUploadedFile(file, path)
		product.Image = path
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}
