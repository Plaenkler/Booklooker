package booklooker

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"time"
)

const testToken = "TEST_TOKEN"

type testAPI struct {
	handlers        map[string]http.HandlerFunc
	authentications int
}

func newTestClient(t *testing.T, handlers map[string]http.HandlerFunc) (*Client, *testAPI) {
	t.Helper()
	api := &testAPI{handlers: handlers}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.serve(t, w, r)
	}))
	t.Cleanup(server.Close)
	client := New("TEST_KEY")
	client.BaseURL = server.URL + "/2.0"
	return client, api
}

func (a *testAPI) serve(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()
	path := strings.TrimPrefix(r.URL.Path, "/2.0/")
	if path == "authenticate" {
		a.authentications++
		if r.Method != http.MethodPost || r.URL.Query().Get("apiKey") != "TEST_KEY" {
			reply(w, "NOK", `"AUTHENTICATION_FAILED"`)
			return
		}
		reply(w, "OK", `"`+testToken+`"`)
		return
	}
	handler, ok := a.handlers[r.Method+" "+path]
	if !ok {
		reply(w, "NOK", `"INVALID_INTERFACE"`)
		return
	}
	if r.URL.Query().Get("token") != testToken {
		reply(w, "NOK", `"TOKEN_UNKNOWN"`)
		return
	}
	handler(w, r)
}

func reply(w http.ResponseWriter, status, value string) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, `{"status":%q,"returnValue":%s}`, status, value)
}

func TestTokenIsReused(t *testing.T) {
	client, api := newTestClient(t, map[string]http.HandlerFunc{
		"GET import_status": func(w http.ResponseWriter, _ *http.Request) { reply(w, "OK", "3") },
	})
	for range 3 {
		count, err := client.ImportStatus()
		if err != nil || count != 3 {
			t.Fatalf("ImportStatus() = %d, %v", count, err)
		}
	}
	if api.authentications != 1 {
		t.Errorf("authenticated %d times, want 1", api.authentications)
	}
	client.used = time.Now().Add(-tokenLifetime)
	_, err := client.ImportStatus()
	if err != nil {
		t.Fatal(err)
	}
	if api.authentications != 2 {
		t.Errorf("an unused token was not renewed, %d authentications", api.authentications)
	}
}

func TestExpiredTokenIsRenewed(t *testing.T) {
	expired := true
	client, api := newTestClient(t, map[string]http.HandlerFunc{
		"GET article_status": func(w http.ResponseWriter, _ *http.Request) {
			if expired {
				expired = false
				reply(w, "NOK", `"TOKEN_EXPIRED"`)
				return
			}
			reply(w, "OK", `"SOLD"`)
		},
	})
	status, err := client.ArticleStatus("HIS0067")
	if err != nil || status != ArticleSold {
		t.Fatalf("ArticleStatus() = %q, %v", status, err)
	}
	if api.authentications != 2 {
		t.Errorf("authenticated %d times, want 2", api.authentications)
	}
}

func TestErrors(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"DELETE article": func(w http.ResponseWriter, _ *http.Request) { reply(w, "NOK", `"NOT_FOUND"`) },
		"GET account":    func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusBadGateway) },
	})
	err := client.DeleteArticle("HIS0067")
	var apiError *Error
	if !errors.As(err, &apiError) || apiError.Code != "NOT_FOUND" || apiError.Interface != "article" {
		t.Errorf("DeleteArticle() error = %v", err)
	}
	_, err = client.Account(Books)
	if err == nil || errors.As(err, &apiError) {
		t.Errorf("Account() should fail with the HTTP status, got %v", err)
	}
	client.APIKey = "WRONG"
	client.reset()
	err = client.DeleteArticle("HIS0067")
	if !errors.As(err, &apiError) || apiError.Code != "AUTHENTICATION_FAILED" {
		t.Errorf("DeleteArticle() with a wrong API key = %v", err)
	}
}

func TestArticleList(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"GET article_list": func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			if query.Get("field") != "isbn" || query.Get("showPrice") != "1" || query.Get("mediaType") != "0" || query.Has("showStock") {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `"9783518276368\t12.50\n9783423214124\t8.00\n"`)
		},
	})
	mediaType := Books
	lines, err := client.ArticleList(ArticleListOptions{Field: FieldISBN, ShowPrice: true, MediaType: &mediaType})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"9783518276368\t12.50", "9783423214124\t8.00"}
	if !slices.Equal(lines, want) {
		t.Errorf("ArticleList() = %q, want %q", lines, want)
	}
}

