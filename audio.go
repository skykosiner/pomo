package main

import (
	_ "embed"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed beep.mp3
var audio []byte

func playSound() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		slog.Error("Error getting cache dir", "error", err)
		return
	}

	beepFile := filepath.Join(cacheDir, "pomo", "beep.mp3")
	if err := os.WriteFile(beepFile, audio, 0644); err != nil {
		slog.Error("Error setting up beep.mp3 sound.")
		return
	}

	cmd := exec.Command("mpv", "--no-video", beepFile)
	if err := cmd.Run(); err != nil {
		slog.Error("Error playing beep.")
		return
	}
}
