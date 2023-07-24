package spotify

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetArtist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_artist.txt")
	defer server.Close()

	artist, err := client.GetArtist("0TnOYISbd1XYRBk9myaseg")
	if err != nil {
		t.Fatal(err)
	}

	if artist != nil {
		if artist.ID != "0TnOYISbd1XYRBk9myaseg" {
			t.Error("Wrong artist ID")
		}
	} else {
		t.Error("Artist is nil")
	}
}

func TestGetSeveralArtists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_artists.txt")
	defer server.Close()

	artist, err := client.GetSeveralArtists([]string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}

	if artist != nil {
		for _, a := range artist.Artists {
			fmt.Println(a.ID)
		}

		fmt.Println(artist)
	} else {
		t.Error("Artist is nil")
	}
}

func TestGetArtistsAlbums(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_artists_albums.txt")
	defer server.Close()

	res, err := client.GetArtistsAlbums("0TnOYISbd1XYRBk9myaseg", &GetArtistsAlbumsParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if res.Items[0].Artists[0].ID != "0TnOYISbd1XYRBk9myaseg" {
			t.Error("Wrong artist ID")
		}
	} else {
		t.Error("Response is nil")
	}
}

func TestGetArtistsTopTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_artists_top_tracks.txt")
	defer server.Close()

	res, err := client.GetArtistsTopTracks("0TnOYISbd1XYRBk9myaseg", &GetArtistsTopTracksParams{
		Market: "es",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if res.Tracks[0].Album.ID != "4rG0MhkU6UojACJxkMHIXB" {
			t.Error("Wrong track ID")
		}
	} else {
		t.Error("Response is nil")
	}
}

func TestGetArtistsRelatedArtists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_artists_related_artists.txt")
	defer server.Close()

	res, err := client.GetArtistsRelatedArtists("0TnOYISbd1XYRBk9myaseg")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}
