package msg

import (
	"github.com/name5566/leaf/network"
)

var Processor network.Processor  //定义路由到哪个模块

func init() {

}

//这里接受消息,然后路由到不同的模块???
//模块之间相互调用的话, GetMod("modeName+modid") ,返回的mod在底层是公用的, 减少直接引用依赖
