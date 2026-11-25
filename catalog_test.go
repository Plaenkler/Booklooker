package booklooker

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestCreateArticle(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"POST article": func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseMultipartForm(1 << 20)
			if err != nil || r.FormValue("orderNo") != "LIT0001" || r.FormValue("price") != "12.50" || r.FormValue("condition") != "2" ||
				r.FormValue("signed") != "1" || r.FormValue("category") != "" {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `{"orderNo":"LIT0001","images":0,"articleStatus":"PENDING"}`)
		},
	})
	result, err := client.CreateArticle(Article{OrderNo: "LIT0001", Title: "Effi Briest", Condition: ConditionLightWear, Price: 1250, Signed: true})
	if err != nil || result.Status != ArticlePending || result.OrderNo != "LIT0001" {
		t.Errorf("CreateArticle() = %+v, %v", result, err)
	}
}

// A rejected article names each field, so that the seller can fix it.
func TestUpdateArticleRejected(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"PUT article": func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil || r.URL.Query().Get("orderNo") != "LIT0001" || r.PostForm.Get("orderNo") != "" || r.PostForm.Get("title") != "Effi Briest" {
				reply(w, "NOK", `"NOT_FOUND"`)
				return
			}
			_, _ = w.Write([]byte(`{"status":"NOK","returnValue":"INVALID_PARAMETERS","errors":[{"field":"price","code":"OUT_OF_RANGE"}]}`))
		},
	})
	_, err := client.UpdateArticle(Article{OrderNo: "LIT0001", Title: "Effi Briest", Price: 10})
	var apiError *Error
	if !errors.As(err, &apiError) || apiError.Code != "INVALID_PARAMETERS" || len(apiError.Fields) != 1 || apiError.Fields[0].Field != "price" {
		t.Errorf("UpdateArticle() = %v", err)
	}
}

// Categories have numbers, languages codes, both are read as text.
func TestValueList(t *testing.T) {
	cases := []struct {
		name   string
		answer string
		want   ValueID
	}{
		{name: "category", answer: `[{"id":12,"name":"Belletristik","categories":[{"id":1201,"name":"Romane"}]}]`, want: "12"},
		{name: "language", answer: `[{"id":"de","name":"Deutsch"}]`, want: "de"},
	}
	for _, current := range cases {
		t.Run(current.name, func(t *testing.T) {
			client, _ := newTestClient(t, map[string]http.HandlerFunc{
				"GET value_list": func(w http.ResponseWriter, _ *http.Request) { reply(w, "OK", current.answer) },
			})
			values, err := client.ValueList(ValueField(current.name), Books)
			if err != nil || len(values) != 1 || values[0].ID != current.want {
				t.Errorf("ValueList() = %+v, %v", values, err)
			}
		})
	}
}

func TestPriceOverview(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"GET price_overview": func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("isbn") != "9783150000014" {
				reply(w, "NOK", `"INVALID_PARAMETERS"`)
				return
			}
			reply(w, "OK", `{"offers":2,"conditions":[{"condition":0,"name":"Neuware","offers":0,"minPrice":null},{"condition":2,"name":"leichte Gebrauchsspuren","offers":2,"minPrice":4.5},{"condition":null,"name":"ohne Angabe","offers":0,"minPrice":null}]}`)
		},
	})
	overview, err := client.PriceOverview("9783150000014")
	if err != nil || overview.Offers != 2 || len(overview.Conditions) != 3 {
		t.Fatalf("PriceOverview() = %+v, %v", overview, err)
	}
	if overview.Conditions[1].MinPrice != "4.5" || overview.Conditions[0].MinPrice != "" || overview.Conditions[2].Condition != "" {
		t.Errorf("conditions %+v", overview.Conditions)
	}
}

// Interval spaces the requests, the authentication included.
func TestInterval(t *testing.T) {
	client, _ := newTestClient(t, map[string]http.HandlerFunc{
		"GET import_status": func(w http.ResponseWriter, _ *http.Request) { reply(w, "OK", "0") },
	})
	client.Interval = 50 * time.Millisecond
	started := time.Now()
	for range 2 {
		_, err := client.ImportStatus()
		if err != nil {
			t.Fatal(err)
		}
	}
	elapsed := time.Since(started)
	if elapsed < 2*client.Interval {
		t.Errorf("three requests took %v", elapsed)
	}
}
