package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	t.Run("league sorted", func(t *testing.T) {
		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Nikita")
		got := store.GetPlayerScore("Nikita")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tempfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempfile.Write([]byte(initialData))

	removeFile := func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}

	return tempfile, removeFile
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error, but got one %v", err)
	}
}
