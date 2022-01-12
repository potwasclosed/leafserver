package msg

import (
	"github.com/name5566/leaf/network"
)

var Processor network.Processor  //定义处理过程,支持路由前拦截,路由到哪个模块 或者不路由, 相当于io层拦截了或者偷窥信息.

func init() {

}

//这里接受消息,然后路由到不同的模块???
//模块之间相互调用的话, GetMod("modeName+modid") ,返回的mod在底层是公用的, 减少直接引用依赖把mod当sess用就好了.
//或者 GetChan()

//type Module struct {  //这个结构要抽象到底层. 不然是私有的结构 包和包之间还是会互相引用
//	*module.Skeleton
//}

//chanrpc的一个容器可以放到一个单独的结构中维护,map[id/chanName]chanRpc而不是直接引用那个模块的变量这样不容易循环依赖.

/*
这个还可以再封装下 api更好用.
func handleRpcMsg(f func(interface{}), s *chanrpc.Server) func(interface{}) {
	return func(_ret interface{}) {
		type AsyncRet struct {
			cb  interface{}
			ret interface{}
		}
		asyncRet := &AsyncRet{
			cb:  f,
			ret: _ret,
		}
		s.Go(asyncRet)
		//A协程call B进程里返回结果,回调函数人在A协程里执行, 把回调和返回值 重新打包投递到chanrpc里执行
		//这个方法适合多  协程跨进程调用别的多协程.
	}
}

example := func() {
	hallses := service.GetRemoteService("hallBackend#1@zone1")
	req := &ssmsg.ServerCallReq{}
	_cb := func(ret interface{}) {
		//回调需要做的事情
	}
	rpc.Call(hallses, req, time.Second*30, handleRpcMsg(_cb, MS[0].ChanRPCServer))
}
	_ = example
*/

/*服务互联
两个服务互相连接,到底哪个作为Client /Server ?  
A调用B的话,A主动创建一个连接  B调用A的话,B主动创建一个连接,    这样调用别人 都从 自己主动想连接的会话管理器中获取对方的连接.
Server线程,想提供连接的话自己连接启动的时候上报给注册中心
client发出event, 服务端收到event处理,然后直接用event.ses回包就好了.
两个game互相连接的话可以这样

感觉还是加一个代理类比较好,处理当前进程??
比如game 当client 连接 网关, 又当server被 api服务器调用.
leaf 怎么处理的 ? 两个模块掉用是不同的clitne ,新建立  chanrpc的时候设置的那个.

每个game 都提供一个对外的game_proxy_zone1

[需要反向代理的]
tcp.svc connected 连接上之后,发一个消息告诉svr我是什么服务,然后svr添加,微服务需要 比如 gate和game 互相连接. 这种明显存在一个CS职责的身份
当两个同类型的服务器需要相互连接时如何处理? 每个svr都创建一个client连接对方? 暂时只想到类似的方法

消息路由没有自动生成代码的话,手动在一个地方注册也行,其余发包的地方就不用写send(msg,msgid了)

*/

