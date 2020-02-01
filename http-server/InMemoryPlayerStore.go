package main

import "sync"

// NewInMemoryPlayerStore initializes InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

// InMemoryPlayerStore implementation
type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

// GetPlayerScore is required to suffice the interface
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

// RecordWin records the win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}
