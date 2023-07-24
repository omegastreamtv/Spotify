package spotify

import (
	"net/http"
	"testing"
)

func TestGetShow(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_a_show.txt")
	defer server.Close()

	res, err := client.GetShow("38bS44xjbVVZ3No3ByF1dJ", "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "38bS44xjbVVZ3No3ByF1dJ" {
			t.Errorf("Expected %s, got %s", "38bS44xjbVVZ3No3ByF1dJ", res.ID)
		}
	}
}

func TestGetSeveralShows(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_shows.txt")
	defer server.Close()

	res, err := client.GetSeveralShows([]string{"5CfCWKI5pZ28U0uOzXkDHe", "5as3aKmN2k11yfDDDSrvaZ"}, "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res.Shows) != 2 {
			t.Errorf("Expected %d, got %d", 2, len(res.Shows))
		}
	} else {
		t.Error("Expected response")
	}
}

func TestGetShowEpisodes(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_show_episodes.txt")
	defer server.Close()

	res, err := client.GetShowEpisodes("38bS44xjbVVZ3No3ByF1dJ", &GetShowEpisodesParams{
		Market: "ES",
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res.Items) != 10 {
			t.Errorf("Expected %d, got %d", 10, len(res.Items))
		}
	} else {
		t.Error("Expected response")
	}
}

func TestGetUsersSavedShows(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_saved_shows.txt")
	defer server.Close()

	res, err := client.GetUsersSavedShows(&GetUsersSavedShowsParams{
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res.Items) != 2 {
			t.Errorf("Expected %d, got %d", 2, len(res.Items))
		}
	} else {
		t.Error("Expected response")
	}
}

func TestSaveShowsForCurrentUser(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "save_shows_for_current_user.txt")
	defer server.Close()

	err := client.SaveShowsForCurrentUser([]string{"5CfCWKI5pZ28U0uOzXkDHe", "5as3aKmN2k11yfDDDSrvaZ"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUsersSavedShows(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "remove_users_saved_shows.txt")
	defer server.Close()

	err := client.RemoveUsersSavedShows([]string{"5CfCWKI5pZ28U0uOzXkDHe", "5as3aKmN2k11yfDDDSrvaZ"}, "ES")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUsersSavedShows(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_users_saved_shows.txt")
	defer server.Close()

	err := client.CheckUsersSavedShows([]string{"5CfCWKI5pZ28U0uOzXkDHe", "5as3aKmN2k11yfDDDSrvaZ"})
	if err != nil {
		t.Fatal(err)
	}
}
