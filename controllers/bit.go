/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/15/18
 * Time: 11:16
 */
package controllers

import (
	"wallet_beego/utils"
)

type BitController struct {
	base
}
//获取最新区块hash
func (b *BitController) GetBestBlockHash()  {
	hash, err := b.bit.GetBestBlockHash()
	if err != nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,hash)
}
//获取区块的信息
func (b *BitController) GetBlockByHash() {
	hash, _ := b.bit.GetBestBlockHash()
	block, err := b.bit.GetBlockByHash(hash)
	if err != nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,utils.DecodeByte(block))
}
//通过区块号获取block信息
func (b *BitController) GetBlockByHeight() {
	bytes, err := b.bit.GetBlockByHeight(118)
	if err!=nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,utils.DecodeByte(bytes))
}
//获取所有账户
func (b *BitController) ListAccounts() {
	accounts, err := b.bit.ListAccounts()
	//fmt.Println(reflect.TypeOf(accounts))
	if err!=nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,accounts)
}
//获取账户子地址
func (b *BitController) GetAddressesByAccount() {
	account, err := b.bit.GetAddressesByAccount("yunlo")
	if err != nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,account)
}
//获取账户接受地址
func (b *BitController) GetAccountAddress() {
	address, err := b.bit.GetAccountAddress("yunlo")
	if err!=nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,address)
}
//转账给某个账户
func (b *BitController) SendToAddress() {
	address,_:= b.bit.GetAccountAddress("yunlo")
	s, err := b.bit.SendToAddress(address.(string), "0.01")
	if err!=nil {
		b.Err(Mfasle,err.Error())
	}
	b.Success(Mtrue,s)
}
//解锁账户
func (b *BitController) WalletPassphrase() {
	result, err := b.bit.WalletPassphrase("liuyunlong", "6000")
	if err != nil {
		b.Err(Mfasle,err)
	}
	b.Success(Mtrue,result)
}
