package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	articlesTestList := GetAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(articlesTestList) != len(ArticleList) {
		t.Fail()
	}
	// Check that each member is identical
	for i, v := range articlesTestList {
		if v.Content != ArticleList[i].Content || v.ID != ArticleList[i].ID || v.Title != ArticleList[i].Title {
			t.Fail()
			break
		}
	}
}
