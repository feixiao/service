package main

import (
	lua "github.com/yuin/gopher-lua"
)

// Golang自定义结构体
type Person struct {
	Name string
}

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
}


func main() {
    L := lua.NewState()
    defer L.Close()
    registerPersonType(L)
    if err := L.DoString(`
        p = person.new("Steeve")
        print(p:name()) -- "Steeve"
        p:name("Alice")
        print(p:name()) -- "Alice"
    `); err != nil {
        panic(err)
    }
}