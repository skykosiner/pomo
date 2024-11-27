package main

import "testing"

func TestGetAudio(t *testing.T) {
	if len(audio) <= 0 {
		t.Fatal("Embeded audio file not wokring.")
	}
}
