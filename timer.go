package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type timer struct {
	StartTime       int64 `json:"StartTime"`
	EndTime         int64 `json:"EndTime"`
	LastUpdated     int64 `json:"LastUpdated"`
	CurrentDuration int   `json:"CurrentDuration"`
}

func NewTimer(length int) timer {
	unixTime := time.Now().Unix()
	return timer{
		StartTime:       unixTime,
		EndTime:         unixTime + int64(length),
		LastUpdated:     unixTime,
		CurrentDuration: length,
	}
}

func (t timer) updateCache() {
	// Create the cache directory if it doesn't exist
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		slog.Error("Error getting cache dir", "error", err)
		return
	}

	pomoCache := filepath.Join(cacheDir, "pomo")
	if err := os.MkdirAll(pomoCache, 0755); err != nil {
		slog.Error("Error creating cache dir.", "error", err, "path", pomoCache)
		return
	}

	bytes, err := json.Marshal(t)
	if err != nil {
		slog.Error("Error marshalling JSON of the timer.", "error", err, "timer", t)
		return
	}

	if err := os.WriteFile(filepath.Join(pomoCache, "pomo.json"), bytes, 0644); err != nil {
		slog.Error("Error updating JSON cache file.", "error", err, "timer", t)
		return
	}
}

func (t *timer) current() {
	now := time.Now().Unix()
	elapsed := now - t.LastUpdated
	t.CurrentDuration -= int(elapsed)

	if t.CurrentDuration < 0 {
		t.CurrentDuration = 0
	}

	t.LastUpdated = now
	t.updateCache()
}

func (t timer) String() string {
	t.current()

	minutes := t.CurrentDuration / 60
	seconds := t.CurrentDuration % 60

	return fmt.Sprintf("🍅 %02d:%02d", minutes, seconds)
}

func (t timer) delete() {
	t.CurrentDuration = 0
	t.current()
}
