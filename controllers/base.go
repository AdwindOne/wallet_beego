/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/13/18
 * Time: 17:32
 */
package controllers

import (
	"github.com/astaxie/beego"
	"wallet_beego/utils"
	"go-web3"
	"go-web3/providers"
	"bitcoin"
)

type base struct {
	beego.Controller
	web  *web3.Web3
	r  *utils.ReUtil
	bit *rpc.BitcoinRPC
}

const Mfasle  = "失败"
const Mtrue  = "成功"
const bit_url string ="http://root:111111@127.0.0.1:44144"
func (e *base) Success(m string,d interface{})   {
	e.r.Code=200
	e.r.Message=m
	e.r.Data = d
	e.Ctx.Output.JSON(e.r,true,true)
}

func (e *base) Err(m string,d interface{})   {
	e.r.Code=400
	e.r.Message=m
	e.r.Data = d
	e.Ctx.Output.JSON(e.r,true,true)
}

func (e *base) Prepare()() {
	url:=beego.AppConfig.String("url")
	e.web = web3.NewWeb3(providers.NewHTTPProvider(url, 10000, false))
	e.r = utils.NewReutil()
	e.bit = rpc.DialHTTP(bit_url)
}


