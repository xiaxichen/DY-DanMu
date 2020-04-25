package lib

import (
	"bytes"
	"encoding/binary"
	"runtime"
	"strconv"
	"unsafe"
)

// IntToBytes:数字转byte组
func IntToBytes(bys int32, byteorder string) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, bys)
	BytesList := bytebuf.Bytes()
	switch byteorder != "" {
	case byteorder == "little":
		for i := 0; i < len(BytesList)/2; i++ {
			li := len(BytesList) - i - 1
			BytesList[i], BytesList[li] = BytesList[li], BytesList[i]
		}
	}
	return BytesList
}

// BytesToInt:byte组转数字
func BytesToInt(bys []byte, byteorder string) int {
	switch byteorder != "" {
	case byteorder == "little":
		for i := 0; i < len(bys)/2; i++ {
			li := len(bys) - i - 1
			bys[i], bys[li] = bys[li], bys[i]
		}
	}
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

// CheckIt:检查并转换interface 为 map string
func CheckIt(item interface{}) map[string]interface{} {
	itemMap, _ := item.(map[string]interface{})
	return itemMap
}

// TrunType:将string转为int32
func TrunType(str string) int64 {
	if str != "" {
		parseInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0
		} else {
			return parseInt
		}
	} else {
		return 0
	}
}

// RunFuncName:获取运行时函数名
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// Str2bytes:字符串转bytes
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
