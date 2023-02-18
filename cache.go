package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ModuleWriter interface {
	storeAllModules(modulesPath string, modulesInfo *ModulesInfo)
	storeModuleData(module *Module, fileName string)
}

type SimpleModuleWriter struct {
}

type RegistryBrowser struct {
	registryUrl string
	modulesPath string
	fetcher     ModuleFetcher
	writer      ModuleWriter
}

func (rb *RegistryBrowser) LoopThroughModules() {
	iterator := rb.fetcher.iterator()

	for iterator.hasNext() {
		modulesInfo := iterator.next()
		rb.storeAllModules(modulesInfo)
	}
}

func (rb *RegistryBrowser) storeAllModules(modulesInfo *ModulesInfo) {
	rb.writer.storeAllModules(rb.modulesPath, modulesInfo)
}

func (writer *SimpleModuleWriter) storeModuleData(module *Module, fileName string) {
	data, _ := json.Marshal(module)

	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", string(data))); err != nil {
		log.Println(err)
	}
}

func (writer *SimpleModuleWriter) storeAllModules(modulesPath string, modulesInfo *ModulesInfo) {
	for i := 0; i < len(modulesInfo.Modules); i++ {
		writer.storeModuleData(&modulesInfo.Modules[i], modulesPath)
	}
}

func RebuildModulesCacheDir(modulesDir string) *ModulesCache {
	mc := ModulesCache{modulesDir}
	if !mc.PathExists() {
		mc.BuildFullPath()
	} else {
		mc.DeleteCacheFile()
	}
	return &mc
}
