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

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRout(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && string.Index(string(p), "<title>Home Pge</title>") > 0

		return statusOK && pageOK
	})
}
