package octadezx

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient("test-key")
	if c.BaseURL != "https://octadezx.com" {
		t.Fatalf("expected BaseURL https://octadezx.com, got %s", c.BaseURL)
	}
}
