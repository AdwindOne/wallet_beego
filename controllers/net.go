/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 11/14/18
 * Time: 18:13
 */
package controllers


type NetController struct {
	base
}
//获取节点数量
func (n *NetController) GetPeerCount()  {
	count, err := n.web.Net.GetPeerCount()
	if err != nil {
		n.Err(Mfasle,err)
	}
	n.Success(Mtrue,count)
}
//是否监听
func (e *NetController) IsListening(){
	isListening, err := e.web.Net.IsListening()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,isListening)
}

//net版本
func (e *NetController) GetVersion(){
	version, err := e.web.Net.GetVersion()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,version)
}
//获取客户端版本
func (e *NetController) ClientVersion(){
	clientVersion, err := e.web.ClientVersion()
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,clientVersion)
}

//sha3
func (e *NetController) Sha3(){
	sha3, err := e.web.Utils.Sha3("test")
	if err!=nil {
		e.Err(Mfasle,err)
	}
	e.Success(Mtrue,sha3)
}

