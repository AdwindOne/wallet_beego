package rpc

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	url = "http://111:111@127.0.0.1:30304"
)

func TestClient(t *testing.T) {
	client, err := DialHTTP(url)
	//var h = "6b4141b10342ff5d941c21d12e63614da7dd90e03359ec5c50df0c1fbf291e89"
	//var blockData []byte
	//client.Call("getblock", h, &blockData)
	//fmt.Println("client.call getblock", h, blockData)

	fmt.Println(json.Marshal(nil))
	fmt.Println(json.Marshal([]byte{}))
	fmt.Println(json.Marshal([0]string{}))
	var result float32
	fmt.Println("client.Call", client.Call("getbalance", nil, &result))
	var bestBlockHash string
	fmt.Println("client.Call", client.Call("getbestblockhash", nil, &bestBlockHash))
	fmt.Println("result", result, err, bestBlockHash, result)
	var blockhash []byte //= make([]byte, 1)
	var req = 1000
	fmt.Println("result2", result, err, blockhash, result)

	fmt.Println("result12", client.Call("getblockhash", req, &blockhash), blockhash)
	fmt.Println("result12", client.Call("getblockhash", req, &blockhash), blockhash)

	var blockData []byte
	fmt.Println("result 222", client.Call("getblock", string(blockhash), &blockData), string(blockData))
	// client.Do()
}
