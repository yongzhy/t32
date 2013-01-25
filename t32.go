package t32

// #cgo linux,amd64 CFLAGS: -DT32HOST_LINUX_X64
// #cgo linux,386 CFLAGS: -DT32HOST_LINUX_X86
// #cgo windows,amd64 CFLAGS: -D_WIN64
// #cgo windows,386 CFLAGS: -D_WIN32
// #cgo windows CFLAGS: -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
// #cgo windows LDFLAGS: -lkernel32 -luser32 -lwsock32
// #include "t32.h"
// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	T32_DEV_OS  = 0
	T32_DEV_ICE = 1
	T32_DEV_ICD = 1 /* similar to ICE but clearer for user */
)

const (
	_INVALID_U64 = 0xFFFFFFFFFFFFFFFF
	_INVALID_S64 = -1
	_INVALID_U32 = 0xFFFFFFFF
	_INVALID_S32 = -1
	_INVALID_U16 = 0xFFFF
	_INVALID_S16 = -1
	_INVALID_U8  = 0xFF
	_INVALID_S8  = -1
)

type BreakPoint struct {
	Address uint32
	Enabled int8
	Type    uint32
	Auxtype uint32
}

func Config(name, value string) error {
	gname, gvalue := C.CString(name), C.CString(value)
	defer C.free(unsafe.Pointer(gname))
	defer C.free(unsafe.Pointer(gvalue))

	code, err := C.T32_Config(gname, gvalue)
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Config Error")
	}
	return nil
}

func Init() error {
	code, err := C.T32_Init()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Init Error")
	}
	return nil
}

func Exit() error {
	code, err := C.T32_Exit()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Exit Error")
	}
	return nil
}

func Attach(dev int) error {
	code, err := C.T32_Attach(C.int(dev))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Attach Error")
	}
	return nil
}

func Nop() error {
	code, err := C.T32_Nop()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Nop Error")
	}
	return nil
}

func Ping() error {
	code, err := C.T32_Ping()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Ping Error")
	}
	return nil
}

func Cmd(command string) error {
	gcmd := C.CString(command)
	defer C.free(unsafe.Pointer(gcmd))
	code, err := C.T32_Cmd(gcmd)
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Cmd Error")
	}
	return nil
}

func CmdWin(win uint32, command string) error {
	gcmd := C.CString(command)
	defer C.free(unsafe.Pointer(gcmd))
	code, err := C.T32_CmdWin(C.dword(win), gcmd)
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_WinCmd Error")
	}
	return nil
}

func Stop() error {
	code, err := C.T32_Stop()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Stop Error")
	}
	return nil
}

func EvalGet() (uint32, error) {
	var eval uint32
	code, err := C.T32_EvalGet((*C.dword)(&eval))
	if err != nil {
		return _INVALID_U32, err
	} else if code != 0 {
		return _INVALID_U32, errors.New("T32_EvalGet Error")
	}
	return eval, nil
}

func GetMessage() (string, uint16, error) {
	var status uint16
	var msg = make([]byte, 256)
	code, err := C.T32_GetMessage((*C.char)(unsafe.Pointer(&msg[0])), (*C.word)(&status))
	if err != nil {
		return "", _INVALID_U16, err
	} else if code != 0 {
		return "", _INVALID_U16, errors.New("T32_GetMessage Error")
	}
	return string(msg), status, nil
}

func Terminate(retval int) error {
	code, err := C.T32_Terminate(C.int(retval))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Terminate Error")
	}
	return nil
}

func GetPracticeState() (int32, error) {
	var pstate int32
	code, err := C.T32_GetPracticeState((*C.int)(&pstate))
	if err != nil {
		return _INVALID_S32, err
	} else if code != 0 {
		return _INVALID_S32, errors.New("T32_GetPracticeState Error")
	}
	return pstate, nil
}

func SetMode(mode int) error {
	code, err := C.T32_SetMode(C.int(mode))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_SetMode Error")
	}
	return nil
}

func GetState() (int32, error) {
	var state int32
	code, err := C.T32_GetState((*C.int)(&state))
	if err != nil {
		return _INVALID_S32, err
	} else if code != 0 {
		return _INVALID_S32, errors.New("T32_GetState Error")
	}
	return state, nil
}

