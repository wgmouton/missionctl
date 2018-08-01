package main

import "github.com/spf13/cobra"

func GitCommand(version Version) *cobra.Command {
	var gitCmd = &cobra.Command{
		Use: "gitflow",
	}

	return gitCmd
}
