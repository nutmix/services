package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/skycoin/services/load-testing/loggy"
)

var cliPath string

func IninCmd(cliPathStr string) {
	cliPath = cliPathStr

}

// runCmd runs the skycoin/skycoin cli () (%GOPATH/bin/cli.exe)
// internal function called by getBalance etc.
func runCmd(args ...string) []byte {
	var c *exec.Cmd

	if runtime.GOOS == "windows" {

		var argSlice1 = []string{"/C", cliPath}
		var argSlice2 = append(argSlice1, args...)

		c = exec.Command("cmd", argSlice2...)

	} else {
		c = exec.Command(cliPath, args...)
	}

	bs, err := c.CombinedOutput()

	if err != nil {
		loggy.Error.Println("call response: %s", bs)
		loggy.Error.Println("runCmd got an err", err)
		os.Exit(1)
	}

	return bs
}
