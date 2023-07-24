package spotify

import (
	"net/http"
	"testing"
)

func TestGetAvailableGenreSeeds(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_available_genre_seeds.txt")
	defer server.Close()

	res, err := client.GetAvailableGenreSeeds()
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Genres) != 126 {
			t.Errorf("Expected %d, got %d", 126, len(res.Genres))
		}
	}
}
