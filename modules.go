package main

import (
	"os"
)

type ModulesCache struct {
	path string
}

func (mc *ModulesCache) PathExists() bool {
	_, error := os.Stat(mc.path)
	if error != nil {
		return false
	}
	return true
}

func (mc *ModulesCache) BuildFullPath() {
	//dir := filepath.Dir(mc.path)
	os.MkdirAll(mc.path, 0755)
}
