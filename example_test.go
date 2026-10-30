package booklooker_test

import (
	"bytes"
	"errors"
	"log"
	"os"
	"time"

	"github.com/plaenkler/booklooker"
)

func ExampleClient_ImportFile() {
	client := booklooker.New(os.Getenv("BOOKLOOKER_API_KEY"))
	content, err := os.ReadFile("books.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = client.ImportFile(booklooker.Upload{
		Name:     "books.txt",
		Content:  bytes.NewReader(content),
		DataType: booklooker.DataReplace,
		FormatID: 1,
		Encoding: booklooker.UTF8,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_Orders() {
	client := booklooker.New(os.Getenv("BOOKLOOKER_API_KEY"))
	orders, err := client.Orders(time.Now())
	if err != nil {
		log.Fatal(err)
	}
	for _, order := range orders {
		log.Printf("%d from %s with %d articles", order.OrderID, order.BuyerUsername, len(order.OrderItems))
	}
}

func ExampleError() {
	client := booklooker.New(os.Getenv("BOOKLOOKER_API_KEY"))
	err := client.DeleteArticle("HIS0067")
	var apiError *booklooker.Error
	if errors.As(err, &apiError) && apiError.Code == "NOT_FOUND" {
		log.Print("the article does not exist")
	}
}
