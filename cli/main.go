package main

import (
	// "fmt"
	// "strings"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Cmd(name string, args ...string) func(bool) (string, error) {
	fmt.Println(name, strings.Join(args, " "))
	output, err := exec.Command(name, args...).CombinedOutput()
	stringOutput := string(output)

	return func(fatal bool) (string, error) {
		if err != nil {
			go log.Print(stringOutput)
			if fatal {
				os.Exit(1)
			}
		} else {
			fmt.Println(stringOutput)
		}
		return stringOutput, err
	}
}

func main() {
	var about About
	about.Parse("../about.yml")

	var version = ParseVersion(about.Version)

	fmt.Println(version.GetVersion())

	version.IncrementMajor()
	fmt.Println(version.GetVersion())

	version.RemoveLabel()
	fmt.Println(version.GetVersion())

	version.RemoveMeta()
	fmt.Println(version.GetVersion())

	var rootCmd = &cobra.Command{
		Use: "missionctl",
	}

	initCmd := initCommand(version)
	versionCmd := VersionCommand(version)
	gitCmd := GitCommand(version)

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(gitCmd)
	rootCmd.Execute()
	/*



		Cmd("git", "status")(false)
		Cmd("git", "head")(false)

		var version = ParseVersion("hi")

		var about About

		fmt.Println(about.Parse("../about.yml").Version)
		fmt.Println(version.GetVersion())

		version.Minor = 20

		fmt.Println(version.GetVersion())*/

}
