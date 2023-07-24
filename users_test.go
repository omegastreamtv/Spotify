package spotify

import (
	"net/http"
	"testing"
)

func TestGetCurrentUsersProfile(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_current_users_profile.txt")
	defer server.Close()

	res, err := client.GetCurrentUsersProfile()
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "testuser" {
			t.Errorf("Expected %s, got %s", "testuser", res.ID)
		}
	}
}

func TestGetUsersTopArtists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_top_artists.txt")
	defer server.Close()

	res, err := client.GetUsersTopArtists(&GetUsersTopItemsParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Items) != 20 {
			t.Errorf("Expected %d, got %d", 20, len(res.Items))
		}
	}
}

func TestGetUsersTopTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_top_tracks.txt")
	defer server.Close()

	res, err := client.GetUsersTopTracks(&GetUsersTopItemsParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Items) != 20 {
			t.Errorf("Expected %d, got %d", 20, len(res.Items))
		}
	}
}

func TestGetUsersProfile(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_profile.txt")
	defer server.Close()

	res, err := client.GetUsersProfile("smedjan")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "smedjan" {
			t.Errorf("Expected %s, got %s", "smedjan", res.ID)
		}
	}
}

func TestFollowPlaylist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "follow_playlist.txt")
	defer server.Close()

	err := client.FollowPlaylist("3cEYpjA9oz9GiPac4AsH4n", &FollowPlaylistBody{
		Public: false,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnfollowPlaylist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "unfollow_playlist.txt")
	defer server.Close()

	err := client.UnfollowPlaylist("3cEYpjA9oz9GiPac4AsH4n")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFollowedArtists(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_followed_artists.txt")
	defer server.Close()

	res, err := client.GetFollowedArtists(&GetFollowedArtistsParams{
		Type: "artist",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Artists.Items) != 1 {
			t.Errorf("Expected %d, got %d", 1, len(res.Artists.Items))
		}
	}
}

func TestFollowArtistsOrUsers(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "follow_artists_or_users.txt")
	defer server.Close()

	err := client.FollowArtistsOrUsers("artist", []string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}

	err = client.FollowArtistsOrUsers("user", []string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnfollowArtistsOrUsers(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "unfollow_artists_or_users.txt")
	defer server.Close()

	err := client.UnfollowArtistsOrUsers("artist", []string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}

	err = client.UnfollowArtistsOrUsers("user", []string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckIfUserFollowsArtistsOrUsers(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_if_user_follows_artists_or_users.txt")
	defer server.Close()

	res, err := client.CheckIfUserFollowsArtistsOrUsers("artist", []string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(*res) != 3 {
			t.Errorf("Expected %d, got %d", 3, len(*res))
		}
	}

	res, err = client.CheckIfUserFollowsArtistsOrUsers("user", []string{"2CIMQHirSU0MQqyYHq0eOx", "57dN52uHvrHOxijzpIgu3E", "1vCWHaC5f2uS3yhpwWbIA6"})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(*res) != 3 {
			t.Errorf("Expected %d, got %d", 3, len(*res))
		}
	}
}

func TestCheckIfUsersFollowPlaylist(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_if_users_follow_playlist.txt")
	defer server.Close()

	res, err := client.CheckIfUsersFollowPlaylist("3cEYpjA9oz9GiPac4AsH4n", []string{"jmperezperez", "thelinmichael", "wizzler"})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(*res) != 3 {
			t.Errorf("Expected %d, got %d", 3, len(*res))
		}
	}
}