func TestImportFile(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"POST file_import": func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			if query.Get("dataType") != "1" || query.Get("mediaType") != "0" || query.Get("formatID") != "1" || query.Get("encoding") != "UTF-8" {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			file, header, err := r.FormFile("file")
			if err != nil {
				reply(w, "NOK", `"FILE_MISSING"`)
				return
			}
			content, _ := io.ReadAll(file)
			if header.Filename != "books.txt" || string(content) != "Titel\tPreis\n" {
				reply(w, "NOK", `"UPLOAD_FAILED"`)
				return
			}
			reply(w, "OK", `"SUCCESS"`)
		},
	})
	err := client.ImportFile(Upload{
		Name:     "books.txt",
		Content:  strings.NewReader("Titel\tPreis\n"),
		FileType: ArticleFile,
		DataType: DataReplace,
		FormatID: 1,
		Encoding: UTF8,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrders(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"GET order": func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			if query.Get("dateFrom") != "2026-10-01" || query.Get("dateTo") != "2026-10-29" {
				reply(w, "NOK", `"INVALID_DATE"`)
				return
			}
			reply(w, "OK", `[{
				"orderId": 4711, "orderDate": "2026-10-29", "orderTime": "12:30:00", "status": "WAITING_FOR_PAYMENT",
				"email": "kunde@example.com", "paymentId": 6, "buyerPositiveRatingPercentage": "100.0",
				"buyerTotalRatingNumber": 12, "buyerUsername": "kunde", "originalProvisionNet": "1.20", "currentProvisionNet": "1.20",
				"invoiceAddress": {"title": "Frau", "name": "Muster", "firstName": "Erika", "street": "Weg 1", "zip": "33602", "city": "Bielefeld", "country": "Deutschland"},
				"orderItems": [{"orderItemId": 815, "amount": 1, "orderNo": "HIS0067", "orderTitle": "Kritik und Krise", "author": "Koselleck, Reinhart",
					"singlePrice": "10.00", "totalPriceRebated": "10.00", "mediaType": 0, "status": "Ordered"}]
			}]`)
		},
	})
	orders, err := client.OrdersBetween(time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), time.Date(2026, 10, 29, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if len(orders) != 1 || len(orders[0].OrderItems) != 1 {
		t.Fatalf("OrdersBetween() = %+v", orders)
	}
	order := orders[0]
	if order.OrderID != 4711 || order.Status != WaitingForPayment || order.PaymentID != PayPal || order.InvoiceAddress.City != "Bielefeld" {
		t.Errorf("unexpected order %+v", order)
	}
	item := order.OrderItems[0]
	if item.OrderItemID != 815 || item.OrderNo != "HIS0067" || item.MediaType != Books || item.SinglePrice != "10.00" {
		t.Errorf("unexpected order item %+v", item)
	}
}

func TestOrderChanges(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"PUT order_status": func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("orderId") != "4711" || r.URL.Query().Get("status") != "SHIPPED_AND_PAID" {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `"SUCCESS"`)
		},
		"PUT order_message": func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("messageType") != "SHIPPING_NOTICE" || r.URL.Query().Get("additionalText") != "Danke!" {
				reply(w, "NOK", `"INVALID_MESSAGE_TYPE"`)
				return
			}
			reply(w, "OK", `"Ihre Bestellung wurde versandt. Danke!"`)
		},
		"PUT order_rating": func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("rating") != "4" || r.URL.Query().Has("comment") {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `"SUCCESS"`)
		},
		"PUT order_item_cancel": func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("orderItemId") != "815" || r.URL.Query().Get("mediaType") != "3" {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `"SUCCESS"`)
		},
	})
	err := client.SetOrderStatus(4711, ShippedAndPaid)
	if err != nil {
		t.Error(err)
	}
	text, err := client.SendOrderMessage(4711, ShippingNotice, "Danke!")
	if err != nil || text != "Ihre Bestellung wurde versandt. Danke!" {
		t.Errorf("SendOrderMessage() = %q, %v", text, err)
	}
	err = client.RateOrder(4711, RatingPerfect, "")
	if err != nil {
		t.Error(err)
	}
	err = client.CancelOrderItem(815, AudioBooks)
	if err != nil {
		t.Error(err)
	}
}

func TestPurchase(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"POST purchase": func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			if !slices.Equal(query["articleId"], []string{"A1", "A2"}) || !slices.Equal(query["paymentId"], []string{"12", "1"}) ||
				query.Get("acceptTerms") != "1" || query.Get("orderToken") != "order-1" || query.Get("addressId") != "0" {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `{"invalidArticles": ["A2"], "orders": [{"sellerId": 42, "status": "OK", "orderId": 4711,
				"price": 10.5, "shippingCost": "2.00", "totalCost": 12.5, "paymentId": 12,
				"articles": [{"articleId": "A1", "status": "OK"}]}]}`)
		},
	})
	result, err := client.Purchase(Purchase{
		ArticleIDs:     []string{"A1", "A2"},
		PaymentMethods: []PaymentMethod{PayPalREST, BankTransfer},
		OrderToken:     "order-1",
		AcceptTerms:    true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(result.InvalidArticles, []string{"A2"}) || len(result.Orders) != 1 {
		t.Fatalf("Purchase() = %+v", result)
	}
	order := result.Orders[0]
	if order.OrderID != 4711 || order.Price.String() != "10.5" || order.ShippingCost.String() != "2.00" || order.Articles[0].ArticleID != "A1" {
		t.Errorf("unexpected purchase order %+v", order)
	}
}
