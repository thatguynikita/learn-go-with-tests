package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemPlayerStore implements Store that can be save to disk
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func initializePlayerDBFile(file *os.File) (err error) {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from %s: %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

// NewFileSystemPlayerStore instantiates our Store
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from %s: %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

// GetLeague returns list of Players sorted from most to least wins
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

// GetPlayerScore returns player's score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin records the win
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