func GetCpuInfo() (string, uint16, uint16, uint16, error) {
	var fpu, pendian, ptype uint16
	var pstring string
	gstr := C.CString(pstring)
	code, err := C.T32_GetCpuInfo((**C.char)(unsafe.Pointer(&gstr)), (*C.word)(&fpu), (*C.word)(&pendian), (*C.word)(&ptype))
	if err != nil {
		return "", _INVALID_U16, _INVALID_U16, _INVALID_U16, err
	} else if code != 0 {
		return "", _INVALID_U16, _INVALID_U16, _INVALID_U16, errors.New("T32_GetCpuInfo Error")
	}
	return pstring, fpu, pendian, ptype, nil
}

func GetRam() (uint32, uint32, uint16, error) {
	var pstart, pend uint32
	var paccess uint16
	code, err := C.T32_GetRam((*C.dword)(&pstart), (*C.dword)(&pend), (*C.word)(&paccess))
	if err != nil {
		return _INVALID_U32, _INVALID_U32, _INVALID_U16, err
	} else if code != 0 {
		return _INVALID_U32, _INVALID_U32, _INVALID_U16, errors.New("T32_GetRam Error")
	}
	return pstart, pend, paccess, nil
}

func ResetCPU() error {
	code, err := C.T32_ResetCPU()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_ResetCPU Error")
	}
	return nil
}

func ReadMemory(address uint32, access, size int) ([]byte, error) {
	var buffer = make([]byte, size)
	code, err := C.T32_ReadMemory(C.dword(address), C.int(access), (*C.byte)(unsafe.Pointer(&buffer[0])), C.int(size))
	if err != nil {
		return nil, err
	} else if code != 0 {
		return nil, errors.New("T32_ReadMemory Error")
	}
	return buffer, nil
}

func WriteMemory(address uint32, access int, buffer []byte, size int) error {
	code, err := C.T32_WriteMemory(C.dword(address), C.int(access), (*C.byte)(unsafe.Pointer(&buffer[0])), C.int(size))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_WriteMemory Error")
	}
	return nil
}

func WriteMemoryPipe(address uint32, access int, buffer []byte, size int) error {
	code, err := C.T32_WriteMemoryPipe(C.dword(address), C.int(access), (*C.byte)(unsafe.Pointer(&buffer[0])), C.int(size))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_WriteMemoryPipe Error")
	}
	return nil
}

func ReadRegister(mask1, mask2 uint32) ([]uint32, error) {
	var buffer = make([]uint32, 64)
	code, err := C.T32_ReadRegister(C.dword(mask1), C.dword(mask2), (*C.dword)(unsafe.Pointer(&buffer[0])))
	if err != nil {
		return nil, err
	} else if code != 0 {
		return nil, errors.New("T32_ReadRegister Error")
	}
	return buffer, nil
}

func ReadRegisterByName(name string) (uint32, uint32, error) {
	var gname = C.CString(name)
	defer C.free(unsafe.Pointer(gname))
	var value, hvalue uint32
	code, err := C.T32_ReadRegisterByName(gname, (*C.dword)(&value), (*C.dword)(&hvalue))
	if err != nil {
		return _INVALID_U32, _INVALID_U32, err
	} else if code != 0 {
		return _INVALID_U32, _INVALID_U32, errors.New("T32_ReadRegisterByName Error")
	}
	return value, hvalue, nil
}

func WriteRegister(mask1, mask2 uint32, buffer []uint32) error {
	code, err := C.T32_WriteRegister(C.dword(mask1), C.dword(mask2), (*C.dword)(unsafe.Pointer(&buffer[0])))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_WriteRegister Error")
	}
	return nil
}

func ReadPP() (uint32, error) {
	var pp uint32
	code, err := C.T32_ReadPP((*C.dword)(&pp))
	if err != nil {
		return _INVALID_U32, err
	} else if code != 0 {
		return _INVALID_U32, errors.New("T32_ReadPP Error")
	}
	return pp, nil
}

