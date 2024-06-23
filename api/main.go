package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type Item struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Valor    string `json:"valor"`
	Quantity int    `json:"quantity"`
}

var inventory = []Item{
	{ID: "1", Nome: "In Search of Lost Time", Valor: "30.00", Quantity: 2},
	{ID: "2", Nome: "The Great Gatsby", Valor: "50.00", Quantity: 5},
	{ID: "3", Nome: "War and Peace", Valor: "25.50", Quantity: 6},
}

func connectDB() *sql.DB {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "uniceub"
	config.DBName = "web"
	conn, err := mysql.NewConnector(config)
	if err != nil {
		panic(err)
	}
	return sql.OpenDB(conn)
}

var db *sql.DB

func getItem(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, inventory)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")
	item, err := getItemByIDFromInventory(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

func checkoutItem(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	item, err := getItemByIDFromInventory(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	if item.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "item not available"})
		return
	}

	item.Quantity--
	c.IndentedJSON(http.StatusOK, item)
}

func returnItem(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	item, err := getItemByIDFromInventory(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	item.Quantity++
	c.IndentedJSON(http.StatusOK, item)
}

func getItemByIDFromInventory(id string) (*Item, error) {
	for i, item := range inventory {
		if item.ID == id {
			return &inventory[i], nil
		}
	}
	return nil, errors.New("item not found")
}

func createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	inventory = append(inventory, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func getInventory(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, inventory)
}

func main() {
    db = connectDB()
    defer db.Close()

    router := gin.Default()
    router.GET("/inventory", getInventory)  // Novo endpoint
    router.GET("/item", getItem)
    router.GET("/item/:id", getItemByID)
    router.POST("/item", createItem)
    router.PATCH("/checkout", checkoutItem)
    router.PATCH("/return", returnItem)

	// Inicia o servidor na porta 8080
	router.Run("localhost:8080")
}
