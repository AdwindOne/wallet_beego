package routers

import (
	"wallet_beego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
	"fmt"
)

func init() {
	beego.AutoRouter(&controllers.PersonalController{})
    beego.Router("/", &controllers.MainController{})
	beego.Router("/eth/acc",&controllers.PersonalController{})
	//personal
    beego.Router("/eth/CreateAccount",&controllers.PersonalController{},"*:CreateAccount")
    beego.Router("/eth/UnlockAccount",&controllers.PersonalController{},"*:UnlockAccount")
    beego.Router("/eth/ListAccounts",&controllers.PersonalController{},"*:ListAccounts")
    beego.Router("/eth/tranfer",&controllers.PersonalController{},"*:SendTransaction")

	//eth
    beego.Router("/eth/blocknumber",&controllers.EthController{},"*:Blocknumber")
	//todo
	beego.Router("/eth/contract",&controllers.EthController{},"*:Contract")
	beego.Router("/eth/estimategas",&controllers.EthController{},"*:EstimateGas")

    beego.Router("/eth/gasprice",&controllers.EthController{},"*:GetGasPrice")
    beego.Router("/eth/getbalance",&controllers.EthController{},"*:GetBalance")
    beego.Router("/eth/GetsBalance",&controllers.EthController{},"*:GetsBalance")
    beego.Router("/eth/GetBlockByHash",&controllers.EthController{},"*:GetBlockByHash")
    beego.Router("/eth/GetBlockByNumber",&controllers.EthController{},"*:GetBlockByNumber")
    beego.Router("/eth/GetBlockTransactionCountByHash",&controllers.EthController{},"*:GetBlockTransactionCountByHash")
    beego.Router("/eth/GetBlockTransactionCountByNumber",&controllers.EthController{},"*:GetBlockTransactionCountByNumber")
    beego.Router("/eth/GetCode",&controllers.EthController{},"*:GetCode")
    beego.Router("/eth/GetTransactionByBlockHashAndIndex",&controllers.EthController{},"*:GetTransactionByBlockHashAndIndex")
    beego.Router("/eth/GetTransactionByBlockNumberAndIndex",&controllers.EthController{},"*:GetTransactionByBlockNumberAndIndex")
    beego.Router("/eth/GetTransactionByHash",&controllers.EthController{},"*:GetTransactionByHash")
    beego.Router("/eth/GetTransactionCount",&controllers.EthController{},"*:GetTransactionCount")
    beego.Router("/eth/GetTransactionReceipt",&controllers.EthController{},"*:GetTransactionReceipt")
    beego.Router("/eth/GetUncleCountByBlockHash",&controllers.EthController{},"*:GetUncleCountByBlockHash")
    beego.Router("/eth/GetUncleCountByBlockNumber",&controllers.EthController{},"*:GetUncleCountByBlockNumber")
    beego.Router("/eth/GetHashRate",&controllers.EthController{},"*:GetHashRate")
    beego.Router("/eth/IsMining",&controllers.EthController{},"*:IsMining")
    beego.Router("/eth/GetProtocolVersion",&controllers.EthController{},"*:GetProtocolVersion")
    beego.Router("/eth/SendTransaction",&controllers.EthController{},"*:SendTransaction")
    beego.Router("/eth/SignTransaction",&controllers.EthController{},"*:SignTransaction")
    beego.Router("/eth/IsSyncing",&controllers.EthController{},"*:IsSyncing")

	//net
    beego.Router("/net/GetPeerCount",&controllers.NetController{},"*:GetPeerCount")
    beego.Router("/net/IsListening",&controllers.NetController{},"*:IsListening")
    beego.Router("/net/GetVersion",&controllers.NetController{},"*:GetVersion")
    beego.Router("/net/ClientVersion",&controllers.NetController{},"*:ClientVersion")
    beego.Router("/net/Sha3",&controllers.NetController{},"*:Sha3")

	//bitcoin
    beego.Router("/bit/GetBestBlockHash",&controllers.BitController{},"*:GetBestBlockHash")
    beego.Router("/bit/GetBlockByHash",&controllers.BitController{},"*:GetBlockByHash")
    beego.Router("/bit/GetBlockByHeight",&controllers.BitController{},"*:GetBlockByHeight")
    beego.Router("/bit/ListAccounts",&controllers.BitController{},"*:ListAccounts")
    beego.Router("/bit/GetAddressesByAccount",&controllers.BitController{},"*:GetAddressesByAccount")
    beego.Router("/bit/GetAccountAddress",&controllers.BitController{},"*:GetAccountAddress")
    beego.Router("/bit/SendToAddress",&controllers.BitController{},"*:SendToAddress")
    beego.Router("/bit/WalletPassphrase",&controllers.BitController{},"*:WalletPassphrase")



    beego.Get("btc", func(c *context.Context) {
		c.Output.Body([]byte("helle btc !"))
	})
    beego.Get("eth", func(ctx *context.Context) {
    	var connect = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))
		accounts, err := connect.Personal.ListAccounts()
		if err != nil {
			fmt.Println(err)
		}
		ctx.Output.JSON(accounts,true,true)
	})
}


