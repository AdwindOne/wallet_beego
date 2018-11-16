package rpc

import (
	"coinlib/rpc"
)

// BitcoinRPC is a warpper of btc/ltc/bcc/usdt.. rpc client.
type BitcoinRPC struct {
	client *rpc.Client
}

// New returns bitcoin rpc client.
func New(rpcClient *rpc.Client) *BitcoinRPC {
	return &BitcoinRPC{
		client: rpcClient,
	}
}

// DialHTTP is a wrapper of rpc.DialHTTP.
func DialHTTP(url string) *BitcoinRPC {
	rpcClient, err := rpc.DialHTTP(url)
	if err != nil {
		return nil
	}
	return &BitcoinRPC{
		client: rpcClient,
	}
}

// GetBestBlockHash returns the bestblockhash.
func (rpc *BitcoinRPC) GetBestBlockHash() (string, error) {
	var (
		bestBlockHash string
		err           error
	)

	err = rpc.client.Call("getbestblockhash", nil, &bestBlockHash)
	return bestBlockHash, err
}

// listaccounts returns the accoutns.
func (rpc *BitcoinRPC) ListAccounts() (interface{}, error) {
	var (
		accounts interface{}
		err           error
	)

	err = rpc.client.Call("listaccounts", nil, &accounts)
	return accounts, err
}

// getaddressesbyaccount returns the addresses.
func (rpc *BitcoinRPC) GetAddressesByAccount(name string) (interface{}, error) {
	var (
		address interface{}
		err           error
	)

	err = rpc.client.Call("getaddressesbyaccount",name , &address)
	return address, err
}

// GetAccountAddress returns the address.
func (rpc *BitcoinRPC) GetAccountAddress(name string) (interface{}, error) {
	var (
		address interface{}
		err           error
	)

	err = rpc.client.Call("getaccountaddress",name , &address)
	return address, err
}

// GetBlockByHash returns block infomations by hash.
func (rpc *BitcoinRPC) GetBlockByHash(h string) ([]byte, error) {
	var (
		blockData []byte
		err       error
	)
	err = rpc.client.Call("getblock", h, &blockData)
	return blockData, err
}

// GetFullBlockByHash returns block full infomations by hash.
func (rpc *BitcoinRPC) GetFullBlockByHash(h string) ([]byte, error) {
	var (
		blockData []byte
		err       error
	)
	err = rpc.client.Call("getblock", []interface{}{h, 2}, &blockData)
	return blockData, err
}

// GetBlockByHeight returns block infomations by height.
func (rpc *BitcoinRPC) GetBlockByHeight(h uint64) ([]byte, error) {
	var (
		blockHash string
		blockData []byte
		err       error
	)
	blockHash, err = rpc.GetBlockHash(h)
	if err != nil {
		return blockData, err
	}
	blockData, err = rpc.GetBlockByHash(blockHash)
	return blockData, err
}

// GetFullBlockByHeight returns block full infomations by height.
func (rpc *BitcoinRPC) GetFullBlockByHeight(h uint64) ([]byte, error) {
	var (
		blockHash string
		blockData []byte
		err       error
	)
	blockHash, err = rpc.GetBlockHash(h)
	if err != nil {
		return blockData, err
	}
	blockData, err = rpc.GetFullBlockByHash(blockHash)
	return blockData, err
}

// GetBlockHash returns block hash with block height.
func (rpc *BitcoinRPC) GetBlockHash(height uint64) (string, error) {
	var (
		blockHash string
		err       error
	)

	err = rpc.client.Call("getblockhash", height, &blockHash)
	return blockHash, err
}

// GetRawTransaction returns raw transaction by transaction hash.
func (rpc *BitcoinRPC) GetRawTransaction(h string) ([]byte, error) {
	var (
		tx  []byte
		err error
	)

	err = rpc.client.Call("getrawtransaction", []interface{}{h, 1}, &tx)
	return tx, err
}

// SendToAddress sends coin to dest address.
func (rpc *BitcoinRPC) SendToAddress(addr, amount string) (interface{}, error) {
	var (
		txid interface{}
		err  error
	)
	rpc.client.Call("sendtoaddress", []interface{}{addr, amount}, &txid)
	return txid, err
}

// unlockwallet return nil.
func (rpc *BitcoinRPC) WalletPassphrase(password, timeout string) (interface{}, error) {
	var (
		result interface{}
		err  error
	)
	rpc.client.Call("walletpassphrase", []interface{}{password, timeout}, &result)
	return result, err
}

// OmniListBlockTransactions returns the omnilayer transactions in block.
func (rpc *BitcoinRPC) OmniListBlockTransactions(height int64) ([]byte, error) {
	var (
		blockTxs []byte
		err      error
	)

	err = rpc.client.Call("omni_listblocktransactions", height, &blockTxs)
	return blockTxs, err
}

// OmniGetTransaction returns omnilayer raw transaction.
func (rpc *BitcoinRPC) OmniGetTransaction(h string) ([]byte, error) {
	var (
		omniTx []byte
		err    error
	)
	err = rpc.client.Call("omni_gettransaction", h, &omniTx)
	return omniTx, err
}

// Close closes rpc connection.
func (rpc *BitcoinRPC) Close() {
	rpc.client.Close()
}
