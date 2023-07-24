package spotify

import (
	"net/http"
	"testing"
)

func TestGetAvailableMarkets(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_available_markets.txt")
	defer server.Close()

	res, err := client.GetAvailableMarkets()
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Markets) != 184 {
			t.Errorf("Expected %d, got %d", 184, len(res.Markets))
		}
	}
}
