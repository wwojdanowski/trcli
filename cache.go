package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const REGISTRY_URL = "https://registry.terraform.io/v1/modules"

type RegistryBrowser struct {
	registryUrl string
	limit       int
	modulesPath string
}

func (rb *RegistryBrowser) fetch(offset int) *ModulesInfo {
	response, err := http.Get(fmt.Sprintf("%s?limit=%d&offset=%d", rb.registryUrl, rb.limit, offset))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	modulesInfo := ModulesInfo{}
	_ = json.Unmarshal(responseBody, &modulesInfo)
	return &modulesInfo
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
	for i := 0; i < len(modulesInfo.Modules); i++ {
		storeModuleData(&modulesInfo.Modules[i], rb.modulesPath)
	}
}

func storeModuleData(module *Module, fileName string) {
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
