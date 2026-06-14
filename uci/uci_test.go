package uci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tamnd/uci-cli/uci"
)

func datasetsPayload(names []string) []byte {
	type wireDataset struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	type resp struct {
		Status     int          `json:"status"`
		StatusText string       `json:"statusText"`
		Data       []wireDataset `json:"data"`
	}
	data := make([]wireDataset, len(names))
	for i, n := range names {
		data[i] = wireDataset{ID: i + 1, Name: n}
	}
	b, _ := json.Marshal(resp{Status: 200, StatusText: "OK", Data: data})
	return b
}

func TestList(t *testing.T) {
	payload := datasetsPayload([]string{"Abalone", "Adult", "Iris"})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") == "" {
			t.Error("request carried no User-Agent")
		}
		_, _ = w.Write(payload)
	}))
	defer srv.Close()

	cfg := uci.DefaultConfig()
	cfg.BaseURL = srv.URL
	cfg.Rate = 0

	c := uci.NewClient(cfg)
	datasets, err := c.List(context.Background(), 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(datasets) != 3 {
		t.Fatalf("got %d datasets, want 3", len(datasets))
	}
	if datasets[0].Name != "Abalone" {
		t.Errorf("name = %q, want Abalone", datasets[0].Name)
	}
	if datasets[0].Rank != 1 {
		t.Errorf("rank = %d, want 1", datasets[0].Rank)
	}
	if datasets[0].ID != 1 {
		t.Errorf("id = %d, want 1", datasets[0].ID)
	}
}

func TestListLimit(t *testing.T) {
	payload := datasetsPayload([]string{"A", "B", "C", "D", "E"})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(payload)
	}))
	defer srv.Close()

	cfg := uci.DefaultConfig()
	cfg.BaseURL = srv.URL
	cfg.Rate = 0

	c := uci.NewClient(cfg)
	datasets, err := c.List(context.Background(), 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(datasets) != 3 {
		t.Fatalf("got %d datasets, want 3 (limit applied)", len(datasets))
	}
}

func TestSearch(t *testing.T) {
	payload := datasetsPayload([]string{"Iris", "Iris Plant", "Abalone"})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(payload)
	}))
	defer srv.Close()

	cfg := uci.DefaultConfig()
	cfg.BaseURL = srv.URL
	cfg.Rate = 0

	c := uci.NewClient(cfg)
	datasets, err := c.Search(context.Background(), "iris", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(datasets) != 2 {
		t.Fatalf("got %d datasets, want 2", len(datasets))
	}
}
