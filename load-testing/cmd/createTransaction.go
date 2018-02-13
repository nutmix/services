package cmd

import (
	"encoding/json"
	"math/big"

	"github.com/skycoin/services/load-testing/loggy"
)

type Response struct {
	Rawtx string `json:rawtx`
}

func createRawTransaction(walletFile string, fromAddress string, toAddress string, amount big.Rat) string {

	var jsonStr []byte

	jsonStr = runCmd("createRawTransaction", "-j", "-f", walletFile, "-1", fromAddress, toAddress, amount.FloatString(6))

	var m Response

	err := json.Unmarshal(jsonStr, &m)

	if err != nil {
		loggy.Error.Println("could not unmarshal json", err, "json:", jsonStr)
	}

	return (m.Rawtx)

}
