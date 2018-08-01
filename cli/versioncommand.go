package main

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use: "version",
}

var setCmd = &cobra.Command{
	Use: "set",
}

var incrementCmd = &cobra.Command{
	Use: "increment",
}

func incrementCommand(version Version) {
	var incrementMajorCmd = &cobra.Command{
		Use: "major",
	}

	var incrementMinorCmd = &cobra.Command{
		Use: "minor",
	}

	var incrementBugFixCmd = &cobra.Command{
		Use: "bugfix",
	}

	var incrementLabelCmd = &cobra.Command{
		Use: "label",
	}

	incrementCmd.AddCommand(incrementMajorCmd)
	incrementCmd.AddCommand(incrementMinorCmd)
	incrementCmd.AddCommand(incrementBugFixCmd)
	incrementCmd.AddCommand(incrementLabelCmd)

}

func VersionCommand(version Version) *cobra.Command {
	incrementCommand(version)

	versionCmd.AddCommand(setCmd)
	versionCmd.AddCommand(incrementCmd)
	return versionCmd
}
