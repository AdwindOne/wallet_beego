/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file eth-getBlockByNumber_test.go
 * @authors:
 *   Jérôme Laurens <jeromelaurens@gmail.com>
 * @date 2017
 */

package test

import (
	"testing"

	"github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

func TestEthGetBlockByNumber(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	blockNumber, err := connection.Eth.GetBlockNumber()

	block, err := connection.Eth.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if block == nil {
		t.Error("Block returned is nil")
		t.FailNow()
	}

	if block.Number.Int64() == 0 {
		t.Error("Block not found")
		t.FailNow()
	}

}
