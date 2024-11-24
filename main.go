package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "pomo - Terminal Pomodoro Tool",
		Use:   "pomo",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(os.Stdout, "test")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
