package luathread

import (
	"bufio"
	"os"
	"testing"
	"text/template/parse"

	lua "github.com/yuin/gopher-lua"
)

// dsds
func LuaServer() {
	pool := newVMPool(nil, 100)
	for {
		l := pool.get()
		_ = l.DoString(`print("hello")`)
		pool.put(l)
	}

}

// func MyWorker() {
// 	L := luaPool.Get()
// 	defer luaPool.Put(L)
// 	/* your code here */
// }

// func main1() {
// 	defer luaPool.Shutdown()
// 	go MyWorker()
// 	go MyWorker()
// 	/* etc... */
// }

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file.
func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

// Example shows how to share the compiled byte code from a lua script between multiple VMs.
func Example() {
	codeToShare, err := CompileLua("mylua.lua")
	a := lua.NewState()
	b := lua.NewState()
	c := lua.NewState()
	DoCompiledFile(a, codeToShare)
	DoCompiledFile(b, codeToShare)
	DoCompiledFile(c, codeToShare)
}

// 虚拟机实例池
func BenchmarkRunWithPool(b *testing.B) {
	pool := newVMPool(nil, 100)
	for i := 0; i < b.N; i++ {
		l := pool.get()
		_ = l.DoString(`a = 1 + 1`)
		pool.put(l)
	}
}
