package main

import (
	"time"
	lua "github.com/yuin/gopher-lua"
)


func receiver(ch, quit chan lua.LValue) {
    L := lua.NewState()
    defer L.Close()
    L.SetGlobal("ch", lua.LChannel(ch))
    L.SetGlobal("quit", lua.LChannel(quit))
    if err := L.DoString(`
    local exit = false
    while not exit do
      channel.select(
        {"|<-", ch, function(ok, v)
          if not ok then
            print("channel closed")
            exit = true
          else
            print("received:", v)
          end
        end},
        {"|<-", quit, function(ok, v)
            print("quit")
            exit = true
        end}
      )
    end
  `); err != nil {
        panic(err)
    }
}

func sender(ch, quit chan lua.LValue) {
    L := lua.NewState()
    defer L.Close()
    L.SetGlobal("ch", lua.LChannel(ch))
    L.SetGlobal("quit", lua.LChannel(quit))
    if err := L.DoString(`
    ch:send("1")
    ch:send("2")
  `); err != nil {
        panic(err)
    }
    ch <- lua.LString("3")
    quit <- lua.LTrue
}

func main() {
    ch := make(chan lua.LValue)
    quit := make(chan lua.LValue)
    go receiver(ch, quit)
    go sender(ch, quit)
    time.Sleep(3 * time.Second)
}