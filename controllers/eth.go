/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/13/18
 * Time: 18:20
 */
package controllers

import (
	"go-web3/eth/block"
	"math/big"
	"go-web3/dto"
)

type EthController struct {
	base
}
//获取区块儿号
func (e *EthController) Blocknumber() {
	number, err := e.web.Eth.GetBlockNumber()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,number)
	}

func (e *EthController) Coinbase() {
	coinbase, err := e.web.Eth.GetCoinbase()
	if err!=nil {
		e.Err(Mfasle,coinbase)
	}
	e.Success(Mtrue,coinbase)
}
//Todo
func (e *EthController) Contract() {

}
//Todo
func (e *EthController) EstimateGas()  {

}

func (e *EthController) GetGasPrice() {
	gasPrice, err := e.web.Eth.GetGasPrice()
	if err!=nil {
		e.Err(Mfasle,gasPrice)
	}
	e.Success(Mtrue,gasPrice)
}
//获取单个地址余额
func (e *EthController) GetBalance() {
	con, err := e.web.Eth.GetCoinbase()
	if err != nil {
		e.Err(Mfasle,err)
	}
	balance, err := e.web.Eth.GetBalance(con, block.LATEST)
	if err !=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,balance)
}
//获取所有余额
func (e *EthController) GetsBalance() {
	accounts, err := e.web.Personal.ListAccounts()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	var acc []*big.Int
	for i:=0;i<len(accounts);i++{
		balance, err := e.web.Eth.GetBalance(accounts[i],block.LATEST)
		if err!=nil {
			e.Err(Mfasle,err)
		}
		acc = append(acc,balance)

	}
	e.Success(Mtrue,acc)
}
//通过hash获取block
func (e *EthController) GetBlockByHash() {
	blockHash:="0x67e8f4385027c871b9586a4dd83d70ee8a508169e1592392c8eec5c01def5623"
	hash, err := e.web.Eth.GetBlockByHash(blockHash, false)
	if err != nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,hash)
}

//通过区块号获取block
func (e *EthController) GetBlockByNumber() {
	number, _ := e.web.Eth.GetBlockNumber()
	getBlockByNumber, err := e.web.Eth.GetBlockByNumber(number, false)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,getBlockByNumber)
}

//通过blockhash获取块里的交易hash数量
func (e *EthController) GetBlockTransactionCountByHash() {
	number, _ := e.web.Eth.GetBlockNumber()
	block, _ := e.web.Eth.GetBlockByNumber(number, false)
	txCount, err := e.web.Eth.GetBlockTransactionCountByHash(block.Hash)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,txCount)
}
//通过blocknumber 获取块儿里的交易数量
func (e *EthController) GetBlockTransactionCountByNumber() {
	number, _ := e.web.Eth.GetBlockNumber()
	txHash, err := e.web.Eth.GetBlockTransactionCountByNumber(block.NUMBER(number))
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,txHash)
}
//获取getcode
func (e *EthController) GetCode(){
	contractAddress:= "ox"
	code, err := e.web.Eth.GetCode(contractAddress, block.LATEST)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}

//
func (e *EthController) GetTransactionByBlockHashAndIndex(){
	var BlockHash  = ""
	var TransactionIndex *big.Int
	code, err := e.web.Eth.GetTransactionByBlockHashAndIndex(BlockHash,TransactionIndex)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}

//
func (e *EthController) GetTransactionByBlockNumberAndIndex(){
	var blockNumber *big.Int
	code, err := e.web.Eth.GetTransactionByBlockNumberAndIndex(blockNumber,big.NewInt(0))
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}
//
func (e *EthController) GetTransactionByHash(){
	var txID string
	code, err := e.web.Eth.GetTransactionByHash(txID)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}
//获取nonce
func (e *EthController) GetTransactionCount(){
	conbase, _ := e.web.Eth.GetCoinbase()
	count, err := e.web.Eth.GetTransactionCount(conbase,block.LATEST)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,count)
}

//获取交易收据
func (e *EthController) GetTransactionReceipt(){
	var hash string
	receipte, err := e.web.Eth.GetTransactionReceipt(hash)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,receipte)
}
//getUncleCount
func (e *EthController) GetUncleCountByBlockHash(){
	var hash string
	code, err := e.web.Eth.GetUncleCountByBlockHash(hash)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}
//getUncleCount
func (e *EthController) GetUncleCountByBlockNumber(){
	var number *big.Int
	code, err := e.web.Eth.GetUncleCountByBlockNumber(number)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}
//getHashRate
func (e *EthController) GetHashRate(){
	code, err := e.web.Eth.GetHashRate()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}
//节点是否在挖矿
func (e *EthController) IsMining(){
	code, err := e.web.Eth.IsMining()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}
//获取节点版本
func (e *EthController) GetProtocolVersion(){
	code, err := e.web.Eth.GetProtocolVersion()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,code)
}

func (e *EthController) SendTransaction(){
	coinbase, err := e.web.Eth.GetCoinbase()

	transaction := new(dto.TransactionParameters)
	transaction.From = coinbase
	transaction.To = coinbase
	transaction.Value = big.NewInt(0).Mul(big.NewInt(500), big.NewInt(1E18))
	transaction.Gas = big.NewInt(40000)

	//transaction.Data = types.ComplexString("p2p transaction")

	//txID, err := e.web.Eth.SendTransaction(transaction)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,coinbase)
}


func (e *EthController) SignTransaction(){
	coinbase, err := e.web.Eth.GetCoinbase()
	transaction := new(dto.TransactionParameters)
	transaction.Nonce = big.NewInt(5)
	transaction.From = coinbase
	transaction.To = coinbase
	transaction.Value = big.NewInt(0).Mul(big.NewInt(500), big.NewInt(1E18))
	transaction.Gas = big.NewInt(40000)
	transaction.GasPrice = big.NewInt(1E9)
	//transaction.Data = types.ComplexString("p2p transaction")

	//code, err := e.web.Eth.SignTransaction(transaction)
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,coinbase)
}

func (e *EthController) IsSyncing(){
	response, err := e.web.Eth.IsSyncing()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,response)
}