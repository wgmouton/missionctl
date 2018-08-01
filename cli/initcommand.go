package main

import "github.com/spf13/cobra"

func initCommand(version Version) *cobra.Command {

	initCmd := &cobra.Command{
		Use: "init",
	}

	return initCmd
}
