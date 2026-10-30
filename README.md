# Booklooker

> This project was built before AI was a thing for me. The current version is a rewrite by Claude Opus 5.5 in Claude Code, based on the OpenAPI specification. Whether that is a mark of quality or the opposite is up to you.

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Checks](https://github.com/Plaenkler/Booklooker/actions/workflows/checks.yml/badge.svg)](https://github.com/Plaenkler/Booklooker/actions/workflows/checks.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/plaenkler/booklooker.svg)](https://pkg.go.dev/github.com/plaenkler/booklooker)

Go client for the [Booklooker REST API 2.0](https://www.booklooker.de/pages/rest_api.php), based on its [OpenAPI specification](https://www.booklooker.de/download/openapi.yaml). It uses only the standard library.

## Installation

```text
go get github.com/plaenkler/booklooker
```

## Usage

```go
client := booklooker.New(os.Getenv("BOOKLOOKER_API_KEY"))
status, err := client.ArticleStatus("HIS0067")
if err != nil {
	log.Fatal(err)
}
log.Print(status) // ACTIVE, SOLD or DELETED
```

The API key is shown in the [personal data](https://www.booklooker.de/app/priv/api_key.php) of the Booklooker account.

- The client authenticates on the first call. Booklooker drops a token after 10 minutes without a call, the client gets a new one after 9 minutes or when Booklooker reports it expired.
- A response with the status `NOK` is returned as `*booklooker.Error`, its `Code` is the value of Booklooker, e.g. `NOT_FOUND` or `QUOTA_EXCEEDED`.
- `BaseURL` and `HTTPClient` of the client can be replaced, the default timeout is 5 minutes because uploads may have 80 MB.
- The client is safe for concurrent use and works without `context`.

## Interfaces

| Interface           | Method                                         | Result                                      |
| ------------------- | ---------------------------------------------- | ------------------------------------------- |
| `account`           | `Account(mediaType)`                           | Seller account and its readiness to sell    |
| `article`           | `DeleteArticle(orderNo)`                       | Deletion within 2 hours                     |
| `article_list`      | `ArticleList(options)`                         | One line per active article                 |
| `article_status`    | `ArticleStatus(orderNo)`                       | `ACTIVE`, `SOLD` or `DELETED`               |
| `search`            | `Search(params)`                               | Undocumented result as raw JSON             |
| `file_import`       | `ImportFile(upload)`                           | Upload of an article or image file          |
| `file_status`       | `FileStatus(filename)`, `FileErrors(filename)` | Processing status or rejected articles      |
| `import_status`     | `ImportStatus()`                               | Number of unprocessed files                 |
| `image`             | `DeleteImage(orderNo, position)`, `DeleteImages(orderNo)` | Deletion of one or all images    |
| `order`             | `Orders(date)`, `OrdersBetween(from, to)`, `Order(orderID)` | Orders of a day, at most one month or one order |
| `order_cancel`      | `CancelOrder(orderID)`                         | Cancellation without commission             |
| `order_item_cancel` | `CancelOrderItem(orderItemID, mediaType)`      | Cancellation of one article                 |
| `order_message`     | `SendOrderMessage(orderID, type, text)`        | Text of the sent mail                       |
| `order_rating`      | `RateOrder(orderID, rating, comment)`          | Rating of the customer                      |
| `order_status`      | `SetOrderStatus(orderID, status)`              | Status for the own organization             |
| `purchase`          | `Purchase(purchase)`                           | Orders per seller and unknown articles      |

- `purchase` has to be enabled by Booklooker via daten@booklooker.de. It only fails if Booklooker did not process the call, each order of the result has a status of its own.
- `SetOrderStatus` with `Canceled` only organizes the orders, cancellations without commission need `CancelOrder` or `CancelOrderItem`.
- Amounts of orders are decimal numbers as text, e.g. `"12.50"`, as Booklooker returns them.

## Example

```go
content, err := os.ReadFile("books.txt")
if err != nil {
	log.Fatal(err)
}
err = client.ImportFile(booklooker.Upload{
	Name:     "books.txt",
	Content:  bytes.NewReader(content),
	DataType: booklooker.DataReplace,
	FormatID: 1, // Booklooker format
	Encoding: booklooker.UTF8,
})
```

More examples are in [`example_test.go`](example_test.go).

## Development

```text
go test ./...
golangci-lint run
```

The tests run against a local test server and need no API key.
