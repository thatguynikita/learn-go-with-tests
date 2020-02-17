package main

import (
	"io/ioutil"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()

	tape := &tape{file}
	tape.Write([]byte("ABC"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "ABC"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
