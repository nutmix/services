// Package provides command line utils to test transfers between wallets.
// Author Simon Hobbs
// code status: draft.  review welcome.
// dependencies: github.com/skycoin/skycoin
// see README.md for more info
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skycoin/services/load-testing/cmd"
	"github.com/skycoin/services/load-testing/loggy"
	"github.com/skycoin/skycoin/src/cipher"
)

const DEBUG = true
const COINDP = 8

var cliPath string

// useage: go run load.go -help
// utility to transfer tokens randomly between a number of accounts specified by a seed.
func main() {
	loggy.InitLogging(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	var goPath = (os.Getenv("GOPATH"))                                                     // this is really a constant.
	cliPath = goPath + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "cli" // used to to call cli

	// Handle simple command line. This can be moved to use urfave/cli later if accepted.
	seedPtr := flag.String("seed", "", "Any set of number of ASCII chars, ideally random for gnerating test addresses")
	numaPtr := flag.Int("numa", 2, "Number of addresses to generate from the seed")
	msPtr := flag.Int("ms", 0, "Number of ms between transfers")
	helpPtr := flag.Bool("help", false, "if set prints usage help text")
	cliPathPtr := flag.String("cliPath", "", "location of cli.exe or unix equivalent")
	fileNamePtr := fileString("file", "", "optional file name for output (stdout used by default)")

	flag.Parse()

	if *helpPtr || *seedPtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *cliPathPtr != "" {
		cliPath = *cliPathPtr
	}

	if cliPath == "" {
		fmt.Println("no path to cli defined, please set GOPATH or use the -cliPath flag")
		os.Exit(1)
	}

	if DEBUG {
		fmt.Println("gopath:", goPath)
		fmt.Println("cliPath:", cliPath)
		fmt.Println("cliPathPtr:", *cliPathPtr)
		fmt.Println("seed:", *seedPtr)
		fmt.Println("numa:", *numaPtr)
		fmt.Println("ms:", *msPtr)
	}

	cmd.IninCmd(cliPath)

	var addresses []cipher.Address
	var secretKeys []cipher.SecKey

	addresses, secretKeys = createAddress(*seedPtr, *numaPtr)

	// hack to allow compilation.
	_ = secretKeys

	// dump out the addressses and their current balances.
	for _, addr := range addresses {
		fmt.Println("Address:", addr.String(), "Balance:", cmd.GetBalance(addr.String()).FloatString(COINDP))
	}
}

// createAddress() creats an arrage of addresses in the format xxxx
func createAddress(seed string, num int) ([]cipher.Address, []cipher.SecKey) {
	var addresses []cipher.Address
	var secretKeys []cipher.SecKey

	addresses = make([]cipher.Address, num)

	secretKeys = cipher.GenerateDeterministicKeyPairs([]byte(seed), num)

	for i := 0; i < num; i++ {
		addresses[i] = cipher.AddressFromSecKey(secretKeys[i])
	}

	return addresses, secretKeys

}
