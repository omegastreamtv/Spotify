package spotify

import (
	"net/http"
	"testing"
)

func TestGetTrack(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_track.txt")
	defer server.Close()

	res, err := client.GetTrack("11dFghVXANMlKmJXsNCbNl", "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "11dFghVXANMlKmJXsNCbNl" {
			t.Errorf("Expected %s, got %s", "11dFghVXANMlKmJXsNCbNl", res.ID)
		}
	}
}

func TestGetSeveralTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_tracks.txt")
	defer server.Close()

	res, err := client.GetSeveralTracks([]string{"7ouMYWpwJ422jRcDASZB7P", "4VqPOruhp5EdPBeR92t6lQ", "2takcwOaAZWiXQijPHIx7B"}, "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Tracks) != 3 {
			t.Errorf("Expected %d, got %d", 3, len(res.Tracks))
		}
	}
}

func TestGetUsersSavedTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_saved_tracks.txt")
	defer server.Close()

	res, err := client.GetUsersSavedTracks(&GetUsersSavedTracksParams{
		Market: "ES",
		Limit:  20,
		Offset: 0,
	})
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

func TestSaveTracksForCurrentUser(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "save_tracks_for_current_user.txt")
	defer server.Close()

	err := client.SaveTracksForCurrentUser([]string{"7ouMYWpwJ422jRcDASZB7P", "4VqPOruhp5EdPBeR92t6lQ", "2takcwOaAZWiXQijPHIx7B"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUsersSavedTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "remove_users_saved_tracks.txt")
	defer server.Close()

	err := client.RemoveUsersSavedTracks([]string{"7ouMYWpwJ422jRcDASZB7P", "4VqPOruhp5EdPBeR92t6lQ", "2takcwOaAZWiXQijPHIx7B"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUsersSavedTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_users_saved_tracks.txt")
	defer server.Close()

	res, err := client.CheckUsersSavedTracks([]string{"7ouMYWpwJ422jRcDASZB7P", "4VqPOruhp5EdPBeR92t6lQ", "2takcwOaAZWiXQijPHIx7B"})
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

func TestGetMultiTracksAudioFeatures(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_multi_tracks_audio_features.txt")
	defer server.Close()

	res, err := client.GetMultiTracksAudioFeatures([]string{"7ouMYWpwJ422jRcDASZB7P", "4VqPOruhp5EdPBeR92t6lQ", "2takcwOaAZWiXQijPHIx7B"})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.AudioFeatures) != 3 {
			t.Errorf("Expected %d, got %d", 3, len(res.AudioFeatures))
		}
	}
}

func TestGetSingleTracksAudioFeatures(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_single_tracks_audio_features.txt")
	defer server.Close()

	res, err := client.GetSingleTracksAudioFeatures("11dFghVXANMlKmJXsNCbNl")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "11dFghVXANMlKmJXsNCbNl" {
			t.Errorf("Expected %s, got %s", "11dFghVXANMlKmJXsNCbNl", res.ID)
		}
	}
}

func TestGetTracksAudioAnalysis(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_tracks_audio_analysis.txt")
	defer server.Close()

	res, err := client.GetTracksAudioAnalysis("11dFghVXANMlKmJXsNCbNl")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.Meta.AnalyzerVersion != "4.0.0" {
			t.Errorf("Expected %s, got %s", "4.0.0", res.Meta.AnalyzerVersion)
		}
	}
}

func TestGetRecommendations(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_recommendations.txt")
	defer server.Close()

	res, err := client.GetRecommendations(GetRecommendationsBody{
		SeedArtists: "4NHQUGzhtTLFvgF5SZesLK",
		SeedGenres:  "classical,country",
		SeedTracks:  "0c6xIDDpzE81m2q797ordA",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Tracks) != 20 {
			t.Errorf("Expected %d, got %d", 20, len(res.Tracks))
		}
	}
}
