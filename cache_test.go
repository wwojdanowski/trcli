package main

import "testing"

type ModuleFetcherMock struct {
}

func (m ModuleFetcherMock) fetch(registryUrl string, limit int, offset int) *ModulesInfo {
	return &ModulesInfo{}
}

type ModuleWriterMock struct {
}

func (m ModuleWriterMock) storeAllModules(modulesPath string, modulesInfo *ModulesInfo) {
}

func (m ModuleWriterMock) storeModuleData(module *Module, fileName string) {
}

func TestAllItemsAreExtractedFromRegistryService(t *testing.T) {
	dummyUrl := "127.0.0.1/registry"
	modulesPath := "/tmp/modules.json"
	fetcherMock := ModuleFetcherMock{}
	writerMock := ModuleWriterMock{}
	rb := RegistryBrowser{dummyUrl,
		3,
		modulesPath,
		fetcherMock, writerMock}
	rb.LoopThroughModules()

}
