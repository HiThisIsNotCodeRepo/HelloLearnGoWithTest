# IO and sorting

## Thread safe naive implementation

```
package main

import (
	"encoding/json"
	"io"
	"sync"
)

type FileSystemPlayerStore struct {
	mu       sync.Mutex
	lock     bool
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	if !f.lock {
		f.mu.Lock()
		defer func() {
			f.lock = false
			f.mu.Unlock()
		}()
	}

	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	f.mu.Lock()
	f.lock = true
	defer func() {
		f.lock = false
		f.mu.Unlock()
	}()

	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	f.mu.Lock()
	f.lock = true
	defer func() {
		f.lock = false
		f.mu.Unlock()
	}()
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

```