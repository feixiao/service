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
    lv := L.ToInt(1)             	/* get argument */
    L.Push(lua.LNumber(lv * 2)) 	/* push result */
    return 1                     	/* number of results */
}