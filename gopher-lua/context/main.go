package main

import (
	"context"
	"time"
	lua "github.com/yuin/gopher-lua"
)

// 超时处理
func main() {
	L := lua.NewState()
	defer L.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// set the context to our LState
	L.SetContext(ctx)
	err := L.DoString(`
		local clock = os.clock
		function sleep(n)  -- seconds
		  local t0 = clock()
		  while clock() - t0 <= n do end
		end
		sleep(3)
	  `)

	if err != nil {
		panic(err)
	}
//	<- ctx.Done()

}