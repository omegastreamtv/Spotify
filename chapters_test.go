package spotify

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAChapter(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_a_chapter.txt")
	defer server.Close()

	res, err := client.GetAChapter("0D5wENdkdwbqlrHoaJ9g29", "ES")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "0D5wENdkdwbqlrHoaJ9g29" {
			t.Errorf("Expected %s, got %s", "0D5wENdkdwbqlrHoaJ9g29", res.ID)
		}
	}
}

func TestGetSeveralChapters(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_chapters.txt")
	defer server.Close()

	res, err := client.GetSeveralChapters([]string{"0IsXVP0JmcB2adSE338GkK", "3ZXb8FKZGU0EHALYX6uCzU", "0D5wENdkdwbqlrHoaJ9g29"}, "ES")
	if err != nil {
		t.Fatal(err)
	}

	if res != nil {
		if len(res.Chapters) != 3 {
			t.Errorf("Expected %d, got %d", 3, len(res.Chapters))
		}
	} else {
		t.Error("Expected response")
	}
}
