package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)[:len(filepath.Dir(b))-5] // cheap hack
)

func Basepath() string {
	return basepath
}

func PrintMyPath() {
	fmt.Println(basepath)
}
