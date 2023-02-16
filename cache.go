package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ModuleFetcher interface {
	fetch(registryUrl string, limit int, offset int) *ModulesInfo
}

type ModuleWriter interface {
	storeAllModules(modulesPath string, modulesInfo *ModulesInfo)
	storeModuleData(module *Module, fileName string)
}

type SimpleModuleFetcher struct {
}

type SimpleModuleWriter struct {
}

type RegistryBrowser struct {
	registryUrl string
	limit       int
	modulesPath string
	fetcher     ModuleFetcher
	writer      ModuleWriter
}

func (rb *RegistryBrowser) fetch(offset int) *ModulesInfo {
	return rb.fetcher.fetch(rb.registryUrl, rb.limit, offset)
}

func (rb *RegistryBrowser) LoopThroughModules() {
	offset := 0

	shouldContinue := true

	for i := 0; shouldContinue; i++ {
		modulesInfo := rb.fetch(offset)
		rb.storeAllModules(modulesInfo)
		shouldContinue = modulesInfo.Meta.NextOffset > 0
		offset = modulesInfo.Meta.NextOffset
	}
}

func (rb *RegistryBrowser) storeAllModules(modulesInfo *ModulesInfo) {
	rb.writer.storeAllModules(rb.modulesPath, modulesInfo)
}

func (fetcher *SimpleModuleFetcher) fetch(registryUrl string, limit int, offset int) *ModulesInfo {
	response, err := http.Get(fmt.Sprintf("%s?limit=%d&offset=%d", registryUrl, limit, offset))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	modulesInfo := ModulesInfo{}
	_ = json.Unmarshal(responseBody, &modulesInfo)
	return &modulesInfo
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
