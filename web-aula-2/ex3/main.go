package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID       int    `binding:"-"`
	Code     string `json:"transactionCode"`
	Currency string
	Amount   float64
	Sender   string
	Receiver string
	Date     time.Time
}

var transactions = []Transaction{}

func FindByID(c *gin.Context) {
	target, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
		return
	}

	for _, transaction := range transactions {
		if transaction.ID == target {
			c.JSON(200, transaction)
			return
		}
	}
	c.Status(404)
}

func GetAllTransactions(c *gin.Context) {
	c.JSON(200, &transactions)
}

func AddNewTransaction(c *gin.Context) {
	var t Transaction
	err := c.BindJSON(&t)

	if err != nil {
		c.JSON(400, err)
		return
	}

	t.ID = len(transactions)
	transactions = append(transactions, t)
	c.Status(201)
}

func main() {
	router := gin.Default()
	group := router.Group("/transactions")
	{
		group.GET("/:id", FindByID)
		group.GET("/", GetAllTransactions)
		group.POST("/", AddNewTransaction)
	}
	router.Run()
}
