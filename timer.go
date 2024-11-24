package main

import (
	"fmt"
	"time"
)

type timer struct {
	startTime       int64
	endTime         int64
	lastUpdated     int64
	currentDuration int
}

func NewTimer(length int) timer {
	unixTime := time.Now().Unix()
	return timer{
		startTime:       unixTime,
		endTime:         unixTime + int64(length),
		lastUpdated:     unixTime,
		currentDuration: 0,
	}
}

func (t timer) updateCache() {
	// Check if cache file exists if not create it
}

func (t *timer) current() {
	t.currentDuration = t.currentDuration - int(t.lastUpdated)
	t.lastUpdated = time.Now().Unix()
}

func (t timer) String() string {
	t.current()
	return fmt.Sprintf("%d", t.currentDuration)
}
