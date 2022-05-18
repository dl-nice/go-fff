// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/fff-chain/go-fff/cmd/utils"
	"github.com/fff-chain/go-fff/core"
	"github.com/fff-chain/go-fff/global_config"
	"os"
	"testing"
)

var customGenesisTests = []struct {
	genesis string
	query   string
	result  string
}{
	// Genesis file with an empty chain configuration (ensure missing fields work)
	{
		genesis: `{
			"alloc"      : {},
			"coinbase"   : "0x0000000000000000000000000000000000000000",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefd8",
			"nonce"      : "0x0000000000001338",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"timestamp"  : "0x00",
			"config"     : {}
		}`,
		query:  "eth.getBlock(0).nonce",
		result: "0x0000000000001338",
	},
	// Genesis file with specific chain configurations
	{
		genesis: `{
			"alloc"      : {},
			"coinbase"   : "0x0000000000000000000000000000000000000000",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefd8",
			"nonce"      : "0x0000000000001339",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"timestamp"  : "0x00",
			"config"     : {
				"homesteadBlock" : 42,
				"daoForkBlock"   : 141,
				"daoForkSupport" : true
			}
		}`,
		query:  "eth.getBlock(0).nonce",
		result: "0x0000000000001339",
	},
}
var test121 =`nohup ./geth --datadir ./data/  --http --http.addr 0.0.0.0 --http.port 8488 --http.api admin,debug,web3,eth,txpool,personal,ethash,miner,net  --http.corsdomain "*"  --ws --ws.addr 0.0.0.0 --ws.port 8588 --ws.api admin,debug,web3,eth,txpool,personal,ethash,miner,net --ws.origins  "*" --http --allow-insecure-unlock  &`

var test111 =`nohup ./geth  --http --http.addr 0.0.0.0 --http.port 8488 --http.api admin,debug,web3,eth,txpool,personal,ethash,miner,net  --http.corsdomain "*"  --ws --ws.addr 0.0.0.0 --ws.port 8588 --ws.api admin,debug,web3,eth,txpool,personal,ethash,miner,net --ws.origins  "*" --unlock FFF3nz4U8Fb3Qkfvt8GsGBH6qyck4pEZxwcouKHWDL5HRXB4Uj3SBrNG7N --password 111.txt --mine --http --allow-insecure-unlock  &`
// Tests that initializing Geth with a custom genesis block and chain definitions
// work properly.
func TestCustomGenesis(t *testing.T) {

	genesis := new(core.Genesis)
	if err := json.Unmarshal([]byte(global_config.GenesisJson), genesis); err != nil {
		utils.Fatalf("111invalid genesis file: %v", err)

	}
	// Query the custom genesis block
	args := []string{"", "--unlock", "FFF3nz4U8Fb3Qkfvt8GsGBH6qyck4pEZxwcouKHWDL5HRXB4Uj3SBrNG7N", "--password", "111.txt", "--mine", "--http", "--allow-insecure-unlock"}
	os.Args = args
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
