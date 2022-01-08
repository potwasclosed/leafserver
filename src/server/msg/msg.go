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
