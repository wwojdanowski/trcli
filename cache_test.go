package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type ModuleFetcherMock struct {
	nextOffset   int
	iteratorMock *ModuleFetcherIteratorMock
}

func (m *ModuleFetcherMock) fetch(offset int) *ModulesInfo {
	delta := 3
	var nextOffset int

	if offset < 7 {
		nextOffset = offset + delta
	} else {
		nextOffset = 0
	}

	value := &ModulesInfo{Meta: struct {
		Limit         int    `json:"limit"`
		CurrentOffset int    `json:"current_offset"`
		NextOffset    int    `json:"next_offset"`
		NextURL       string `json:"next_url"`
	}{
		Limit:         3,
		CurrentOffset: offset,
		NextOffset:    nextOffset,
		NextURL:       "",
	},
	}
	return value
}

func (m *ModuleFetcherMock) iterator() ModuleFetcherIterator {
	return m.iteratorMock
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
	fetcherMock := &ModuleFetcherMock{iteratorMock: &ModuleFetcherIteratorMock{count: 0}}
	writerMock := &ModuleWriterMock{}
	rb := RegistryBrowser{dummyUrl,
		modulesPath,
		fetcherMock,
		writerMock}

	assert.Equal(t, 0, fetcherMock.iteratorMock.count)
	assert.Equal(t, true, fetcherMock.iteratorMock.hasNext())

	rb.LoopThroughModules()

	assert.Equal(t, false, fetcherMock.iteratorMock.hasNext())
	assert.Equal(t, 5, fetcherMock.iteratorMock.count)
}

func TestRebuildModulesCacheDir(t *testing.T) {
	path, _ := os.MkdirTemp("", "modules_rebuild_cache")
	path = path + "/additional/subdirectories"

	RebuildModulesCacheDir(path)
	assert.DirExists(t, path)
	os.Create(path + "/modules.json")

	RebuildModulesCacheDir(path)
	assert.NoFileExists(t, path+"/modules.json")
}
