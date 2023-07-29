package spotify

import (
	"net/http"
	"testing"
)

func TestGetPlaylist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_playlist.txt")
	defer server.Close()

	res, err := client.GetPlaylist("37i9dQZF1DXcBWIGoYBM5M", &GetPlaylistParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestChangePlaylistDetails(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "change_playlist_details.txt")
	defer server.Close()

	err := client.ChangePlaylistDetails("3cEYpjA9oz9GiPac4AsH4n", &ChangePlaylistDetailsBody{
		Name:        "New Playlist Name",
		Public:      false,
		Description: "New playlist description",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPlaylistItems(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_playlist_items.txt")
	defer server.Close()

	res, err := client.GetPlaylistItems("3cEYpjA9oz9GiPac4AsH4n", &GetPlaylistItemsParams{
		Market: MarketSpain,
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestUpdatePlaylistItems(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "update_playlist_items.txt")
	defer server.Close()

	res, err := client.UpdatePlaylistItems("3cEYpjA9oz9GiPac4AsH4n", &UpdatePlaylistItemsBody{
		RangeStart:   1,
		InsertBefore: 3,
		RangeLength:  2,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}
