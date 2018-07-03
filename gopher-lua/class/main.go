package main

import (
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	meta := L.NewState()
	L.SetGlobal("Counter",meta)

	// 注册函数
	L.SetField(meta,"new",L.NewFunction(newCounter))
	L.SetField(meta,"incr",L.NewFunction(incrCounter))
	L.SetField(meta,"get",L.NewFunction(getCounter))

	err := L.DoString(`
		counter = Counter:new(100)
		for i=1,10 do
			print(counter:incr(i))
		end
		print(counter:get())`
	)

	if err != nil {
		panic(err)
	}
}

func newCounter(L *lua.LState) int {
	c := L.NewState()
	self := L.CheckTable()
	value := lua.LNumber(0)
	// 第二个为可选参数
	if L.GetTop() >= 2 {
		value = L.CheckTable(2)
	}
	L.SetField(c,"value",value)
	L.SetMetatable(c,self)
	L.SetField(self,"__index",self)
	L.Push(c)
	// 返回[函数返回值的个数]
	return 1
}


func incrCounter(L *lua.LState) int {
	self := L.CheckTable(1)
	value := lua.LNumber(0)
	if L.GetTop() >= 2 {
		value = L.CheckNumber(2)
	}
	current := L.GetField(self,"value").(lua.LNumber)
	current += value
	L.SetField(self,"value",current)
	L.Push(current)
	return 1 // 返回参数的个数
}


func getCounter(L *lua.LState) {
	self := L.CheckTable(1)
	value := L.GetField(self,"value").(lua.LNumber)
	L.Push(value)
	return 1
}