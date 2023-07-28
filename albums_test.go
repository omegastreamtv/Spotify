package spotify

import (
	"net/http"
	"testing"
)

func TestGetAlbum(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_album.txt")
	defer server.Close()

	album, err := client.GetAlbum("0sNOF9WDwhWunNAHPD3Baj", "es")
	if err != nil {
		t.Fatal(err)
	}

	if album == nil {
		t.Error("Album is nil")
	}
}

func TestGetSeveralAlbums(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_albums.txt")
	defer server.Close()

	res, err := client.GetSeveralAlbums([]string{"382ObEPsp2rxGrnsizN5TX", "1A2GTWGtFfWp7KSQTwWOyo", "2noRn2Aes5aoNVsU6iWThc"}, "es")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}

func TestGetAlbumTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_album_tracks.txt")
	defer server.Close()

	res, err := client.GetAlbumTracks("4aawyAB9vmqN3uQ7FjRGTy", &GetAlbumTracksParams{
		Market: MarketSpain,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}

func TestGetUsersSavedAlbums(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_saved_albums.txt")
	defer server.Close()

	res, err := client.GetUsersSavedAlbums(&GetUsersSavedAlbumsParams{
		Limit:  20,
		Offset: 0,
		Market: MarketSpain,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	} else {
		if len(res.Items) != 20 {
			t.Errorf("Expected 20 items, got %d", len(res.Items))
		}
	}
}

func TestSaveAlbumsForCurrentUser(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "save_albums_for_current_user.txt")
	defer server.Close()

	err := client.SaveAlbumsForCurrentUser([]string{"382ObEPsp2rxGrnsizN5TX", "1A2GTWGtFfWp7KSQTwWOyo", "2noRn2Aes5aoNVsU6iWThc"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveAlbumsForCurrentUser(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "remove_albums_for_current_user.txt")
	defer server.Close()

	err := client.RemoveAlbumsForCurrentUser([]string{"382ObEPsp2rxGrnsizN5TX", "1A2GTWGtFfWp7KSQTwWOyo", "2noRn2Aes5aoNVsU6iWThc"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUsersSavedAlbums(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_users_saved_albums.txt")
	defer server.Close()

	res, err := client.CheckUsersSavedAlbums([]string{"382ObEPsp2rxGrnsizN5TX", "1A2GTWGtFfWp7KSQTwWOyo", "2noRn2Aes5aoNVsU6iWThc"})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	} else {
		if len(res) != 2 {
			t.Errorf("Expected 2 items, got %d", len(res))
		}
	}
}
