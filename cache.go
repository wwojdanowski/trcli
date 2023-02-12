package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const REGISTRY_URL = "https://registry.terraform.io/v1/modules"

func loopThroughModules(modulesPath string) {
	offset := 0
	limit := 100

	shouldContinue := true

	for i := 0; shouldContinue; i++ {
		response, err := http.Get(fmt.Sprintf("%s?limit=%d&offset=%d", REGISTRY_URL, limit, offset))
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		responseBody, _ := ioutil.ReadAll(response.Body)
		modulesInfo := ModulesInfo{}
		_ = json.Unmarshal([]byte(responseBody), &modulesInfo)
		storeAllModules(modulesPath, &modulesInfo)
		shouldContinue = modulesInfo.Meta.NextOffset > 0
		offset = modulesInfo.Meta.NextOffset

	}
}

func storeAllModules(modulesFile string, modulesInfo *ModulesInfo) {
	for i := 0; i < len(modulesInfo.Modules); i++ {
		storeModuleData(&modulesInfo.Modules[i], modulesFile)
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
