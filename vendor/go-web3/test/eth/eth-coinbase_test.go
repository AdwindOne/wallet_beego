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
 * @file eth-coinbase_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package test

import (
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
	"github.com/regcostajr/go-web3/providers"
)

func TestEthCoinbase(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	coinbase, err := connection.Eth.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_, err = connection.Eth.GetBalance(coinbase, block.LATEST)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

}
