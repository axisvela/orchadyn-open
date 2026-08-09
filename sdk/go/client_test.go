package orchadyn

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost || request.URL.Path != "/plans:generate" {
			t.Fatalf("unexpected request: %s %s", request.Method, request.URL.Path)
		}
		var payload PlanningRequest
		if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
			t.Fatal(err)
		}
		if len(payload.Goals) != 1 || payload.Goals[0].ID != "goal-1" {
			t.Fatalf("unexpected payload: %#v", payload)
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"plan":{"id":"plan-1"}}`))
	}))
	defer server.Close()

	client, err := NewClient(server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	var response struct {
		Plan struct {
			ID string `json:"id"`
		} `json:"plan"`
	}
	err = client.Generate(context.Background(), PlanningRequest{Goals: []Goal{{ID: "goal-1"}}}, &response)
	if err != nil {
		t.Fatal(err)
	}
	if response.Plan.ID != "plan-1" {
		t.Fatalf("unexpected response: %#v", response)
	}
}

func TestReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = writer.Write([]byte(`{"error":"invalid plan"}`))
	}))
	defer server.Close()

	client, err := NewClient(server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Health(context.Background())
	apiError, ok := err.(*APIError)
	if !ok || apiError.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("expected API error, got %v", err)
	}
}
