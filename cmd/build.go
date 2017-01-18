package cmd

import "fmt"

var CmdBuild = &Command{
	Run:       runBuild,
	UsageLine: "build ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdBuild.Flag.BoolVar(&flagA, "a", false, "")
}

// runBuild executes build command and return exit code.
func runBuild(args []string) int {
	fmt.Println("not yet implemented")
	return 0
}
