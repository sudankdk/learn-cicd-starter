package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-api-key")

	got, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got != "my-api-key" {
		t.Fatalf("expected 'my-api-key', got %v", got)
	}

}

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	headers := http.Header{}
	got, err := GetAPIKey(headers)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if got != "" {
		t.Fatalf("expected empty string, got %v", got)
	}
}