func ReadBreakpoint(address uint32, access, size int) ([]uint16, error) {
	var buffer = make([]uint16, size)
	code, err := C.T32_ReadBreakpoint(C.dword(address), C.int(access), (*C.word)(unsafe.Pointer(&buffer[0])), C.int(size))
	if err != nil {
		return nil, err
	} else if code != 0 {
		return nil, errors.New("T32_ReadBreakpoint Error")
	}
	return buffer, nil
}

func WriteBreakpoint(address uint32, access, breakpoint, size int) error {
	code, err := C.T32_WriteBreakpoint(C.dword(address), C.int(access), C.int(breakpoint), C.int(size))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_WriteBreakpoint Error")
	}
	return nil
}

func GetBreakpointList(max int) (int32, []BreakPoint, error) {
	bps := make([]_Ctype_T32_Breakpoint, max)
	// bps := make([]C.struct_T32_Breakpoint, max)
	var numbps int32
	code, err := C.T32_GetBreakpointList((*C.int)(&numbps), (*_Ctype_T32_Breakpoint)(unsafe.Pointer(&bps[0])), C.int(max))
	// code, err := C.T32_GetBreakpointList((*C.int)(&numbps), (*C.struct_T32_Breakpoint)(unsafe.Pointer(&bps[0])), C.int(max))
	if err != nil {
		return _INVALID_S32, nil, err
	} else if code != 0 {
		return _INVALID_S32, nil, errors.New("T32_GetBreakpointList Error")
	}
	if numbps > 0 {
		var gbps = make([]BreakPoint, numbps)
		for i := 0; i < int(numbps); i++ {
			gbps[i].Address = uint32(bps[i].address)
			gbps[i].Auxtype = uint32(bps[i].auxtype)
			gbps[i].Enabled = int8(bps[i].enabled)
			gbps[i].Type = uint32(bps[i]._type)
		}
		return numbps, gbps, nil
	}
	return 0, nil, nil
}

func Step() error {
	code, err := C.T32_Step()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Step Error")
	}
	return nil
}

func StepMode(mode int) error {
	code, err := C.T32_StepMode(C.int(mode))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_StepMode Error")
	}
	return nil
}

func Go() error {
	code, err := C.T32_Go()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Go Error")
	}
	return nil
}

func Break() error {
	code, err := C.T32_Break()
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Break Error")
	}
	return nil
}

func GetTriggerMessage() (string, error) {
	var msg = make([]byte, 256)
	code, err := C.T32_GetTriggerMessage((*C.char)(unsafe.Pointer(&msg[0])))
	if err != nil {
		return "", err
	} else if code != 0 {
		return "", errors.New("T32_GetTriggerMessage Error")
	}
	return string(msg), nil
}

func GetSymbol(name string) (uint32, uint32, uint32, error) {
	var address, size, access uint32
	gname := C.CString(name)
	defer C.free(unsafe.Pointer(gname))
	code, err := C.T32_GetSymbol(gname, (*C.dword)(&address), (*C.dword)(&size), (*C.dword)(&access))
	if err != nil {
		return _INVALID_U32, _INVALID_U32, _INVALID_U32, err
	} else if code != 0 {
		return _INVALID_U32, _INVALID_U32, _INVALID_U32, errors.New(fmt.Sprintf("T32_GetSymbol Error : %s", name))
	}
	return address, size, access, nil
}

func ReadVariableValue(name string) (uint32, uint32, error) {
	var value, hvalue uint32
	gname := C.CString(name)
	defer C.free(unsafe.Pointer(gname))
	code, err := C.T32_ReadVariableValue(gname, (*C.dword)(&value), (*C.dword)(&hvalue))
	if err != nil {
		return _INVALID_U32, _INVALID_U32, err
	} else if code != 0 {
		return _INVALID_U32, _INVALID_U32, errors.New(fmt.Sprintf("T32_ReadVariableValue Error : %s", name))
	}
	return value, hvalue, nil
}

