package spotify

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAnAudiobook(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_an_audiobook.txt")
	defer server.Close()

	audiobook, err := client.GetAnAudiobook("7iHfbu1YPACw6oZPAFJtqe", "es")
	if err != nil {
		t.Fatal(err)
	}

	if audiobook == nil {
		t.Error("Audiobook is nil")
	}
}

func TestGetSeveralAudiobooks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_audiobooks.txt")
	defer server.Close()

	res, err := client.GetSeveralAudiobooks([]string{"18yVqkdbdRvS24c0Ilj2ci", ",1HGw3J3NxZO1TP1BTtVhpZ", "7iHfbu1YPACw6oZPAFJtqe"}, "es")
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}

func TestGetAudiobookChapters(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_audiobook_chapters.txt")
	defer server.Close()

	res, err := client.GetAudiobookChapters("7iHfbu1YPACw6oZPAFJtqe", GetAudiobookChaptersParams{})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}

func TestGetUsersSavedAudioBooks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_users_saved_audiobooks.txt")
	defer server.Close()

	res, err := client.GetUsersSavedAudioBooks(10, 5)
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Response is nil")
	}
}

func TestSaveAudiobooksForCurrentUser(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "save_audiobooks_for_current_user.txt")
	defer server.Close()

	err := client.SaveAudiobooksForCurrentUser([]string{"18yVqkdbdRvS24c0Ilj2ci", "1HGw3J3NxZO1TP1BTtVhpZ", "7iHfbu1YPACw6oZPAFJtqe"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveUsersSavedAudiobooks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "remove_users_saved_audiobooks.txt")
	defer server.Close()

	err := client.RemoveUsersSavedAudiobooks([]string{"18yVqkdbdRvS24c0Ilj2ci", "1HGw3J3NxZO1TP1BTtVhpZ", "7iHfbu1YPACw6oZPAFJtqe"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckUsersSavedAudiobooks(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "check_users_saved_audiobooks.txt")
	defer server.Close()

	bools, err := client.CheckUsersSavedAudiobooks([]string{"18yVqkdbdRvS24c0Ilj2ci", "1HGw3J3NxZO1TP1BTtVhpZ", "7iHfbu1YPACw6oZPAFJtqe"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(bools)
}
