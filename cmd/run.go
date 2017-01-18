package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/patdhlk/bf/compile"
	"github.com/patdhlk/bf/vm"
)

var CmdRun = &Command{
	Run:       runRun,
	UsageLine: "run ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	//cmdRun.Flag.StringVar(&flagA, "a", false, "")
}

// runRun executes build command and return exit code.
func runRun(args []string) int {
	//fmt.Println(args)
	fileName := args[0]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := compile.NewCompiler(string(code))
	instructions := compiler.Compile()

	m := vm.NewMachine(instructions, os.Stdin, os.Stdout)
	m.Execute()
	return 0
}
