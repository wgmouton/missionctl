package main

import (
	"github.com/spf13/cobra"
)

var startRelease = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		Cmd("git", "fetch")
		Cmd("git", "fetch")
	},
}

var releaseCmd = &cobra.Command{
	Use: "release",
}

func GitCommand(version Version) *cobra.Command {
	var gitCmd = &cobra.Command{
		Use: "gitflow",
	}

	releaseCmd.AddCommand(startRelease)
	gitCmd.AddCommand(releaseCmd)
	return gitCmd
}
