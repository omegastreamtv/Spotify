package spotify

import (
	"net/http"
	"testing"
)

func TestSearchForItem(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "search_for_item.txt")
	defer server.Close()

	res, err := client.SearchForItem(&SearchForItemParams{
		Q:      "remaster,track:Doxy,artist:Miles,Davis",
		Type:   []string{"album"},
		Market: MarketSpain,
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		t.Error(err)
	}

	if len(res.Albums.Items) == 0 {
		t.Error("Expected at least one album")
	}
}
