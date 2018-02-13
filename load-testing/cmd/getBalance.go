package cmd

import (
	"encoding/json"
	"math/big"

	"github.com/skycoin/services/load-testing/loggy"
)

type BalanceResponse struct {
	Confirmed Balance `json:confimed`
	Spendable Balance `json:spendable`
	Expected  Balance `json:expected`
}

type Balance struct {
	Coins string `json:coins`
	Hours string `json:hours`
}
type Address struct {
	confirmed Balance
	spendable Balance
	expected  BalanceResponse
	adddress  []string
}

// getBalanace expects and address string such as "23J2eDUqw6toyHjwEkZQuWW8knVXbvCxhKu"
// It returns the confirmed balance.
func GetBalance(address string) *big.Rat {

	var jsonStr []byte

	jsonStr = runCmd("addressBalance", address)

	var m BalanceResponse

	err := json.Unmarshal(jsonStr, &m)

	if err != nil {
		loggy.Error.Println("could not unmarshal coins", err, "json:", jsonStr)
	}

	var coins = new(big.Rat)
	var worked bool = false

	coins, worked = coins.SetString(m.Confirmed.Coins)

	if !worked {
		loggy.Error.Println("could not convert coins", m.Spendable)
	}

	// TODO handle parse error

	return (coins)

}
