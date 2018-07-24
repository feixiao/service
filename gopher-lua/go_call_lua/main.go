package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"github.com/fatih/structs"
	"github.com/corpix/lua/mapper"
	
)
// Golang自定义结构体
type Person struct {
	Name string
}

/*
const luaPersonTypeName = "person"

func registerPersonType(L *lua.LState) {
	// 使用名字创建
	mt := L.NewTypeMetatable(luaPersonTypeName)
	L.SetGlobal(luaPersonTypeName,mt)
	// Lua中的new方法，对应Golang中的newPerson
	L.SetField(mt,"new",L.NewFunction(newPerson))
	// 管理内部函数(通过map映射)
	L.SetField(mt,"__index",L.SetFuncs(L.NewTable(),personMethods))
}

func newPerson(L *lua.LState) int {
	person := &Person{L.CheckString(1)}
	ud := L.NewUserData()
	ud.Value = person
	L.SetMetatable(ud,L.GetTypeMetatable(luaPersonTypeName))
	L.Push(ud)
	return 1
}

var personMethods = map[string]lua.LGFunction {
	"name":personGetSetName, 		// name方法关联personGetSetName方法
}

func checkPerson(L *lua.LState) *Person {
	ud := L.CheckUserData(1)
	if v,ok := ud.Value.(*Person); ok {
		return v
	}
	L.ArgError(1,"person expected")
	return nil
}

func personGetSetName(L *lua.LState) int {
	p := checkPerson(L)
	if L.GetTop() == 2 {
        p.Name = L.CheckString(2)
        return 0
    }
    L.Push(lua.LString(p.Name))
    return 1
} */


func main() {
	L := lua.NewState()
	defer L.Close()
	// 加载fib.lua
	if err := L.DoFile("set.lua"); err != nil {
		panic(err)
	}

	person := &Person{
		Name:"frank",
	}

	m := structs.Map(person)
	inputValue, err := mapper.ToValue(m)
	if err != nil {
		panic(err)
    }

	err = L.CallByParam( lua.P{
		Fn: L.GetGlobal("setPerson"),	// 获取fib函数引用
		NRet: 1,						// 指定返回值数量
		Protect: true,					// 如果出现异常，是panic还是返回err
	}, inputValue)						// 传递输入参数
	if err != nil {
		panic(err)
	}
	// 获取返回结果
	outputValue := L.Get(-1)
	L.Pop(1)

	output, err := mapper.FromValue(outputValue)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}