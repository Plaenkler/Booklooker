# Booklooker

> This project is a rewrite by Claude Opus 5.5 in Claude Code, based on the OpenAPI specification. Whether that is a mark of quality or the opposite is up to you.

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Checks](https://github.com/Plaenkler/Booklooker/actions/workflows/checks.yml/badge.svg)](https://github.com/Plaenkler/Booklooker/actions/workflows/checks.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/plaenkler/booklooker/v2.svg)](https://pkg.go.dev/github.com/plaenkler/booklooker/v2)

Go client for the [Booklooker REST API 2.0](https://www.booklooker.de/pages/rest_api.php), based on its [OpenAPI specification](https://www.booklooker.de/download/openapi.yaml). It uses only the standard library.

## Installation

```text
go get github.com/plaenkler/booklooker/v2
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
