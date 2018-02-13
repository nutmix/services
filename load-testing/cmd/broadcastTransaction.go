package cmd

// broadcastTransaction takes the raw transaction from createTransqaction and returns teh transaction hash (no json)
func broadcastTransaction(trx string) []byte {

	var raw []byte

	raw = runCmd("broadcastTransaction", trx)

	return (raw)

}
