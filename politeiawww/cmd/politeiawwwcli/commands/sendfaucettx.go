// Copyright (c) 2017-2019 The Fonero developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	"github.com/fonero-project/politeia/util"
)

// SendFaucetTxCmd uses the Fonero testnet faucet to send the specified amount
// of FNO (in atoms) to the specified address.
type SendFaucetTxCmd struct {
	Args struct {
		Address       string `positional-arg-name:"address" required:"true"` // FNO address
		Amount        uint64 `positional-arg-name:"amount" required:"true"`  // Amount in atoms
		OverrideToken string `positional-arg-name:"overridetoken"`           // Faucet override token
	} `positional-args:"true"`
}

// Execute executes the send faucet tx command.
func (cmd *SendFaucetTxCmd) Execute(args []string) error {
	address := cmd.Args.Address
	atoms := cmd.Args.Amount
	fno := float64(atoms) / 1e8

	if address == "" && atoms == 0 {
		return fmt.Errorf("Invalid arguments. Unable to pay %v FNO to %v",
			fno, address)
	}

	txID, err := util.PayWithTestnetFaucet(cfg.FaucetHost, address, atoms,
		cmd.Args.OverrideToken)
	if err != nil {
		return err
	}

	switch {
	case cfg.Silent:
		// Keep quite
	case cfg.RawJSON:
		fmt.Printf(`{"txid":"%v"}`, txID)
		fmt.Printf("\n")
	default:
		fmt.Printf("Paid %v FNO to %v with txID %v\n",
			fno, address, txID)
	}

	return nil
}

// sendFaucetTxHelpMsg is the output for the help command when 'sendfaucettx'
// is specified.
const sendFaucetTxHelpMsg = `sendfaucettx "address" "amount" "overridetoken"

Use the Fonero testnet faucet to send FNO (in atoms) to an address. One atom is
one hundred millionth of a single FNO (0.00000001 FNO).

Arguments:
1. address          (string, required)   Receiving address
2. amount           (uint64, required)   Amount to send (atoms)
3. overridetoken    (string, optional)   Override token for testnet faucet

Result:
Paid [amount] FNO to [address] with txID [transaction id]`
