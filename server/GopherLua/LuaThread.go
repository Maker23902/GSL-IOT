package main

import lua "github.com/yuin/gopher-lua"

// @ lua
// @err
func LuaServer() {
	L := lua.NewState() // 创建一个lua解释器实例
	defer L.Close()
	// 执行字符串语句
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}
}
