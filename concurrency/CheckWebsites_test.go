package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://sdfasf.qwe" {
		return false
	}
	return true
}
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://nikita.sh",
		"http://sdfasf.qwe",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://nikita.sh":  true,
		"http://sdfasf.qwe": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %v, but expected %v", got, want)
	}
}
