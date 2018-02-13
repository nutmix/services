package cmd

import (
	"encoding/json"

	"github.com/skycoin/services/load-testing/loggy"
)

type StatusResponse struct {
	Running               bool   `json:running`
	Num_of_blocks         int64  `json:num_of_blocks`
	Hash_of_last_block    string `json:running`
	Time_since_last_block string `json:hash_of_last_block`
	Webrpc_address        string `json:webrpc_address`
}

func GetStatus() bool {

	var jsonStr []byte

	jsonStr = runCmd("status")

	var m StatusResponse

	err := json.Unmarshal(jsonStr, &m)

	if err != nil {
		loggy.Error.Println("could not unmarshal status", err, "json:", jsonStr)
	}

	var running bool = false

	running = m.Running

	return (running)

}
