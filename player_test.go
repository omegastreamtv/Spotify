package spotify

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetPlaybackState(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_playback_state.txt")
	defer server.Close()

	res, err := client.GetPlaybackState(&GetPlaybackStateParams{
		Market:          MarketUnitedStates,
		AdditionalTypes: "track",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.Item.isLeft {
			t.Logf("Is a track: %v", res.Item.isLeft)
		} else {
			t.Logf("Is a episode: %v", res.Item.right.Name)
		}

		if res.Timestamp != 1690379634686 {
			t.Errorf("Expected %d, got %d", 1690379634686, res.Timestamp)
		}
	}
}

func TestTransferPlayback(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "transfer_playback.txt")
	defer server.Close()

	err := client.TransferPlayback(&TransferPlaybackBody{
		DeviceIDs: []string{"74ASZWbe4lXaubB36ztrGX"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAvailableDevices(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_available_devices.txt")
	defer server.Close()

	res, err := client.GetAvailableDevices()
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if len(res.Devices) != 2 {
			t.Errorf("Expected %d, got %d", 2, len(res.Devices))
		}
	}
}

func TestGetCurrentlyPlayingTrack(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_currently_playing_track.txt")
	defer server.Close()

	res, err := client.GetCurrentlyPlayingTrack(&GetCurrentlyPlayingTrackParams{
		Market:          MarketUnitedStates,
		AdditionalTypes: "track",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.Item.isLeft {
			fmt.Println(res.Item)
		} else {
			fmt.Println(res.Item)
		}

		if res.Timestamp != 1690381580786 {
			t.Errorf("Expected %d, got %d", 1690381580786, res.Timestamp)
		}
	}
}

func TestStartResumePlayback(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "start_resume_playback.txt")
	defer server.Close()

	err := client.StartResumePlayback(&StartResumePlaybackParams{
		DeviceID: "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	}, &StartResumePlaybackBody{
		ContextURI: "spotify:album:1Je1IMUlBXcx1Fz0WE7oPT",
		PositionMS: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPausePlayback(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "pause_playback.txt")
	defer server.Close()

	err := client.PausePlayback(&PausePlaybackParams{
		DeviceID: "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSkipToNext(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "skip_to_next.txt")
	defer server.Close()

	err := client.SkipToNext(&SkipToNextParams{
		DeviceID: "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSkipToPrevious(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "skip_to_previous.txt")
	defer server.Close()

	err := client.SkipToPrevious(&SkipToPreviousParams{
		DeviceID: "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSeekToPosition(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "seek_to_position.txt")
	defer server.Close()

	err := client.SeekToPosition(&SeekToPositionParams{
		DeviceID:   "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
		PositionMS: 25000,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetRepeatMode(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "set_repeat_mode.txt")
	defer server.Close()

	err := client.SetRepeatMode(&SetRepeatModeParams{
		State:    "off",
		DeviceID: "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetPlaybackVolume(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "set_playback_volume.txt")
	defer server.Close()

	err := client.SetPlaybackVolume(&SetPlaybackVolumeParams{
		VolumePercent: 50,
		DeviceID:      "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTogglePlaybackShuffle(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "toggle_playback_shuffle.txt")
	defer server.Close()

	err := client.TogglePlaybackShuffle(&TogglePlaybackShuffleParams{
		State:    true,
		DeviceID: "0d1841b0976bae2a3a310dd74c0f3df354899bc8",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRecentlyPlayedTracks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_recently_played_tracks.txt")
	defer server.Close()

	res, err := client.GetRecentlyPlayedTracks(&GetRecentlyPlayedTracksParams{
		Limit: 20,
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

func TestGetTheUsersQueue(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_the_users_queue.txt")
	defer server.Close()

	res, err := client.GetTheUsersQueue()
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		fmt.Println(res.Queue, res.CurrentlyPlaying)
	}
}

func TestAddItemToPlaybackQueue(t *testing.T) {
	client, server := testClientFile(http.StatusNoContent, "add_item_to_playback_queue.txt")
	defer server.Close()

	err := client.AddItemToPlaybackQueue(&AddItemToPlaybackQueueParams{
		URI: "spotify:track:4iV5W9uYEdYUVa79Axb7Rh",
	})
	if err != nil {
		t.Fatal(err)
	}
}
