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
