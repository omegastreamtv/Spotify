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

func TestAddItemsToPlaylist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "add_items_to_playlist.txt")
	defer server.Close()

	res, err := client.AddItemsToPlaylist("3cEYpjA9oz9GiPac4AsH4n", &AddItemsToPlaylistBody{
		URIs: []string{
			"spotify:track:4iV5W9uYEdYUVa79Axb7Rh",
			"spotify:track:1301WleyT98MSxVHPZCA6M",
		},
		Position: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestRemovePlaylistItems(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "remove_playlist_items.txt")
	defer server.Close()

	res, err := client.RemovePlaylistItems("3cEYpjA9oz9GiPac4AsH4n", &RemovePlaylistItemsBody{
		Tracks: []struct {
			URI string "json:\"uri\""
		}{
			{
				URI: "spotify:track:4iV5W9uYEdYUVa79Axb7Rh",
			},
		},
		SnapshotID: "abc",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestGetCurrentUsersPlaylists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_current_users_playlists.txt")
	defer server.Close()

	res, err := client.GetCurrentUsersPlaylists(&GetCurrentUsersPlaylistsParams{
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

func TestGetUsersPlaylists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_playlists.txt")
	defer server.Close()

	res, err := client.GetUsersPlaylists("smedjan", &GetUsersPlaylistsParams{
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

func TestCreatePlaylist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "create_playlist.txt")
	defer server.Close()

	res, err := client.CreatePlaylist("smedjan", &CreatePlaylistBody{
		Name:        "New Playlist",
		Public:      false,
		Description: "New playlist description",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestGetFeaturedPlaylists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_featured_playlists.txt")
	defer server.Close()

	res, err := client.GetFeaturedPlaylists(&GetFeaturedPlaylistsParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestGetCategorysPlaylists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_categorys_playlists.txt")
	defer server.Close()

	res, err := client.GetCategorysPlaylists("party", &GetCategorysPlaylistsParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestGetPlaylistCoverImage(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_playlist_cover_image.txt")
	defer server.Close()

	res, err := client.GetPlaylistCoverImage("3cEYpjA9oz9GiPac4AsH4n")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	}
}

func TestAddCustomPlaylistCoverImage(t *testing.T) {
	client, server := testClientFile(http.StatusAccepted, "add_custom_playlist_cover_image.txt")
	defer server.Close()

	err := client.AddCustomPlaylistCoverImage("3cEYpjA9oz9GiPac4AsH4n", []byte("/9j/2wCEABoZGSccJz4lJT5CLy8vQkc9Ozs9R0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0cBHCcnMyYzPSYmPUc9Mj1HR0dEREdHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR0dHR//dAAQAAf/uAA5BZG9iZQBkwAAAAAH/wAARCAABAAEDACIAAREBAhEB/8QASwABAQAAAAAAAAAAAAAAAAAAAAYBAQAAAAAAAAAAAAAAAAAAAAAQAQAAAAAAAAAAAAAAAAAAAAARAQAAAAAAAAAAAAAAAAAAAAD/2gAMAwAAARECEQA/AJgAH//Z"))
	if err != nil {
		t.Fatal(err)
	}
}
