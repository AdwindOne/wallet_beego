/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/13/18
 * Time: 11:58
 */
package controllers

type PersonalController struct {
	base
}
//列出节点账户地址
func (e *PersonalController) ListAccounts() {
	accounts, err := e.web.Personal.ListAccounts()
	if err != nil {
		e.Err("内部异常",err)
	}
	e.Success("账户列表请求成功!",accounts)
}
//创建钱包地址
func (e *PersonalController) CreateAccount() {
	address, err := e.web.Personal.NewAccount("123")
	if err != nil {
		e.Err("创建异常",err)
	}
	e.Success("账户创建成功 !",address)
}
//解锁钱包地址
func (e *PersonalController) UnlockAccount(){
	var address string = "0x9f04aea9cf11c15c74ad5b1e93e556c8f02c786f"
	re, err := e.web.Personal.UnlockAccount(address, "123", 60)
	if err != nil {
		e.Err("解锁异常",err)
	}
	e.Success("解锁成功！",re)
}

//发起交易
func (e *PersonalController) SendTransaction() {

	//coinbase, err := e.web.Eth.GetCoinbase()
	//if err != nil {
	//	e.Err("coinbase 获取失败",err)
	//}
	////coinbase := "0x246cb955901d36e6be379c911221aac4ac022722"
	//transaction := new(*dto.TransactionParameters)
	//transaction.From = coinbase
	//transaction.To = coinbase
	//transaction.Value = big.NewInt(10)
	//transaction.Gas = big.NewInt(40000)
	//txhash, err := e.web.Personal.SendTransaction(transaction,"")
	//if err!=nil {
	//	e.Err("交易发起异常",err)
	//}
	//e.Success("交易hash获取成功",txhash)
}
