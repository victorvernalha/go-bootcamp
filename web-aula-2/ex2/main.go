package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Transaction struct {
	ID       int
	Code     string    `json:"transactionCode"`
	Currency string    `binding:"required"`
	Amount   float64   `binding:"required"`
	Sender   string    `binding:"required"`
	Receiver string    `binding:"required"`
	Date     time.Time `binding:"required"`
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

func parseError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			errorMessages[i] = e.Field()
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("The field %s must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}

func AddNewTransaction(c *gin.Context) {
	var t Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		msgs := parseError(err)
		msg := strings.Join(msgs, "\n")
		fmt.Printf("%T: %v\n", err, err)
		c.JSON(400, msg)
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
