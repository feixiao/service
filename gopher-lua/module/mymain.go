package main

import (
    "github.com/feixiao/learning/gopher-lua/module/mymodule"
    "github.com/yuin/gopher-lua"
)

func main() {
    L := lua.NewState()
    defer L.Close()
    L.PreloadModule("mymodule", mymodule.Loader)
    if err := L.DoFile("main.lua"); err != nil {
        panic(err)
    }
}
