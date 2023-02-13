package main

import "testing"

func TestAllItemsAreExtractedFromRegistryService(t *testing.T) {
	dummyUrl := "127.0.0.1/registry"
	modulesPath := "/tmp/modules.json"
	rb := RegistryBrowser{dummyUrl,
		3,
		modulesPath}
	rb.LoopThroughModules()

}
