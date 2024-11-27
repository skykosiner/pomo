package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func loadTimer() (timer, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return timer{}, fmt.Errorf("failed to get cache dir: %w", err)
	}

	pomoCache := filepath.Join(cacheDir, "pomo", "pomo.json")

	if _, err := os.Stat(pomoCache); os.IsNotExist(err) {
		return timer{}, fmt.Errorf("no saved timer state found")
	}

	data, err := os.ReadFile(pomoCache)
	if err != nil {
		return timer{}, fmt.Errorf("failed to read timer state: %w", err)
	}

	var t timer
	if err := json.Unmarshal(data, &t); err != nil {
		return timer{}, fmt.Errorf("failed to unmarshal timer state: %w", err)
	}

	t.current()

	return t, nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "pomo",
		Short: "pomo - Terminal Pomodoro Tool",
		Run: func(cmd *cobra.Command, args []string) {
			t, err := loadTimer()
			if err != nil {
				return
			}

			t.print()
		},
	}

	commands := []cobra.Command{
		{
			Use:   "new",
			Short: "Start a new timer",
			Run: func(cmd *cobra.Command, args []string) {
				t, err := loadTimer()
				if err == nil && t.CurrentDuration > 0 {
					return
				}

				// 25 minutes
				length := 1500
				if len(args) > 0 {
					switch args[0] {
					case "hour":
						length = (60 - time.Now().Minute()) * 60
					default:
						l, err := strconv.Atoi(args[0])
						if err != nil {
							return
						}

						length = l
					}
				}

				t = NewTimer(length)
				t.current()
			},
		},
		{
			Use:   "stop",
			Short: "Stop the current timer",
			Run: func(cmd *cobra.Command, args []string) {
				t, err := loadTimer()
				if err != nil {
					return
				}

				t.delete()
			},
		},
		{
			Use:   "pause",
			Short: "Pause the timer",
			Run: func(cmd *cobra.Command, args []string) {
				t, err := loadTimer()
				if err != nil {
					return
				}

				t.pause()
			},
		},
		{
			Use:   "resume",
			Short: "Resume the timer",
			Run: func(cmd *cobra.Command, args []string) {
				t, err := loadTimer()
				if err != nil {
					return
				}

				t.resume()
			},
		},
	}

	for _, command := range commands {
		rootCmd.AddCommand(&command)
	}

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
