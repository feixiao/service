package main 

import (
	lua "github.com/yuin/gopher-lua"
)


func main() {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("double", L.NewFunction(Double)) /* Original lua_setglobal uses stack... */
	// 在Lua中调用double(20)，实际是调用Double函数
	if err:= L.DoString(`print(double(20))`); err != nil {
		panic(err)
	}
}


func Double(L *lua.LState) int {
    lv := L.ToInt(1)             	// 获取第一个参数
    L.Push(lua.LNumber(lv * 2)) 	// 写入返回值
    return 1                     	// 返回
}