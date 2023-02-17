package main

import "testing"

type ModuleFetcherMock struct {
	nextOffset int
}

func (m *ModuleFetcherMock) fetch(offset int) *ModulesInfo {

	if offset < 7 {
		m.nextOffset = offset + 3
	}

	value := &ModulesInfo{Meta: struct {
		Limit         int    `json:"limit"`
		CurrentOffset int    `json:"current_offset"`
		NextOffset    int    `json:"next_offset"`
		NextURL       string `json:"next_url"`
	}{
		Limit:         3,
		CurrentOffset: offset,
		NextOffset:    m.nextOffset,
		NextURL:       "",
	},
	}
	return value
}

func (m *ModuleFetcherMock) iterator() ModuleFetcherIterator {
	return &ModuleFetcherIteratorMock{count: 0}
}

type ModuleWriterMock struct {
}

func (m *ModuleWriterMock) storeAllModules(modulesPath string, modulesInfo *ModulesInfo) {
}

func (m *ModuleWriterMock) storeModuleData(module *Module, fileName string) {
}

type ModuleFetcherIteratorMock struct {
	count int
}

func (it *ModuleFetcherIteratorMock) hasNext() bool {
	return it.count < 5
}

func (it *ModuleFetcherIteratorMock) next() *ModulesInfo {
	it.count += 1
	return &ModulesInfo{}
}

func TestAllItemsAreExtractedFromRegistryService(t *testing.T) {
	dummyUrl := "127.0.0.1/registry"
	modulesPath := "/tmp/modules.json"
	fetcherMock := &ModuleFetcherMock{}
	writerMock := &ModuleWriterMock{}
	rb := RegistryBrowser{dummyUrl,
		3,
		modulesPath,
		fetcherMock, writerMock}
	rb.LoopThroughModules()

}
