package main

import (
	"nbstore/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := SetupRouter()
	router.Run("127.0.0.1:8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/", heartbeat)
	router.GET("/products", getProducts)
	router.POST("/products", addProduct)
	router.GET("/products/:id", getProduct)

	return router
}

func heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm alive!",
	})
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items.GetProducts())
}

func addProduct(c *gin.Context) {
	var newProduct items.Product

	//BindJSON deserializes the recieved JSON to the strcut Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		return
	}

	insertedProduct, _ := items.AddProduct(newProduct)
	c.IndentedJSON(http.StatusCreated, insertedProduct)

}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := items.GetProduct(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, product)

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})

}
