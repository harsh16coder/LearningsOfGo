package concurrrency

import (
	"reflect"
	"testing"
)

func TestUrls(t *testing.T) {

	urls := []string{
		"google.com",
		"facebook.com",
		"fdegahia.in",
	}

	want := map[string]bool{
		"google.com":   true,
		"facebook.com": true,
		"fdegahia.in":  false,
	}

	got := fetchurls(urls, mockwebsiteChecker)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func mockwebsiteChecker(url string) bool {
	return url != "fdegahia.in"
}
