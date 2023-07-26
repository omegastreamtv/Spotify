package spotify

import (
	"net/http"
	"testing"
)

func TestGetEpisode(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_an_episode.txt")
	defer server.Close()

	res, err := client.GetEpisode("512ojhOuo1ktJprKbVcKyQ", "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "5Xt5DXGzch68nYYamXrNxZ" {
			t.Errorf("Expected %s, got %s", "5Xt5DXGzch68nYYamXrNxZ", res.ID)
		}
	}
}

func TestGetSeveralEpisodes(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_episodes.txt")
	defer server.Close()

	res, err := client.GetSeveralEpisodes([]string{"77o6BIVlYM3msb4MMIL1jH", "0Q86acNRm6V9GYx55SXKwf"}, "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res.Episodes) != 2 {
			t.Errorf("Expected %d, got %d", 2, len(res.Episodes))
		}
	} else {
		t.Error("Expected response")
	}
}

func TestGetUsersSavedEpisodes(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_saved_episodes.txt")
	defer server.Close()

	res, err := client.GetUsersSavedEpisodes(&GetUsersSavedEpisodesParams{
		Market: MarketSpain,
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res.Items) != 1 {
			t.Errorf("Expected %d, got %d", 1, len(res.Items))
		}
	} else {
		t.Error("Expected response")
	}
}

func TestSaveEpisodesForCurrentUser(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "save_episodes_for_current_user.txt")
	defer server.Close()

	err := client.SaveEpisodesForCurrentUser([]string{"77o6BIVlYM3msb4MMIL1jH", "0Q86acNRm6V9GYx55SXKwf"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUsersSavedEpisodes(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "remove_users_saved_episodes.txt")
	defer server.Close()

	err := client.RemoveUsersSavedEpisodes([]string{"7ouMYWpwJ422jRcDASZB7P", "4VqPOruhp5EdPBeR92t6lQ", "2takcwOaAZWiXQijPHIx7B"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUsersSavedEpisodes(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_users_saved_episodes.txt")
	defer server.Close()

	res, err := client.CheckUsersSavedEpisodes([]string{"77o6BIVlYM3msb4MMIL1jH", "0Q86acNRm6V9GYx55SXKwf"})
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res) != 2 {
			t.Errorf("Expected %d, got %d", 2, len(res))
		}
	} else {
		t.Error("Expected response")
	}
}
