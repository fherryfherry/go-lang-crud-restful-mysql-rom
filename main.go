package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	ID          uint    `json:"id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float32 `json:"price"`
}

func connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/sample_shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func listData(c *gin.Context) {
	db := connect()

	var products[] Product

	db.Find(&products)

	c.JSON(200, gin.H{"status":1,"message":"ok","data":products})
}

func detailData(c *gin.Context) {
	db := connect()

	var product Product

	db.First(&product, c.Param("id"))

	c.JSON(200, gin.H{"status":1,"message":"ok","data":product})
}

func updateData(c *gin.Context) {
	db := connect()
	var product Product
	db.First(&product, c.PostForm("id"))

	product.Sku = c.PostForm("sku")
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")

	stock, err := strconv.Atoi(c.PostForm("stock"))
	if err == nil {
		product.Stock = stock
	}

	price, err := strconv.ParseFloat(c.PostForm("price"),32)
	if err == nil {
		product.Price = float32(price)
	}
	db.Save(&product)

	c.JSON(200, gin.H{"status":1,"message":"ok"})
}

func createData(c *gin.Context) {
	db := connect()

	var product Product
	product.Sku = c.PostForm("sku")
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")

	stock, err := strconv.Atoi(c.PostForm("stock"))
	if err == nil {
		product.Stock = stock
	}

	price, err := strconv.ParseFloat(c.PostForm("price"),32)
	if err == nil {
		product.Price = float32(price)
	}
	db.Save(&product)

	c.JSON(200, gin.H{"status":1,"message":"ok"})
}

func deleteData(c *gin.Context) {
	db := connect()

	db.Delete(&Product{}, c.PostForm("id"))

	c.JSON(200, gin.H{"status":1,"message":"ok"})
}

func main() {
	fmt.Println("Go CRUD With MySQL + ORM")

	router := gin.Default()
	router.GET("/products", listData)
	router.GET("/products/detail/:id", detailData)
	router.POST("/products/update", updateData)
	router.POST("/products/delete", deleteData)
	router.POST("/products/create", createData)

	router.Run("localhost:8080")
}