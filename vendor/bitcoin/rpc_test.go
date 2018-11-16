package rpc

import (
	"fmt"
	"testing"
	"sharex_08_10/chaincode/utils"
)

const (
	//url = "http://111:111@127.0.0.1:8332"
	url = "http://root:111111@127.0.0.1:44144"
)

func TestClient(t *testing.T) {
	client := DialHTTP(url)
	fmt.Println(client.GetBlockHash(4))
	 hash,error := client.GetBlockByHash("11c14228f368627354917dba69609ecf8d3efc45bf693f178d7184bf05a2ef15")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(utils.ToString(hash))

	fmt.Println(client.GetBlockByHash("11c14228f368627354917dba69609ecf8d3efc45bf693f178d7184bf05a2ef15"))
	fmt.Println(client.GetRawTransaction("cf9aed205810e71907cffdcc9f4afd52def2b3f65c0c04cbf73723e2ab5f7082"))
	// client.Do()
}
