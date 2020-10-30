package luathread

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

// dsds
func LuaServer() {
	luaPool := newVMPool(100)

	defer luaPool.Shutdown()
	go MyWorker(luaPool)

}

// luaPool
func MyWorker(pool *lStatePool) {
	for {
		L := pool.Get()
		//defer pool.Put(L)
		/* your code here */
		// 加载fib.lua
		if err := L.DoFile("fib.lua"); err != nil {
			panic(err)
		}
		// 调用fib(n)
		err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("fib"), // 获取fib函数引用
			NRet:    1,                  // 指定返回值数量
			Protect: true,               // 如果出现异常，是panic还是返回err
		}, lua.LNumber(10)) // 传递输入参数n=10
		if err != nil {
			panic(err)
		}
		// 获取返回结果
		ret := L.Get(-1)
		// 从堆栈中扔掉返回结果
		L.Pop(-1)
		// 打印结果
		res, ok := ret.(lua.LNumber)
		if ok {
			fmt.Println(int(res))
		} else {
			fmt.Println("unexpected result")
		}
		pool.Put(L)

	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*
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

*/

func GetSensorData() //获取传感器数据

func SendDataToContollor() //发送数据
