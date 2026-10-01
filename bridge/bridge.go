package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"os"
	"unsafe"

	"github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/hub"
	"github.com/metacubex/mihomo/hub/executor"
	"github.com/metacubex/mihomo/log"
)

//export InitCore
func InitCore(homeDir *C.char) C.int {
	home := C.GoString(homeDir)
	constant.SetHomeDir(home)
	log.Infoln("[LuminClashCore] Mihomo core initialized with home: %s", home)
	return 0
}

//export StartCore
func StartCore(configPath *C.char) *C.char {
	cfg := C.GoString(configPath)
	rawCfg, err := os.ReadFile(cfg)
	if err != nil {
		return C.CString(err.Error())
	}

	if err := hub.Parse(rawCfg); err != nil {
		return C.CString(err.Error())
	}

	log.Infoln("[LuminClashCore] Mihomo core started successfully with external controller")
	return nil
}

//export StopCore
func StopCore() {
	executor.Shutdown()
	log.Infoln("[LuminClashCore] Mihomo core stopped")
}

//export GetCoreVersion
func GetCoreVersion() *C.char {
	return C.CString(constant.Version)
}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func main() {}
