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
		Market: "es",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}
