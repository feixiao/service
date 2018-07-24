package mymodule

import (
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
    "myfunc": myfunc,
}

func myfunc(L *lua.LState) int {
    return 0	// 函数返回值个数
}

func Loader(L *lua.LState) int {
	 // 注册函数
	 mod := L.SetFuncs(L.NewTable(), exports)
	 // register other stuff
	 L.SetField(mod, "name", lua.LString("value"))
 
	 // returns the module
	 L.Push(mod)
	 return 1
}