func ReadVariableString(symbol string, maxlen int) (string, error) {
	var gsymbol = C.CString(symbol)
	defer C.free(unsafe.Pointer(gsymbol))
	var buffer = make([]byte, maxlen)
	code, err := C.T32_ReadVariableString(gsymbol, (*C.char)(unsafe.Pointer(&buffer[0])), C.int(maxlen))
	if err != nil {
		return "", err
	} else if code != 0 {
		return "", errors.New("T32_ReadVariableString Error")
	}
	return string(buffer), nil
}

func GetSource(address uint32) (string, uint32, error) {
	var filename = make([]byte, 256)
	var line uint32
	code, err := C.T32_GetSource(C.dword(address), (*C.char)(unsafe.Pointer(&filename[0])), (*C.dword)(&line))
	if err != nil {
		return "", _INVALID_U32, err
	} else if code != 0 {
		return "", _INVALID_U32, errors.New("T32_GetSource Error")
	}
	return string(filename), line, nil
}

func GetSelectedSource() (string, uint32, error) {
	var filename = make([]byte, 256)
	var line uint32
	code, err := C.T32_GetSelectedSource((*C.char)(unsafe.Pointer(&filename[0])), (*C.dword)(&line))
	if err != nil {
		return "", _INVALID_U32, err
	} else if code != 0 {
		return "", _INVALID_U32, errors.New("T32_GetSelectedSource Error")
	}
	return string(filename), line, nil
}

func AnaStatusGet() (uint8, int32, int32, int32, error) {
	var state uint8
	var size, min, max int32

	code, err := C.T32_AnaStatusGet((*C.byte)(&state), (*C.long)(&size), (*C.long)(&min), (*C.long)(&max))
	if err != nil {
		return _INVALID_U8, _INVALID_S32, _INVALID_S32, _INVALID_S32, err
	} else if code != 0 {
		return _INVALID_U8, _INVALID_S32, _INVALID_S32, _INVALID_S32, errors.New("T32_AnaStatusGet Error")
	}
	return state, size, min, max, nil
}

func AnaRecordGet(recordnr int32, length int) ([]byte, error) {
	var buffer = make([]byte, length)
	code, err := C.T32_AnaRecordGet(C.long(recordnr), (*C.byte)(unsafe.Pointer(&buffer[0])), C.int(length))
	if err != nil {
		return nil, err
	} else if code != 0 {
		return nil, errors.New("T32_AnaRecordGet Error")
	}
	return buffer, nil
}

func GetTraceState(tracetype int) (int32, int32, int32, int32, error) {
	var state int32
	var size, min, max int32
	code, err := C.T32_GetTraceState(C.int(tracetype), (*C.int)(&state), (*C.long)(&size), (*C.long)(&min), (*C.long)(&max))
	if err != nil {
		return _INVALID_U8, _INVALID_S32, _INVALID_S32, _INVALID_S32, err
	} else if code != 0 {
		return _INVALID_U8, _INVALID_S32, _INVALID_S32, _INVALID_S32, errors.New("T32_GetTraceState Error")
	}
	return state, size, min, max, nil
}

func GetSocketHandle() (int32, error) {
	var soc int32
	C.T32_GetSocketHandle((*C.int)(&soc))
	return soc, nil
}

func NotifyStateEnable(event int, fun func()) error {
	code, err := C.T32_NotifyStateEnable(C.int(event), (C.T32_NotificationCallback_t)(unsafe.Pointer(&fun)))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_Go Error")
	}
	return nil
}

func CheckStateNotify(param1 uint32) error {
	code, err := C.T32_CheckStateNotify(C.unsigned(param1))
	if err != nil {
		return err
	} else if code != 0 {
		return errors.New("T32_CheckStateNotify Error")
	}
	return nil
}

func ReadTrace(tracetype int, record int32, n int, mask uint32) ([]byte, error) {
	var buffer = make([]byte, 256*n)
	code, err := C.T32_ReadTrace(C.int(tracetype), C.long(record), C.int(n), C.ulong(mask), (*C.byte)(unsafe.Pointer(&buffer[0])))
	if err != nil {
		return nil, err
	} else if code != 0 {
		return nil, errors.New("T32_ReadTrace Error")
	}
	return buffer, nil
}
