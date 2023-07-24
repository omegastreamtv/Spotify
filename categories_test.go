package spotify

import (
	"net/http"
	"testing"
)

func TestGetSeveralBrowseCategories(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_several_browse_categories.txt")
	defer server.Close()

	res, err := client.GetSeveralBrowseCategories(&GetSeveralBrowseCategoriesParams{
		Country: "es",
		Locale:  "sv_EN",
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.Categories.Items[0].ID != "toplists" {
			t.Errorf("Expected %s, got %s", "toplists", res.Categories.Items[0].ID)
		}
	}
}

func TestGetSingleBrowseCategory(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "get_single_browse_category.txt")
	defer server.Close()

	res, err := client.GetSingleBrowseCategory("toplists", &GetSingleBrowseCategoryParams{
		Country: "es",
		Locale:  "sv_EN",
	})
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Error("Expected response")
	} else {
		if res.ID != "0JQ5DAqbMKFRY5ok2pxXJ0" {
			t.Errorf("Expected %s, got %s", "0JQ5DAqbMKFRY5ok2pxXJ0", res.ID)
		}
	}
}
