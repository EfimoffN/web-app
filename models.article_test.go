package main
import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := TestGetAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content || 
			v.ID != articleList[i].ID || 
			v.Title != articleList[i].Title{
				t.Fail()
				break
			}
	}
}
