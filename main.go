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
			if err != nil || t.CurrentDuration == 0 {
				overTime := int(time.Now().Unix() - t.EndTime)

				if overTime <= 8 {
					visible := true
					for range 10 {
						if visible {
							fmt.Printf("\r%s", "00:00")
						} else {
							fmt.Print("\r      ")
						}

						visible = !visible
						time.Sleep(1 * time.Second)

						fmt.Print("\r")
					}
					return
				}

				fmt.Fprintln(os.Stderr, "There is currently no timer running. Use pomo new to create a new timer.")
				return
			}

			fmt.Println(t)
		},
	}

	commands := []cobra.Command{
		{
			Use:   "new",
			Short: "Start a new timer",
			Run: func(cmd *cobra.Command, args []string) {
				t, err := loadTimer()
				if err == nil && t.CurrentDuration > 0 {
					fmt.Println(t)
					return
				}

				// 25 minutes
				length := 1500
				if len(args) > 0 {
					l, err := strconv.Atoi(args[0])
					if err != nil {
						return
					}

					length = l
				}

				t = NewTimer(length)
				fmt.Println(t)
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
	}

	for _, command := range commands {
		rootCmd.AddCommand(&command)
	}

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
